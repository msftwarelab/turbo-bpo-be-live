package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Adjustment struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	UserID          string              `bson:"userId"`
	Category        string              `bson:"category"`
	Order           int                 `bson:"order"`
	Label           string              `bson:"label"`
	From            float64             `bson:"from"`
	To              float64             `bson:"to"`
	Value           float64             `bson:"value"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Adjustment) ToModels() *models.Adjustment {
	return &models.Adjustment{
		ID:       u.ID.Hex(),
		Category: u.Category,
		Order:    u.Order,
		Label:    u.Label,
		From:     u.From,
		To:       u.To,
		Value:    u.Value,
	}
}
