package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineDoc struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	PipelineId      *string             `bson:"pipelineId,omitempty"`
	Type            string              `bson:"type,omitempty"`
	Filename        string              `bson:"fileName,omitempty"`
	Url             string              `bson:"url,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PipelineDoc) ToModels() *models.PipelineDoc {
	return &models.PipelineDoc{
		ID:              strings.ToObject(u.ID.Hex()),
		Type:            strings.ToObject(u.Type),
		FileName:        strings.ToObject(u.Filename),
		URL:             strings.ToObject(u.Url),
		CreatedBy:       strings.ToObject(u.CreatedBy),
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}

type GetCountGroupbyId struct {
	PipelineId string `bson:"_id,omitempty"`
	Count      int    `bson:"count,omitempty"`
}
