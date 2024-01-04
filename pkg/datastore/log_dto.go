package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	Datetime   primitive.DateTime `bson:"datetime"`
	Action     string             `bson:"action"`
	Value      string             `bson:"value"`
	ModifiedBy string             `bson:"modifiedBy"`
}

type IformHistoryLog struct {
	CreatedDate primitive.DateTime `bson:"createdDate"`
	UpdatedDate primitive.DateTime `bson:"updatedDate"`
	Url         string             `bson:"url"`
	fileName    string             `bson:"fileName"`
	ModifiedBy  string             `bson:"modifiedBy"`
}

func (u *Log) ToModels() *models.Log {
	return &models.Log{
		Datetime:   TimeConversion(&u.Datetime),
		Action:     u.Action,
		Value:      u.Value,
		ModifiedBy: u.ModifiedBy,
	}
}

func (u *IformHistoryLog) ToModels() *models.IformHistory {
	return &models.IformHistory{
		CreatedDate: strings.ToObject(TimeConversion(&u.CreatedDate)),
		UpdatedDate: strings.ToObject(TimeConversion(&u.UpdatedDate)),
		ModifiedBy:  strings.ToObject(u.ModifiedBy),
		URL:         strings.ToObject(u.Url),
	}
}
