package services

import (
	"context"
	"errors"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/excel"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/paypal"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func AllBilling(ctx context.Context, filter *models.BillingFilterInput, myRoles []*string, myUserID string) (*models.BillingResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)

	if filter == nil {
		filter = &models.BillingFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	if contains(myRoles, constants.UserRoleClient) {
		filter.UserID = strings.ToObject(myUserID)
	}

	rawBilling, err := datastore.SearchBillings(ctx, *filter.Offset, *filter.Limit, filter.DateFrom, filter.DateTo, filter.UserID, nil, filter.OrderNumber)
	if err != nil {
		log.Debug("Failed to retrieve loginLog: %v", err)
		return nil, err
	}

	count, err := datastore.GetBillingsCount(ctx, filter.DateFrom, filter.DateTo, filter.UserID, filter.OrderNumber)
	if err != nil {
		log.Debug("Failed to retrieve count of billing: %v", err)
		return nil, err
	}

	billings := make([]*models.Billing, 0)
	for _, u := range rawBilling {
		billings = append(billings, u.ToModels())
	}
	toInt := int(*count)
	return &models.BillingResult{
		TotalCount: &toInt,
		Results:    billings,
	}, nil
}

func SaveBilling(ctx context.Context, input models.SaveBillingInput, cratedBy string) (string, error) {
	return datastore.SaveBilling(ctx, input, cratedBy)
}

func UpdateBilling(ctx context.Context, id string, input models.UpdateBillingInput, updatedBy string) (bool, error) {
	return datastore.UpdateBilling(ctx, id, input)
}

func DeleteBilling(ctx context.Context, id string, deletedBy string) (bool, error) {
	return datastore.DeleteBilling(ctx, id)
}

func VerifyPaypalTransaction(ctx context.Context, paypalOrderID string, billingOrderID *string, myUserFullName string) (bool, error) {

	if billingOrderID != nil {
		err := paypal.ValidateTransaction(paypalOrderID)
		if err != nil {
			return false, err
		}
		_, err = datastore.UpdateBillingStatus(ctx, *billingOrderID, "PAID")
		if err != nil {
			return false, err
		}
		//Update pipelineStatus to Paid
		billingFilter := datastore.FilterById(*billingOrderID)
		billingRaw, err := datastore.GetBilling(ctx, billingFilter)
		if err != nil {
			return false, err
		}
		if billingRaw == nil {
			return false, errs.InvalidBillingId
		}
		orderNumbers := billingRaw.GetOrdernumbers()
		if len(orderNumbers) > 0 {
			pipelineFilter := datastore.FilterByOrderNumbers(orderNumbers)
			_, err := datastore.UpdatePipelineStatuses(ctx, pipelineFilter, constants.PipelineStatusPaid, nil)
			if err != nil {
				return false, err
			}
		}
		return true, nil

	} else {
		return true, nil
	}

}

func CreateBillingExcel(ctx context.Context, billingID string) (string, error) {

	billingFilter := datastore.FilterById(billingID)
	billingRaw, err := datastore.GetBilling(ctx, billingFilter)
	if err != nil || billingRaw.ID.IsZero() {
		return "", errors.New("invalid billing id")
	}

	fileLocation, err := excel.CreateBilling(ctx, billingRaw)
	if err != nil {
		return "", err
	}
	fileUrl, err := awsS3.S3UploaderFromFileExcel(fileLocation)
	if err != nil {
		return "", err
	}

	return strings.ObjectTOString(fileUrl), nil
}
