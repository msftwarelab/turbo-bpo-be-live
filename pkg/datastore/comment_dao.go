package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	constants "github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddComment(ctx context.Context, userId string, input models.CommentInput) (string, error) {

	newUser := &Comment{
		UserID:          userId,
		Category:        input.Category,
		Label:           input.Label,
		Value:           input.Value,
		Section:         strings.ToObject(input.Section),
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Comments.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateComment(ctx context.Context, id string, value string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"value":          value,
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Comments.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteComment(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Comments.DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
func DeleteComments(ctx context.Context, userId string) error {
	filter := bson.D{{"userId", userId}}

	_, err := DbCollections.Comments.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func GetComments(ctx context.Context, userId string) ([]*models.Comment, error) {
	filter := bson.D{{"userId", userId}}

	cur, err := DbCollections.Comments.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query adjustment: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Comment, 0)
	for cur.Next(ctx) {
		a := &Comment{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode adjustment entry: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil
}

func GetComment(ctx context.Context, userId string) (*Comment, error) {
	filter := bson.D{{"userId", userId}}
	a := &Comment{}
	err := DbCollections.Comments.FindOne(ctx, filter).Decode(a)
	if err != nil {
		log.Error("Failed to query comment: %v", err)
		return nil, errs.DbError
	}
	return a, nil
}
func SaveMultiComment(ctx context.Context, inputs []*Comment) (bool, error) {

	docs := make([]interface{}, 0)
	for _, comment := range inputs {
		doc := bson.D{
			{"userId", comment.UserID},
			{"category", comment.Category},
			{"label", comment.Label},
			{"section", constants.CommentSectionDefaultValue},
			{"value", comment.Value},
			{"createdDateTime", primitive.DateTime(millis.NowInMillis())},
		}
		docs = append(docs, doc)
	}
	_, err := DbCollections.Comments.InsertMany(ctx, docs)
	if err != nil {
		return false, err
	}
	return true, nil
}
