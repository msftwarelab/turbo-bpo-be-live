package datastore

import (
	"context"
	"strings"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	stringsUtils "github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PipelineState struct {
	MaxDailyVolume      *int                `bson:"maxDailyVolume"`
	StandByAutoComplete *int                `bson:"standByAutoComplete"`
	IsRush              *bool               `bson:"isRush"`
	IsNewOrder          *bool               `bson:"isNewOrder"`
	OrderMessage        *string             `bson:"orderMessage"`
	TTSlow              *int                `bson:"tTSlow"`
	TTModerate          *int                `bson:"tTModerate"`
	TTBusy              *int                `bson:"tTBusy"`
	TTMax               *int                `bson:"tTMax"`
	TLSlow              *string             `bson:"tLSlow"`
	TLModerate          *string             `bson:"tLModerate"`
	TLBusy              *string             `bson:"tLBusy"`
	OPInterior          *float64            `bson:"oPInterior"`
	OPExterior          *float64            `bson:"oPExterior"`
	OPDataEntry         *float64            `bson:"oPDataEntry"`
	OPRush              *float64            `bson:"oPRush"`
	OPSuperRush         *float64            `bson:"oPSuperRush"`
	OPConditionReport   *float64            `bson:"oPConditionReport"`
	OPRentalAddendum    *float64            `bson:"oPRentalAddendum"`
	OPInitialBpo        *float64            `bson:"oPInitialBPO"`
	OPInspection        *float64            `bson:"oPInspection"`
	PCIsAcceptOrder     *bool               `bson:"pCIsAcceptOrder"`
	PCcatchTime         *int                `bson:"pCcatchTime"`
	OAOfferLimitInMin   *int                `bson:"oAOfferLimitInMin"`
	OAIsAutoAssign      *bool               `bson:"oAIsAutoAssign"`
	QCElapseTime        *int                `bson:"qCElapseTime"`
	LastUpdateTime      *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

func isBetween(value, from, to float64) bool {
	if value >= from && value <= to {
		return true
	}
	return false
}

func (u *PipelineState) TurnAoundTime(ctx context.Context) (*int, error) {
	todayOrderCount, _ := GetTodayOrderCount(ctx)

	if todayOrderCount == nil {
		todayOrderCount = pointers.Int(0)
	}

	TlSlowSlice := strings.Split(stringsUtils.ObjectTOString(u.TLSlow), "-")
	if len(TlSlowSlice) != 2 {
		return nil, errs.InvalidTurboLoad
	}
	if isBetween(float64(*todayOrderCount), utils.Str2Float64(TlSlowSlice[0]), utils.Str2Float64(TlSlowSlice[1])) {
		return u.TTSlow, nil
	}

	TlModerateSlice := strings.Split(stringsUtils.ObjectTOString(u.TLModerate), "-")
	if len(TlModerateSlice) != 2 {
		return nil, errs.InvalidTurboLoadModerate
	}
	if isBetween(float64(*todayOrderCount), utils.Str2Float64(TlModerateSlice[0]), utils.Str2Float64(TlModerateSlice[1])) {
		return u.TTModerate, nil
	}

	TlBusySlice := strings.Split(stringsUtils.ObjectTOString(u.TLBusy), "-")
	if len(TlBusySlice) != 2 {
		return nil, errs.InvalidTurboLoadBusy
	}
	if isBetween(float64(*todayOrderCount), utils.Str2Float64(TlBusySlice[0]), utils.Str2Float64(TlBusySlice[1])) {
		return u.TTBusy, nil
	}
	return u.TTMax, nil

}

func (u *PipelineState) ToModels() *models.PipelineState {
	return &models.PipelineState{
		MaxDailyVolume:      u.MaxDailyVolume,
		StandByAutoComplete: u.StandByAutoComplete,
		IsRush:              u.IsRush,
		IsNewOrder:          u.IsNewOrder,
		OrderMessage:        u.OrderMessage,
		TTSlow:              u.TTSlow,
		TTModerate:          u.TTModerate,
		TTBusy:              u.TTBusy,
		TTMax:               u.TTMax,
		TLSlow:              u.TLSlow,
		TLModerate:          u.TLModerate,
		TLBusy:              u.TLBusy,
		OPInterior:          u.OPInterior,
		OPExterior:          u.OPExterior,
		OPDataEntry:         u.OPDataEntry,
		OPRush:              u.OPRush,
		OPSuperRush:         u.OPSuperRush,
		OPConditionReport:   u.OPConditionReport,
		OPRentalAddendum:    u.OPRentalAddendum,
		OPInitialBpo:        u.OPInitialBpo,
		OPInspection:        u.OPInspection,
		PCIsAcceptOrder:     u.PCIsAcceptOrder,
		PCcatchTime:         u.PCcatchTime,
		OAOfferLimitInMin:   u.OAOfferLimitInMin,
		OAIsAutoAssign:      u.OAIsAutoAssign,
		QCElapseTime:        u.QCElapseTime,
	}
}
