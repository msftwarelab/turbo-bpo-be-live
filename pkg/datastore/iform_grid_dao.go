package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveIformGrid(ctx context.Context, pipelineID string, input models.SaveIformGridInput, createdBy string) (string, error) {

	newIformGrid := &IformGrid{
		PipelineID:           pipelineID,
		Address:              input.Address,
		Age:                  input.Age,
		BasementFinishedSqFt: input.BasementFinishedSqFt,
		BasementSquareFeet:   input.BasementSquareFeet,
		BasementType:         input.BasementType,
		Bathrooms:            input.Bathrooms,
		Bedrooms:             input.Bedrooms,
		Carport:              input.Carport,
		City:                 input.City,
		Construction:         input.Construction,
		DaysOnMarket:         input.DaysOnMarket,
		Exterior:             input.Exterior,
		ExteriorFeatures:     input.ExteriorFeatures,
		Fireplace:            input.Fireplace,
		FullBaths:            input.FullBaths,
		Garage:               input.Garage,
		GarageDescription:    input.GarageDescription,
		HalfBaths:            input.HalfBaths,
		HOAFee:               input.HOAFee,
		ListPrice:            input.ListPrice,
		LotSize:              input.LotSize,
		MLSNumber:            input.MlsNumber,
		MLSComments:          input.MlsComments,
		OriginalListPrice:    input.OriginalListPrice,
		ParkingSpacesCarport: input.ParkingSpacesCarport,
		ParkingSpacesGarage:  input.ParkingSpacesGarage,
		Pool:                 input.Pool,
		Porch:                input.Porch,
		PriceClosed:          input.PriceClosed,
		PriceList:            input.PriceList,
		PropertyStyle:        input.PropertyStyle,
		Proplmg:              input.Proplmg,
		Proximity:            input.Proximity,
		RealEstateOwned:      input.RealEstateOwned,
		SalePrice:            input.SalePrice,
		SaleType:             input.SaleType,
		SelType:              input.SelType,
		ShortSale:            input.ShortSale,
		SquareFootage:        input.SquareFootage,
		Status:               input.Status,
		StreetDirection:      input.StreetDirection,
		StreetName:           input.StreetName,
		StreetNumber:         input.StreetNumber,
		StreetType:           input.StreetType,
		Subdivision:          input.Subdivision,
		TermsofSale:          input.TermsOfSale,
		TotalRooms:           input.TotalRooms,
		TotalUnits:           input.TotalUnits,
		UnitNumber:           input.UnitNumber,
		View:                 input.View,
		Waterfront:           input.Waterfront,
		YearBuilt:            input.YearBuilt,
		Zip:                  input.Zip,
		CreatedDateTime:      primitive.DateTime(millis.NowInMillis()),
	}

	if input.ListDate != nil {
		ListDate, err := pointers.StringTimeToPrimitive(*input.ListDate)
		if err != nil {
			return "", err
		}
		newIformGrid.ListDate = ListDate
	}
	if input.OriginalListDate != nil {
		OriginalListDate, err := pointers.StringTimeToPrimitive(*input.OriginalListDate)
		if err != nil {
			return "", err
		}
		newIformGrid.OriginalListDate = OriginalListDate
	}
	if input.SaleDate != nil {
		SaleDate, err := pointers.StringTimeToPrimitive(*input.OriginalListDate)
		if err != nil {
			return "", err
		}
		newIformGrid.SaleDate = SaleDate
	}

	res, err := DbCollections.IformGrids.InsertOne(ctx, newIformGrid)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func DeleteIformGrid(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.IformGrids.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchIformGrids(ctx context.Context, offset, limit int, pipelineID string, search *string) ([]*IformGrid, error) {

	filter := bson.M{
		"pipelineId": pipelineID,
	}
	if search != nil {
		filter["mlsNumber"] = *search
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.IformGrids.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query IformGrid:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawIformGrids := make([]*IformGrid, 0)
	for cur.Next(ctx) {
		a := &IformGrid{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode IformGrids entry: %v", err)
			return nil, err
		}
		rawIformGrids = append(rawIformGrids, a)
	}

	return rawIformGrids, nil
}

func GetIformGridsCount(ctx context.Context, pipelineID string, search *string) (*int64, error) {

	filter := bson.M{
		"pipelineId": pipelineID,
	}
	if search != nil {
		filter["mlsNumber"] = *search
	}

	count, err := DbCollections.IformGrids.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query IformGrid: %v", err)
		return nil, err
	}

	return &count, nil
}
