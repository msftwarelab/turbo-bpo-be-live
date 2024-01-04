package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateDefault(ctx context.Context, userId string, input models.DefaultInput) (bool, error) {

	filter := bson.D{{"userId", userId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.ListingType != nil {
		setDoc = append(setDoc, bson.E{"listingType", *input.ListingType})
	}
	if input.AlwayssubmitOrder != nil {
		setDoc = append(setDoc, bson.E{"alwayssubmitOrder", *input.AlwayssubmitOrder})
	}
	if input.AutoCompleteStandbyOrder != nil {
		setDoc = append(setDoc, bson.E{"autoCompleteStandbyOrder", *input.AutoCompleteStandbyOrder})
	}
	if input.InitialSearchGla != nil {
		setDoc = append(setDoc, bson.E{"initialSearchGla", *input.InitialSearchGla})
	}
	if input.InitialSearchAge != nil {
		setDoc = append(setDoc, bson.E{"initialSearchAge", *input.InitialSearchAge})
	}
	if input.InitialSearchProximity != nil {
		setDoc = append(setDoc, bson.E{"initialSearchProximity", *input.InitialSearchProximity})
	}
	if input.SecondSearchGla != nil {
		setDoc = append(setDoc, bson.E{"secondSearchGla", *input.SecondSearchGla})
	}
	if input.SecondSearchAge != nil {
		setDoc = append(setDoc, bson.E{"secondSearchAge", *input.SecondSearchAge})
	}
	if input.SecondSearchProximity != nil {
		setDoc = append(setDoc, bson.E{"secondSearchProximity", *input.SecondSearchProximity})
	}
	if input.SecondSearchSaleDates != nil {
		setDoc = append(setDoc, bson.E{"secondSearchSaleDates", *input.SecondSearchSaleDates})
	}
	if input.ThirdSearchGla != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchGla", *input.ThirdSearchGla})
	}
	if input.ThirdSearchAge != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchAge", *input.ThirdSearchAge})
	}
	if input.ThirdSearchProximity != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchProximity", *input.ThirdSearchProximity})
	}
	if input.ThirdSearchSaleDates != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchSaleDates", *input.ThirdSearchSaleDates})
	}
	if input.ThirdSearchFilterByComplexName != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchFilterByComplexName", *input.ThirdSearchFilterByComplexName})
	}
	if input.ThirdSearchFilterByCity != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchFilterByCity", *input.ThirdSearchFilterByCity})
	}
	if input.ThirdSearchFilterByZip != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchFilterByZip", *input.ThirdSearchFilterByZip})
	}
	if input.ThirdSearchFilterByCountry != nil {
		setDoc = append(setDoc, bson.E{"thirdSearchFilterByCountry", *input.ThirdSearchFilterByCountry})
	}
	if input.UseDefaults != nil {
		setDoc = append(setDoc, bson.E{"useDefaults", *input.UseDefaults})
	}
	if input.UseIformValidations != nil {
		setDoc = append(setDoc, bson.E{"useIformValidations", *input.UseIformValidations})
	}
	if input.SubjectType != nil {
		setDoc = append(setDoc, bson.E{"subjectType", *input.SubjectType})
	}
	if input.StyleDesign != nil {
		setDoc = append(setDoc, bson.E{"styleDesign", *input.StyleDesign})
	}
	if input.ExteriorFinish != nil {
		setDoc = append(setDoc, bson.E{"exteriorFinish", *input.ExteriorFinish})
	}
	if input.Condition != nil {
		setDoc = append(setDoc, bson.E{"condition", *input.Condition})
	}
	if input.Quality != nil {
		setDoc = append(setDoc, bson.E{"quality", *input.Quality})
	}
	if input.View != nil {
		setDoc = append(setDoc, bson.E{"view", *input.View})
	}
	if input.Pool != nil {
		setDoc = append(setDoc, bson.E{"pool", *input.Pool})
	}
	if input.PorchPatioDeck != nil {
		setDoc = append(setDoc, bson.E{"porchPatioDeck", *input.PorchPatioDeck})
	}
	if input.FirePlace != nil {
		setDoc = append(setDoc, bson.E{"firePlace", *input.FirePlace})
	}
	if input.Basement != nil {
		setDoc = append(setDoc, bson.E{"basement", *input.Basement})
	}
	if input.Condo != nil {
		setDoc = append(setDoc, bson.E{"condo", *input.Condo})
	}
	if input.MultiUnit != nil {
		setDoc = append(setDoc, bson.E{"multiUnit", *input.MultiUnit})
	}
	if input.MobileHome != nil {
		setDoc = append(setDoc, bson.E{"mobileHome", *input.MobileHome})
	}
	if input.Sfd != nil {
		setDoc = append(setDoc, bson.E{"sfd", *input.Sfd})
	}
	if input.SfaTownhouse != nil {
		setDoc = append(setDoc, bson.E{"sfaTownhouse", *input.SfaTownhouse})
	}
	if input.Theme != nil {
		setDoc = append(setDoc, bson.E{"theme", *input.Theme})
	}
	if input.Theme != nil {
		setDoc = append(setDoc, bson.E{"theme", *input.Theme})
	}
	if input.IsEnableEmailNotification != nil {
		setDoc = append(setDoc, bson.E{"isEnableEmailNotification", *input.IsEnableEmailNotification})
		userFilter := FilterById(userId)
		updateData := UpdateUserDataIsEnableNotif(*input.IsEnableEmailNotification)
		UpdateUserByDefined(ctx, userFilter, updateData)
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.Defaults.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetDefault(ctx context.Context, userId string) (*Default, error) {
	filter := bson.D{{"userId", userId}}
	log.Debug("@default filter %s", PrettyPrint(filter))
	a := &Default{}
	err := DbCollections.Defaults.FindOne(ctx, filter).Decode(a)
	if err != nil {
		log.Error("Failed to query default: %v", err)
		return nil, errs.DbError
	}

	return a, nil
}

func SaveDefault(ctx context.Context, input *Default) (string, error) {
	res, err := DbCollections.Defaults.InsertOne(ctx, input)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}
