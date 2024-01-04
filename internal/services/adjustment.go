package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
)

func UpdateAdjustment(ctx context.Context, id string, value float64) (bool, error) {
	return datastore.UpdateAdjustment(ctx, id, value)
}

func SetAdjustmentDefault(ctx context.Context, userId string) (bool, error) {
	err := datastore.DeleteAdjustments(ctx, userId)
	if err != nil {
		return false, err
	}
	err = initAdjustments(ctx, userId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func AllAdjustment(ctx context.Context, myID string, userID *string, roles []*string) ([]*models.Adjustment, error) {

	// if role == constants.UserRoleAdmin {
	// 	userId = constants.SuperAdminId
	// }
	if contains(roles, constants.UserRoleAdmin) {
		myID = constants.SuperAdminId
	}
	if userID != nil {
		myID = *userID
	}
	return datastore.GetAdjustments(ctx, myID)
}
