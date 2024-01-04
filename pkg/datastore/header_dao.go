package datastore

import (
	"context"
	"fmt"

	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveHeader(ctx context.Context, name string, parentId *string) (string, error) {

	newHeader := &Header{
		Name:            name,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	if parentId != nil {
		newHeader.ParentId = *parentId
	}

	res, err := DbCollections.Headers.InsertOne(ctx, newHeader)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateHeader(ctx context.Context, id string, name string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"name", name},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Headers.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, errs.NoRecordUpdate
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteHeader(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Headers.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchHeaders(ctx context.Context, offset, limit int, Name *string, parentId *string) ([]*Header, error) {

	filter := bson.M{}

	if Name != nil {
		filter["name"] = bson.M{
			"$regex":   *Name,
			"$options": "i",
		}
	}

	if parentId != nil {
		filter["parentId"] = *parentId
	} else {
		filter["parentId"] = nil
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Headers.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query header:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawHeaders := make([]*Header, 0)
	for cur.Next(ctx) {
		a := &Header{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode header entry: %v", err)
			return nil, err
		}
		rawHeaders = append(rawHeaders, a)
	}

	return rawHeaders, nil
}

func GetHeadersCount(ctx context.Context, name *string, parentId *string) (*int64, error) {

	filter := bson.M{}
	if name != nil {
		filter["name"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	if parentId != nil {
		filter["parentId"] = *parentId
	} else {
		filter["parentId"] = nil
	}

	count, err := DbCollections.Headers.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query header: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetHeadersByCode(ctx context.Context, codes []*string, isParent *bool) ([]*Header, error) {

	filter := bson.M{
		"name": bson.M{"$in": codes},
		//	"IsParent": false,
	}

	if isParent != nil {
		if *isParent {
			filter["parentId"] = bson.M{
				"$exists": false,
			}
		} else {
			filter["parentId"] = bson.M{
				"$exists": true,
			}
		}
	}
	cur, err := DbCollections.Headers.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query header: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*Header, 0)
	for cur.Next(ctx) {
		a := &Header{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode header entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func GetHeadersByID(ctx context.Context, ids []string) ([]*Header, error) {

	filter := FilterByIds(ids)
	cur, err := DbCollections.Headers.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query header: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*Header, 0)
	for cur.Next(ctx) {
		a := &Header{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode header entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
