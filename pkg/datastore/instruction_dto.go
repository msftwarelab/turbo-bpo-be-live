package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Instruction struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Tag             string              `bson:"tag,omitempty"`
	Client          string              `bson:"client,omitempty"`
	ClientId        string              `bson:"clientId,omitempty"`
	Company         string              `bson:"company,omitempty"`
	CompanyId       string              `bson:"companyId,omitempty"`
	FileName        string              `bson:"fileName,omitempty"`
	Url             string              `bson:"url,omitempty"`
	Comment         string              `bson:"comment,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Instruction) ToModels() *models.Instruction {

	return &models.Instruction{
		ID:              strings.ToObject(u.ID.Hex()),
		Tag:             strings.ToObject(u.Tag),
		Client:          strings.ToObject(u.Client),
		ClientID:        strings.ToObject(u.ClientId),
		Company:         strings.ToObject(u.Company),
		CompanyID:       strings.ToObject(u.CompanyId),
		URL:             strings.ToObject(u.Url),
		FileName:        strings.ToObject(u.FileName),
		Comment:         strings.ToObject(u.Comment),
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}
