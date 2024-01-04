package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveAccount(ctx context.Context, userId string, input models.AccountInput, userName string) (string, error) {
	return datastore.SaveAccount(ctx, userId, input, userName)
}

func UpdateAccount(ctx context.Context, id string, input models.AccountInput, userName string) (bool, error) {
	return datastore.UpdateAccount(ctx, id, input, userName)
}

func DeleteAccount(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteAccount(ctx, id)
}

func AllAccount(ctx context.Context, userId string, filter *models.AccountFilterInput) (*models.AccountResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.AccountFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	if filter.UserID != nil {
		userId = *filter.UserID
	}
	rawAccounts, err := datastore.SearchAccounts(ctx, userId, *filter.Offset, *filter.Limit, filter.Username)
	if err != nil {
		log.Debug("Failed to retrieve account: %v", err)
		return nil, err
	}

	count, err := datastore.GetAccountsCount(ctx, userId, filter.Username)
	if err != nil {
		log.Debug("Failed to retrieve count of account: %v", err)
		return nil, err
	}

	accounts := make([]*models.Account, 0)
	for _, u := range rawAccounts {
		accounts = append(accounts, u.ToModels())
	}
	toInt := int(*count)
	return &models.AccountResult{
		TotalCount: &toInt,
		Results:    accounts,
	}, nil
}
