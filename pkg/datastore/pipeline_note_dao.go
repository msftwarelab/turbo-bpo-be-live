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

func AddPipelineNote(ctx context.Context, pipelineId string, input models.SavePipelineNoteInput, createdBy string) (string, error) {
	newPipelineNote := &PipelineNote{
		PipelineId:                  pipelineId,
		Type:                        "NOTE",
		OrderNotes:                  input.OrderNotes,
		InspectionNotes:             input.InspectionNotes,
		ExteriorRepairDescription1:  input.ExteriorRepairDescription1,
		ExteriorRepairPrice1:        input.ExteriorRepairPrice1,
		ExteriorRepairDescription2:  input.ExteriorRepairDescription2,
		ExteriorRepairPrice2:        input.ExteriorRepairPrice2,
		ExteriorRepairDescription3:  input.ExteriorRepairDescription3,
		ExteriorRepairPrice3:        input.ExteriorRepairPrice3,
		ExteriorRepairDescription4:  input.ExteriorRepairDescription4,
		ExteriorRepairPrice4:        input.ExteriorRepairPrice4,
		ExteriorRepairDescription5:  input.ExteriorRepairDescription5,
		ExteriorRepairPrice5:        input.ExteriorRepairPrice5,
		ExteriorRepairDescription6:  input.ExteriorRepairDescription6,
		ExteriorRepairPrice6:        input.ExteriorRepairPrice6,
		ExteriorRepairDescription7:  input.ExteriorRepairDescription7,
		ExteriorRepairPrice7:        input.ExteriorRepairPrice7,
		ExteriorRepairDescription8:  input.ExteriorRepairDescription8,
		ExteriorRepairPrice8:        input.ExteriorRepairPrice8,
		ExteriorRepairDescription9:  input.ExteriorRepairDescription9,
		ExteriorRepairPrice9:        input.ExteriorRepairPrice9,
		ExteriorRepairDescription10: input.ExteriorRepairDescription10,
		ExteriorRepairPrice10:       input.ExteriorRepairPrice10,
		ExteriorRepairPriceTotal:    input.ExteriorRepairPriceTotal,
		InteriorRepairDescription1:  input.InteriorRepairDescription1,
		InteriorRepairPrice1:        input.InteriorRepairPrice1,
		InteriorRepairDescription2:  input.InteriorRepairDescription2,
		InteriorRepairPrice2:        input.InteriorRepairPrice2,
		InteriorRepairDescription3:  input.InteriorRepairDescription3,
		InteriorRepairPrice3:        input.InteriorRepairPrice3,
		InteriorRepairDescription4:  input.InteriorRepairDescription4,
		InteriorRepairPrice4:        input.InteriorRepairPrice4,
		InteriorRepairDescription5:  input.InteriorRepairDescription5,
		InteriorRepairPrice5:        input.InteriorRepairPrice5,
		InteriorRepairDescription6:  input.InteriorRepairDescription6,
		InteriorRepairPrice6:        input.InteriorRepairPrice6,
		InteriorRepairDescription7:  input.InteriorRepairDescription7,
		InteriorRepairPrice7:        input.InteriorRepairPrice7,
		InteriorRepairDescription8:  input.InteriorRepairDescription8,
		InteriorRepairPrice8:        input.InteriorRepairPrice8,
		InteriorRepairDescription9:  input.InteriorRepairDescription9,
		InteriorRepairPrice9:        input.InteriorRepairPrice9,
		InteriorRepairDescription10: input.InteriorRepairDescription10,
		InteriorRepairPrice10:       input.InteriorRepairPrice10,
		InteriorRepairPriceTotal:    input.InteriorRepairPriceTotal,

		CreatedBy:       createdBy,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}
	if input.OrderNotes != nil {
		newPipelineNote.Message = *input.OrderNotes
	}

	res, err := DbCollections.PipeLineNotes.InsertOne(ctx, newPipelineNote)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func GetPipelineNotes(ctx context.Context, pipelineIds []string) ([]*PipelineNote, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineIds}
	filter["type"] = "NOTE"

	cur, err := DbCollections.PipeLineNotes.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query PipelineNotes: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*PipelineNote, 0)
	for cur.Next(ctx) {
		a := &PipelineNote{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode PipelineNotes entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func SearchPipelineNotes(ctx context.Context, pipelineId string, offset, limit int) ([]*PipelineNote, error) {

	//skip := (offset - 1) * limit
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	filter["type"] = "NOTE"

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PipeLineNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipelineNotes:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelineNotes := make([]*PipelineNote, 0)
	for cur.Next(ctx) {
		a := &PipelineNote{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineNotes entry: %v", err)
			return nil, err
		}
		if a.OrderNotes != nil || a.InspectionNotes != nil {
			rawPipelineNotes = append(rawPipelineNotes, a)
		}

	}

	return rawPipelineNotes, nil
}

func GetPipelineNotesCount(ctx context.Context, pipelineId string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = pipelineId
	filter["type"] = "NOTE"
	count, err := DbCollections.PipeLineNotes.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query pipelineNotes: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetPipelineNotesCountGroupByPipelineId(ctx context.Context, pipelineIds []string) ([]*GetCountGroupbyId, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["pipelineId"] = bson.M{"$in": pipelineIds}
	filter["type"] = "NOTE"

	group := bson.M{
		"_id":   "$pipelineId",
		"count": bson.M{"$sum": 1},
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$group": group},
	}

	cur, err := DbCollections.PipeLineNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query pepilineNotes count: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)
	rawGetCountGroupbyId := make([]*GetCountGroupbyId, 0)
	for cur.Next(ctx) {
		cos := &GetCountGroupbyId{}
		if err := cur.Decode(cos); err != nil {
			log.Debug("Failed to decode pepilineNotes total entry: %v", err)
			return nil, err
		}
		rawGetCountGroupbyId = append(rawGetCountGroupbyId, cos)
	}

	return rawGetCountGroupbyId, nil
}

func GetAllEmptyPipelineNotes(ctx context.Context) ([]*PipelineNote, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["type"] = "NOTE"
	filter["orderNotes"] = ""

	pipe := []bson.M{
		{"$match": filter},
	}
	cur, err := DbCollections.PipeLineNotes.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query pepilineNotes count: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)
	rawPipelineNotes := make([]*PipelineNote, 0)
	for cur.Next(ctx) {
		a := &PipelineNote{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineNotes entry: %v", err)
			return nil, err
		}
		if a.OrderNotes != nil || a.InspectionNotes != nil {

			rawPipelineNotes = append(rawPipelineNotes, a)
		}

	}

	return rawPipelineNotes, nil

}

func DeletePipelineNote(ctx context.Context, id primitive.ObjectID) (bool, error) {
	filter := bson.D{{"_id", id}}

	updatePipeline := bson.M{
		"status": "DELETED",
	}
	updateDoc := bson.M{
		"$set": updatePipeline,
	}
	res, err := DbCollections.PipeLineNotes.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}
