package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SaveReview(ctx context.Context, createdBy string, input models.SaveReviewInput) (string, error) {

	//var fileUrls []*string
	// for _, v := range input.Attachements {
	// 	fileUrl, err := s3Uploader(*v)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	fileUrls = append(fileUrls, fileUrl)
	// }

	fileUrl := strings.ToObject("")
	if input.Attachment != nil {
		var err error
		fileUrl, err = awsS3.S3Uploader(*input.Attachment)
		if err != nil {
			return "", err
		}
	}

	//Todo Get Pipeline
	pipelineFilter := datastore.FilterById(input.PipelineID)
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

	return datastore.SaveReview(ctx, createdBy, input, fileUrl, pipelineRaw[0])
}

func UpdateReview(ctx context.Context, id string, input models.UpdateReviewInput, updatedBy string) (bool, error) {

	fileUrl := strings.ToObject("")
	if input.Attachment != nil {
		var err error
		fileUrl, err = awsS3.S3Uploader(*input.Attachment)
		if err != nil {
			return false, err
		}
	}

	return datastore.UpdateReview(ctx, id, input, fileUrl, updatedBy)
}

func DeleteReview(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteReview(ctx, id)
}

func AllReview(ctx context.Context, filter *models.FilterInput) (*models.ReviewResult, error) {

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

	rawCompanys, err := datastore.SearchReviews(ctx, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve company: %v", err)
		return nil, err
	}

	count, err := datastore.GetReviewsCount(ctx)
	if err != nil {
		log.Debug("Failed to retrieve count of company: %v", err)
		return nil, err
	}

	reviewsModel := make([]*models.Review, 0)
	for _, u := range rawCompanys {
		reviewsModel = append(reviewsModel, u.ToModels())
	}
	toInt := int(*count)
	return &models.ReviewResult{
		TotalCount: &toInt,
		Results:    reviewsModel,
	}, nil
}
