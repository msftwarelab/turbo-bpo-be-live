package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func UploadProfileDoc(ctx context.Context, userId string, input models.ProfileDocInput) (string, error) {

	fileUrl, err := awsS3.S3Uploader(input.Doc)
	if err != nil {
		return "", err
	}
	return datastore.SaveProfileDoc(ctx, userId, input, input.Doc.Filename, *fileUrl)
}

func DeleteProfileDoc(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteProfileDoc(ctx, id)
}

func AllProfileDoc(ctx context.Context, userId string, filterUserId *string, filter *models.FilterInput) (*models.ProfileDocResult, error) {

	if filterUserId != nil {
		userId = *filterUserId
	}

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

	rawProfileDocs, err := datastore.SearchProfileDocs(ctx, userId, *filter.Offset, *filter.Limit)
	if err != nil {
		log.Debug("Failed to retrieve profileDoc: %v", err)
		return nil, err
	}

	count, err := datastore.GetProfileDocsCount(ctx, userId)
	if err != nil {
		log.Debug("Failed to retrieve count of profileDoc: %v", err)
		return nil, err
	}

	profileDocs := make([]*models.Doc, 0)
	for _, u := range rawProfileDocs {
		profileDocs = append(profileDocs, u.ToModels())
	}
	toInt := int(*count)
	return &models.ProfileDocResult{
		TotalCount: &toInt,
		Results:    profileDocs,
	}, nil

}
