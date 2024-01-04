package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Header struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string              `bson:"name,omitempty"`
	IsParent        bool                `bson:"isParent,omitempty"`
	ParentId        string              `bson:"parentId,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Header) ToModels() *models.Header {

	return &models.Header{
		ID:   strings.ToObject(u.ID.Hex()),
		Name: strings.ToObject(u.Name),
	}
}
