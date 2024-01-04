package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveLoginLog(ctx context.Context, username, ipAddress string) (string, error) {

	newAccount := &LoginLog{
		Username:  username,
		IPAddress: ipAddress,
		Datetime:  primitive.DateTime(millis.NowInMillis()),
	}
	res, err := DbCollections.LoginLogs.InsertOne(ctx, newAccount)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func SearchLoginLogs(ctx context.Context, offset, limit int, userId, dateFrom, dateTo *string) ([]*LoginLog, error) {

	filter := bson.M{}

	layout := "2006-01-02"
	daoFilter2 := bson.M{}
	dateFilter := bson.M{}
	if dateFrom != nil {
		varDateFrom, err := time.Parse(layout, *dateFrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		dateFilter["datetime"] = daoFilter2
		filter["datetime"] = daoFilter2
	}

	if dateTo != nil {
		varDateTo, err := time.Parse(layout, *dateTo)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		filter["datetime"] = daoFilter2
	}
	if userId != nil {
		filterById := FilterById(*userId)
		user, err := GetUser(ctx, filterById)
		if err != nil {
			return nil, err
		}

		idFilter := bson.M{"username": user.Email}
		andMatch := []bson.M{
			idFilter,
			dateFilter,
		}
		filter["$and"] = andMatch
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"datetime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.LoginLogs.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query userLoginLogs:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawLoginLogs := make([]*LoginLog, 0)
	for cur.Next(ctx) {
		a := &LoginLog{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode loginLog entry: %v", err)
			return nil, err
		}
		rawLoginLogs = append(rawLoginLogs, a)
	}

	return rawLoginLogs, nil
}

func GetLoginLogsCount(ctx context.Context, userId, dateFrom, dateTo *string) (*int64, error) {

	filter := bson.M{}

	layout := "2006-01-02"
	daoFilter2 := bson.M{}
	dateFilter := bson.M{}
	if dateFrom != nil {
		varDateFrom, err := time.Parse(layout, *dateFrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		filter["datetime"] = daoFilter2
		dateFilter["datetime"] = daoFilter2
	}

	if dateTo != nil {
		varDateTo, err := time.Parse(layout, *dateTo)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		filter["datetime"] = daoFilter2
		dateFilter["datetime"] = daoFilter2
	}

	if userId != nil {
		filterById := FilterById(*userId)
		user, err := GetUser(ctx, filterById)
		if err != nil {
			return nil, err
		}

		idFilter := bson.M{"username": user.Email}
		andMatch := []bson.M{
			idFilter,
			dateFilter,
		}
		filter["$and"] = andMatch
	}

	count, err := DbCollections.LoginLogs.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query loginLog: %v", err)
		return nil, err
	}

	return &count, nil
}
