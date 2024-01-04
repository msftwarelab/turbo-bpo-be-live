package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func UpdateDefault(ctx context.Context, userId string, roles []*string, input models.DefaultInput) (bool, error) {

	if contains(roles, constants.UserRoleAdmin) {
		userId = constants.SuperAdminId
	}
	return datastore.UpdateDefault(ctx, userId, input)
}

func Default(ctx context.Context, userId string, filterUserId *string) (*models.Default, error) {
	if filterUserId != nil {
		userId = *filterUserId
	}
	defaultRaw, err := datastore.GetDefault(ctx, userId)
	return defaultRaw.ToModels(), err
}

func contains(slice []*string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[*s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
