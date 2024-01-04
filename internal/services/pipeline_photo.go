package services

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SavePipelinePhoto(ctx context.Context, pipelineId string, input models.PipelinePhotoInput, userId string, userName string) (string, error) {
	fileUrl, err := awsS3.S3Uploader(input.Doc)
	if err != nil {
		return "", err
	}
	// myDefault, err := datastore.GetDefault(ctx, userId)
	// if err != nil || myDefault == nil {
	// 	return "", errs.NoDefaultdata
	// }
	id, err := datastore.AddPipelinePhoto(ctx, pipelineId, input, userId, *fileUrl, false, userName)
	if err != nil {
		return "", err
	}
	pipelineHistoryValue := fmt.Sprintf("added pipeline photo picture URL: %s", *fileUrl)
	datastore.AddSetPipelineHistory(ctx, pipelineId, "ADD_PHOTO", pipelineHistoryValue, userName)
	return id, nil
}

func SubmitPipelinePhoto(ctx context.Context, id string, userId string, userName string, isSubmitPipelinePhoto bool) (bool, error) {
	isSubmit, err := datastore.SubmitPipelinePhoto(ctx, id, userId, isSubmitPipelinePhoto)
	if err != nil {
		return false, err
	}
	//pipelineHistoryValue := fmt.Sprintf("pipeline photo submitted")
	//datastore.AddSetPipelineHistory(ctx, pipelineId, "SUBMIT_PHOTO", pipelineHistoryValue, userName)
	return isSubmit, nil
}

func DeletePipelinePhoto(ctx context.Context, pipelinePhotoID string, userId string, userName string) (bool, error) {

	pipelinePhotoRaw, err := datastore.SearchPipelinePhotos(ctx, nil, 0, 1, strings.ToObject(pipelinePhotoID))
	if err != nil || len(pipelinePhotoRaw) == 0 {
		return false, errs.InvalidId
	}
	isDeleted, err := datastore.DeletePipelinePhoto(ctx, pipelinePhotoID, userId)
	if err != nil {
		return false, err
	}
	pipelineHistoryValue := fmt.Sprintf("pipeline photo deleted id : %s", pipelinePhotoID)
	var pipelineId string = ""
	if pipelinePhotoRaw[0].PipelineId != nil {
		pipelineId = *pipelinePhotoRaw[0].PipelineId
	}
	err = datastore.AddSetPipelineHistory(ctx, pipelineId, "DELETE_PHOTO", pipelineHistoryValue, userName)
	if err != nil {
		log.Debug("AddSetPipelineHistory error %v", err)
	}
	return isDeleted, nil
}

func AllPipelinePhoto(ctx context.Context, pipelineId string, filter *models.FilterInput) (*models.PipelinePhotoResult, error) {
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

	rawPipelineQualityPhotos, err := datastore.SearchPipelinePhotos(ctx, strings.ToObject(pipelineId), *filter.Offset, *filter.Limit, nil)
	if err != nil {
		log.Debug("Failed to retrieve pipelinePhoto: %v", err)
		return nil, err
	}

	count, err := datastore.GetPipelinePhotosCount(ctx, pipelineId)
	if err != nil {
		log.Debug("Failed to retrieve count of pipelinePhoto: %v", err)
		return nil, err
	}

	pipelineQualityPhotos := make([]*models.PipelinePhoto, 0)
	for _, u := range rawPipelineQualityPhotos {
		pipelineQualityPhotos = append(pipelineQualityPhotos, u.ToModels())
	}
	toInt := int(*count)
	return &models.PipelinePhotoResult{
		TotalCount: &toInt,
		Results:    pipelineQualityPhotos,
	}, nil
}
