package datastore

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const pipelineStateId string = "5dce086c7f342e4680e9691d"

func UpdatePiplineState(ctx context.Context, input models.UpdatePipelineStateInput) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(pipelineStateId)
	filter := &bson.M{"_id": objId}

	updatePipelineState := &PipelineState{
		MaxDailyVolume:      input.MaxDailyVolume,
		StandByAutoComplete: input.StandByAutoComplete,
		IsRush:              input.IsRush,
		IsNewOrder:          input.IsNewOrder,
		OrderMessage:        input.OrderMessage,
		TTSlow:              input.TTSlow,
		TTModerate:          input.TTModerate,
		TTBusy:              input.TTBusy,
		TTMax:               input.TTMax,
		TLSlow:              input.TLSlow,
		TLModerate:          input.TLModerate,
		TLBusy:              input.TLBusy,
		OPInterior:          input.OPInterior,
		OPExterior:          input.OPExterior,
		OPDataEntry:         input.OPDataEntry,
		OPRush:              input.OPRush,
		OPSuperRush:         input.OPSuperRush,
		OPConditionReport:   input.OPConditionReport,
		OPRentalAddendum:    input.OPRentalAddendum,
		OPInitialBpo:        input.OPInitialBpo,
		OPInspection:        input.OPInspection,
		PCIsAcceptOrder:     input.PCIsAcceptOrder,
		PCcatchTime:         input.PCcatchTime,
		OAOfferLimitInMin:   input.OAOfferLimitInMin,
		OAIsAutoAssign:      input.OAIsAutoAssign,
		QCElapseTime:        input.QCElapseTime,
		LastUpdateTime:      pointers.PrimitiveDateTime(nil),
	}

	updateDoc := bson.M{
		"$set": updatePipelineState,
	}
	res, err := DbCollections.PipelineStates.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetPipelineState(ctx context.Context) (*models.PipelineState, error) {
	objId, _ := primitive.ObjectIDFromHex(pipelineStateId)
	filter := &bson.M{"_id": objId}

	cur, err := DbCollections.PipelineStates.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query pipelineState: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		a := &PipelineState{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineState entry: %v", err)
			return nil, err
		}
		todayOrderCount, _ := GetTodayOrderCount(ctx)
		pipelinetate := a.ToModels()
		pipelinetate.TodayOrderCount = todayOrderCount
		return pipelinetate, nil
	}
	return nil, errs.DbError
}

func GetPipelineStateDataStore(ctx context.Context) (*PipelineState, error) {
	objId, _ := primitive.ObjectIDFromHex(pipelineStateId)
	filter := &bson.M{"_id": objId}

	cur, err := DbCollections.PipelineStates.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query pipelineState: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		a := &PipelineState{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineState entry: %v", err)
			return nil, err
		}
		return a, nil
	}
	return nil, errs.DbError
}
