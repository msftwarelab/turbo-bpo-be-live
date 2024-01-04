package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionGroup struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            *string             `bson:"name,omitempty"`
	Permissions     []*string           `bson:"permissions,omitempty"`
	CreatedBy       string              `bson:"createdBy,omitempty"`
	Status          string              `bson:"status,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *PermissionGroup) ToModels() *models.PermissionGroup {
	return &models.PermissionGroup{
		ID:          strings.ToObject(u.ID.Hex()),
		Name:        u.Name,
		Permissions: u.Permissions,
	}
}
