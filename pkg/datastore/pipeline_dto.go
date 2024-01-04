package datastore

import (
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pipeline struct {
	ID                        primitive.ObjectID         `bson:"_id,omitempty"`
	Status                    *string                    `bson:"status,omitempty"`
	UserId                    *string                    `bson:"userId,omitempty"`
	UserName                  *string                    `bson:"userName,omitempty"`
	OrderNumber               *string                    `bson:"orderNumber,omitempty"`
	Address                   *string                    `bson:"address,omitempty"`
	Country                   *string                    `bson:"country,omitempty"`
	County                    *string                    `bson:"county,omitempty"`
	Location                  *string                    `bson:"location,omitempty"`
	ZipCode                   *string                    `bson:"zipCode,omitempty"`
	Company                   *string                    `bson:"company,omitempty"`
	CompanyID                 *string                    `bson:"companyId,omitempty"`
	PremiumCompany            *string                    `bson:"premiumCompany,omitempty"`
	PremiumCompanyID          *string                    `bson:"premiumCompanyId,omitempty"`
	OtherCompany              *string                    `bson:"otherCompany,omitempty"`
	Type                      *string                    `bson:"type,omitempty"`
	OrderType                 *string                    `bson:"orderType,omitempty"`
	Objective                 *string                    `bson:"objective,omitempty"`
	Assign                    *string                    `bson:"assign,omitempty"`
	AssignId                  *string                    `bson:"assignId,omitempty"`
	Mls                       *string                    `bson:"mls,omitempty"`
	IsRushOrder               *bool                      `bson:"isRushOrder,omitempty"`
	IsSuperRush               *bool                      `bson:"isSuperRush,omitempty"`
	IsInspection              *bool                      `bson:"isInspection,omitempty"`
	IsInitialBpo              *bool                      `bson:"isInitialBpo,omitempty"`
	OrderFee                  *float64                   `bson:"orderFee,omitempty"`
	TotalFee                  *float64                   `bson:"totalFee,omitempty"`
	IsSyncedToTurboBpo        *bool                      `bson:"isSyncedToTurboBpo,omitempty"`
	RatingOverAll             *int                       `bson:"ratingOverAll,omitempty"`
	RatingTimeliness          *int                       `bson:"ratingTimeliness,omitempty"`
	RatingQuality             *int                       `bson:"ratingQuality,omitempty"`
	RatingFeedback            *string                    `bson:"ratingFeedback,omitempty"`
	IsProcessIform            *bool                      `bson:"isProcessIform,omitempty"`
	ProcessIformModifiedDate  *primitive.DateTime        `bson:"processIformModifiedDate,omitempty"`
	IsProcessIfill            *bool                      `bson:"isProcessIfill,omitempty"`
	IfillProcessModifiedDate  *primitive.DateTime        `bson:"ifillProcessModifiedDate,omitempty"`
	IsProcessReview           *bool                      `bson:"isProcessReview,omitempty"`
	ProcessReviewModifiedDate *primitive.DateTime        `bson:"processReviewModifiedDate,omitempty"`
	IsForQa                   *bool                      `bson:"isForQa,omitempty"`
	Coordinator               *string                    `bson:"coordinator,omitempty"`
	CoordinatorId             *string                    `bson:"coordinatorId,omitempty"`
	CreatedDateTime           primitive.DateTime         `bson:"createdDateTime,omitempty"`
	DueDateTime               *primitive.DateTime        `bson:"dueDateTime,omitempty"`
	HoldDateTime              *primitive.DateTime        `bson:"holdDateTime,omitempty"`
	LastUpdateTime            *primitive.DateTime        `bson:"lastUpdateTime,omitempty"`
	PipelineHistory           []*PipelineHistory         `bson:"history,omitempty"`
	AssignedHistory           []*PipelineAssignedHistory `bson:"assignedHistory,omitempty"`
	ActivationDateTime        *primitive.DateTime        `bson:"activationDate,omitempty"`
	AssignDateTime            *primitive.DateTime        `bson:"assignDateTime,omitempty"`
	PhotosCount               *int                       `bson:"photosCount,omitempty"`
	DocsCount                 *int                       `bson:"docsCount,omitempty"`
	Ishold                    *bool                      `bson:"ishold,omitempty"`
	HoldRemarks               *string                    `bson:"holdRemarks,omitempty"`
	UnHoldRemarks             *string                    `bson:"unHoldRemarks,omitempty"`
	CancelRemarks             *string                    `bson:"cancelRemarks,omitempty"`
	IsBilled                  *bool                      `bson:"isBilled,omitempty"`
	IsQC                      *bool                      `bson:"isQC ,omitempty"`
}

type PipelineHistory struct {
	LogDateTime primitive.DateTime `bson:"logDateTime,omitempty"`
	Action      *string            `bson:"action,omitempty"`
	Value       *string            `bson:"value,omitempty"`
	ModifiedBy  *string            `bson:"modifiedBy,omitempty"`
}

type PipelineAssignedHistory struct {
	LogDateTime  primitive.DateTime `bson:"logDateTime,omitempty"`
	Action       *string            `bson:"action,omitempty"`
	Assignee     *string            `bson:"assignee,omitempty"`
	AssigneeID   *string            `bson:"assigneeID,omitempty"`
	AssignedBy   *string            `bson:"assignedBy,omitempty"`
	AssignedByID *string            `bson:"assignedByID,omitempty"`
	ModifiedBy   *string            `bson:"modifiedBy,omitempty"`
}

type SalesAnalytics struct {
	ID             SalesAnalyticsID `bson:"_id"`
	CompletedOrder int              `bson:"count"`
}

type SalesAnalyticsID struct {
	Day   int `bson:"day"`
	Month int `bson:"month"`
	Year  int `bson:"year"`
}

type OrderAnalytics struct {
	ID     OrderAnalyticsID `bson:"_id"`
	Unpaid float32          `bson:"unpaid"`
	Paid   float32          `bson:"paid"`
}

type OrderAnalyticsID struct {
	Month    int    `bson:"month"`
	Year     int    `bson:"year"`
	ClientId string `bson:"userId"`
}

type ClientOrder struct {
	ID    ClientOrderId `bson:"_id"`
	Count *int          `bson:"count"`
}

type AssignedCount struct {
	ID    int32 `bson:"_id"`
	Count *int  `bson:"count"`
}

type ClientOrderId struct {
	ClientId string `bson:"userId"`
}

type ContractorAssignOrder struct {
	ID    ContractorAssignOrderId `bson:"_id"`
	Count *int                    `bson:"count"`
}

type ContractorAssignOrderId struct {
	ContactorId string `bson:"assignId"`
	Month       *int   `bson:"month"`
	Year        *int   `bson:"year"`
}

type OrderSubmitted struct {
	ID    OrderSubmittedId `bson:"_id"`
	Count *int             `bson:"count"`
}

type OrderSubmittedId struct {
	Coordinator string `bson:"coordinator"`
	Month       *int   `bson:"month"`
	Year        *int   `bson:"year"`
}

type ClientBalance struct {
	ClientId ClientBalanceId `bson:"_id"`
	Total    float64         `bson:"total"`
	Paid     float64         `bson:"paid"`
	Other    float64         `bson:"otjer"`
	Unpaid   float64         `bson:"unpaid"`
}

type ClientBalanceId struct {
	ClientId string `bson:"userId"`
}

func (u *OrderAnalytics) ToModels() *models.OrderAnalytics {
	return &models.OrderAnalytics{
		Month:  &u.ID.Month,
		Paid:   pointers.Float64(float64(u.Paid)),
		Unpaid: pointers.Float64(float64(u.Paid)),
		Client: &u.ID.ClientId,
	}
}

func TimeConversion(d *primitive.DateTime) string {
	if d != nil {
		datetime := (time.Unix(int64(*d)/1000, int64(*d)%1000*1000000))
		dateUtc, _ := datetime.MarshalText()
		return string(dateUtc)
	}
	return ""
}

func (u *SalesAnalytics) ToModels() *models.SalesAnalytics {
	return &models.SalesAnalytics{
		Day:            &u.ID.Day,
		CompletedOrder: &u.CompletedOrder,
	}
}

func (u *PipelineHistory) ToModels() *models.PipelineHistory {
	return &models.PipelineHistory{
		LogDateTime: strings.ToObject(TimeConversion(&u.LogDateTime)),
		Action:      u.Action,
		Value:       u.Value,
		ModifiedBy:  u.ModifiedBy,
	}
}

func (u *PipelineAssignedHistory) ToModels() *models.PipelineAssignedHistory {
	return &models.PipelineAssignedHistory{
		LogDateTime:  strings.ToObject(TimeConversion(&u.LogDateTime)),
		Action:       u.Action,
		Assignee:     u.Assignee,
		AssigneeID:   u.AssigneeID,
		AssignedBy:   u.AssignedBy,
		AssignedByID: u.AssignedByID,
		ModifiedBy:   u.ModifiedBy,
	}
}

func (u *Pipeline) ToModels() *models.Pipeline {

	historyList := make([]*models.PipelineHistory, 0)
	for _, v := range u.PipelineHistory {
		historyList = append(historyList, v.ToModels())
	}
	assignedHistory := make([]*models.PipelineAssignedHistory, 0)
	for _, v := range u.AssignedHistory {
		assignedHistory = append(assignedHistory, v.ToModels())

	}
	return &models.Pipeline{
		ID:                        u.ID.Hex(),
		OrderNumber:               u.OrderNumber,
		Status:                    u.Status,
		Address:                   u.Address,
		Country:                   u.Country,
		County:                    u.County,
		Location:                  u.Location,
		ZipCode:                   u.ZipCode,
		Company:                   u.Company,
		CompanyID:                 u.CompanyID,
		PremiumCompany:            u.PremiumCompany,
		PremiumCompanyID:          u.PremiumCompanyID,
		Type:                      u.Type,
		OrderType:                 u.OrderType,
		Objective:                 u.Objective,
		Assign:                    u.Assign,
		AssignID:                  u.AssignId,
		Mls:                       u.Mls,
		IsRushOrder:               u.IsRushOrder,
		IsSuperRush:               u.IsSuperRush,
		IsInspection:              u.IsInspection,
		IsInitialBpo:              u.IsInitialBpo,
		OrderFee:                  u.OrderFee,
		TotalFee:                  u.TotalFee,
		IsSyncedToTurboBpo:        u.IsSyncedToTurboBpo,
		RatingOverAll:             u.RatingOverAll,
		RatingTimeliness:          u.RatingTimeliness,
		RatingQuality:             u.RatingQuality,
		RatingFeedback:            u.RatingFeedback,
		IsProcessIform:            u.IsProcessIform,
		ProcessIformModifiedDate:  strings.ToObject(TimeConversion(u.ProcessIformModifiedDate)),
		IsProcessIfill:            u.IsProcessIfill,
		IfillProcessModifiedDate:  strings.ToObject(TimeConversion(u.IfillProcessModifiedDate)),
		IsProcessReview:           u.IsProcessReview,
		ProcessReviewModifiedDate: strings.ToObject(TimeConversion(u.ProcessReviewModifiedDate)),
		AuthorID:                  u.UserId,
		AuthorName:                u.UserName,
		CreatedDateTime:           strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		DueDateTime:               strings.ToObject(TimeConversion(u.DueDateTime)),
		HoldDateTime:              strings.ToObject(TimeConversion(u.HoldDateTime)),
		PipelineHistory:           historyList,
		AssignedHistory:           assignedHistory,
		IsHold:                    u.Ishold,
		HoldRemarks:               u.HoldRemarks,
		UnHoldRemarks:             u.UnHoldRemarks,
		CancelRemarks:             u.CancelRemarks,
		AssignDateTime:            strings.ToObject(TimeConversion(u.AssignDateTime)),
		OtherCompany:              u.OtherCompany,
		PipelinePhotoTotal:        u.PhotosCount,
		PipelineDocTotal:          u.DocsCount,
		IsBilled:                  u.IsBilled,
		IsQc:                      u.IsQC,
	}
}
