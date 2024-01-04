package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Default struct {
	ID                             *primitive.ObjectID `bson:"_id,omitempty"`
	UserId                         string              `bson:"userId"`
	ListingType                    string              `bson:"listingType"`
	AlwayssubmitOrder              bool                `bson:"alwayssubmitOrder"`
	AutoCompleteStandbyOrder       bool                `bson:"autoCompleteStandbyOrder"`
	InitialSearchGla               string              `bson:"initialSearchGla"`
	InitialSearchAge               string              `bson:"initialSearchAge"`
	InitialSearchProximity         string              `bson:"initialSearchProximity"`
	SecondSearchGla                string              `bson:"secondSearchGla"`
	SecondSearchAge                string              `bson:"secondSearchAge"`
	SecondSearchProximity          string              `bson:"secondSearchProximity"`
	SecondSearchSaleDates          string              `bson:"secondSearchSaleDates"`
	ThirdSearchGla                 string              `bson:"thirdSearchGla"`
	ThirdSearchAge                 string              `bson:"thirdSearchAge"`
	ThirdSearchProximity           string              `bson:"thirdSearchProximity"`
	ThirdSearchSaleDates           string              `bson:"thirdSearchSaleDates"`
	ThirdSearchFilterByComplexName bool                `bson:"thirdSearchFilterByComplexName"`
	ThirdSearchFilterByCity        bool                `bson:"thirdSearchFilterByCity"`
	ThirdSearchFilterByZip         bool                `bson:"thirdSearchFilterByZip"`
	ThirdSearchFilterByCountry     bool                `bson:"thirdSearchFilterByCountry"`
	UseDefaults                    bool                `bson:"useDefaults"`
	UseIformValidations            bool                `bson:"useIformValidations"`
	SubjectType                    string              `bson:"subjectType"`
	StyleDesign                    string              `bson:"styleDesign"`
	ExteriorFinish                 string              `bson:"exteriorFinish"`
	Condition                      string              `bson:"condition"`
	Quality                        string              `bson:"quality"`
	View                           string              `bson:"view"`
	Pool                           string              `bson:"pool"`
	PorchPatioDeck                 string              `bson:"porchPatioDeck"`
	FirePlace                      bool                `bson:"firePlace"`
	Basement                       string              `bson:"basement"`
	Condo                          string              `bson:"condo"`
	MultiUnit                      string              `bson:"multiUnit"`
	MobileHome                     string              `bson:"mobileHome"`
	Sfd                            string              `bson:"sfd"`
	SfaTownhouse                   string              `bson:"sfaTownhouse"`
	Theme                          string              `bson:"theme"`
	IsEnableEmailNotification      *bool               `json:"isEnableEmailNotification"`
	CreatedDateTime                primitive.DateTime  `bson:"createdDateTime"`
	LastUpdateTime                 *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Default) ToModels() *models.Default {

	if u != nil {
		return &models.Default{
			ListingType:                    u.ListingType,
			AlwayssubmitOrder:              u.AlwayssubmitOrder,
			AutoCompleteStandbyOrder:       u.AutoCompleteStandbyOrder,
			InitialSearchGla:               u.InitialSearchGla,
			InitialSearchAge:               u.InitialSearchAge,
			InitialSearchProximity:         u.InitialSearchProximity,
			SecondSearchGla:                u.SecondSearchGla,
			SecondSearchAge:                u.SecondSearchAge,
			SecondSearchProximity:          u.SecondSearchProximity,
			SecondSearchSaleDates:          u.SecondSearchSaleDates,
			ThirdSearchGla:                 u.ThirdSearchGla,
			ThirdSearchAge:                 u.ThirdSearchAge,
			ThirdSearchProximity:           u.ThirdSearchProximity,
			ThirdSearchSaleDates:           u.ThirdSearchSaleDates,
			ThirdSearchFilterByComplexName: u.ThirdSearchFilterByComplexName,
			ThirdSearchFilterByCity:        u.ThirdSearchFilterByCity,
			ThirdSearchFilterByZip:         u.ThirdSearchFilterByZip,
			ThirdSearchFilterByCountry:     u.ThirdSearchFilterByCountry,
			UseDefaults:                    u.UseDefaults,
			UseIformValidations:            u.UseIformValidations,
			SubjectType:                    u.SubjectType,
			StyleDesign:                    u.StyleDesign,
			ExteriorFinish:                 u.ExteriorFinish,
			Condition:                      u.Condition,
			Quality:                        u.Quality,
			View:                           u.View,
			Pool:                           u.Pool,
			PorchPatioDeck:                 u.PorchPatioDeck,
			FirePlace:                      u.FirePlace,
			Basement:                       u.Basement,
			Condo:                          u.Condo,
			MultiUnit:                      u.MultiUnit,
			MobileHome:                     u.MobileHome,
			Sfd:                            u.Sfd,
			SfaTownhouse:                   u.SfaTownhouse,
			Theme:                          &u.Theme,
			IsEnableEmailNotification:      u.IsEnableEmailNotification,
		}
	} else {
		return &models.Default{}
	}
}
