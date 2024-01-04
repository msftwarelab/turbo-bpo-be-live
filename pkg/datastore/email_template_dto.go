package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailTemplate struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Type            string              `bson:"type,omitempty"`
	Template        string              `bson:"template,omitempty"`
	Subject         string              `bson:"subject,omitempty"`
	Message         string              `bson:"message,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *EmailTemplate) ToModels() *models.EmailTemplate {

	return &models.EmailTemplate{
		ID:       strings.ToObject(u.ID.Hex()),
		Type:     &u.Type,
		Template: &u.Template,
		Subject:  &u.Subject,
		Message:  &u.Message,
	}
}
