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

func AddCredit(ctx context.Context, myID string, input models.SaveCreditInput, myName string) (string, error) {

	newCredit := &Credit{
		UserId:          myID,
		PaypalOrderID:   &input.PaypalOrderID,
		PaypalToken:     &input.PaypalToken,
		Amount:          &input.Amount,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Credits.InsertOne(ctx, newCredit)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	strings.TrimObjectChar(lastInsertedIDStr)

	newCreditLedger := &CreditLedger{
		ClientID:        strings.ToObject(myID),
		ClientName:      strings.ToObject(myName),
		PaypalOrderID:   strings.ToObject(input.PaypalOrderID),
		Type:            strings.ToObject("CREDIT"),
		CreatedDateTime: pointers.PrimitiveDateTime(nil),
		Amount:          pointers.Float64(input.Amount),
	}

	res, err = DbCollections.CreditLedgers.InsertOne(ctx, newCreditLedger)
	if err != nil {
		return "", err
	}

	lastInsertedIDStr = fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func AddCreditLedger(ctx context.Context, myID string, input CreditLedger, myName string) (string, error) {

	res, err := DbCollections.CreditLedgers.InsertOne(ctx, input)
	if err != nil {
		return "", err
	}

	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func SaveCreditLedger(ctx context.Context, input models.AddCreditLedgerInput) (string, error) {
	filterById := FilterById(*input.UserID)
	user, err := GetUser(ctx, filterById)
	if err != nil {
		return "", err
	}
	newCreditLedger := &CreditLedger{
		ClientID:        strings.ToObject(*input.UserID),
		ClientName:      strings.ToObject(user.FirstName + " " + user.LastName),
		Type:            strings.ToObject("CREDIT"),
		CreatedDateTime: pointers.PrimitiveDateTime(nil),
		Amount:          pointers.Float64(input.Amount),
	}

	res, err := DbCollections.CreditLedgers.InsertOne(ctx, newCreditLedger)
	if err != nil {
		return "", err
	}

	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func GetCredits(ctx context.Context, userId string) ([]*models.Credit, error) {
	filter := bson.D{{"userId", userId}}

	cur, err := DbCollections.Credits.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query credit: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Credit, 0)
	for cur.Next(ctx) {
		a := &Credit{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode credit entry: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil
}

func GetCreditSum(ctx context.Context, userId string) (*float64, error) {

	filter := bson.M{"clientId": userId}

	group := bson.M{
		"_id":         "$clientId",
		"totalAmount": bson.M{"$sum": "$amount"},
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$group": group},
	}

	cur, err := DbCollections.CreditLedgers.Aggregate(ctx, pipe)
	if err != nil {
		log.Error("Failed to query credit: %v", err)
		return nil, errs.DbError
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		cos := &CreditTotal{}
		if err := cur.Decode(cos); err != nil {
			log.Debug("Failed to decode credit total entry: %v", err)
			return nil, err
		}
		return &cos.TotalAmount, nil
	}

	cnt := float64(0)
	return &cnt, nil

}

func SearchCredits(ctx context.Context, offset, limit int, datefrom, dateto, invoiceNumber *string) ([]*Credit, error) {

	filter := bson.M{}
	layout := "2006-01-02"
	daoFilter2 := bson.M{}

	if datefrom != nil {
		varDateFrom, err := time.Parse(layout, *datefrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
			return nil, err
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		filter["createdDateTime"] = daoFilter2
	}

	if dateto != nil {
		varDateTo, err := time.Parse(layout, *dateto)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
			return nil, err
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		filter["createdDateTime"] = daoFilter2
	}

	if invoiceNumber != nil {
		filter["invoiceNubmer"] = bson.M{
			"$regex":   *invoiceNumber,
			"$options": "i",
		}
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Credits.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query credit:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawCredits := make([]*Credit, 0)
	for cur.Next(ctx) {
		a := &Credit{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode credit entry: %v", err)
			return nil, err
		}
		rawCredits = append(rawCredits, a)
	}

	return rawCredits, nil
}

func GetCreditsCount(ctx context.Context, datefrom, dateto, invoiceNumber *string) (*int64, error) {

	filter := bson.M{}
	layout := "2006-01-02"
	daoFilter2 := bson.M{}

	if datefrom != nil {
		varDateFrom, err := time.Parse(layout, *datefrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
			return nil, err
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		filter["createdDateTime"] = daoFilter2
	}

	if dateto != nil {
		varDateTo, err := time.Parse(layout, *dateto)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
			return nil, err
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		filter["createdDateTime"] = daoFilter2
	}

	if invoiceNumber != nil {
		filter["invoiceNubmer"] = bson.M{
			"$regex":   *invoiceNumber,
			"$options": "i",
		}
	}

	count, err := DbCollections.Credits.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query credit: %v", err)
		return nil, err
	}

	return &count, nil
}

func SearchCreditLedgers(ctx context.Context, offset, limit int, myID string) ([]*CreditLedger, error) {

	filter := bson.M{
		"clientId": myID,
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.CreditLedgers.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query credit ledger:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawCreditLedgers := make([]*CreditLedger, 0)
	for cur.Next(ctx) {
		a := &CreditLedger{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode credit ledger entry: %v", err)
			return nil, err
		}
		rawCreditLedgers = append(rawCreditLedgers, a)
	}

	return rawCreditLedgers, nil
}

func GetCreditLedgersCount(ctx context.Context, myID string) (*int64, error) {

	filter := bson.M{
		"clientId": myID,
	}

	count, err := DbCollections.CreditLedgers.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query credit ledger: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetCreditLedger(ctx context.Context, filter *bson.M) (*CreditLedger, error) {
	var resCreditLedger CreditLedger
	daoFilter := bson.M{}
	if filter != nil {
		daoFilter = *filter
	}
	err := DbCollections.CreditLedgers.FindOne(ctx, daoFilter).Decode(&resCreditLedger)
	if err != nil {
		return nil, err
	}
	return &resCreditLedger, nil
}
