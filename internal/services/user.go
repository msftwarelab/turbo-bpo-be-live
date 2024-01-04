package services

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/passwords"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/token"
	"go.mongodb.org/mongo-driver/bson"
)

const adminUserId = "5dc12cf84f26a8f5c2501d92"

func RegisterUser(ctx context.Context, input models.RegisterInput) (string, error) {

	filterByEmail := datastore.FilterByEmail(input.Email)
	isUserExsist, err := datastore.GetUsers(ctx, filterByEmail)
	if err != nil {
		return "", err
	}
	if len(isUserExsist) > 0 {
		return "", errs.EmailAreadyExist
	}
	userId, err := datastore.RegisterUser(ctx, input)
	if err != nil {
		return "", err
	}
	err = newAccountInit(ctx, userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func PriceModule(ctx context.Context, userID string) (*models.PriceModule, error) {
	filter := datastore.FilterById(userID)
	user, err := datastore.GetUser(ctx, filter)
	if err != nil || user == nil {
		return nil, errs.InvalidUserId
	}

	priceModule := &models.PriceModule{
		Orderinterior:            user.PriceModule.Orderinterior,
		Orderexterior:            user.PriceModule.Orderexterior,
		OrderdataEntry:           user.PriceModule.OrderdataEntry,
		Orderrush:                user.PriceModule.Orderrush,
		OrdersuperRush:           user.PriceModule.OrdersuperRush,
		OrderconditionReport:     user.PriceModule.OrderconditionReport,
		OrderrentalAddendum:      user.PriceModule.OrderrentalAddendum,
		PhotoExterior:            user.PriceModule.PhotoExterior,
		PhotoInteriorVacantLb:    user.PriceModule.PhotoInteriorVacantLb,
		PhotoInteriorAppointment: user.PriceModule.PhotoInteriorAppointment,
	}
	return priceModule, nil

}

func ForgetPassword(ctx context.Context, email string) (bool, error) {

	filterByEmail := datastore.FilterByEmail(email)
	isUserExsist, err := datastore.GetUsers(ctx, filterByEmail)
	if err != nil {
		return false, err
	}
	if len(isUserExsist) == 0 {
		return false, errs.EmailDoesNotExist
	}

	token := uuid.New().String()
	updateFilter := datastore.FilterById(isUserExsist[0].ID.Hex())
	updateData := datastore.User{
		Password:                             "",
		ResetPasswordToken:                   strings.ToObject(token),
		IsResetPasswordTokenUsed:             pointers.Bool(false),
		ResetPasswordTokenExpirationDateTime: pointers.PrimitiveDateTimeAddHr(1),
	}
	updatePayload := datastore.UpdateUserByDatastore(updateData)
	datastore.UpdateUserByDefined(ctx, updateFilter, updatePayload)
	datastore.SendEmailAccountForgetPassword(ctx, isUserExsist[0].Email, token)

	return true, nil
}

func ResetPassword(ctx context.Context, resetPasswordToken, newPassword string) (bool, error) {

	filter := datastore.FilterByResetPasswordToken(resetPasswordToken)
	user, err := datastore.GetUser(ctx, filter)
	if err != nil || user == nil {
		return false, errs.InvalidToken
	}

	if pointers.ObjectTOBool(user.IsResetPasswordTokenUsed) {
		return false, errs.TokenAlreadyUsed
	}

	if user.ResetPasswordTokenExpirationDateTime != nil {
		duedatetime := pointers.PrimativeToDateTime(*user.ResetPasswordTokenExpirationDateTime)
		if duedatetime.Before(time.Now().Local()) {
			return false, errs.TokenExpired
		}
	}

	hashNewPassword, err := passwords.PasswordHashAndSalt([]byte(newPassword))
	if err != nil {
		return false, err
	}

	updateFilter := datastore.FilterById(user.ID.Hex())
	updateData := datastore.User{
		IsResetPasswordTokenUsed: pointers.Bool(true),
		Password:                 strings.ObjectTOString(hashNewPassword),
	}
	updatePayload := datastore.UpdateUserByDatastore(updateData)
	isupdate, err := datastore.UpdateUserByDefined(ctx, updateFilter, updatePayload)
	if isupdate {
		datastore.SendEmailAccountChangePasswordSuccess(ctx, user.Email)
	}
	return isupdate, err

}

func SaveUser(ctx context.Context, input models.SaveUserInput) (string, error) {

	filterByEmail := datastore.FilterByEmail(input.Email)
	isUserExsist, err := datastore.GetUsers(ctx, filterByEmail)
	if err != nil {
		return "", err
	}
	if len(isUserExsist) > 0 {
		return "", errs.EmailAreadyExist
	}
	userId, err := datastore.SaveUser(ctx, input)
	if err != nil {
		return "", err
	}
	err = newAccountInit(ctx, userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func ManualNewAccountInit(ctx context.Context, userId string) error {

	err := initAdjustments(ctx, userId)
	if err != nil {
		return err
	}

	err = initComments(ctx, userId)
	if err != nil {
		return err
	}

	defaultRaw, err := datastore.GetDefault(ctx, adminUserId)
	if err != nil {
		return err
	}
	(defaultRaw).UserId = userId
	(defaultRaw).ID = nil

	// _, err = datastore.SaveDefault(ctx, defaultRaw)
	// if err != nil {
	// 	return err
	// }
	return nil
}
func newAccountInit(ctx context.Context, userId string) error {

	err := initAdjustments(ctx, userId)
	if err != nil {
		return err
	}

	err = initComments(ctx, userId)
	if err != nil {
		return err
	}

	defaultRaw, err := datastore.GetDefault(ctx, adminUserId)
	if err != nil {
		return err
	}
	(defaultRaw).UserId = userId
	(defaultRaw).ID = nil

	_, err = datastore.SaveDefault(ctx, defaultRaw)
	if err != nil {
		return err
	}
	return nil
}

func initComments(ctx context.Context, userId string) error {
	comments, err := datastore.GetComments(ctx, adminUserId)
	if err != nil {
		return err
	}

	inputComments := make([]*datastore.Comment, 0)
	for _, comment := range comments {
		r := &datastore.Comment{
			UserID:   userId,
			Category: comment.Category,
			Label:    comment.Label,
			Value:    comment.Value,
		}
		inputComments = append(inputComments, r)
	}
	_, err = datastore.SaveMultiComment(ctx, inputComments)
	if err != nil {
		return err
	}
	return nil
}

func initAdjustments(ctx context.Context, userId string) error {
	adjustments, err := datastore.GetAdjustments(ctx, adminUserId)
	if err != nil {
		return err
	}
	inputAdjustments := make([]*datastore.Adjustment, 0)
	for _, adjustment := range adjustments {
		r := &datastore.Adjustment{
			UserID:   userId,
			Category: adjustment.Category,
			Order:    adjustment.Order,
			Label:    adjustment.Label,
			From:     adjustment.From,
			To:       adjustment.To,
			Value:    adjustment.Value,
		}
		inputAdjustments = append(inputAdjustments, r)
	}
	_, err = datastore.SaveMultiAdjustment(ctx, inputAdjustments)
	if err != nil {
		return err
	}
	return nil
}

func AllUser(ctx context.Context, filter *models.UserFilterInput) (*models.UserResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.UserFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawUsers, err := datastore.SearchUsers(ctx, *filter.Offset, *filter.Limit, filter.Name, filter.Status, filter.Roles)
	if err != nil {
		log.Debug("Failed to retrieve user: %v", err)
		return nil, err
	}

	count, err := datastore.GetUsersCount(ctx, filter.Name, filter.Status, filter.Roles)
	if err != nil {
		log.Debug("Failed to retrieve count of user: %v", err)
		return nil, err
	}

	var userIds []string
	users := make([]*models.User, 0)
	for _, u := range rawUsers {
		users = append(users, u.ToModels())
		userIds = append(userIds, u.ID.Hex())

	}
	if len(filter.Roles) > 0 && filter.OrderMonth != nil && filter.OrderYear != nil {
		isRolehasClient := contains(filter.Roles, constants.UserRoleClient)
		if isRolehasClient && len(userIds) > 0 {
			clientOrders, err := datastore.GetPipelineCountOrderPerMonthAndYearFilterByUserIdsGroupByUserId(ctx, *filter.OrderMonth, *filter.OrderYear, userIds)
			if err != nil {
				log.Debug("Failed tretrieve client order: %v", err)
				return nil, err
			}
			orderTotal := 0
			for _, v := range clientOrders {
				if v != nil {
					orderTotal += *v
				}
			}
			for i, _ := range users {
				if clientOrders[users[i].ID] != nil {
					users[i].OrderTotal = clientOrders[users[i].ID]
					var perc float64
					perc = float64(*users[i].OrderTotal) / float64(orderTotal)
					users[i].AssignmentPercentage = pointers.Float64(perc)
				}
			}

		}
	}

	for i, _ := range users {
		filterById := datastore.FilterById(users[i].ID)
		var profModelMap map[string]interface{}
		data, _ := json.Marshal(users[i])

		json.Unmarshal(data, &profModelMap)

		userCreds := profModelMap["priceModule"].(map[string]interface{})["credits"]
		totalAmount, err := datastore.GetCreditSum(ctx, users[i].ID)
		if err != nil {
			return nil, err
		}
		if userCreds == nil {
			if totalAmount != nil {
				update := bson.M{
					"lastUpdateTime": pointers.PrimitiveDateTime(nil),
				}
				update["priceModule.credits"] = totalAmount
				updateDoc := bson.M{
					"$set": update,
				}
				_, err := datastore.DbCollections.Users.UpdateOne(ctx, filterById, updateDoc)
				if err != nil {
					return nil, err
				}
			}

		} else {
			*users[i].PriceModule.Credits = *totalAmount
		}

	}

	toInt := int(*count)
	return &models.UserResult{
		TotalCount: &toInt,
		Results:    users,
	}, nil
}

func GetUser(ctx context.Context, userId string) (*models.User, error) {
	filterById := datastore.FilterById(userId)
	user, err := datastore.GetUser(ctx, filterById)
	if err != nil {
		return nil, err
	}

	totalAmount, err := datastore.GetCreditSum(ctx, userId)
	if err != nil {
		return nil, err
	}
	permissionGroups, err := datastore.SearchPermissionGroups(ctx, 0, 1, nil, user.PermissionGroupID)
	if err != nil {
		return nil, err
	}
	profileModel := user.ToModels()
	profileModel.Credit = totalAmount
	profileModel.PermissionList = permissionGroups[0].Permissions
	return profileModel, nil
}

//Todo, refactor to pkg helper
func findIP(input string) string {
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

	regEx := regexp.MustCompile(regexPattern)
	return regEx.FindString(input)
}

func Login(ctx context.Context, email string, password string, ipAddress string) (*models.Token, error) {
	filterByEmail := datastore.FilterByEmail(email)
	user, err := datastore.GetUser(ctx, filterByEmail)
	if user == nil {
		return nil, errs.EmailDoesNotExist
	}
	if err != nil {
		return nil, err
	}

	isCorrect := passwords.ComparePasswords(user.Password, []byte(password))
	isMasterPasswordCorrect := compareMasterPassword(password)
	isValidPAsswords := hasTrue(isCorrect, isMasterPasswordCorrect)
	if !isValidPAsswords {
		return nil, errs.InvalidPassword
	}
	fullName := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	jwtToken, err := token.New(email, user.ID.Hex(), fullName, user.Roles)
	if user == nil {
		return nil, errs.DecodeError
	}

	//logging ip address if role is not client
	if !contains(user.Roles, constants.UserRoleClient) {
		_, _ = datastore.SaveLoginLog(ctx, email, findIP(ipAddress))
	}
	return &models.Token{Token: jwtToken}, nil
}

func compareMasterPassword(typedPassword string) bool {
	masterPassword := config.AppConfig.GetString("masterPassword")
	if typedPassword == masterPassword {
		return true
	}
	return false
}

func hasTrue(vars ...bool) bool {
	for _, v := range vars {
		if v {
			return true
		}
	}
	return false
}

func UpdateProfile(ctx context.Context, userId string, input models.ProfileInput) (bool, error) {

	profilePicUrl := new(string)
	profilePicUrl = nil
	if input.ProfilePicture != nil {
		profilePic, err := awsS3.S3Uploader(*input.ProfilePicture)
		if err != nil {
			return false, err

		}
		profilePicUrl = profilePic
	}

	return datastore.UpdateProfile(ctx, userId, input, profilePicUrl)
}

func UpdateUser(ctx context.Context, userId string, input models.UpdateUserInput, updatedBy string) (bool, error) {

	return datastore.UpdateUser(ctx, userId, input, updatedBy)
}
