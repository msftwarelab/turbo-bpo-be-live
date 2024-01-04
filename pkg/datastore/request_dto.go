package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	Status          *string             `bson:"status,omitempty"`
	PipelineId      *string             `bson:"pepilineId,omitempty"`
	OrderNumber     *string             `bson:"orderNumber,omitempty"`
	OrderType       *string             `bson:"orderType,omitempty"`
	Address         *string             `bson:"address,omitempty"`
	Company         *string             `bson:"company,omitempty"`
	HasPhotos       *bool               `bson:"hasPhotos,omitempty"`
	ConditionType   *string             `bson:"conditionType,omitempty"`
	Type            *string             `bson:"type,omitempty"`
	RequestedBy     *string             `bson:"requestedBy,omitempty"`
	RequestedByID   *string             `bson:"requestedById,omitempty"`
	Remarks         *string             `bson:"remarks,omitempty"`
	UpdatedBy       *string             `bson:"updatedBy,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Request) ToModels() *models.Request {
	return &models.Request{
		ID:              strings.ToObject(u.ID.Hex()),
		OrderNumber:     u.OrderNumber,
		Address:         u.Address,
		Company:         u.Company,
		ConditionType:   u.ConditionType,
		RequestedByID:   u.RequestedByID,
		RequestedBy:     u.RequestedBy,
		PipelineID:      u.PipelineId,
		Status:          u.Status,
		Type:            u.Type,
		OrderType:       u.OrderType,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}

type RequestHistory struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	PipelineId      string             `bson:"pipelineId,omitempty"`
	Status          string             `bson:"status"`
	ClientId        string             `bson:"clientId,omitempty"`
	EmployeeId      string             `bson:"employeeId,omitempty"`
	Type            *string            `bson:"type,omitempty"`
	OrderNumber     *string            `bson:"orderNumber,omitempty"`
	Address         *string            `bson:"address,omitempty"`
	Company         *string            `bson:"company,omitempty"`
	Remarks         *string            `bson:"remarks,omitempty"`
	CreatedDateTime primitive.DateTime `bson:"createdDateTime,omitempty"`
}
