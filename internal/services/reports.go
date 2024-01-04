package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func AllBalance(ctx context.Context, filter *models.BalanceFilterInput) (*models.BalanceResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.BalanceFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	roleClient := []*string{strings.ToObject(constants.UserRoleClient)}

	rawUsers, err := datastore.SearchUsers(ctx, *filter.Offset, *filter.Limit, filter.ClientName, nil, roleClient)
	if err != nil {
		log.Debug("Failed to retrieve balance: %v", err)
		return nil, err
	}

	var userIds []string
	for _, v := range rawUsers {
		userIds = append(userIds, v.ID.Hex())
	}
	orderBalance := make(map[string]*datastore.ClientBalance)
	if len(userIds) > 0 {
		orderBalance, err = datastore.GetPipelineOrderByUserIdSumPaidUnpaid(ctx, userIds)
		if err != nil {
			log.Debug("Failed to retrieve orderblance: %v", err)
			return nil, err
		}
	}

	count, err := datastore.GetUsersCount(ctx, filter.ClientName, nil, roleClient)
	if err != nil {
		log.Debug("Failed to retrieve count of balance: %v", err)
		return nil, err
	}

	balances := make([]*models.Balance, 0)
	for _, user := range rawUsers {

		balancesModel := &models.Balance{
			Client: strings.ToObject(user.FullName()),
			Other:  pointers.Float64(0.0),
		}

		if orderBalance[user.ID.Hex()] != nil {
			balancesModel.Total = pointers.Float64(orderBalance[user.ID.Hex()].Total)
			balancesModel.Unpaid = pointers.Float64(orderBalance[user.ID.Hex()].Unpaid)
			balancesModel.PaidAmount = pointers.Float64(orderBalance[user.ID.Hex()].Paid)
		}
		balances = append(balances, balancesModel)
	}
	toInt := int(*count)
	return &models.BalanceResult{
		TotalCount: &toInt,
		Results:    balances,
	}, nil
}

func AllCheckout(ctx context.Context, filter *models.CheckoutFilterInput) (*models.CheckoutResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.CheckoutFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawUsers, err := datastore.SearchUsers(ctx, *filter.Offset, *filter.Limit, nil, nil, nil)
	if err != nil {
		log.Debug("Failed to retrieve balance: %v", err)
		return nil, err
	}

	var userIds []string
	for _, v := range rawUsers {
		userIds = append(userIds, v.ID.Hex())
	}

	count, err := datastore.GetUsersCount(ctx, nil, nil, nil)
	if err != nil {
		log.Debug("Failed to retrieve count of balance: %v", err)
		return nil, err
	}

	balances := make([]*models.Checkout, 0)
	for _, user := range rawUsers {

		balancesModel := &models.Checkout{
			ClientName: strings.ToObject(user.FullName()),
			Invoice:    strings.ToObject("123213-dummy"),
			Total:      pointers.Float64(123),
			Status:     strings.ToObject("COMPLETE"),
			Date:       strings.ToObject("2019-09-31"),
		}
		balances = append(balances, balancesModel)
	}
	toInt := int(*count)
	return &models.CheckoutResult{
		TotalCount: &toInt,
		Results:    balances,
	}, nil
}

