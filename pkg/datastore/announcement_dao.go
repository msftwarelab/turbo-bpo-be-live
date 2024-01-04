package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveAnnouncement(ctx context.Context, createdBy string, input models.AnnouncementInput) (string, error) {

	varStartDate, err := time.Parse(time.RFC3339, *input.StartDate)
	if err != nil {
		return "", err
	}
	varEndDate, err := time.Parse(time.RFC3339, *input.EndDate)
	if err != nil {
		return "", err
	}
	newAnnouncement := &Announcement{
		Subject:         input.Subject,
		StartDate:       pointers.PrimitiveDateTime(&varStartDate),
		EndDate:         pointers.PrimitiveDateTime(&varEndDate),
		Message:         input.Message,
		CreatedBy:       createdBy,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Announcements.InsertOne(ctx, newAnnouncement)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateAnnouncement(ctx context.Context, id string, input models.AnnouncementInput, updatedBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.Subject != nil {
		setDoc = append(setDoc, bson.E{"subject", *input.Subject})
	}
	if input.StartDate != nil {
		varStartDate, err := time.Parse(time.RFC3339, *input.StartDate)
		if err != nil {
			return false, err
		}
		setDoc = append(setDoc, bson.E{"startDate", pointers.PrimitiveDateTime(&varStartDate)})
	}

	if input.EndDate != nil {
		varEndDate, err := time.Parse(time.RFC3339, *input.EndDate)
		if err != nil {
			return false, err
		}
		setDoc = append(setDoc, bson.E{"endDate", pointers.PrimitiveDateTime(&varEndDate)})
	}

	if input.Message != nil {
		setDoc = append(setDoc, bson.E{"message", *input.Message})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Announcements.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteAnnouncement(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Announcements.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchAnnouncements(ctx context.Context, offset, limit int, search *string, isActive *bool) ([]*Announcement, error) {

	filter := bson.M{}
	if search != nil {
		filter["name"] = bson.M{
			"$regex":   *search,
			"$options": "i",
		}
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"name": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Announcements.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query announcement:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawAnnouncements := make([]*Announcement, 0)
	for cur.Next(ctx) {
		a := &Announcement{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode announcement entry: %v", err)
			return nil, err
		}
		rawAnnouncements = append(rawAnnouncements, a)
	}

	return rawAnnouncements, nil
}

func GetAnnouncementsCount(ctx context.Context, search *string, isActive *bool) (*int64, error) {

	filter := bson.M{}
	if search != nil {
		filter["$or"] = []bson.M{
			{"message": bson.M{
				"$regex":   *search,
				"$options": "i",
			}},
			{"subject": bson.M{
				"$regex":   *search,
				"$options": "i",
			}},
		}
	}

	if isActive != nil {
		if *isActive {
			//TODO, refactor this area
			// daoGTEFilter := bson.M{}
			// currentDateTime := primitive.DateTime(millis.NowInMillis())
			// daoGTEFilter["$gte"] = currentDateTime
			// filter["startDate"] = daoFilter2

			// daoLTEFilter := bson.M{}
			// daoLTEFilter["$lte"] = currentDateTime

		}
	}
	count, err := DbCollections.Announcements.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query announcement: %v", err)
		return nil, err
	}

	return &count, nil
}
