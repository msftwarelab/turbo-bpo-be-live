package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPipelineDoc(ctx context.Context, pipelineId string, input models.PipelineDocInput, fileUrl string, createdBy string) (string, error) {
	newPipelineDoc := &PipelineDoc{
		PipelineId:      &pipelineId,
		Type:            input.Type,
		Filename:        input.Doc.Filename,
		Url:             fileUrl,
		CreatedBy:       createdBy,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.PipeLineDocs.InsertOne(ctx, newPipelineDoc)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func AutoAddPipelineDoc(ctx context.Context, pipelineId string, instructions []*Instruction, createdBy string) (bool, error) {
	docs := make([]interface{}, 0)

	for _, v := range instructions {
		doc := bson.D{
			{"pipelineId", pipelineId},
			{"fileName", v.FileName},
			{"url", v.Url},
			{"createdBy", createdBy},
			{"createdDateTime", primitive.DateTime(millis.NowInMillis())},
		}
		docs = append(docs, doc)
	}
	_, err := DbCollections.PipeLineDocs.InsertMany(ctx, docs)
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeletePipelineDoc(ctx context.Context, id string, userId string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"status":         "DELETED",
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.PipeLineDocs.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetPipelineDocs(ctx context.Context, pipelineIds []string) ([]*PipelineDoc, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineIds}

	cur, err := DbCollections.PipeLineDocs.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query pipelineDocs: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*PipelineDoc, 0)
	for cur.Next(ctx) {
		a := &PipelineDoc{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineDocs entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func SearchPipelineDocs(ctx context.Context, pipelineId string, offset, limit int) ([]*PipelineDoc, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipeLineDocs.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipelineDocs:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelineDoc := make([]*PipelineDoc, 0)
	for cur.Next(ctx) {
		a := &PipelineDoc{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineDocs entry: %v", err)
			return nil, err
		}
		rawPipelineDoc = append(rawPipelineDoc, a)
	}

	return rawPipelineDoc, nil
}

func GetPipelineDocsCount(ctx context.Context, pipelineId string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId

	count, err := DbCollections.PipeLineDocs.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query pipelineDocs: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetPipelineDocsCountGroupByPipelineId(ctx context.Context, pipelineIds []string) ([]*GetCountGroupbyId, error) {
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

	cur, err := DbCollections.PipeLineDocs.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query pepilineDocs count: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)
	rawGetCountGroupbyId := make([]*GetCountGroupbyId, 0)
	for cur.Next(ctx) {
		cos := &GetCountGroupbyId{}
		if err := cur.Decode(cos); err != nil {
			log.Debug("Failed to decode pepilineDocs total entry: %v", err)
			return nil, err
		}
		rawGetCountGroupbyId = append(rawGetCountGroupbyId, cos)
	}

	return rawGetCountGroupbyId, nil
}
