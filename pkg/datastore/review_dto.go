package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PipelineId        string             `bson:"name,omitempty"`
	OrderNumber       *string            `bson:"orderNumber,omitempty"`
	Address           *string            `bson:"address,omitempty"`
	FileName          *string            `bson:"fileName,omitempty"`
	AssignedTo        *string            `bson:"assignedTo,omitempty"`
	ReviewDescription *string            `bson:"reviewDescription,omitempty"`
	ReviewBy          *string            `bson:"reviewBy,omitempty"`
	Attachement       *string            `bson:"attachements,omitempty"`
	CreatedDateTime   primitive.DateTime `bson:"createdDateTime"`
}

func (u *Review) ToModels() *models.Review {

	return &models.Review{
		ID:                strings.ToObject(u.ID.Hex()),
		OrderNumber:       u.OrderNumber,
		Address:           u.Address,
		AssignedTo:        u.AssignedTo,
		ReviewDescription: u.ReviewDescription,
		ReviewDate:        strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		ReviewBy:          u.ReviewBy,
		URL:               u.Attachement,
		FileName:          u.FileName,
		PipelineID:        strings.ToObject(u.PipelineId),
	}
}
