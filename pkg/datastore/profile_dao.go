package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveProfileDoc(ctx context.Context, userId string, input models.ProfileDocInput, fileName string, url string) (string, error) {

	newProfileDoc := &ProfileDoc{
		Type:            input.Type,
		FileName:        fileName,
		Url:             url,
		UserID:          userId,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.ProfileDoc.InsertOne(ctx, newProfileDoc)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func DeleteProfileDoc(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.ProfileDoc.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetProfileDocs(ctx context.Context, userId string) ([]*models.Doc, error) {
	filter := bson.D{{"userID", userId}}

	cur, err := DbCollections.ProfileDoc.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query user: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Doc, 0)
	for cur.Next(ctx) {
		a := &ProfileDoc{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode profileDocs entry: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil
}

func SearchProfileDocs(ctx context.Context, userId string, offset, limit int) ([]*ProfileDoc, error) {

	//skip := (offset - 1) * limit
	filter := bson.D{{"userID", userId}}

	pipe := []bson.M{
		{"$match": filter},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.ProfileDoc.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query userAccounts:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawProfileDocs := make([]*ProfileDoc, 0)
	for cur.Next(ctx) {
		a := &ProfileDoc{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode userAccounts entry: %v", err)
			return nil, err
		}
		rawProfileDocs = append(rawProfileDocs, a)
	}

	return rawProfileDocs, nil
}

func GetProfileDocsCount(ctx context.Context, userId string) (*int64, error) {

	filter := bson.M{"userID": userId}
	count, err := DbCollections.ProfileDoc.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query profileDocs: %v", err)
		return nil, err
	}

	return &count, nil
}
