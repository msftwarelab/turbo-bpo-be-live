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

func SaveCompany(ctx context.Context, userId string, input models.CompanyInput) (string, error) {

	newCompany := &Company{
		Name:            strings.ToObject(input.Name),
		CreatedBy:       userId,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}
	if input.WebSite != nil {
		newCompany.WebSite = *input.WebSite
	}
	if input.IsAdmin != nil {
		newCompany.IsAdmin = *input.IsAdmin
	}
	if input.IsClient != nil {
		newCompany.IsClient = *input.IsClient
	}
	if input.IsPremium != nil {
		newCompany.IsPremium = *input.IsPremium
	}

	res, err := DbCollections.Companies.InsertOne(ctx, newCompany)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateCompany(ctx context.Context, id string, input models.CompanyInput) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	updateCompany := bson.M{
		"name":           input.Name,
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}

	if input.WebSite != nil {
		updateCompany["webSite"] = *input.WebSite
	}
	if input.IsAdmin != nil {
		updateCompany["isAdmin"] = *input.IsAdmin
	}
	if input.IsClient != nil {
		updateCompany["isClient"] = *input.IsClient
	}
	if input.IsPremium != nil {
		updateCompany["isPremium"] = *input.IsPremium
	}

	if len(input.Forms) > 0 {
		forms := []*CompanyForm{}
		for _, v := range input.Forms {
			entry := &CompanyForm{
				Name:  v.Name,
				Style: v.Style,
			}
			forms = append(forms, entry)
		}
		updateCompany["forms"] = forms
	}

	updateDoc := bson.M{
		"$set": updateCompany,
	}
	res, err := DbCollections.Companies.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteCompany(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Companies.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchCompanys(ctx context.Context, userId string, offset, limit int, name *string, isAdmin, isClient, isPremium *bool) ([]*Company, error) {

	filter := bson.M{}

	if name != nil {
		filter["name"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	filterAnd := []bson.M{}

	if isAdmin != nil {
		if *isAdmin {
			filter["isAdmin"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isAdmin": false},
				{"isAdmin": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)

		}
	}

	if isClient != nil {
		if *isClient {
			filter["isClient"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isClient": false},
				{"isClient": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)
		}
	}
	if isPremium != nil {
		if *isPremium {
			filter["isPremium"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isPremium": false},
				{"isPremium": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)
		}
	}

	if len(filterAnd) > 0 {
		filter["$and"] = filterAnd
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"name": 1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Companies.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query company:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawCompanys := make([]*Company, 0)
	for cur.Next(ctx) {
		a := &Company{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode company entry: %v", err)
			return nil, err
		}
		rawCompanys = append(rawCompanys, a)
	}

	return rawCompanys, nil
}

func GetCompanysCount(ctx context.Context, userId string, name *string, isAdmin, isClient, isPremium *bool) (*int64, error) {

	filter := bson.M{}

	if name != nil {
		filter["name"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	filterAnd := []bson.M{}

	if isAdmin != nil {
		if *isAdmin {
			filter["isAdmin"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isAdmin": false},
				{"isAdmin": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)
		}
	}

	if isClient != nil {
		if *isClient {
			filter["isClient"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isClient": false},
				{"isClient": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)
		}
	}
	if isPremium != nil {
		if *isPremium {
			filter["isPremium"] = true
		} else {
			daoOrFilter := bson.M{}
			daoOrFilter["$or"] = []bson.M{
				{"isPremium": false},
				{"isPremium": bson.M{
					"$exists": false,
				}},
			}
			filterAnd = append(filterAnd, daoOrFilter)
		}
	}

	if len(filterAnd) > 0 {
		filter["$and"] = filterAnd
	}

	count, err := DbCollections.Companies.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query company: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetCompany(ctx context.Context, filter *bson.M) (*Company, error) {
	var resCompany Company
	daoFilter := bson.M{}
	if filter != nil {
		daoFilter = *filter
	}
	err := DbCollections.Companies.FindOne(ctx, daoFilter).Decode(&resCompany)
	if err != nil {
		return nil, err
	}
	return &resCompany, nil
}
