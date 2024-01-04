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

func AddRequest(ctx context.Context, pipelineID string, requestedBy, requestedByID string, pipelineRaw *Pipeline, hasPhotos bool) (string, error) {

	newRequest := &Request{
		OrderNumber:     pipelineRaw.OrderNumber,
		RequestedBy:     &requestedBy,
		RequestedByID:   &requestedByID,
		PipelineId:      &pipelineID,
		OrderType:       pipelineRaw.OrderType,
		Type:            strings.ToObject("INVOICE"),
		Status:          strings.ToObject("PENDING"),
		Company:         pipelineRaw.Company,
		Address:         pipelineRaw.Address,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Requests.InsertOne(ctx, newRequest)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateRequest(ctx context.Context, id string, input models.UpdateRequestInput, modifiedBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := &Request{
		UpdatedBy:      &modifiedBy,
		LastUpdateTime: pointers.PrimitiveDateTime(nil),
	}

	if input.Action != nil {
		update.Status = input.Action
	}

	if input.ConditionType != nil {
		update.ConditionType = input.ConditionType
	}

	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Requests.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchRequests(ctx context.Context, offset, limit int, filter *models.RequestFilterInput) ([]*Request, error) {

	daoFilter := bson.M{}

	if filter.Type != nil {
		daoFilter["type"] = *filter.Type
	}

	if filter.IsPending == nil {
		daoFilter["status"] = "PENDING"
	} else {
		if *filter.IsPending {
			daoFilter["status"] = "PENDING"
		}
	}

	if filter.IsPhoto != nil {
		if *filter.IsPhoto {
			daoFilter["hasPhotos"] = true
		}
	}

	if filter.RequestedByID != nil {
		daoFilter["requestedById"] = *filter.RequestedByID
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

	pipe := []bson.M{
		{"$match": daoFilter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Requests.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query request:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawRequests := make([]*Request, 0)
	for cur.Next(ctx) {
		a := &Request{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode request entry: %v", err)
			return nil, err
		}
		rawRequests = append(rawRequests, a)
	}

	return rawRequests, nil
}

func GetRequestsCount(ctx context.Context, filter *models.RequestFilterInput) (*int64, error) {

	daoFilter := bson.M{}

	if filter.Type != nil {
		daoFilter["type"] = *filter.Type
	}

	if filter.IsPending == nil {
		daoFilter["status"] = "PENDING"
	} else {
		if *filter.IsPending {
			daoFilter["status"] = "PENDING"
		}
	}

	if filter.IsPhoto != nil {
		if *filter.IsPhoto {
			daoFilter["hasPhotos"] = true
		}
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
	count, err := DbCollections.Requests.CountDocuments(ctx, daoFilter)
	if err != nil {
		log.Debug("Failed to query request: %v", err)
		return nil, err
	}

	return &count, nil
}

func AddRequestHistory(ctx context.Context, newRequestHistory RequestHistory) (string, error) {

	res, err := DbCollections.RequestsHistory.InsertOne(ctx, newRequestHistory)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}
