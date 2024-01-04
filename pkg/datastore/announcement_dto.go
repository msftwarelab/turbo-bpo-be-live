package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Announcement struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty" bson:"_id,omitempty"`
	Subject         *string             `bson:"subject,omitempty"`
	StartDate       *primitive.DateTime `bson:"startDate,omitempty"`
	EndDate         *primitive.DateTime `bson:"endDate,omitempty"`
	Message         *string             `bson:"message,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Announcement) ToModels() *models.Announcement {

	return &models.Announcement{
		ID:              strings.ToObject(u.ID.Hex()),
		Subject:         u.Subject,
		Message:         u.Message,
		StartDate:       strings.ToObject(TimeConversion(u.StartDate)),
		EndDate:         strings.ToObject(TimeConversion(u.EndDate)),
		CreatedBy:       &u.CreatedBy,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}
