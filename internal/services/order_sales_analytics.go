package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func AllSalesAnalytics(ctx context.Context, filter models.SalesAnalyticsFilterInput) ([]*models.SalesAnalytics, error) {
	return datastore.GetPipelineCompleteOrderFilterByMonthYear(ctx, filter)

}

func AllOrderAnalytics(ctx context.Context, filter models.OrderAnalyticsFilterInput) ([]*models.OrderAnalytics, error) {

	rawreports, err := datastore.GetPipelineSumByPaymentStatusPerMonth(ctx, filter)

	var userIds []string
	for _, v := range rawreports {
		if v.Client != nil {
			userIds = append(userIds, *v.Client)
		}
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

	for i, v := range rawreports {
		if clientsMap[*v.Client] != nil {
			rawreports[i].Client = strings.ToObject(clientsMap[*v.Client].FullName())
		}
	}
	return rawreports, err
}
