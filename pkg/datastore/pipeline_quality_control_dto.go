package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineQualityControl struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	PipelineId      string              `bson:"pipelineId,omitempty"`
	Message         string              `bson:"message,omitempty"`
	RequestType     string              `bson:"requestType,omitempty"`
	Type            string              `bson:"type,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	Status          *string             `bson:"status,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PipelineQualityControl) ToModels() *models.PipelineQualityControl {
	return &models.PipelineQualityControl{
		OrderNotes:      &u.Message,
		CreatedBy:       &u.CreatedBy,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}

func (u *PipelineQualityControl) ToModelAndNotes() *models.PipelineQualityControlAndNote {
	return &models.PipelineQualityControlAndNote{
		Message:   &u.Message,
		Date:      strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		CreatedBy: &u.CreatedBy,
		Category:  &u.Type,
	}
}
