package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string              `bson:"userId"`
	InvoiceDate     primitive.DateTime  `bson:"invoiceDate"`
	Isrunning       bool                `bson:"isRunning"`
	Start           primitive.DateTime  `bson:"start"`
	End             *primitive.DateTime `bson:"end,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Session) ToModels() *models.Session {
	return &models.Session{
		ID:          strings.ToObject(u.ID.Hex()),
		InvoiceDate: strings.ToObject(TimeConversion(&u.InvoiceDate)),
		Start:       strings.ToObject(TimeConversion(&u.Start)),
		End:         strings.ToObject(TimeConversion(u.End)),
	}
}
