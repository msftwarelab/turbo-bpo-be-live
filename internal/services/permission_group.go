package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SavePermissionGroup(ctx context.Context, input models.PermissionGroupInput, cratedBy string) (string, error) {
	return datastore.SavePermissionGroup(ctx, cratedBy, input)
}

func UpdatePermissionGroup(ctx context.Context, id string, input models.PermissionGroupInput, updatedBy string) (bool, error) {
	return datastore.UpdatePermissionGroup(ctx, id, input)
}

func DeletePermissionGroup(ctx context.Context, id string, deletedBy string) (bool, error) {
	return datastore.DeletePermissionGroup(ctx, id)
}

func AllPermissionGroup(ctx context.Context, filter *models.PermissionGroupFilterInput) (*models.PermissionGroupResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.PermissionGroupFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawPermissionGroups, err := datastore.SearchPermissionGroups(ctx, *filter.Offset, *filter.Limit, filter.Name, nil)
	if err != nil {
		log.Debug("Failed to retrieve permissionGroup: %v", err)
		return nil, err
	}

	count, err := datastore.GetPermissionGroupsCount(ctx, filter.Name)
	if err != nil {
		log.Debug("Failed to retrieve count of permissionGroup: %v", err)
		return nil, err
	}

	Companies := make([]*models.PermissionGroup, 0)
	for _, u := range rawPermissionGroups {
		Companies = append(Companies, u.ToModels())
	}
	toInt := int(*count)
	return &models.PermissionGroupResult{
		TotalCount: &toInt,
		Results:    Companies,
	}, nil
}
