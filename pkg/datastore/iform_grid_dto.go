package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IformGrid struct {
	ID                   primitive.ObjectID  `bson:"_id,omitempty" bson:"_id,omitempty,omitempty,omitempty"`
	PipelineID           string              `bson:"pipelineId,omitempty"`
	Address              *string             `bson:"Address,omitempty"`
	Age                  *string             `bson:"Age,omitempty"`
	BasementFinishedSqFt *string             `bson:"basementFinishedSqFt,omitempty"`
	BasementSquareFeet   *string             `bson:"basementSquareFeet,omitempty"`
	BasementType         *string             `bson:"basementType,omitempty"`
	Bathrooms            *string             `bson:"bathrooms,omitempty"`
	Bedrooms             *string             `bson:"bedrooms,omitempty"`
	Carport              *string             `bson:"carport,omitempty"`
	City                 *string             `bson:"city,omitempty"`
	Construction         *string             `bson:"construction,omitempty"`
	DaysOnMarket         *string             `bson:"daysOnMarket,omitempty"`
	Exterior             *string             `bson:"exterior,omitempty"`
	ExteriorFeatures     *string             `bson:"exteriorFeatures,omitempty"`
	Fireplace            *string             `bson:"fireplace,omitempty"`
	FullBaths            *string             `bson:"fullBaths,omitempty"`
	Garage               *string             `bson:"garage,omitempty"`
	GarageDescription    *string             `bson:"garageDescription,omitempty"`
	HalfBaths            *string             `bson:"halfBaths,omitempty"`
	HOAFee               *string             `bson:"hOAFee,omitempty"`
	ListDate             *primitive.DateTime `bson:"listDate,omitempty"`
	ListPrice            *string             `bson:"listPrice,omitempty"`
	LotSize              *string             `bson:"lotSize,omitempty"`
	MLSNumber            *string             `bson:"mLSNumber,omitempty"`
	MLSComments          *string             `bson:"mLSComments,omitempty"`
	OriginalListDate     *primitive.DateTime `bson:"originalListDate,omitempty"`
	OriginalListPrice    *string             `bson:"originalListPrice,omitempty"`
	ParkingSpacesCarport *string             `bson:"ParkingSpacesCarport,omitempty"`
	ParkingSpacesGarage  *string             `bson:"parkingSpacesGarage,omitempty"`
	Pool                 *string             `bson:"pool,omitempty"`
	Porch                *string             `bson:"porch,omitempty"`
	PriceClosed          *string             `bson:"priceClosed,omitempty"`
	PriceList            *string             `bson:"priceList,omitempty"`
	PropertyStyle        *string             `bson:"propertyStyle,omitempty"`
	Proplmg              *string             `bson:"proplmg,omitempty"`
	Proximity            *string             `bson:"proximity,omitempty"`
	RealEstateOwned      *string             `bson:"realEstateOwned,omitempty"`
	SaleDate             *primitive.DateTime `bson:"saleDate,omitempty"`
	SalePrice            *string             `bson:"salePrice,omitempty"`
	SaleType             *string             `bson:"saleType,omitempty"`
	SelType              *string             `bson:"selType,omitempty"`
	ShortSale            *string             `bson:"shortSale,omitempty"`
	SquareFootage        *string             `bson:"squareFootage,omitempty"`
	Status               *string             `bson:"status,omitempty"`
	StreetDirection      *string             `bson:"streetDirection,omitempty"`
	StreetName           *string             `bson:"streetName,omitempty"`
	StreetNumber         *string             `bson:"streetNumber,omitempty"`
	StreetType           *string             `bson:"streetType,omitempty"`
	Subdivision          *string             `bson:"subdivision,omitempty"`
	TermsofSale          *string             `bson:"termsofSale,omitempty"`
	TotalRooms           *string             `bson:"totalRooms,omitempty"`
	TotalUnits           *string             `bson:"totalUnits,omitempty"`
	UnitNumber           *string             `bson:"unitNumber,omitempty"`
	View                 *string             `bson:"view,omitempty"`
	Waterfront           *string             `bson:"waterfront,omitempty"`
	YearBuilt            *string             `bson:"yearBuilt,omitempty"`
	Zip                  *string             `bson:"zip,omitempty"`
	CreatedDateTime      primitive.DateTime  `bson:"createdDateTime,omitempty"`
}

func (u *IformGrid) ToModels() *models.IformGrid {
	return &models.IformGrid{
		ID:                   strings.ToObject(u.ID.Hex()),
		Address:              u.Address,
		Age:                  u.Age,
		BasementFinishedSqFt: u.BasementFinishedSqFt,
		BasementSquareFeet:   u.BasementSquareFeet,
		BasementType:         u.BasementType,
		Bathrooms:            u.Bathrooms,
		Bedrooms:             u.Bedrooms,
		Carport:              u.Carport,
		City:                 u.City,
		Construction:         u.Construction,
		DaysOnMarket:         u.DaysOnMarket,
		Exterior:             u.Exterior,
		ExteriorFeatures:     u.ExteriorFeatures,
		Fireplace:            u.Fireplace,
		FullBaths:            u.FullBaths,
		Garage:               u.Garage,
		GarageDescription:    u.GarageDescription,
		HalfBaths:            u.HalfBaths,
		HOAFee:               u.HOAFee,
		ListDate:             strings.ToObject(TimeConversion(u.ListDate)),
		ListPrice:            u.ListPrice,
		LotSize:              u.LotSize,
		MlsNumber:            u.MLSNumber,
		MlsComments:          u.MLSComments,
		OriginalListDate:     strings.ToObject(TimeConversion(u.OriginalListDate)),
		OriginalListPrice:    u.OriginalListPrice,
		ParkingSpacesCarport: u.ParkingSpacesCarport,
		ParkingSpacesGarage:  u.ParkingSpacesGarage,
		Pool:                 u.Pool,
		Porch:                u.Porch,
		PriceClosed:          u.PriceClosed,
		PriceList:            u.PriceList,
		PropertyStyle:        u.PropertyStyle,
		Proplmg:              u.Proplmg,
		Proximity:            u.Proximity,
		RealEstateOwned:      u.RealEstateOwned,
		SaleDate:             strings.ToObject(TimeConversion(u.SaleDate)),
		SalePrice:            u.SalePrice,
		SaleType:             u.SaleType,
		SelType:              u.SelType,
		ShortSale:            u.ShortSale,
		SquareFootage:        u.SquareFootage,
		Status:               u.Status,
		StreetDirection:      u.StreetDirection,
		StreetName:           u.StreetName,
		StreetNumber:         u.StreetNumber,
		StreetType:           u.StreetType,
		Subdivision:          u.Subdivision,
		TermsOfSale:          u.TermsofSale,
		TotalRooms:           u.TotalRooms,
		TotalUnits:           u.TotalUnits,
		UnitNumber:           u.UnitNumber,
		View:                 u.View,
		Waterfront:           u.Waterfront,
		YearBuilt:            u.YearBuilt,
		Zip:                  u.Zip,
	}
}
