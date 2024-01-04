package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelinePhoto struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	PipelineId      *string             `bson:"pipelineId,omitempty"`
	IsSubmitted     bool                `bson:"isSubmitted,omitempty"`
	Filename        string              `bson:"fileName,omitempty"`
	Url             string              `bson:"url,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	FileSize        *int64              `bson:"fileSize,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PipelinePhoto) ToModels() *models.PipelinePhoto {
	return &models.PipelinePhoto{
		ID:              strings.ToObject(u.ID.Hex()),
		IsSubmitted:     pointers.Bool(u.IsSubmitted),
		FileName:        strings.ToObject(u.Filename),
		FileSize:        pointers.Int64ToInt(u.FileSize),
		CreatedBy:       strings.ToObject(u.CreatedBy),
		URL:             strings.ToObject(u.Url),
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}
