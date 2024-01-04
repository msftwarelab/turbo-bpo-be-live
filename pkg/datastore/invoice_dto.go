package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty"`
	PipelineId            string              `bson:"pipelineId,omitempty"`
	Status                string              `bson:"status"`
	ClientId              *string             `bson:"clientId,omitempty"`
	EmployeeId            *string             `bson:"employeeId,omitempty"`
	Type                  *string             `bson:"type"`
	Name                  *string             `bson:"name"`
	OrderNumber           *string             `bson:"orderNumber"`
	Address               *string             `bson:"address"`
	Company               *string             `bson:"company"`
	Client                *string             `bson:"client"`
	OrderType             *string             `bson:"orderType"`
	IsSuperRush           *bool               `bson:"isSuperRush,omitempty"`
	SuperRushRemarks      *string             `bson:"superRushRemarks,omitempty"`
	IsRush                *bool               `bson:"isRush,omitempty"`
	RushRemarks           *string             `bson:"rushRemarks,omitempty"`
	IsInterior            *bool               `bson:"isInterior,omitempty"`
	InteriorRemarks       *string             `bson:"interiorRemarks,omitempty"`
	IsRentalAddendum      *bool               `bson:"isRentalAddendum,omitempty"`
	RentalAddendumRemarks *string             `bson:"RentalAddendumRemarks,omitempty"`
	IsInitialBpo          *bool               `bson:"isInitialBpo,omitempty"`
	InitialBpoRemarks     *string             `bson:"initialBpoRemarks,omitempty"`
	IsInspection          *bool               `bson:"isInspection,omitempty"`
	InspectionRemarks     *string             `bson:"inspectionRemarks,omitempty"`
	IsNoCsv               *bool               `bson:"isNoCsv,omitempty"`
	NoCsvRemarks          *string             `bson:"noCsvRemarks,omitempty"`
	IsNoIFill             *bool               `bson:"isNoIFill,omitempty"`
	NoIFillRemarks        *string             `bson:"noIFillRemarks,omitempty"`
	IsOtherPremium        *bool               `bson:"isOtherPremium,omitempty"`
	OtherPremiumRemarks   *string             `bson:"otherPremiumRemarks,omitempty"`
	CreatedBy             string              `bson:"createdBy,omitempty"`
	Updatedby             *string             `bson:"updatedBy,omitempty"`
	Remarks               *string             `bson:"remarks,omitempty"`
	QcType                *string             `bson:"qcType,omitempty"`
	CreatedDateTime       primitive.DateTime  `bson:"createdDateTime,omitempty"`
	LastUpdateTime        *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func (u *Invoice) ToModels() *models.Invoice {

	return &models.Invoice{
		ID:                    strings.ToObject(u.ID.Hex()),
		Type:                  u.Type,
		Name:                  u.Name,
		EmployeeID:            u.EmployeeId,
		OrderNumber:           u.OrderNumber,
		Address:               u.Address,
		Company:               u.Company,
		Client:                u.Client,
		OrderType:             u.OrderType,
		IsSuperRush:           u.IsSuperRush,
		SuperRushRemarks:      u.SuperRushRemarks,
		IsRush:                u.IsRush,
		RushRemarks:           u.RushRemarks,
		IsInterior:            u.IsInterior,
		InteriorRemarks:       u.InteriorRemarks,
		IsRentalAddendum:      u.IsRentalAddendum,
		RentalAddendumRemarks: u.RentalAddendumRemarks,
		IsInitialBpo:          u.IsInitialBpo,
		InitialBpoRemarks:     u.InitialBpoRemarks,
		IsInspection:          u.IsInspection,
		InspectionRemarks:     u.InspectionRemarks,
		IsNoCsv:               u.IsNoCsv,
		NoCsvRemarks:          u.NoCsvRemarks,
		IsNoIFill:             u.IsNoIFill,
		NoIFillRemarks:        u.NoIFillRemarks,
		IsOtherPremium:        u.IsOtherPremium,
		OtherPremiumRemarks:   u.OtherPremiumRemarks,
		QcType:                u.QcType,
		Date:                  strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}
