package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPipelineQualityControl(ctx context.Context, pipelineId string, orderNotes string, createdBy string) (string, error) {
	newPipelineQualityControl := &PipelineQualityControl{
		PipelineId:      pipelineId,
		Message:         orderNotes,
		CreatedBy:       createdBy,
		Type:            "QUALITY_CONTROL",
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.PipeLineQualityControlAndNotes.InsertOne(ctx, newPipelineQualityControl)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func GetPipelineQualityControls(ctx context.Context, pipelineId []string) ([]*PipelineQualityControl, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineId}
	filter["type"] = "QUALITY_CONTROL"

	cur, err := DbCollections.PipeLineQualityControlAndNotes.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query PipelineQualityControlAndNotes: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*PipelineQualityControl, 0)
	for cur.Next(ctx) {
		a := &PipelineQualityControl{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode PipelineQualityControlAndNotes entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}

func SearchPipelineQualityControl(ctx context.Context, pipelineId string, offset, limit int) ([]*PipelineQualityControl, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	filter["type"] = "QUALITY_CONTROL"

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipeLineQualityControlAndNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query PipelineQualityControlAndNotes:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelineQualityControl := make([]*PipelineQualityControl, 0)
	for cur.Next(ctx) {
		a := &PipelineQualityControl{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineQualityCotrol entry: %v", err)
			return nil, err
		}
		rawPipelineQualityControl = append(rawPipelineQualityControl, a)
	}

	return rawPipelineQualityControl, nil
}

func GetPipelineQualityControlCount(ctx context.Context, pipelineId string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	filter["type"] = "QUALITY_CONTROL"
	count, err := DbCollections.PipeLineQualityControlAndNotes.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query PipelineQualityControlAndNotes: %v", err)
		return nil, err
	}

	return &count, nil
}

func SearchPipelineQualityControlAndNotes(ctx context.Context, pipelineId string, offset, limit int) ([]*PipelineQualityControl, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipeLineQualityControlAndNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query PipelineQualityControlAndNotes:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelineQualityControl := make([]*PipelineQualityControl, 0)
	for cur.Next(ctx) {
		a := &PipelineQualityControl{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode PipelineQualityControlAndNotes entry: %v", err)
			return nil, err
		}
		rawPipelineQualityControl = append(rawPipelineQualityControl, a)
	}

	return rawPipelineQualityControl, nil
}

func GetPipelineQualityControlAndNotesCount(ctx context.Context, pipelineId string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	count, err := DbCollections.PipeLineQualityControlAndNotes.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query pipelineQualityControlAndNotes: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetPipelineQualityControlsCountGroupByPipelineId(ctx context.Context, pipelineIds []string) ([]*GetCountGroupbyId, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["type"] = "QUALITY_CONTROL"
	filter["pipelineId"] = bson.M{"$in": pipelineIds}

	group := bson.M{
		"_id":   "$pipelineId",
		"count": bson.M{"$sum": 1},
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$group": group},
	}

	cur, err := DbCollections.PipeLineQualityControlAndNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query pepilineQualityControls count: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)
	rawGetCountGroupbyId := make([]*GetCountGroupbyId, 0)
	for cur.Next(ctx) {
		cos := &GetCountGroupbyId{}
		if err := cur.Decode(cos); err != nil {
			log.Debug("Failed to decode pepilineQualityControlss total entry: %v", err)
			return nil, err
		}
		rawGetCountGroupbyId = append(rawGetCountGroupbyId, cos)
	}

	return rawGetCountGroupbyId, nil
}

func SavePipelineQualityControlAndNote(ctx context.Context, pipelineId string, input models.SavePipelineQualityControlAndNoteInput, myName string) (string, error) {

	newPipelineQualityControl := &PipelineQualityControl{
		PipelineId:      pipelineId,
		Message:         strings.ObjectTOString(input.Message),
		CreatedBy:       myName,
		RequestType:     strings.ObjectTOString(input.RequestType),
		Type:            strings.ObjectTOString(input.Category),
		Status:          input.Status,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.PipeLineQualityControlAndNotes.InsertOne(ctx, newPipelineQualityControl)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func CheckQualityControlStatus(ctx context.Context, pipelineId string) (bool, error) {
	isQc := false
	pipeline := bson.M{"pipelineId": pipelineId}
	filter := []bson.M{pipeline, FilterQA()}

	pipe := []bson.M{
		{"$match": bson.M{"$and": filter}},
	}

	cur, err := DbCollections.QualityControls.Aggregate(ctx, pipe)
	if err != nil {
		// log.Debug("Failed to query quality controls:  %v", err)
		return false, err
	}
	defer cur.Close(ctx)

	rawPipeline := make([]*Pipeline, 0)
	for cur.Next(ctx) {
		a := &Pipeline{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode PipelineQualityControl status entry: %v", err)
			return false, err
		}
		rawPipeline = append(rawPipeline, a)
	}

	if len(rawPipeline) > 0 {
		isQc = true
	}

	return isQc, err
}
