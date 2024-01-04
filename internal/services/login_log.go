package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func AllLoginLog(ctx context.Context, filter *models.LoginLogFilterInput) (*models.LoginLogResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.LoginLogFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawLoginLogs, err := datastore.SearchLoginLogs(ctx, *filter.Offset, *filter.Limit, filter.ID, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve loginLog: %v", err)
		return nil, err
	}

	count, err := datastore.GetLoginLogsCount(ctx, filter.ID, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve count of loginLog: %v", err)
		return nil, err
	}

	loginLogs := make([]*models.LoginLog, 0)
	for _, u := range rawLoginLogs {
		loginLogs = append(loginLogs, u.ToModels())
	}
	toInt := int(*count)
	return &models.LoginLogResult{
		TotalCount: &toInt,
		Results:    loginLogs,
	}, nil
}
