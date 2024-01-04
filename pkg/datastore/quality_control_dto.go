package datastore

import (
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QualityControl struct {
	ID               primitive.ObjectID       `bson:"_id,omitempty"`
	PipelineId       *string                  `bson:"pipelineId,omitempty"`
	OrderNumber      *string                  `bson:"orderNumber,omitempty"`
	OrderAddress     *string                  `bson:"OrderAddress,omitempty"`
	OrderCompany     *string                  `bson:"OrderCompany,omitempty"`
	OrderType        *string                  `bson:"OrderType,omitempty"`
	QcId             int64                    `bson:"qcId,omitempty"`
	Requests         *int                     `bson:"requests,omitempty"`
	RequestStatus    *string                  `bson:"requestStatus,omitempty"`
	RequestType      *string                  `bson:"requestType,omitempty"`
	Status           *string                  `bson:"status,omitempty"`
	Assignee         *string                  `bson:"assignee,omitempty"`
	AssigneeName     *string                  `bson:"assigneeName,omitempty"`
	ContractorId     *string                  `bson:"contractorId,omitempty"`
	ContractorName   *string                  `bson:"orderAssignee,omitempty"` //name of contractor from pipline
	History          []*QualityControlHistory `bson:"history,omitempty"`
	CreatedBy        *string                  `bson:"createdBy,omitempty"`
	Reason           *string                  `bson:"reason,omitempty"`
	IsAccepted       *bool                    `bson:"isAccepted,omitempty"`
	AcceptedDatetime *primitive.DateTime      `bson:"acceptedDateTime,omitempty"`
	CreatedDateTime  primitive.DateTime       `bson:"createdDateTime,omitempty"`
	LastUpdateTime   *primitive.DateTime      `bson:"LastUpdateTime,omitempty"`
}

type QualityControlHistory struct {
	Status          *string             `bson:"status,omitempty""`
	Reason          *string             `bson:"reason,omitempty""`
	Date            *primitive.DateTime `bson:"date,omitempty""`
	By              *string             `bson:"by,omitempty""`
	CurrentAssignee *string             `bson:"currentAssignee,omitempty""`
	NewAssignee     *string             `bson:"newAssignee,omitempty""`
}

type QcCompleted struct {
	ID    QcCompletedID `bson:"_id"`
	Count *int          `bson:"count"`
}

type QcCompletedID struct {
	Assignee string `bson:"assigneeName"`
	Month    *int   `bson:"month"`
	Year     *int   `bson:"year"`
}

func (u *QualityControlHistory) ToModels() *models.QualityControlHistory {
	return &models.QualityControlHistory{
		Status:          u.Status,
		Reason:          u.Reason,
		Date:            strings.ToObject(TimeConversion(u.Date)),
		Cratedby:        u.By,
		CurrentAssignee: u.CurrentAssignee,
		NewAssignee:     u.NewAssignee,
	}
}

func (u *QualityControl) ToModels() *models.QualityControl {

	logList := make([]*models.QualityControlHistory, 0)
	for _, v := range u.History {

		logList = append(logList, v.ToModels())
	}

	return &models.QualityControl{

		ID:                strings.ToObject(int64ToString(u.QcId)),
		PipelineID:        u.PipelineId,
		Requests:          u.Requests,
		Status:            u.Status,
		Assignee:          u.AssigneeName,
		History:           logList,
		AssigneeID:        u.Assignee,
		OrderNumber:       u.OrderNumber,
		OrderContractor:   u.ContractorName,
		OrderContractorID: u.ContractorId,
		CreatedDateTime:   strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		LastUpdateTime:    strings.ToObject(TimeConversion(u.LastUpdateTime)),
	}
}

func int64ToString(i int64) string {
	return fmt.Sprintf("%v", i)

}
