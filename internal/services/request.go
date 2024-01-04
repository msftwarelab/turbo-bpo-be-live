package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveRequest(ctx context.Context, pipelineId string, requestedBy string, requestedByID string) (string, error) {

	pipelineFilter := datastore.FilterById(pipelineId)
	pipelineRaw, err := datastore.SearchPipelines(ctx, *pipelineFilter, 0, 1)
	if err != nil {
		return "", err
	}
	if len(pipelineRaw) == 0 {
		return "", errs.InvalidPipelineId
	}
	if pipelineRaw[0] == nil {
		return "", errs.InvalidPipelineId
	}
	pipelinePhotosTotals, err := datastore.GetPipelinePhotosCountGroupByPipelineId(ctx, []string{pipelineId})

	hasPhotos := false
	if len(pipelinePhotosTotals) > 0 {
		hasPhotos = true
	}

	return datastore.AddRequest(ctx, pipelineId, requestedBy, requestedByID, pipelineRaw[0], hasPhotos)
}

func UpdateRequest(ctx context.Context, id string, input models.UpdateRequestInput, modifiedBy string) (bool, error) {
	return datastore.UpdateRequest(ctx, id, input, modifiedBy)
}

func AllRequest(ctx context.Context, filter *models.RequestFilterInput) (*models.RequestResult, error) {
	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.RequestFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawRequests, err := datastore.SearchRequests(ctx, *filter.Offset, *filter.Limit, filter)
	if err != nil {
		log.Debug("Failed to retrieve request: %v", err)
		return nil, err
	}

	count, err := datastore.GetRequestsCount(ctx, filter)
	if err != nil {
		log.Debug("Failed to retrieve count of request: %v", err)
		return nil, err
	}

	requestsModel := make([]*models.Request, 0)
	for _, u := range rawRequests {
		requestsModel = append(requestsModel, u.ToModels())
	}
	toInt := int(*count)
	return &models.RequestResult{
		TotalCount: &toInt,
		Results:    requestsModel,
	}, nil
}