func AllCredits(ctx context.Context, filter *models.CreditsFilterInput) (*models.CreditsResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.CreditsFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawCredits, err := datastore.SearchCredits(ctx, *filter.Offset, *filter.Limit, filter.DateFrom, filter.DateTo, nil)
	if err != nil {
		log.Debug("Failed to retrieve credits: %v", err)
		return nil, err
	}

	count, err := datastore.GetCreditsCount(ctx, filter.DateFrom, filter.DateTo, nil)
	if err != nil {
		log.Debug("Failed to retrieve count of credits: %v", err)
		return nil, err
	}

	var userIds []string
	for _, v := range rawCredits {
		userIds = append(userIds, v.UserId)
	}
	clientFilter := datastore.FilterByIds(userIds)
	clientRaws, err := datastore.GetUsers(ctx, clientFilter)
	if err != nil {
		log.Debug("Failed to retrieve clients record: %v", err)
		return nil, err
	}
	// convert to map
	clientsMap := make(map[string]*datastore.User)
	for _, v := range clientRaws {
		clientsMap[v.ID.Hex()] = v
	}
	credits := make([]*models.Credits, 0)
	for _, u := range rawCredits {
		credit := &models.Credits{
			Invoice:   strings.ToObject("123431"),
			Credits:   u.Amount,
			ExpiresAt: strings.ToObject(datastore.TimeConversion(&u.CreatedDateTime)),
			Date:      strings.ToObject(datastore.TimeConversion(&u.CreatedDateTime)),
		}
		if clientsMap[u.UserId] != nil {
			credit.ClientName = strings.ToObject(clientsMap[u.UserId].FullName())
		}
		credits = append(credits, credit)
	}
	toInt := int(*count)
	return &models.CreditsResult{
		TotalCount: &toInt,
		Results:    credits,
	}, nil
}

// func AllOrderSubmit(ctx context.Context, year int) ([]*models.OrderSubmit, error) {
// 	contractorUsers, err := datastore.GetUsers(ctx, datastore.FilterByRole(constants.UserRoleContractor))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var contractorUsersIds []string
// 	for _, v := range contractorUsers {
// 		contractorUsersIds = append(contractorUsersIds, v.ID.Hex())
// 	}

// 	// //Todo, get pipeline where contractor id equal to contactorsId
// 	pipelinesRaw, err := datastore.GetPipelineCountOrderPerMonthAndYearFilterByUserIdsAndyearGroupByUserId(ctx, year, contractorUsersIds)
// 	if err != nil {
// 		return nil, err
// 	}

// 	orderSubmits := make([]*models.OrderSubmit, 0)
// 	for _, v := range contractorUsers {

// 		orderSubmit := &models.OrderSubmit{
// 			CoordinatorName: strings.ToObject(v.FullName()),
// 		}
// 		if pipelinesRaw[v.ID.Hex()] != nil {
// 			orderSubmit.Year = pipelinesRaw[v.ID.Hex()].ID.Year
// 			orderSubmit.Month = pipelinesRaw[v.ID.Hex()].ID.Month
// 			orderSubmit.Count = pipelinesRaw[v.ID.Hex()].Count
// 		}
// 		orderSubmits = append(orderSubmits, orderSubmit)
// 	}

// 	return orderSubmits, nil
// }

func AllOrderSubmit(ctx context.Context, year int) ([]*models.OrderSubmit, error) {

	mapQcRaws, err := datastore.GetOrderSubmitted(ctx, year)
	if err != nil {
		return nil, err
	}
	orderSubmits := make([]*models.OrderSubmit, 0)
	for _, v := range mapQcRaws {

		orderSubmit := &models.OrderSubmit{
			CoordinatorName: strings.ToObject(v.ID.Coordinator),
			Month:           v.ID.Month,
			Year:            v.ID.Year,
			Count:           v.Count,
		}
		orderSubmits = append(orderSubmits, orderSubmit)
	}

	return orderSubmits, nil
}

func AllQcCompleted(ctx context.Context, year int) ([]*models.QcCompleted, error) {

	mapQcRaws, err := datastore.GetCompletedQcCountOrderPerMonthAndYearFilterByUserIdsAndyearGroupByUserId(ctx, year)
	if err != nil {
		return nil, err
	}
	completedQcs := make([]*models.QcCompleted, 0)
	for _, v := range mapQcRaws {

		completedQc := &models.QcCompleted{
			Month:              v.ID.Month,
			Year:               v.ID.Year,
			Normal:             v.Count,
			FullRec:            pointers.Int(0),
			Dd:                 pointers.Int(0),
			Total:              v.Count,
			QualityControlName: strings.ToObject(v.ID.Assignee),
		}
		completedQcs = append(completedQcs, completedQc)
	}

	return completedQcs, nil
}
