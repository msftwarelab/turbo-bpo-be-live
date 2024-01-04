package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveSession(ctx context.Context, userId string, invoiceDate string, createdBy string) (string, error) {
	return datastore.SaveSession(ctx, userId, invoiceDate, createdBy)
}

func StopSession(ctx context.Context, id string, updatedBy string) (bool, error) {
	return datastore.StopSession(ctx, id, updatedBy)
}

func UpdateSession(ctx context.Context, id string, input models.UpdateSessionInput, updatedBy string) (bool, error) {
	return datastore.UpdateSession(ctx, id, input, updatedBy)
}

func AllSession(ctx context.Context, userId string, filter *models.SessionFilterInput) (*models.SessionResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.SessionFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawSessions, err := datastore.SearchSessions(ctx, userId, *filter.Offset, *filter.Limit, filter.UserID, nil, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve session: %v", err)
		return nil, err
	}

	count, err := datastore.GetSessionsCount(ctx, userId, filter.UserID, nil, filter.DateFrom, filter.DateTo)
	if err != nil {
		log.Debug("Failed to retrieve count of session: %v", err)
		return nil, err
	}

	Companies := make([]*models.Session, 0)
	for _, u := range rawSessions {
		Companies = append(Companies, u.ToModels())
	}
	toInt := int(*count)
	return &models.SessionResult{
		TotalCount: &toInt,
		Results:    Companies,
	}, nil
}
func ContinueSession(ctx context.Context, userID string) (bool, error) {
	return datastore.ContinueSession(ctx, userID)
}
