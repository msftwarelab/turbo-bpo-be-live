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

func SavePermissionGroup(ctx context.Context, createdBy string, input models.PermissionGroupInput) (string, error) {

	newPermissionGroup := &PermissionGroup{
		Name:            input.Name,
		Permissions:     input.Permissions,
		CreatedBy:       createdBy,
		Status:          "ACTIVE",
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.PermissionGroups.InsertOne(ctx, newPermissionGroup)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdatePermissionGroup(ctx context.Context, id string, input models.PermissionGroupInput) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}
	if input.Name != nil {
		setDoc = append(setDoc, bson.E{"name", *input.Name})
	}
	if len(input.Permissions) > 0 {
		setDoc = append(setDoc, bson.E{"permissions", input.Permissions})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.PermissionGroups.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeletePermissionGroup(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"status", "DELETED"},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.PermissionGroups.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchPermissionGroups(ctx context.Context, offset, limit int, name *string, permissionGroupId *string) ([]*PermissionGroup, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	if name != nil {
		filter["name"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}
	if permissionGroupId != nil {
		objId, _ := primitive.ObjectIDFromHex(*permissionGroupId)
		filter["_id"] = objId
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"name": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.PermissionGroups.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query permissionGroup:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPermissionGroups := make([]*PermissionGroup, 0)
	for cur.Next(ctx) {
		a := &PermissionGroup{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode permissionGroup entry: %v", err)
			return nil, err
		}
		rawPermissionGroups = append(rawPermissionGroups, a)
	}

	return rawPermissionGroups, nil
}

func GetPermissionGroupsCount(ctx context.Context, name *string) (*int64, error) {

	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}

	if name != nil {
		filter["name"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	count, err := DbCollections.PermissionGroups.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query permissionGroup: %v", err)
		return nil, err
	}

	return &count, nil
}
