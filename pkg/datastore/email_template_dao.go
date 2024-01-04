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

func SaveEmailTemplate(ctx context.Context, input models.SaveEmailTemplateInput) (string, error) {

	newEmailTemplate := &EmailTemplate{
		Type:            input.Type,
		Template:        input.Template,
		Subject:         input.Subject,
		Message:         input.Message,
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.EmailTemplates.InsertOne(ctx, newEmailTemplate)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateEmailTemplate(ctx context.Context, id string, input models.UpdateEmailTemplateInput) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"type", input.Type},
		{"template", input.Template},
		{"subject", input.Subject},
		{"message", input.Message},
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.EmailTemplates.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteEmailTemplate(ctx context.Context, id string) (bool, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	_, err := DbCollections.EmailTemplates.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func SearchEmailTemplates(ctx context.Context, offset, limit int, subject *string) ([]*EmailTemplate, error) {

	filter := bson.M{}
	if subject != nil {
		filter["subject"] = bson.M{
			"$regex":   *subject,
			"$options": "i",
		}
	}

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.EmailTemplates.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query email templates:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawEmailTemplates := make([]*EmailTemplate, 0)
	for cur.Next(ctx) {
		a := &EmailTemplate{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode email templates entry: %v", err)
			return nil, err
		}
		rawEmailTemplates = append(rawEmailTemplates, a)
	}

	return rawEmailTemplates, nil
}

func GetEmailTemplatesCount(ctx context.Context, subject *string) (*int64, error) {

	filter := bson.M{}
	if subject != nil {
		filter["username"] = bson.M{
			"$regex":   *subject,
			"$options": "i",
		}
	}
	count, err := DbCollections.EmailTemplates.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query email template: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetEmailTemplateByTemplateCode(ctx context.Context, templateCode string) (*EmailTemplate, error) {
	var resEmailTemplate EmailTemplate
	daoFilter := bson.M{
		"template": templateCode,
	}

	err := DbCollections.EmailTemplates.FindOne(ctx, daoFilter).Decode(&resEmailTemplate)
	if err != nil {
		return nil, err
	}
	return &resEmailTemplate, nil
}
