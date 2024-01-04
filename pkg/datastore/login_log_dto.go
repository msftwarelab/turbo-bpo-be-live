package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginLog struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Datetime  primitive.DateTime `bson:"datetime"`
	IPAddress string             `bson:"ipAddress"`
}

func (u *LoginLog) ToModels() *models.LoginLog {
	return &models.LoginLog{
		Username:  strings.ToObject(u.Username),
		Datetime:  strings.ToObject(TimeConversion(&u.Datetime)),
		IPAddress: strings.ToObject(u.IPAddress),
	}
}
