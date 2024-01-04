package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveCompany(ctx context.Context, userId string, input models.CompanyInput) (string, error) {
	return datastore.SaveCompany(ctx, userId, input)
}

func UpdateCompany(ctx context.Context, id string, input models.CompanyInput) (bool, error) {
	return datastore.UpdateCompany(ctx, id, input)
}

func DeleteCompany(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteCompany(ctx, id)
}

func AllCompany(ctx context.Context, userId string, filter *models.CompanyFilterInput) (*models.CompanyResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.CompanyFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawCompanys, err := datastore.SearchCompanys(ctx, userId, *filter.Offset, *filter.Limit, filter.Name, filter.IsAdmin, filter.IsClient, filter.IsPremium)
	if err != nil {
		log.Debug("Failed to retrieve company: %v", err)
		return nil, err
	}

	count, err := datastore.GetCompanysCount(ctx, userId, filter.Name, filter.IsAdmin, filter.IsClient, filter.IsPremium)
	if err != nil {
		log.Debug("Failed to retrieve count of company: %v", err)
		return nil, err
	}

	Companies := make([]*models.Company, 0)
	for _, u := range rawCompanys {
		Companies = append(Companies, u.ToModels())
	}
	toInt := int(*count)
	return &models.CompanyResult{
		TotalCount: &toInt,
		Results:    Companies,
	}, nil
}

func Company(ctx context.Context, id string) (*models.Company, error) {

	filterByID := datastore.FilterById(id)
	company, err := datastore.GetCompany(ctx, filterByID)
	if err != nil {
		return nil, err
	}
	return company.ToModels(), nil
}
