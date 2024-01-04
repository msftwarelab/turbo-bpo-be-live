package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineNeighborhood struct {
	ID                   string              `bson:"id,omitempty"`
	PipelineId           string              `bson:"pipelineId,omitempty"`
	IsReoDriven          bool                `bson:"isReoDriven"`
	MarketTrend          *string             `bson:"marketTrend,omitempty"`
	MonthlyPercentage    *int                `bson:"monthlyPercentage,omitempty"`
	SixmonthPercentage   *int                `bson:"_6monthPercentage,omitempty"`
	AnnualPercentage     *int                `bson:"annualPercentage,omitempty"`
	TotalListings        *int                `bson:"totalListings,omitempty"`
	Supply               *string             `bson:"supply,omitempty"`
	ListingsMinValue     *int                `bson:"listingsMinValue,omitempty"`
	ListingsMedValue     *int                `bson:"listingsMedValue,omitempty"`
	ListingsMaxValue     *int                `bson:"listingsMaxValue,omitempty"`
	ListingsDomAve       *int                `bson:"listingsDomAve,omitempty"`
	ListingsDomRangeFrom *int                `bson:"listingsDomRangeFrom,omitempty"`
	ListingsDomRangeTo   *int                `bson:"listingsDomRangeTo,omitempty"`
	Fm                   *int                `bson:"fm,omitempty"`
	Ss                   *int                `bson:"ss,omitempty"`
	Reo                  *int                `bson:"reo,omitempty"`
	Distressed           *int                `bson:"distressed,omitempty"`
	TotalSales           *float64            `bson:"totalSales,omitempty"`
	Demand               *string             `bson:"demand,omitempty"`
	SalesMinValue        *int                `bson:"salesMinValue,omitempty"`
	SalesMedValue        *int                `bson:"salesMedValue,omitempty"`
	SalesMaxValue        *int                `bson:"salesMaxValue,omitempty"`
	SalesDomRangeFrom    *int                `bson:"salesDomRangeFrom,omitempty"`
	SalesDomRangeTo      *int                `bson:"salesDomRangeTo,omitempty"`
	ZntComments          *string             `bson:"zntComments,omitempty"`
	NtComments           *string             `bson:"ntComments,omitempty"`
	LastUpdateTime       *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
	CreatedDateTime      primitive.DateTime  `bson:"createdDateTime,omitempty"`
}

func (u *PipelineNeighborhood) ToModels() *models.PipelineNeighborhood {
	return &models.PipelineNeighborhood{
		IsReoDriven:          pointers.Bool(u.IsReoDriven),
		MarketTrend:          u.MarketTrend,
		MonthlyPercentage:    u.MonthlyPercentage,
		SixmonthPercentage:   u.SixmonthPercentage,
		AnnualPercentage:     u.AnnualPercentage,
		TotalListings:        u.TotalListings,
		Supply:               u.Supply,
		ListingsMinValue:     u.ListingsMinValue,
		ListingsMedValue:     u.ListingsMedValue,
		ListingsMaxValue:     u.ListingsMaxValue,
		ListingsDomAve:       u.ListingsDomAve,
		ListingsDomRangeFrom: u.ListingsDomRangeFrom,
		ListingsDomRangeTo:   u.ListingsDomRangeTo,
		Fm:                   u.Fm,
		Ss:                   u.Ss,
		Reo:                  u.Reo,
		Distressed:           u.Distressed,
		TotalSales:           u.TotalSales,
		Demand:               u.Demand,
		SalesMinValue:        u.SalesMinValue,
		SalesMedValue:        u.SalesMedValue,
		SalesMaxValue:        u.SalesMaxValue,
		SalesDomRangeFrom:    u.SalesDomRangeFrom,
		SalesDomRangeTo:      u.SalesDomRangeTo,
		ZntComments:          u.ZntComments,
		NtComments:           u.NtComments,
	}
}
