package services

import (
	"context"
	"strings"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SavePipelineDoc(ctx context.Context, pipelineID string, input models.PipelineDocInput, createdBy string) (string, error) {
	fileUrl, err := awsS3.S3Uploader(input.Doc)
	if err != nil {
		return "", err
	}

	newID, err := datastore.AddPipelineDoc(ctx, pipelineID, input, *fileUrl, createdBy)

	//Todo, codition to add duetime if orderType==data entry
	if strings.ToUpper(input.Type) == constants.PipelineDocTypeCompsMls {

		currentPipeline, _ := datastore.GetPipelineByIdDataStore(ctx, pipelineID)

		if currentPipeline != nil {
			if currentPipeline.OrderType != nil {
				if strings.ToUpper(*currentPipeline.OrderType) == constants.PipelineOrderTypeDataEntry {
					datastore.UpdatePipelineDueDateTimeAndAddDocsCount(ctx, pipelineID, *currentPipeline.OrderType, currentPipeline.IsRushOrder, currentPipeline.IsSuperRush)
					return newID, err
				}
			}
		}
	}

	datastore.UpdatePipelineDocsCount(ctx, pipelineID, +1)
	return newID, err

}

func DeletePipelineDoc(ctx context.Context, pipelineId string, userId string) (bool, error) {
	isTrue, err := datastore.DeletePipelineDoc(ctx, pipelineId, userId)
	datastore.UpdatePipelineDocsCount(ctx, pipelineId, -1)
	return isTrue, err
}

func AllPipelineDoc(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelineDocResult, error) {

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

	rawPipelineQualityDocs, err := datastore.SearchPipelineDocs(ctx, pipelineId, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve pipelineDoc: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelineDocsCount(ctx, pipelineId)
	if err != nil {
		log.Debug("Failed to retrieve count of pipelineDoc: %v", err)
		return nil, err
	}

	pipelineQualityDocs := make([]*models.PipelineDoc, 0)
	for _, u := range rawPipelineQualityDocs {
		pipelineQualityDocs = append(pipelineQualityDocs, u.ToModels())
	}
	toInt := int(*count)
	return &models.PipelineDocResult{
		TotalCount: &toInt,
		Results:    pipelineQualityDocs,
	}, nil
}
