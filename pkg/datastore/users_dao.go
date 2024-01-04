package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/passwords"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterByEmail(email string) *bson.M {
	return &bson.M{"email": email}
}

func FilterById(id string) *bson.M {
	objId, _ := primitive.ObjectIDFromHex(id)
	return &bson.M{"_id": objId}
}

func FilterByOrderNumber(orderNumber string) *bson.M {
	return &bson.M{"orderNumber": orderNumber}
}

func FilterByOrderNumbers(orderNumbers []string) *bson.M {
	if len(orderNumbers) > 0 {
		return &bson.M{"orderNumber": bson.M{"$in": orderNumbers}}
	}
	return &bson.M{}
}

func FilterByRole(role string) *bson.M {
	filter := bson.M{}
	filter["role"] = role
	return &filter
}

func FilterByIds(ids []string) *bson.M {
	var objectIds []primitive.ObjectID
	for _, v := range ids {
		objId, _ := primitive.ObjectIDFromHex(v)
		objectIds = append(objectIds, objId)
	}
	if len(objectIds) > 0 {
		return &bson.M{"_id": bson.M{"$in": objectIds}}
	}
	return &bson.M{}
}

func GetUsers(ctx context.Context, filter *bson.M) ([]*User, error) {

	if filter == nil {
		filter = &bson.M{}
	}
	cur, err := DbCollections.Users.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query user: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*User, 0)
	for cur.Next(ctx) {
		a := &User{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode user entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func FilterByCollectionIds(ids []string) *bson.M {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		obj, _ := primitive.ObjectIDFromHex(id)
		objIds = append(objIds, obj)
	}
	return &bson.M{"_id": bson.M{"$in": objIds}}
}

func FilterByResetPasswordToken(token string) *bson.M {
	filter := bson.M{}
	filter["resetPasswordToken"] = token
	return &filter
}

func GetUser(ctx context.Context, filter *bson.M) (*User, error) {
	var resUser User
	daoFilter := bson.M{}
	if filter != nil {
		daoFilter = *filter
	}
	err := DbCollections.Users.FindOne(ctx, daoFilter).Decode(&resUser)
	if err != nil {
		return nil, err
	}
	return &resUser, nil
}

func RegisterUser(ctx context.Context, input models.RegisterInput) (string, error) {

	hashPassword, err := passwords.PasswordHashAndSalt([]byte(input.Password))
	if err != nil {
		return "", err
	}

	roleClient := constants.UserRoleClient
	roles := []*string{&roleClient}
	newUser := &User{
		Email:                input.Email,
		Password:             *hashPassword,
		FirstName:            input.FirstName,
		LastName:             input.LastName,
		Company:              &input.Company,
		PhoneNumber:          &input.PhoneNumber,
		Address:              &input.Address,
		City:                 &input.City,
		State:                &input.State,
		Zipcode:              &input.ZipCode,
		Hdyfu:                input.Hdyfu,
		PhoneConsultation:    input.PhoneConsultation,
		Roles:                roles,
		Status:               constants.UserStatusActive,
		OrderTotal:           0,
		AssignmentPercentage: 0,
		PermissionGroupID:    strings.ToObject(constants.DefaultClientPermissionGroupID),
		Theme:                strings.ToObject(constants.DefaultTheme),
		CreatedDateTime:      primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Users.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func SaveUser(ctx context.Context, input models.SaveUserInput) (string, error) {

	hashPassword, err := passwords.PasswordHashAndSalt([]byte(input.Password))
	if err != nil {
		return "", err
	}

	if len(input.Roles) == 0 {
		return "", errs.RoleIsEmpty
	}

	var roles []*string
	for _, v := range input.Roles {
		roles = append(roles, &v)
	}
	newUser := &User{
		Roles:                     roles,
		Email:                     input.Email,
		Password:                  *hashPassword,
		FirstName:                 input.FirstName,
		LastName:                  input.LastName,
		Company:                   &input.Company,
		PhoneNumber:               &input.PhoneNumber,
		Address:                   &input.Address,
		City:                      &input.City,
		State:                     &input.State,
		Zipcode:                   &input.ZipCode,
		Status:                    input.Status,
		OrderTotal:                0,
		AssignmentPercentage:      0,
		PermissionGroupID:         &input.PermissionGroupID,
		Theme:                     strings.ToObject(constants.DefaultTheme),
		IsEnableEmailNotification: input.IsEnableEmailNotification,
		CreatedDateTime:           primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Users.InsertOne(ctx, newUser)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func AddUserOrder(ctx context.Context, userId string, pipelineTotals *int64) error {

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": objId}

	user, err := GetUser(ctx, &filter)
	if err != nil {
		return err
	}

	var assignPercentage float64
	if user.OrderTotal == 0 {
		assignPercentage = 0
	} else {
		assignPercentage = float64(int(*pipelineTotals) / user.OrderTotal)
	}
	updateDoc := bson.M{
		"$set": bson.M{"assignmentPercentage": assignPercentage},
		"$inc": bson.M{"orderTotal": 1},
	}

	res, err := DbCollections.Users.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	if res.ModifiedCount < 1 {
		return errs.NoRecordUpdate
	}
	return nil
}

func UpdateProfile(ctx context.Context, userId string, input models.ProfileInput, profilePicUrl *string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"email":                 input.Email,
		"firstName":             input.FirstName,
		"lastName":              input.LastName,
		"company":               input.Company,
		"phoneNumber":           input.PhoneNumber,
		"address":               input.Address,
		"city":                  input.City,
		"state":                 input.State,
		"imABroker":             input.ImABroker,
		"broker":                input.Broker,
		"brokerLicense":         input.BrokerLicense,
		"licenseDate":           input.LicenseDate,
		"licenseExpirationDate": input.LicenseExpirationDate,
		"brokerage":             input.Brokerage,
		"agent":                 input.Agent,
		"agentLicense":          input.AgentLicense,
		"yearOfExperience":      input.YearOfExperience,
		"lastUpdateTime":        pointers.PrimitiveDateTime(nil),
		"disclaimer":            input.Disclaimer,
	}
	if profilePicUrl != nil {
		update["profilePicture"] = profilePicUrl

	}
	if input.Theme != nil {
		update["theme"] = *input.Theme
	}
	if input.Password != nil {
		hashPassword, err := passwords.PasswordHashAndSalt([]byte(*input.Password))
		if err != nil {
			return false, err
		}
		update["password"] = hashPassword
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Users.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func UpdateUser(ctx context.Context, userId string, input models.UpdateUserInput, updatedBy string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}
	if input.Status != nil {
		update["status"] = *input.Status
	}

	if input.Email != nil {
		update["email"] = *input.Email
	}
	if input.FirstName != nil {
		update["firstName"] = *input.FirstName
	}
	if input.LastName != nil {
		update["lastName"] = *input.LastName
	}
	if input.Password != nil {
		hashPassword, err := passwords.PasswordHashAndSalt([]byte(*input.Password))
		if err != nil {
			return false, err
		}
		update["password"] = hashPassword
	}
	if input.PhoneNumber != nil {
		update["phoneNumber"] = *input.PhoneNumber
	}
	if input.Broker != nil {
		update["broker"] = *input.Broker
	}
	if input.BrokerLicense != nil {
		update["brokerLicense"] = *input.BrokerLicense
	}
	if input.LicenseDate != nil {
		update["licenseDate"] = *input.LicenseDate
	}
	if input.LicenseExpirationDate != nil {
		update["licenseExpirationDate"] = *input.LicenseExpirationDate
	}
	if input.Brokerage != nil {
		update["brokerage"] = *input.Brokerage
	}
	if input.Agent != nil {
		update["agent"] = *input.Agent
	}
	if input.AgentLicense != nil {
		update["agentLicense"] = *input.AgentLicense
	}
	if input.YearOfExperience != nil {
		update["yearOfExperience"] = *input.YearOfExperience
	}
	if input.Address != nil {
		update["address"] = *input.Address
	}
	if input.City != nil {
		update["city"] = *input.City
	}
	if input.State != nil {
		update["state"] = *input.State
	}
	if input.ZipCode != nil {
		update["zipCode"] = *input.ZipCode
	}
	if input.Company != nil {
		update["company"] = *input.Company
	}
	if len(input.CompanyList) > 0 {
		update["companyList"] = input.CompanyList
	}

	if input.PermissionGroupID != nil {
		update["permissionGroupId"] = input.PermissionGroupID
	}
	if len(input.Roles) > 0 {
		update["role"] = input.Roles
	}

	if input.Theme != nil {
		update["theme"] = *input.Theme
	}
	if input.IsEnableEmailNotification != nil {
		update["isEnableEmailNotification"] = *input.IsEnableEmailNotification
	}
	//Price module
	if input.PriceModule != nil {
		if input.PriceModule.Credits != nil {
			update["priceModule.credits"] = input.PriceModule.Credits
		}
		if input.PriceModule.Orderinterior != nil {
			update["priceModule.orderinterior"] = input.PriceModule.Orderinterior
		}
		if input.PriceModule.Orderexterior != nil {
			update["priceModule.orderexterior"] = *input.PriceModule.Orderexterior
		}
		if input.PriceModule.OrderdataEntry != nil {
			update["priceModule.orderdataEntry"] = *input.PriceModule.OrderdataEntry
		}
		if input.PriceModule.Orderrush != nil {
			update["priceModule.orderrush"] = *input.PriceModule.Orderrush
		}
		if input.PriceModule.OrdersuperRush != nil {
			update["priceModule.ordersuperRush"] = *input.PriceModule.OrdersuperRush
		}
		if input.PriceModule.OrderconditionReport != nil {
			update["priceModule.orderconditionReport"] = *input.PriceModule.OrderconditionReport
		}
		if input.PriceModule.OrderrentalAddendum != nil {
			update["priceModule.orderrentalAddendum"] = *input.PriceModule.OrderrentalAddendum
		}
		if input.PriceModule.PhotoExterior != nil {
			update["priceModule.photoExterior"] = *input.PriceModule.PhotoExterior
		}
		if input.PriceModule.PhotoInteriorVacantLb != nil {
			update["priceModule.photoInteriorVacantLB"] = *input.PriceModule.PhotoInteriorVacantLb
		}
		if input.PriceModule.PhotoInteriorAppointment != nil {
			update["priceModule.photoInteriorAppointment"] = *input.PriceModule.PhotoInteriorAppointment
		}
	}

	//end price module

	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Users.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func UpdateUserOrderCount(ctx context.Context, userId string, assignActive, assignRush, assignStandby *int) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{"_id", objId}}

	update := bson.M{
		"assignDate": pointers.PrimitiveDateTime(nil),
	}

	updateInc := bson.M{}
	if assignActive != nil {
		updateInc["assignActive"] = *assignActive
	}
	if assignRush != nil {
		updateInc["assignRush"] = *assignRush
	}
	if assignStandby != nil {
		updateInc["assignStandby"] = *assignStandby
	}

	updateDoc := bson.M{
		"$set": update,
		"$inc": updateInc,
	}
	_, err := DbCollections.Users.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}

	return true, nil
}

func SearchUsers(ctx context.Context, offset, limit int, name *string, status *string, roles []*string) ([]*User, error) {

	filter := bson.M{}
	if status != nil {
		filter["status"] = *status

	}
	if roles != nil {
		if len(roles) > 0 {
			filter["role"] = bson.M{"$in": roles}
		}
	}

	if name != nil {
		filter["fullName"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	addFields := bson.M{
		"fullName": bson.M{"$concat": []string{"$firstName", " ", "$lastName"}},
	}

	pipe := []bson.M{
		{"$addFields": addFields},
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Users.Aggregate(ctx, pipe)

	if err != nil {
		log.Debug("Failed to query user:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawUsers := make([]*User, 0)
	for cur.Next(ctx) {
		a := &User{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode user entry: %v", err)
			return nil, err
		}

		rawUsers = append(rawUsers, a)
	}

	return rawUsers, nil
}

func GetUsersCount(ctx context.Context, name *string, status *string, roles []*string) (*int64, error) {

	filter := bson.M{}
	if status != nil {
		filter["status"] = *status

	}
	if roles != nil {
		if len(roles) > 0 {
			filter["role"] = bson.M{"$in": roles}
		}
	}

	if name != nil {
		filter["fullName"] = bson.M{
			"$regex":   *name,
			"$options": "i",
		}
	}

	addFields := bson.M{
		"fullName": bson.M{"$concat": []string{"$firstName", " ", "$lastName"}},
	}

	group := bson.M{
		"_id":   nil,
		"count": bson.M{"$sum": 1},
	}

	pipe := []bson.M{
		{"$addFields": addFields},
		{"$match": filter},
		{"$group": group},
		{"$sort": bson.M{"createdDateTime": -1}},
	}
	cursor, err := DbCollections.Users.Aggregate(ctx, pipe)
	var docCount int32
	if err != nil {
		log.Debug("Failed to query user: %v", err)
		return nil, err
		// If the API call was a success
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {

			// declare a result BSON object
			var result bson.M
			err := cursor.Decode(&result)

			// If there is a cursor.Decode error
			if err != nil {
				log.Debug("Failed to decode user entry: %v", err)
				return nil, err
				// If there are no cursor.Decode errors
			} else {
				docCount = result["count"].(int32)
			}
		}
	}
	return pointers.Int64(int64(docCount)), nil
}

func UpdateUserByDatastore(input User) bson.M {
	update := bson.M{
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}

	if input.Company != nil {
		update["company"] = *input.Company
	}
	if input.Password != "" {
		update["password"] = input.Password
	}
	if input.PhoneNumber != nil {
		update["phoneNumber"] = *input.PhoneNumber
	}
	if input.Address != nil {
		update["address"] = *input.Address
	}
	if input.City != nil {
		update["city"] = *input.City
	}
	if input.State != nil {
		update["state"] = *input.State
	}
	if input.Zipcode != nil {
		update["zipcode"] = *input.Zipcode
	}
	if input.Title != nil {
		update["title"] = *input.Title
	}
	if input.About != nil {
		update["about"] = *input.About
	}
	if input.Hdyfu != nil {
		update["hdyfu"] = *input.Hdyfu
	}
	if input.PhoneConsultation != nil {
		update["phoneConsultation"] = *input.PhoneConsultation
	}
	if input.ImABroker != nil {
		update["imABroker"] = *input.ImABroker
	}
	if input.Broker != nil {
		update["broker"] = *input.Broker
	}
	if input.BrokerLicense != nil {
		update["brokerLicense"] = *input.BrokerLicense
	}
	if input.Agent != nil {
		update["agent"] = *input.Agent
	}
	// if input.AgentLicense != nil {
	// 	update["agentLicense"] = *input.AgentLicense
	// }
	// if input.LicenseDate != nil {
	// 	update["licenseDate"] = *input.LicenseDate
	// }
	if input.LicenseExpirationDate != nil {
		update["licenseExpirationDate"] = *input.LicenseExpirationDate
	}
	if input.Brokerage != nil {
		update["brokerage"] = *input.Brokerage
	}
	// if input.YearOfExperience != nil {
	// 	update["yearOfExperience"] = *input.YearOfExperience
	// }
	if input.ProfilePicture != nil {
		update["profilePicture"] = *input.ProfilePicture
	}
	if input.LastUpdateTime != nil {
		update["lastUpdateTime"] = *input.LastUpdateTime
	}
	if input.PermissionGroupID != nil {
		update["permissionGroupID"] = *input.PermissionGroupID
	}
	if input.ResetPasswordToken != nil {
		update["resetPasswordToken"] = *input.ResetPasswordToken
	}
	if input.ResetPasswordTokenExpirationDateTime != nil {
		update["resetPasswordTokenExpirationDateTime"] = *input.ResetPasswordTokenExpirationDateTime
	}
	if input.IsResetPasswordTokenUsed != nil {
		update["isResetPasswordTokenUsed"] = *input.IsResetPasswordTokenUsed
	}

	return update
}

func UpdateUserDataIsEnableNotif(isEnableEmailNotification bool) bson.M {
	update := bson.M{
		"IsEnableEmailNotification": isEnableEmailNotification,
	}
	return update
}

func FilterUsers(ctx context.Context) []string {
	filter := bson.M{}

	roles := []string{constants.UserRoleAdmin, constants.UserRoleContractor, constants.UserRoleQualityControl}
	filter["role"] = bson.M{"$in": roles}
	cur, err := DbCollections.Users.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query user: %v", err)
		return nil
	}

	defer cur.Close(ctx)

	list := make([]*User, 0)
	for cur.Next(ctx) {
		a := &User{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode user entry: %v", err)
			return nil
		}
		list = append(list, a)

	}
	var listID []string
	for _, v := range list {

		listID = append(listID, v.ID.Hex())
	}
	return listID

}

func UpdateUserByDefined(ctx context.Context, filter *bson.M, update bson.M) (bool, error) {

	updateDoc := bson.M{
		"$set": update,
	}
	_, err := DbCollections.Users.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	return true, nil
}
