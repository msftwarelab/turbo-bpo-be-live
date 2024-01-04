package datastore

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateAdjustment(ctx context.Context, id string, value float64) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"value":          value,
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Adjustments.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}
func SaveMultiAdjustment(ctx context.Context, inputs []*Adjustment) (bool, error) {

	docs := make([]interface{}, 0)
	for _, adjustment := range inputs {
		doc := bson.D{
			{"userId", adjustment.UserID},
			{"category", adjustment.Category},
			{"order", adjustment.Order},
			{"label", adjustment.Label},
			{"from", adjustment.From},
			{"to", adjustment.To},
			{"value", adjustment.Value},
			{"createdDateTime", primitive.DateTime(millis.NowInMillis())},
		}
		docs = append(docs, doc)
	}
	_, err := DbCollections.Adjustments.InsertMany(ctx, docs)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetAdjustment(ctx context.Context, userId string) (*models.Adjustment, error) {
	filter := bson.D{{"userId", userId}}
	a := &Adjustment{}
	err := DbCollections.Adjustments.FindOne(ctx, filter).Decode(a)
	if err != nil {
		log.Error("Failed to query adjustment: %v", err)
		return nil, errs.DbError
	}
	return a.ToModels(), nil
}

func GetAdjustments(ctx context.Context, userId string) ([]*models.Adjustment, error) {
	filter := bson.D{{"userId", userId}}

	cur, err := DbCollections.Adjustments.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query adjustment: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Adjustment, 0)
	for cur.Next(ctx) {
		a := &Adjustment{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode adjustment entry: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil
}

func DeleteAdjustments(ctx context.Context, userId string) error {
	filter := bson.D{{"userId", userId}}

	_, err := DbCollections.Adjustments.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
