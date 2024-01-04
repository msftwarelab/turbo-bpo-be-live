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

func UpdatePiplineNeighborhood(ctx context.Context, pipelineId string, input models.UpdatePipelineNeighborhoodInput) (bool, error) {
	filter := &bson.M{"pipelineId": pipelineId}

	updatePipelineState := &PipelineNeighborhood{
		MarketTrend:          input.MarketTrend,
		MonthlyPercentage:    input.MonthlyPercentage,
		SixmonthPercentage:   input.SixmonthPercentage,
		AnnualPercentage:     input.AnnualPercentage,
		TotalListings:        input.TotalListings,
		Supply:               input.Supply,
		ListingsMinValue:     input.ListingsMinValue,
		ListingsMedValue:     input.ListingsMedValue,
		ListingsMaxValue:     input.ListingsMaxValue,
		ListingsDomAve:       input.ListingsDomAve,
		ListingsDomRangeFrom: input.ListingsDomRangeFrom,
		ListingsDomRangeTo:   input.ListingsDomRangeTo,
		Fm:                   input.Fm,
		Ss:                   input.Ss,
		Reo:                  input.Reo,
		Distressed:           input.Distressed,
		TotalSales:           input.TotalSales,
		Demand:               input.Demand,
		SalesMinValue:        input.SalesMinValue,
		SalesMedValue:        input.SalesMedValue,
		SalesMaxValue:        input.SalesMaxValue,
		SalesDomRangeFrom:    input.SalesDomRangeFrom,
		SalesDomRangeTo:      input.SalesDomRangeTo,
		ZntComments:          input.ZntComments,
		NtComments:           input.NtComments,
	}
	if input.IsReoDriven != nil {
		updatePipelineState.IsReoDriven = *input.IsReoDriven
	}
	updateDoc := bson.M{
		"$set": updatePipelineState,
	}
	res, err := DbCollections.PipeLineNeighborhood.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetPipelineNeighborhood(ctx context.Context, pipelineId string) (*models.PipelineNeighborhood, error) {
	filter := &bson.M{"pipelineId": pipelineId}

	cur, err := DbCollections.PipeLineNeighborhood.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query pipelineNeighborhood: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		a := &PipelineNeighborhood{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipelineNeighborhood entry: %v", err)
			return nil, err
		}
		return a.ToModels(), nil
	}
	return nil, errs.DbError
}

func SavePipelineNeighborhood(ctx context.Context, pipelineId string) (string, error) {

	newPipelineNeighborhood := &PipelineNeighborhood{
		PipelineId:      pipelineId,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.PipeLineNeighborhood.InsertOne(ctx, newPipelineNeighborhood)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}
