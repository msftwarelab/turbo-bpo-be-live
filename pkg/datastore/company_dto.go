package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            *string             `bson:"name,omitempty"`
	WebSite         string              `bson:"website,omitempty"`
	IsAdmin         bool                `bson:"isAdmin,omitempty"`
	IsClient        bool                `bson:"isClient,omitempty"`
	IsPremium       bool                `bson:"isPremium,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
	Forms           []*CompanyForm      `bson:"forms,omitempty"`
}

type CompanyForm struct {
	Name  *string `bson:"name,omitempty"`
	Style *string `bson:"stype,omitempty"`
}

func (u *Company) ToModels() *models.Company {

	forms := make([]*models.CompanyForm, 0)
	for _, v := range u.Forms {
		forms = append(forms, &models.CompanyForm{v.Name, v.Style})
	}

	return &models.Company{
		ID:        strings.ToObject(u.ID.Hex()),
		Name:      u.Name,
		WebSite:   strings.ToObject(u.WebSite),
		IsAdmin:   &u.IsAdmin,
		IsClient:  &u.IsClient,
		IsPremium: &u.IsPremium,
		Forms:     forms,
	}
}
