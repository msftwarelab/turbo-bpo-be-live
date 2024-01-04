package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	constants "github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	UserID          string              `bson:"userId"`
	Category        string              `bson:"category"`
	Label           string              `bson:"label"`
	Value           string              `bson:"value"`
	Section         *string             `bson:"section,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Comment) ToModels() *models.Comment {
	res := &models.Comment{
		ID:       u.ID.Hex(),
		Category: u.Category,
		Label:    u.Label,
		Value:    u.Value,
		Section:  strings.ObjectTOString(u.Section),
	}
	if u.Section == nil {
		res.Section = constants.CommentSectionDefaultValue
	}
	return res

}
