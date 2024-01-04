package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"

	"go.mongodb.org/mongo-driver/bson"
)

func AddPipelineRepair(ctx context.Context, pipelineId string, input models.SavePipelineNoteInput) (string, error) {

	newPipelineRepair := &PipelineRepair{
		PipelineId:                  strings.ToObject(pipelineId),
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
	}

	res, err := DbCollections.PipeLineRepairs.InsertOne(ctx, newPipelineRepair)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdatePipelineRepair(ctx context.Context, pipelineId string, input models.UpdatePipelineRepairInput, updatedBy string) (bool, error) {

	filter := bson.D{{"pipelineId", pipelineId}}

	update := PipelineRepair{
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
	}

	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.PipeLineRepairs.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetPipelineRepair(ctx context.Context, pipelineId string) (*models.PipelineRepair, error) {
	filter := bson.M{"pipelineId": pipelineId}
	var pipelineRepair PipelineRepair
	err := DbCollections.PipeLineRepairs.FindOne(ctx, filter).Decode(&pipelineRepair)
	if err != nil {
		return nil, err
	}
	return pipelineRepair.ToModels(), nil
}
