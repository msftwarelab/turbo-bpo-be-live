package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveInstruction(ctx context.Context, input models.SaveInstructionInput, url *string) (string, error) {

	newInstruction := &Instruction{
		Tag:             input.Tag,
		Comment:         input.Comment,
		FileName:        input.File.Filename,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}
	if input.Client != nil {
		newInstruction.Client = *input.Client
	}
	if input.ClientID != nil {
		newInstruction.ClientId = *input.ClientID
	}
	if input.Company != nil {
		newInstruction.Company = *input.Company
	}
	if input.CompanyID != nil {
		newInstruction.CompanyId = *input.CompanyID
	}
	if url != nil {
		newInstruction.Url = *url
	}

	res, err := DbCollections.Instructions.InsertOne(ctx, newInstruction)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func DeleteInstruction(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.Instructions.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
func SearchInstructions(ctx context.Context, offset, limit int, tag *string, clientID, companyID *string) ([]*Instruction, error) {

	filter := bson.M{}
	if tag != nil {
		filter["tag"] = *tag
	}

	filterOR := []bson.M{}
	if clientID != nil {
		filterOR = append(filterOR, bson.M{"clientId": *clientID})
	}
	if companyID != nil {
		filterOR = append(filterOR, bson.M{"companyId": *companyID})
	}
	if clientID != nil || companyID != nil {
		filter["$or"] = filterOR
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Instructions.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query instruction:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawInstructions := make([]*Instruction, 0)
	for cur.Next(ctx) {
		a := &Instruction{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode instruction entry: %v", err)
			return nil, err
		}
		rawInstructions = append(rawInstructions, a)
	}

	return rawInstructions, nil
}

func GetInstructionsCount(ctx context.Context, tag *string) (*int64, error) {

	filter := bson.M{}
	if tag != nil {
		filter["tag"] = *tag
	}

	count, err := DbCollections.Instructions.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query instruction: %v", err)
		return nil, err
	}

	return &count, nil
}
