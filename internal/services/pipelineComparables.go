package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func AllPipelineComparable(ctx context.Context, pipelineID string, filter *models.PipelineComparableFilterInput) (*models.PipelineComparableResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.PipelineComparableFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawPipelineComparable, err := datastore.SearchPipelineComparables(ctx, *filter.Offset, *filter.Limit, pipelineID)
	if err != nil {
		log.Debug("Failed to retrieve loginLog: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelineComparablesCount(ctx, pipelineID)
	if err != nil {
		log.Debug("Failed to retrieve count of loginLog: %v", err)
		return nil, err
	}

	pipelineComparables := make([]*models.PipelineComparable, 0)
	for _, u := range rawPipelineComparable {
		pipelineComparables = append(pipelineComparables, u.ToModels())
	}
	toInt := int(*count)
	return &models.PipelineComparableResult{
		TotalCount: &toInt,
		Results:    pipelineComparables,
	}, nil
}

func SavePipelineComparable(ctx context.Context, pipelineID string, input models.SavePipelineComparableInput, cratedBy string) (string, error) {
	return datastore.SavePipelineComparable(ctx, cratedBy, pipelineID, input)
}

func UpdatePipelineComparable(ctx context.Context, id string, mls string, updatedBy string) (bool, error) {
	return datastore.UpdatePipelineComparable(ctx, id, mls)
}

func DeletePipelineComparable(ctx context.Context, id string, deletedBy string) (bool, error) {
	return datastore.DeletePipelineComparable(ctx, id)
}
