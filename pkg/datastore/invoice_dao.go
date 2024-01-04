package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveInvoice(ctx context.Context, input models.SaveInvoiceInput, myFullName string, pipelineraw *Pipeline, assignedQCId, assignedQCName *string) (string, error) {

	newInvoice := &Invoice{
		Status:           "ACTIVE",
		PipelineId:       pipelineraw.ID.Hex(),
		ClientId:         pipelineraw.UserId,
		EmployeeId:       pipelineraw.AssignId,
		Type:             strings.ToObject(input.Type),
		Name:             pipelineraw.Assign,
		OrderNumber:      pipelineraw.OrderNumber,
		Address:          pipelineraw.Address,
		Company:          pipelineraw.Company,
		Client:           pipelineraw.UserName,
		OrderType:        pipelineraw.OrderType,
		IsSuperRush:      pipelineraw.IsSuperRush,
		IsRush:           pipelineraw.IsRushOrder,
		IsInterior:       pointers.Bool(false),
		IsRentalAddendum: pointers.Bool(false),
		IsInitialBpo:     pipelineraw.IsInitialBpo,
		IsInspection:     pipelineraw.IsInspection,
		IsNoCsv:          pointers.Bool(false),
		IsNoIFill:        pointers.Bool(false),
		IsOtherPremium:   pointers.Bool(false),
		CreatedBy:        myFullName,
		Remarks:          input.Remarks,
		QcType:           input.QcType,
		CreatedDateTime:  primitive.DateTime(millis.NowInMillis()),
	}
	//if has assignedQC
	if assignedQCId != nil {
		newInvoice.EmployeeId = assignedQCId
	}

	if assignedQCName != nil {
		newInvoice.Name = assignedQCName
	}

	res, err := DbCollections.Invoices.InsertOne(ctx, newInvoice)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateInvoice(ctx context.Context, id string, input models.UpdateInvoiceInput, updatedBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"updatedBy", updatedBy},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.IsSuperRush != nil {
		setDoc = append(setDoc, bson.E{"isSuperRush", *input.IsSuperRush})
	}
	if input.SuperRushRemarks != nil {
		setDoc = append(setDoc, bson.E{"superRushRemarks", *input.SuperRushRemarks})
	}
	if input.IsRush != nil {
		setDoc = append(setDoc, bson.E{"isRush", *input.IsRush})
	}
	if input.RushRemarks != nil {
		setDoc = append(setDoc, bson.E{"rushRemarks", *input.RushRemarks})
	}
	if input.IsInterior != nil {
		setDoc = append(setDoc, bson.E{"isInterior", *input.IsInterior})
	}
	if input.InteriorRemarks != nil {
		setDoc = append(setDoc, bson.E{"interiorRemarks", *input.InteriorRemarks})
	}
	if input.IsRentalAddendum != nil {
		setDoc = append(setDoc, bson.E{"isRentalAddendum", *input.IsRentalAddendum})
	}
	if input.RentalAddendumRemarks != nil {
		setDoc = append(setDoc, bson.E{"rentalAddendumRemarks", *input.RentalAddendumRemarks})
	}
	if input.IsInitialBpo != nil {
		setDoc = append(setDoc, bson.E{"isInitialBpo", *input.IsInitialBpo})
	}
	if input.InitialBpoRemarks != nil {
		setDoc = append(setDoc, bson.E{"initialBpoRemarks", *input.InitialBpoRemarks})
	}
	if input.IsInspection != nil {
		setDoc = append(setDoc, bson.E{"isInspection", *input.IsInspection})
	}
	if input.InspectionRemarks != nil {
		setDoc = append(setDoc, bson.E{"inspectionRemarks", *input.InspectionRemarks})
	}
	if input.IsNoCsv != nil {
		setDoc = append(setDoc, bson.E{"isNoCsv", *input.IsNoCsv})
	}
	if input.NoCsvRemarks != nil {
		setDoc = append(setDoc, bson.E{"noCsvRemarks", *input.NoCsvRemarks})
	}
	if input.IsNoIFill != nil {
		setDoc = append(setDoc, bson.E{"isNoIFill", *input.IsNoIFill})
	}
	if input.NoIFillRemarks != nil {
		setDoc = append(setDoc, bson.E{"noIFillRemarks", *input.NoIFillRemarks})
	}
	if input.IsOtherPremium != nil {
		setDoc = append(setDoc, bson.E{"isOtherPremium", *input.IsOtherPremium})
	}
	if input.OtherPremiumRemarks != nil {
		setDoc = append(setDoc, bson.E{"otherPremiumRemarks", *input.OtherPremiumRemarks})
	}
	if input.OrderType != nil {
		setDoc = append(setDoc, bson.E{"orderType", *input.OrderType})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Invoices.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func CancelInvoice(ctx context.Context, id string, reason *string, cancelBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"status", "CANCEL"},
		{"cancelBy", cancelBy},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}
	if reason != nil {
		setDoc = append(setDoc, bson.E{"cancel_reason", *reason})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Invoices.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

// func DeleteInvoice(ctx context.Context, id string) (bool, error) {
// 	objId, _ := primitive.ObjectIDFromHex(id)
// 	filter := bson.D{{"_id", objId}}
// 	_, err := DbCollections.Companies.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

func GetInvoices(ctx context.Context, filter *models.InvoiceFilterInput) ([]*models.Invoice, error) {

	var usersID []string = FilterUsers(ctx)
	daoFilter := bson.M{}
	daoFilter["employeeId"] = bson.M{
		"$in": usersID,
	}
	if filter != nil {

		if filter.EmployeeID != nil {
			daoFilter["employeeId"] = *filter.EmployeeID
		}

		if filter.IsCancelled != nil {
			if *filter.IsCancelled {
				daoFilter["status"] = "CANCEL"
			}
		} else {
			daoFilter["status"] = "ACTIVE"
		}

		layout := "2006-01-02"
		daoFilter2 := bson.M{}
		if filter.DateFrom != nil {
			varDateFrom, err := time.Parse(layout, *filter.DateFrom)
			if err != nil {
				log.Debug("failed time parse error : %v", err)
			}
			daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
			daoFilter["createdDateTime"] = daoFilter2
		}

		if filter.DateTo != nil {
			varDateTo, err := time.Parse(layout, *filter.DateTo)
			if err != nil {
				log.Debug("failed time parse error : %v", err)
			}
			daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
			daoFilter["createdDateTime"] = daoFilter2
		}
	}

	cur, err := DbCollections.Invoices.Find(ctx, daoFilter)
	if err != nil {
		log.Error("Failed to query invoice: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Invoice, 0)
	for cur.Next(ctx) {
		a := &Invoice{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode invoice entry: %v", err)
			return nil, err
		}

		list = append(list, a.ToModels())
	}
	fmt.Println(PrettyPrint(list))
	return list, nil

}

func SearchInvoiceRequests(ctx context.Context, offset, limit int, myUserID string, dateFrom, dateTo *string) ([]*Invoice, error) {

	filter := bson.M{
		"employeeId": myUserID,
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Invoices.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query invoice:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawInvoices := make([]*Invoice, 0)
	for cur.Next(ctx) {
		a := &Invoice{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode invoice entry: %v", err)
			return nil, err
		}
		rawInvoices = append(rawInvoices, a)
	}

	return rawInvoices, nil
}

func GetInvoiceRequestsCount(ctx context.Context, myUserID string, dateFrom, dateTo *string) (*int64, error) {

	filter := bson.M{
		"employeeId": myUserID,
	}

	count, err := DbCollections.Invoices.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query invoice request: %v", err)
		return nil, err
	}

	return &count, nil
}

func SearchInvoiceRequestHistories(ctx context.Context, offset, limit int, myUserID string) ([]*Invoice, error) {

	filter := bson.M{
		"employeeId": myUserID,
		"type":       "INVOICE",
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"name": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.RequestsHistory.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query request history:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawInvoices := make([]*Invoice, 0)
	for cur.Next(ctx) {
		a := &RequestHistory{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode request history entry: %v", err)
			return nil, err
		}
		invoice := &Invoice{
			ID:              a.ID,
			PipelineId:      a.PipelineId,
			Status:          a.Status,
			ClientId:        strings.ToObject(a.ClientId),
			EmployeeId:      strings.ToObject(a.EmployeeId),
			Type:            a.Type,
			OrderNumber:     a.OrderNumber,
			Address:         a.Address,
			Company:         a.Company,
			Remarks:         a.Remarks,
			CreatedDateTime: a.CreatedDateTime,
		}
		rawInvoices = append(rawInvoices, invoice)
	}

	return rawInvoices, nil
}

func GetInvoiceRequestHistoriesCount(ctx context.Context, myUserID string) (*int64, error) {

	filter := bson.M{
		"employeeId": myUserID,
		"type":       "INVOICE",
	}
	count, err := DbCollections.RequestsHistory.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query invoice request: %v", err)
		return nil, err
	}

	return &count, nil
}

func FilterInvoiceByType(ctx context.Context, invoiceType string) ([]*Invoice, error) {
	filter := bson.M{"type": invoiceType}
	return GetInvoicesDao(ctx, filter)
}
func GetInvoicesDao(ctx context.Context, filter bson.M) ([]*Invoice, error) {

	cur, err := DbCollections.Invoices.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query invoice: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*Invoice, 0)
	for cur.Next(ctx) {
		a := &Invoice{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode invoice entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}

func UpdateInvoiceForBatch(ctx context.Context, employeeName string, filter *bson.M) (bool, error) {

	setDoc := bson.D{
		{"name", employeeName},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Invoices.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}
