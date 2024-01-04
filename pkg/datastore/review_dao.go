package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveReview(ctx context.Context, createdBy string, input models.SaveReviewInput, fileUrls *string, pipelineData *Pipeline) (string, error) {

	var fileName *string

	if input.Attachment != nil {
		fileName = strings.ToObject(input.Attachment.Filename)
	}

	newReview := &Review{
		ReviewDescription: input.ReviewDescription,
		ReviewBy:          strings.ToObject(createdBy),
		CreatedDateTime:   primitive.DateTime(millis.NowInMillis()),
		Attachement:       fileUrls,
		FileName:          fileName,
		PipelineId:        pipelineData.ID.Hex(),
		OrderNumber:       pipelineData.OrderNumber,
		Address:           pipelineData.Address,
		AssignedTo:        pipelineData.Assign,
	}

	res, err := DbCollections.Reviews.InsertOne(ctx, newReview)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateReview(ctx context.Context, id string, input models.UpdateReviewInput, url *string, updatedBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.ReviewDescription != nil {
		setDoc = append(setDoc, bson.E{"reviewDescription", *input.ReviewDescription})
	}
	if input.Attachment != nil {
		setDoc = append(setDoc, bson.E{"attachement", url})
		setDoc = append(setDoc, bson.E{"fileName", input.Attachment.Filename})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Reviews.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteReview(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Reviews.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchReviews(ctx context.Context, offset, limit int) ([]*Review, error) {

	filter := bson.M{}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Reviews.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query reviews:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawReviews := make([]*Review, 0)
	for cur.Next(ctx) {
		a := &Review{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode reviews entry: %v", err)
			return nil, err
		}
		rawReviews = append(rawReviews, a)
	}

	return rawReviews, nil
}

func GetReviewsCount(ctx context.Context) (*int64, error) {

	filter := bson.M{}
	count, err := DbCollections.Reviews.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query review: %v", err)
		return nil, err
	}

	return &count, nil
}
