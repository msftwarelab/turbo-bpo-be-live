package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SavePipelineQualityControl(ctx context.Context, pipelineId string, orderNotes string, createdBy string) (string, error) {

	//pipelineStatus validation
	pipelineRaw, err := datastore.GetPipelineById(ctx, pipelineId)
	if err != nil {
		return "", err
	}
	if pipelineRaw.Status == nil {
		return "", errs.EmptyPipelineStatus
	}

	allowedRoles := []*string{
		strings.ToObject(constants.PipelineStatusComplete),
		strings.ToObject(constants.PipelineStatusStandBy),
		strings.ToObject(constants.PipelineStatusPaid),
	}
	if !contains(allowedRoles, *pipelineRaw.Status) {
		return "", errs.InvalidPipelineStatus
	}

	pipelineQc, err := datastore.AddPipelineQualityControl(ctx, pipelineId, orderNotes, createdBy)
	if err != nil {
		return "", err
	}

	updatePipelinedata := models.UpdatePipelineInput{IsForQa: pointers.Bool(true)}
	isSuccess, err := UpdatePipeline(ctx, pipelineId, updatePipelinedata, createdBy, "")
	if err != nil || isSuccess == false {
		return "", errs.InvalidPipelineId
	}

	_, err = datastore.SaveQualityControl(ctx, pipelineId, createdBy)

	if err != nil {
		return "", err
	}
	return pipelineQc, nil

}

func AllPipelineQualityControl(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineQualityControlResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.FilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawPipelineQualityControls, err := datastore.SearchPipelineQualityControl(ctx, pipelineId, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve pipelineQualityControl: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelineQualityControlCount(ctx, pipelineId)
	if err != nil {
		log.Debug("Failed to retrieve count of pipelineQualityControl: %v", err)
		return nil, err
	}

	pipelineQualityControls := make([]*models.PipelineQualityControl, 0)
	for _, u := range rawPipelineQualityControls {
		pipelineQualityControls = append(pipelineQualityControls, u.ToModels())
	}
	toInt := int(*count)
	return &models.PipelineQualityControlResult{
		TotalCount: &toInt,
		Results:    pipelineQualityControls,
	}, nil

}

func AllPipelineQualityControlAndNote(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineQualityControlAndNoteResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.FilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawPipelineQualityControls, err := datastore.SearchPipelineQualityControlAndNotes(ctx, pipelineId, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve pipelineQualityControl: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelineQualityControlAndNotesCount(ctx, pipelineId)
	if err != nil {
		log.Debug("Failed to retrieve count of pipelineQualityControl: %v", err)
		return nil, err
	}

	pipelineQualityControlAndNotes := make([]*models.PipelineQualityControlAndNote, 0)
	for _, u := range rawPipelineQualityControls {
		pipelineQualityControlAndNotes = append(pipelineQualityControlAndNotes, u.ToModelAndNotes())
	}
	toInt := int(*count)
	return &models.PipelineQualityControlAndNoteResult{
		TotalCount: &toInt,
		Results:    pipelineQualityControlAndNotes,
	}, nil

}
