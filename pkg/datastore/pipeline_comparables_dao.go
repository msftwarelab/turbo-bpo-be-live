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

func SavePipelineComparable(ctx context.Context, createdBy, pipelineID string, input models.SavePipelineComparableInput) (string, error) {

	newPipelineComparable := &PipelineComparable{
		PipelineID:      pipelineID,
		Mls:             strings.ToObject(input.Mls),
		Order:           input.Order,
		Status:          strings.ToObject(input.Status),
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}
	updatePipeline := models.UpdatePipelineInput{
		Status: strings.ToObject(constants.PipelineStatusActive),
	}

	res, err := DbCollections.PipelineComparables.InsertOne(ctx, newPipelineComparable)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)

	//Update pipeline to Active if added mls and set activeationTime
	isActivePipeline := pointers.Bool(true)
	UpdatePipeline(ctx, pipelineID, updatePipeline, createdBy, "", isActivePipeline)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdatePipelineComparable(ctx context.Context, id string, mls string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"mls", mls},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.PipelineComparables.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeletePipelineComparable(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"status", "DELETED"},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.PipelineComparables.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchPipelineComparables(ctx context.Context, offset, limit int, pipelineID string) ([]*PipelineComparable, error) {

	filter := bson.M{"pipelineId": pipelineID}
	filter["status"] = bson.M{"$ne": "DELETED"}
	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"name": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipelineComparables.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query PipelineComparable:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelineComparables := make([]*PipelineComparable, 0)
	for cur.Next(ctx) {
		a := &PipelineComparable{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode PipelineComparables entry: %v", err)
			return nil, err
		}
		rawPipelineComparables = append(rawPipelineComparables, a)
	}

	return rawPipelineComparables, nil
}

func GetPipelineComparablesCount(ctx context.Context, pipelineID string) (*int64, error) {

	filter := bson.M{"pipelineId": pipelineID}
	filter["status"] = bson.M{"$ne": "DELETED"}
	count, err := DbCollections.PipelineComparables.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query PipelineComparable: %v", err)
		return nil, err
	}

	return &count, nil
}

