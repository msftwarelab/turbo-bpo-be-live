package services

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SaveInvoice(ctx context.Context, input models.SaveInvoiceInput, myID, myUserFullName string, assignedQCId, assignedQCName *string) (string, error) {
	pipelineRaws, err := datastore.GetPipelineByOrderNumberDataStore(ctx, input.OrderNumber)
	if err != nil {
		return "", err
	}
	if len(pipelineRaws) == 0 {
		return "", errs.InvalidOrderNumber
	}
	newID, err := datastore.SaveInvoice(ctx, input, myUserFullName, pipelineRaws[0], assignedQCId, assignedQCName)
	if err != nil {
		return "", err
	}

	// Add entry on request
	pipelinePhotosTotals, err := datastore.GetPipelinePhotosCountGroupByPipelineId(ctx, []string{pipelineRaws[0].ID.Hex()})

	hasPhotos := false
	if len(pipelinePhotosTotals) > 0 {
		hasPhotos = true
	}

	datastore.AddRequest(ctx, pipelineRaws[0].ID.Hex(), myUserFullName, myID, pipelineRaws[0], hasPhotos)

	//Add invoice request history!
	requestHistory := datastore.RequestHistory{
		PipelineId:      pipelineRaws[0].ID.Hex(),
		Status:          constants.RequestStatusPending,
		ClientId:        strings.ObjectTOString(pipelineRaws[0].UserId),
		EmployeeId:      myID,
		Type:            strings.ToObject("INVOICE"),
		OrderNumber:     pipelineRaws[0].OrderNumber,
		Address:         pipelineRaws[0].Address,
		Company:         pipelineRaws[0].Company,
		Remarks:         input.Remarks,
		CreatedDateTime: pipelineRaws[0].CreatedDateTime,
	}

	datastore.AddRequestHistory(ctx, requestHistory)
	return newID, nil
}

func UpdateInvoice(ctx context.Context, id string, input models.UpdateInvoiceInput, updatedBy string) (bool, error) {
	return datastore.UpdateInvoice(ctx, id, input, updatedBy)
}

func CancelInvoice(ctx context.Context, id string, reason *string, cancelBy string) (bool, error) {
	return datastore.CancelInvoice(ctx, id, reason, cancelBy)
}

func AllInvoice(ctx context.Context, filter *models.InvoiceFilterInput, myRoles []*string, myUserId string) ([]*models.Invoice, error) {

	if contains(myRoles, constants.UserRoleContractor) && filter.EmployeeID == nil {
		filter.EmployeeID = &myUserId
	}
	return datastore.GetInvoices(ctx, filter)
}

func AllInvoiceRequest(ctx context.Context, myUserID string, filter *models.InvoiceRequestFilterInput) (*models.InvoiceRequestResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.InvoiceRequestFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawInvoices, err := datastore.SearchInvoiceRequests(ctx, *filter.Offset, *filter.Limit, myUserID, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve invoice: %v", err)
		return nil, err
	}

	count, err := datastore.GetInvoiceRequestsCount(ctx, myUserID, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve count of invoice: %v", err)
		return nil, err
	}

	invoiceRequests := make([]*models.Invoice, 0)
	for _, u := range rawInvoices {
		invoiceRequests = append(invoiceRequests, u.ToModels())
	}
	toInt := int(*count)
	return &models.InvoiceRequestResult{
		TotalCount: &toInt,
		Results:    invoiceRequests,
	}, nil
}

func AllInvoiceRequestHistory(ctx context.Context, myUserID string, filter *models.FilterInput) (*models.InvoiceRequestHistoryResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.FilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawInvoices, err := datastore.SearchInvoiceRequestHistories(ctx, *filter.Offset, *filter.Limit, myUserID)
	if err != nil {
		log.Debug("Failed to retrieve invoice: %v", err)
		return nil, err
	}

	count, err := datastore.GetInvoiceRequestHistoriesCount(ctx, myUserID)
	if err != nil {
		log.Debug("Failed to retrieve count of invoice: %v", err)
		return nil, err
	}

	invoiceRequests := make([]*models.InvoiceRequestHistory, 0)
	for _, u := range rawInvoices {
		invoiceRequest := &models.InvoiceRequestHistory{
			ID:            strings.ToObject(u.ID.Hex()),
			DateRequested: strings.ToObject(datastore.TimeConversion(&u.CreatedDateTime)),
			OrderNumber:   u.OrderNumber,
			Address:       u.Address,
			Company:       u.Company,
			Remarks:       u.Remarks,
			Status:        strings.ToObject(u.Status),
		}
		invoiceRequests = append(invoiceRequests, invoiceRequest)
	}
	toInt := int(*count)
	return &models.InvoiceRequestHistoryResult{
		TotalCount: &toInt,
		Results:    invoiceRequests,
	}, nil
}

func BatchUpdateInvoice(ctx context.Context) {
	log.Info("running on services")
	invoiceList, err := datastore.FilterInvoiceByType(ctx, "QC")
	if err != nil {
		panic(err)
	}
	log.Info("done  getting invoiceList")
	for _, v := range invoiceList {
		if v != nil {

			if v.EmployeeId != nil {
				userFilter := datastore.FilterById(*v.EmployeeId)
				userRaw, _ := datastore.GetUser(ctx, userFilter)

				if userRaw == nil {
					continue
				}

				filterUpdateInvoice := datastore.FilterById(v.ID.Hex())
				isSuccess, err := datastore.UpdateInvoiceForBatch(ctx, userRaw.FullName(), filterUpdateInvoice)
				if err != nil {
					log.Error("err : %v", err)
					continue
				}
				if isSuccess {
					log.Info("update success on invoice - ID: %+v  | employeeID - %+v | employeeName - %+v ", v.ID.Hex(), fmt.Sprint(v.EmployeeId), userRaw.FullName())
				}
			}

		}

	}

}
