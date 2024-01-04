package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveBilling(ctx context.Context, input models.SaveBillingInput, createdBy string) (string, error) {

	dueDate, err := pointers.StringTimeToPrimitive(input.DueDate)
	if err != nil {
		return "", errs.InvalidDueDate
	}

	varDate, err := pointers.StringTimeToPrimitive(input.Date)
	if err != nil {
		return "", errs.InvalidDate
	}

	filter := FilterById(input.UserID)
	client, err := GetUser(ctx, filter)
	if err != nil || client == nil {
		return "", errs.InvalidClientId
	}

	newBilling := &Billing{
		InvoiceNumber:   strings.ToObject(input.InvoiceNumber),
		Status:          strings.ToObject(constants.BillingStatusPending),
		Date:            varDate,
		DueDate:         dueDate,
		UserID:          strings.ToObject(input.UserID),
		UserName:        strings.ToObject(client.FullName()),
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	//set dateFrom
	if input.DateFrom != nil {
		dateFrom, err := pointers.StringTimeToPrimitive(*input.DateFrom)
		if err != nil {
			return "", errs.InvalidDateFrom
		}
		newBilling.DateFrom = dateFrom
	}
	//set dateTo
	if input.DateTo != nil {
		dateTo, err := pointers.StringTimeToPrimitive(*input.DateTo)
		if err != nil {
			return "", errs.InvalidDateTo
		}
		newBilling.DateTo = dateTo
	}

	entries := []*BillingEntry{}

	if len(input.Entries) > 0 {

		for _, v := range input.Entries {
			entry := &BillingEntry{
				OrderNumber: v.OrderNumber,
				Description: strings.ToObject(v.Description),
				Amount:      &v.Amount,
				Type:        strings.ToObject(v.Type),
			}
			entries = append(entries, entry)
		}

	}
	newBilling.Entries = entries

	res, err := DbCollections.Billings.InsertOne(ctx, newBilling)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateBilling(ctx context.Context, id string, input models.UpdateBillingInput) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	updateBilling := &Billing{}

	if input.UserID != nil {
		filter := FilterById(*input.UserID)
		client, err := GetUser(ctx, filter)
		if err != nil || client == nil {
			return false, errs.InvalidClientId
		}
		updateBilling.UserID = input.UserID
		updateBilling.UserName = strings.ToObject(client.FullName())
	}

	if input.InvoiceNumber != nil {
		updateBilling.InvoiceNumber = input.InvoiceNumber
	}

	if input.Date != nil {
		varDate, err := pointers.StringTimeToPrimitive(*input.Date)
		if err != nil {
			return false, errs.InvalidDate
		}
		updateBilling.Date = varDate

	}

	if input.DateFrom != nil {
		dateFrom, err := pointers.StringTimeToPrimitive(*input.DateFrom)
		if err != nil {
			return false, errs.InvalidDateFrom
		}
		updateBilling.DateFrom = dateFrom

	}

	if input.DateTo != nil {
		dateTo, err := pointers.StringTimeToPrimitive(*input.DateTo)
		if err != nil {
			return false, errs.InvalidDateTo
		}
		updateBilling.DateTo = dateTo

	}

	if input.DueDate != nil {
		dueDate, err := pointers.StringTimeToPrimitive(*input.DueDate)
		if err != nil {
			return false, errs.InvalidDueDate
		}
		updateBilling.DueDate = dueDate
	}

	if input.Status != nil {
		updateBilling.Status = input.Status
	}

	entries := []*BillingEntry{}
	if input.Entries != nil {
		for _, v := range input.Entries {
			entry := &BillingEntry{
				OrderNumber: v.OrderNumber,
				Description: &v.Description,
				Amount:      &v.Amount,
				Type:        &v.Type,
			}
			entries = append(entries, entry)
		}
	}
	updateBilling.Entries = entries
	updateDoc := bson.M{
		"$set": updateBilling,
	}
	res, err := DbCollections.Billings.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteBilling(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"status", "DELETED"},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Billings.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func UpdateBillingStatus(ctx context.Context, id string, newStatus string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"status", newStatus},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Billings.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchBillings(ctx context.Context, offset, limit int, dateFrom, dateTo, clientID *string, isHasDueDate *bool, orderNumber *string) ([]*Billing, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	daoFilter2 := bson.M{}

	if dateFrom != nil {
		premDateFrom, err := pointers.StringTimeToPrimitive(*dateFrom)
		if err != nil {
			return nil, errs.InvalidDateFrom
		}
		daoFilter2["$gte"] = premDateFrom
		filter["date"] = daoFilter2
	}

	if dateTo != nil {
		premDateTo, err := pointers.StringTimeToPrimitive(*dateTo)
		if err != nil {
			return nil, errs.InvalidDateTo
		}
		daoFilter2["$lte"] = premDateTo
		filter["date"] = daoFilter2
	}

	if clientID != nil {
		filter["userId"] = *clientID
	}

	if pointers.ObjectTOBool(isHasDueDate) {
		daoFilter2["$lt"] = pointers.PrimitiveDateTime(nil)
		filter["dueDate"] = daoFilter2
		filter["status"] = constants.BillingStatusPending

	}

	if orderNumber != nil {
		filter["entries.orderNumber"] = *orderNumber
	}
	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"date": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Billings.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query Billing:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawBillings := make([]*Billing, 0)
	for cur.Next(ctx) {
		a := &Billing{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode Billings entry: %v", err)
			return nil, err
		}
		rawBillings = append(rawBillings, a)
	}

	return rawBillings, nil
}

func GetBillingsCount(ctx context.Context, dateFrom, dateTo, clientID, orderNumber *string) (*int64, error) {
	filter := bson.M{}
	daoFilter2 := bson.M{}

	if dateFrom != nil {
		premDateFrom, err := pointers.StringTimeToPrimitive(*dateFrom)
		if err != nil {
			return nil, errs.InvalidDateFrom
		}
		daoFilter2["$gte"] = premDateFrom
		filter["date"] = daoFilter2
	}

	if dateTo != nil {
		premDateTo, err := pointers.StringTimeToPrimitive(*dateTo)
		if err != nil {
			return nil, errs.InvalidDateTo
		}
		daoFilter2["$lte"] = premDateTo
		filter["date"] = daoFilter2
	}

	if clientID != nil {
		filter["userId"] = *clientID
	}

	if orderNumber != nil {
		filter["entries.orderNumber"] = *orderNumber
	}

	count, err := DbCollections.Billings.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query Billing: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetBilling(ctx context.Context, filter *bson.M) (*Billing, error) {
	var resBilling Billing
	daoFilter := bson.M{}
	if filter != nil {
		daoFilter = *filter
	}
	err := DbCollections.Billings.FindOne(ctx, daoFilter).Decode(&resBilling)
	if err != nil {
		return nil, err
	}
	return &resBilling, nil
}
