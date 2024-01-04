package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func AllIformGrid(ctx context.Context, pipelinID string, filter *models.IformGridFilterInput) (*models.IformGridResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.IformGridFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawIformGrid, err := datastore.SearchIformGrids(ctx, *filter.Offset, *filter.Limit, pipelinID, filter.Search)
	if err != nil {
		log.Debug("Failed to retrieve loginLog: %v", err)
		return nil, err
	}

	count, err := datastore.GetIformGridsCount(ctx, pipelinID, filter.Search)
	if err != nil {
		log.Debug("Failed to retrieve count of iformGrid: %v", err)
		return nil, err
	}

	iformGrids := make([]*models.IformGrid, 0)
	for _, u := range rawIformGrid {
		iformGrids = append(iformGrids, u.ToModels())
	}
	toInt := int(*count)
	return &models.IformGridResult{
		TotalCount: &toInt,
		Results:    iformGrids,
	}, nil
}

func SaveIformGrid(ctx context.Context, pipelineIdD string, input models.SaveIformGridInput, cratedBy string) (string, error) {
	return datastore.SaveIformGrid(ctx, pipelineIdD, input, cratedBy)
}

// func UpdateIformGrid(ctx context.Context, id string, input models.UpdateIformGridInput, updatedBy string) (bool, error) {
// 	return datastore.UpdateIformGrid(ctx, id, input)
// }

func DeleteIformGrid(ctx context.Context, id string, deletedBy string) (bool, error) {
	return datastore.DeleteIformGrid(ctx, id)
}
