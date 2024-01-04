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

func SaveSession(ctx context.Context, userId string, invoiceDate string, createdBy string) (string, error) {

	layout := "2006-01-02"
	invoiceDateTime, err := time.Parse(layout, invoiceDate)
	if err != nil {
		log.Debug("failed time parse error : %v", err)
		return "", err
	}

	newSession := &Session{
		UserID:          userId,
		InvoiceDate:     *pointers.PrimitiveDateTime(&invoiceDateTime),
		Isrunning:       true,
		Start:           primitive.DateTime(millis.NowInMillis()),
		CreatedBy:       createdBy,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	//TODO, valite if there is running session
	log.Debug("@debug session userID %s", userId)
	sessionACtiveCount, err := GetSessionsCount(ctx, userId, nil, pointers.Bool(true), nil, nil)
	if err != nil {
		return "", err
	}
	if *sessionACtiveCount > 0 {
		return "", errs.OpenSessionExist
	}

	res, err := DbCollections.Sessions.InsertOne(ctx, newSession)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func StopSession(ctx context.Context, userId string, modifiedBy string) (bool, error) {

	filter := bson.D{
		{"userId", userId},
		{"isRunning", true},
	}

	setDoc := bson.D{
		{"isRunning", false},
		{"end", pointers.PrimitiveDateTime(nil)},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Sessions.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

// setDoc := bson.D{
// 	{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
// }

// if input.Subject != nil {
// 	setDoc = append(setDoc, bson.E{"subject", *input.Subject})
// }

func UpdateSession(ctx context.Context, id string, input models.UpdateSessionInput, myName string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.InvoiceDate != nil {
		invoiceDate, err := pointers.StringTimeToPrimitive(*input.InvoiceDate)
		if err != nil {
			return false, errs.InvalidInvoiceDate
		}
		setDoc = append(setDoc, bson.E{"invoiceDate", invoiceDate})
	}

	if input.Start != nil {
		startDate, err := pointers.StringTDateimeToPrimitive(*input.Start)
		if err != nil {
			return false, errs.InvalidStartDate
		}
		setDoc = append(setDoc, bson.E{"start", startDate})
	}

	if input.End != nil {
		endDate, err := pointers.StringTDateimeToPrimitive(*input.End)
		if err != nil {
			return false, errs.InvalidEndDate
		}
		setDoc = append(setDoc, bson.E{"end", endDate})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Sessions.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchSessions(ctx context.Context, userId string, offset, limit int, selectedUserId *string, isRunning *bool, dateFrom, dateTo *string) ([]*Session, error) {

	filter := bson.M{
		"userId": userId,
	}
	daoFilter2 := bson.M{}

	if isRunning != nil {
		filter["isRunning"] = *isRunning
	}

	if selectedUserId != nil {
		filter["userId"] = *selectedUserId
	}

	if dateFrom != nil {
		premDateFrom, err := pointers.StringTimeToPrimitive(*dateFrom)
		if err != nil {
			return nil, errs.InvalidDateFrom
		}
		daoFilter2["$gte"] = premDateFrom
		filter["invoiceDate"] = daoFilter2
	}

	if dateTo != nil {
		premDateTo, err := pointers.StringTimeToPrimitive(*dateTo)
		if err != nil {
			return nil, errs.InvalidDateTo
		}
		daoFilter2["$lte"] = premDateTo
		filter["invoiceDate"] = daoFilter2
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Sessions.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query session:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawSessions := make([]*Session, 0)
	for cur.Next(ctx) {
		a := &Session{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode session entry: %v", err)
			return nil, err
		}
		rawSessions = append(rawSessions, a)
	}

	return rawSessions, nil
}

func GetSessionsCount(ctx context.Context, userId string, selectedUserId *string, isRunning *bool, dateFrom, dateTo *string) (*int64, error) {

	filter := bson.M{
		"userId": userId,
	}
	daoFilter2 := bson.M{}

	if isRunning != nil {
		filter["isRunning"] = *isRunning
	}

	if selectedUserId != nil {
		filter["userId"] = *selectedUserId
	}

	if dateFrom != nil {
		premDateFrom, err := pointers.StringTimeToPrimitive(*dateFrom)
		if err != nil {
			return nil, errs.InvalidDateFrom
		}
		daoFilter2["$gte"] = premDateFrom
		filter["invoiceDate"] = daoFilter2
	}

	if dateTo != nil {
		premDateTo, err := pointers.StringTimeToPrimitive(*dateTo)
		if err != nil {
			return nil, errs.InvalidDateTo
		}
		daoFilter2["$lte"] = premDateTo
		filter["invoiceDate"] = daoFilter2
	}

	count, err := DbCollections.Sessions.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query session: %v", err)
		return nil, err
	}

	return &count, nil
}

func ContinueSession(ctx context.Context, userID string) (bool, error) {
	filter := bson.M{
		"userId": userID,
	}
	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$limit": 1},
	}

	cur, err := DbCollections.Sessions.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query session:  %v", err)
		return false, err
	}
	defer cur.Close(ctx)
	rawSessions := make([]*Session, 0)
	for cur.Next(ctx) {
		a := &Session{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode session entry: %v", err)
			return false, err
		}
		rawSessions = append(rawSessions, a)
	}
	//unset end element
	if len(rawSessions) == 0 {
		return false, nil
	}
	var session_id = rawSessions[0].ID

	filter = bson.M{
		"_id": session_id,
	}

	updateDoc := bson.M{
		"$unset": bson.M{
			"end": "-1",
		},
		"$set": bson.M{
			"isRunning": true,
		},
	}
	res, err := DbCollections.Sessions.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}
