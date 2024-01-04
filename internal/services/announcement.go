package services

import (
	"context"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
)

func AddAnnouncement(ctx context.Context, input models.AnnouncementInput, createdBy string) (string, error) {
	return datastore.SaveAnnouncement(ctx, createdBy, input)
}

func UpdateAnnouncement(ctx context.Context, id string, input models.AnnouncementInput, updatedBy string) (bool, error) {
	return datastore.UpdateAnnouncement(ctx, id, input, updatedBy)
}

func DeleteAnnouncement(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteAnnouncement(ctx, id)
}

func AllAnnouncement(ctx context.Context, filter *models.AnnouncementFilterInput) (*models.AnnouncementResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.AnnouncementFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawAnnouncements, err := datastore.SearchAnnouncements(ctx, *filter.Offset, *filter.Limit, filter.Search, filter.IsActive)
	if err != nil {
		log.Debug("Failed to retrieve announcement: %v", err)
		return nil, err
	}

	count, err := datastore.GetAnnouncementsCount(ctx, filter.Search, filter.IsActive)
	if err != nil {
		log.Debug("Failed to retrieve count of announcement: %v", err)
		return nil, err
	}
	timeToday := time.Now()
	Companies := make([]*models.Announcement, 0)
	for _, u := range rawAnnouncements {
		endtime := pointers.PrimativeToDateTime(*u.EndDate)
		if endtime.Unix() >= timeToday.Unix() {
			Companies = append(Companies, u.ToModels())
		} else {

		}
	}
	toInt := int(*count)
	return &models.AnnouncementResult{
		TotalCount: &toInt,
		Results:    Companies,
	}, nil
}
