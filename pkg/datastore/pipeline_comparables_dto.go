package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineComparable struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Mls             *string             `bson:"mls,omitempty"`
	Status          *string             `bson:"status,omitempty"`
	PipelineID      string              `bson:"pipelineId,omitempty"`
	Order           int                 `bson:"order,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PipelineComparable) ToModels() *models.PipelineComparable {
	return &models.PipelineComparable{
		ID:     strings.ToObject(u.ID.Hex()),
		Mls:    u.Mls,
		Status: u.Status,
		Order:  &u.Order,
	}
}
