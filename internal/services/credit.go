package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
)

func SaveCredit(ctx context.Context, userId string, input models.SaveCreditInput, myName string) (string, error) {
	return datastore.AddCredit(ctx, userId, input, myName)
	//Todo update userAccount credit
}

func SaveCreditLedger(ctx context.Context, input models.AddCreditLedgerInput) (string, error) {
	return datastore.SaveCreditLedger(ctx, input)
}

func AllCredit(ctx context.Context, userId string) ([]*models.Credit, error) {
	return datastore.GetCredits(ctx, userId)
}

func AllCreditLedger(ctx context.Context, filter *models.FilterInput, myUserID string, filterUserId *string) (*models.CreditLedgerResult, error) {
	if filterUserId != nil {
		myUserID = *filterUserId
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

	rawCreditLedgers, err := datastore.SearchCreditLedgers(ctx, *filter.Offset, *filter.Limit, myUserID)
	if err != nil {
		log.Debug("Failed to retrieve credit ledger: %v", err)
		return nil, err
	}

	count, err := datastore.GetCreditLedgersCount(ctx, myUserID)
	if err != nil {
		log.Debug("Failed to retrieve count of credit ledger: %v", err)
		return nil, err
	}

	creditLedgers := make([]*models.CreditLedger, 0)
	for _, u := range rawCreditLedgers {
		creditLedgers = append(creditLedgers, u.ToModels())
	}
	toInt := int(*count)
	return &models.CreditLedgerResult{
		TotalCount: &toInt,
		Results:    creditLedgers,
	}, nil
}
