package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveAccount(ctx context.Context, userId string, input models.AccountInput, userName string) (string, error) {

	newAccount := &Account{
		UserID:          userId,
		RecordType:      input.RecordType,
		Company:         input.Company,
		WebSite:         input.WebSite,
		Username:        input.Username,
		Password:        input.Password,
		Question1:       input.Question1,
		Answer1:         input.Answer1,
		Question2:       input.Question2,
		Answer2:         input.Answer2,
		Question3:       input.Question3,
		Answer3:         input.Answer3,
		Others:          input.Others,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}
	logRaw, err := makeLog(constants.AppActionSave, accountLogBuilder(input), userName)
	if err != nil {
		return "", err
	}
	(newAccount).Logs = logRaw

	res, err := DbCollections.Accounts.InsertOne(ctx, newAccount)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func accountLogBuilder(a models.AccountInput) interface{} {
	logRaw := make(map[string]string)
	logRaw["recordType"] = a.RecordType
	logRaw["company"] = a.Company
	logRaw["website"] = a.WebSite
	logRaw["username"] = a.Username
	logRaw["password"] = a.Password
	return logRaw
}

func UpdateAccount(ctx context.Context, id string, input models.AccountInput, userName string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	update := &Account{
		RecordType:     input.RecordType,
		Company:        input.Company,
		WebSite:        input.WebSite,
		Username:       input.Username,
		Password:       input.Password,
		Question1:      input.Question1,
		Answer1:        input.Answer1,
		Question2:      input.Question2,
		Answer2:        input.Answer2,
		Question3:      input.Question3,
		Answer3:        input.Answer3,
		Others:         input.Others,
		LastUpdateTime: pointers.PrimitiveDateTime(nil),
	}

	logRaw, err := makeLogRaw(constants.AppActionUpdate, accountLogBuilder(input), userName)
	if err != nil {
		return false, err
	}

	updateDoc := bson.M{
		"$set": update,
		"$addToSet": bson.M{
			"logs": logRaw,
		},
	}
	res, err := DbCollections.Accounts.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteAccount(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Accounts.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetAccounts(ctx context.Context, userId string) ([]*models.Account, error) {
	filter := bson.D{{"userID", userId}}

	cur, err := DbCollections.Accounts.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query account: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Account, 0)
	for cur.Next(ctx) {
		a := &Account{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode account entry: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil
}

func SearchAccounts(ctx context.Context, userId string, offset, limit int, userName *string) ([]*Account, error) {

	// filter := bson.D{{"userID", userId}}
	// filter := bson.D{{"userID", userId}}
	filter := bson.M{}
	filter["userID"] = userId
	if userName != nil {
		filter["username"] = bson.M{
			"$regex":   *userName,
			"$options": "i",
		}
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Accounts.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query userAccounts:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawAccounts := make([]*Account, 0)
	for cur.Next(ctx) {
		a := &Account{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode account entry: %v", err)
			return nil, err
		}
		rawAccounts = append(rawAccounts, a)
	}

	return rawAccounts, nil
}

func GetAccountsCount(ctx context.Context, userId string, userName *string) (*int64, error) {

	filter := bson.M{}
	filter["userID"] = userId
	if userName != nil {
		filter["username"] = bson.M{
			"$regex":   *userName,
			"$options": "i",
		}
	}
	count, err := DbCollections.Accounts.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query account: %v", err)
		return nil, err
	}

	return &count, nil
}
