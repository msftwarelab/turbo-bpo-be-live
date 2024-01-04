package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileDoc struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	UserID          string             `bson:"userID"`
	Type            string             `bson:"type"`
	FileName        string             `bson:"fileName"`
	Url             string             `bson:"url"`
	CreatedDateTime primitive.DateTime `bson:"createdDateTime"`
}

func (u *ProfileDoc) ToModels() *models.Doc {
	return &models.Doc{
		ID:              u.ID.Hex(),
		Type:            u.Type,
		FileName:        u.FileName,
		URL:             u.Url,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}
