package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveEmailTemplate(ctx context.Context, input models.SaveEmailTemplateInput) (string, error) {
	return datastore.SaveEmailTemplate(ctx, input)
}

func UpdateEmailTemplate(ctx context.Context, id string, input models.UpdateEmailTemplateInput) (bool, error) {
	return datastore.UpdateEmailTemplate(ctx, id, input)
}

func DeleteEmailTemplate(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteEmailTemplate(ctx, id)
}

func AllEmailTemplate(ctx context.Context, filter *models.EmailTemplateFilterInput) (*models.EmailTemplateResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.EmailTemplateFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawEmailTemplates, err := datastore.SearchEmailTemplates(ctx, *filter.Offset, *filter.Limit, filter.Subject)
	if err != nil {
		log.Debug("Failed to retrieve email templates: %v", err)
		return nil, err
	}

	count, err := datastore.GetEmailTemplatesCount(ctx, filter.Subject)
	if err != nil {
		log.Debug("Failed to retrieve count of email templates: %v", err)
		return nil, err
	}

	emailTemplates := make([]*models.EmailTemplate, 0)
	for _, u := range rawEmailTemplates {
		emailTemplates = append(emailTemplates, u.ToModels())
	}
	toInt := int(*count)
	return &models.EmailTemplateResult{
		TotalCount: &toInt,
		Results:    emailTemplates,
	}, nil
}
