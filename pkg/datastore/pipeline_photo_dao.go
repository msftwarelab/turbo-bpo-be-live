package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPipelinePhoto(ctx context.Context, pipelineId string, input models.PipelinePhotoInput, userId string, fileUrl string, alwayssubmitOrder bool, createdBy string) (string, error) {
	newPipelinePhoto := &PipelinePhoto{
		PipelineId:      &pipelineId,
		Filename:        input.Doc.Filename,
		Url:             fileUrl,
		IsSubmitted:     alwayssubmitOrder,
		FileSize:        pointers.Int64(input.Doc.Size),
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
		CreatedBy:       createdBy,
	}

	res, err := DbCollections.PipeLinePhotos.InsertOne(ctx, newPipelinePhoto)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func SubmitPipelinePhoto(ctx context.Context, id string, userId string, isSubmitPipelinePhoto bool) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"isSubmitted":    isSubmitPipelinePhoto,
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.PipeLinePhotos.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeletePipelinePhoto(ctx context.Context, id string, userId string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"status":         "DELETED",
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.PipeLinePhotos.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetPipelinePhotos(ctx context.Context, pipelineIds []string) ([]*PipelinePhoto, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineIds}

	cur, err := DbCollections.PipeLinePhotos.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query pipelinePhotos: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*PipelinePhoto, 0)
	for cur.Next(ctx) {
		a := &PipelinePhoto{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelinePhotos entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}

func SearchPipelinePhotos(ctx context.Context, pipelineId *string, offset, limit int, pipelinePhotoID *string) ([]*PipelinePhoto, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}

	if pipelineId != nil {
		filter["pipelineId"] = *pipelineId
	}

	if pipelinePhotoID != nil {
		objId, _ := primitive.ObjectIDFromHex(*pipelinePhotoID)
		filter["_id"] = objId
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipeLinePhotos.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipelinePhotos:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelinePhotos := make([]*PipelinePhoto, 0)
	for cur.Next(ctx) {
		a := &PipelinePhoto{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelinePhotos entry: %v", err)
			return nil, err
		}
		rawPipelinePhotos = append(rawPipelinePhotos, a)
	}

	return rawPipelinePhotos, nil
}

func GetPipelinePhotosCount(ctx context.Context, pipelineId string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	count, err := DbCollections.PipeLinePhotos.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query pipelinePhotos: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetPipelinePhotosCountGroupByPipelineId(ctx context.Context, pipelineIds []string) ([]*GetCountGroupbyId, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineIds}

	group := bson.M{
		"_id":   "$pipelineId",
		"count": bson.M{"$sum": 1},
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$group": group},
	}

	cur, err := DbCollections.PipeLinePhotos.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query pepilinePhotos count: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)
	rawGetCountGroupbyId := make([]*GetCountGroupbyId, 0)
	for cur.Next(ctx) {
		cos := &GetCountGroupbyId{}
		if err := cur.Decode(cos); err != nil {
			log.Debug("Failed to decode pepilinePhotos total entry: %v", err)
			return nil, err
		}
		rawGetCountGroupbyId = append(rawGetCountGroupbyId, cos)
	}

	return rawGetCountGroupbyId, nil
}
