package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/email"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	stringUtls "github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const orderTypeInteriorDueHr = 12
const orderTypeExteriorDueHr = 24
const orderConditionReportDueHr = 24
const orderTypeDataEntryDueHr = 24
const orderIsRush = 6
const orderIsSuperRush = 2

func CalculateDueDate(orderType string, turnAroundtime float64) float64 {

	switch strings.ToUpper(orderType) {
	case constants.PipelineOrderTypeInterior:
		return orderTypeInteriorDueHr
	case constants.PipelineOrderTypeExterior:
		return orderTypeExteriorDueHr
	case constants.PipelineOrderTypeConditionReport:
		return orderConditionReportDueHr
	case constants.PipelineOrderTypeDataEntry:
		return orderTypeDataEntryDueHr
	}
	return 24
}

func getPipelineAuthor(ctx context.Context, userId string) (*User, error) {
	authorUserFilter := FilterById(userId)
	authorUserRaw, err := GetUser(ctx, authorUserFilter)
	if err != nil {
		return nil, err
	}
	if authorUserRaw == nil {
		return nil, errs.InvalidCreatorId
	}
	return authorUserRaw, nil
}

//Todo, Refactor to add  this on util pkg
func contains(slice []*string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[*s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func AddPipeline(ctx context.Context, userId string, input models.PipelineInput) (string, error) {

	authorUserRaw, err := getPipelineAuthor(ctx, userId)
	if err != nil {
		return "", err
	}

	if contains(authorUserRaw.Roles, constants.UserRoleAdmin) || contains(authorUserRaw.Roles, constants.UserRoleContractor) {
		//Validate MAO/ maximum 8 order per day
		count, err := GetCreatedOrderPerday(ctx, &userId)
		if err != nil {
			return "", err
		}
		log.Debug("the created count %s", count)
		if count != nil {
			if *count > 8 {
				return "", errs.MaxOrderLimit
			}
		}
	}

	newPipeline := &Pipeline{
		OrderNumber:        &input.OrderNumber,
		UserId:             &userId,
		UserName:           stringUtls.ToObject(authorUserRaw.FullName()),
		Status:             input.Status,
		Address:            &input.Address,
		ZipCode:            input.ZipCode,
		Country:            &input.Country,
		County:             input.County,
		Location:           &input.Location,
		Company:            &input.Company,
		CompanyID:          input.CompanyID,
		Type:               &input.Type,
		OrderType:          &input.OrderType,
		Objective:          input.Objective,
		Assign:             &input.Assign,
		AssignId:           input.AssignID,
		Mls:                input.Mls,
		IsRushOrder:        input.IsRushOrder,
		IsSuperRush:        input.IsSuperRush,
		IsInspection:       input.IsInspection,
		IsInitialBpo:       input.IsInitialBpo,
		OrderFee:           input.OrderFee,
		TotalFee:           input.TotalFee,
		RatingOverAll:      pointers.Int(5),
		RatingTimeliness:   pointers.Int(5),
		RatingQuality:      pointers.Int(5),
		IsSyncedToTurboBpo: input.IsSyncedToTurboBpo,
		OtherCompany:       input.OtherCompany,
		CreatedDateTime:    primitive.DateTime(millis.NowInMillis()),
	}

	if input.AssignID != nil {
		newPipeline.AssignDateTime = pointers.PrimitiveDateTime(nil)
	}

	if input.Status == nil {
		newPipeline.Status = stringUtls.ToObject(constants.PipelineStatusActive)
		newPipeline.ActivationDateTime = pointers.PrimitiveDateTime(nil)
	}

	pipelineState, err := GetPipelineStateDataStore(ctx)
	if err != nil || pipelineState == nil {
		log.Error("%s", "something wrong on pipeline state")
		return "", errs.PipelineStateError
	}
	//dueDatetime read turnAroundtime
	turnAroundtime, err := pipelineState.TurnAoundTime(ctx)
	if err != nil || turnAroundtime == nil {
		return "", errs.InvalidTurnAroundTime
	}

	dueDateTime := pointers.PrimitiveDateTimeAddHr(CalculateDueDate(input.OrderType, float64(*turnAroundtime)))
	//pipelineState, err := GetPipelineStateDataStore(ctx)
	if err != nil {
		return "", err
	}
	if input.IsRushOrder != nil {
		if *input.IsRushOrder {
			if pipelineState.OPRush != nil {
				dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsRush)
			}
		}
	}
	if input.IsSuperRush != nil {
		if *input.IsSuperRush {
			if pipelineState.OPSuperRush != nil {
				dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsSuperRush)
			}
		}
	}

	if input.Status != nil {
		//Todo, refactor use case statement
		if *input.Status == constants.PipelineStatusActive {
			newPipeline.ActivationDateTime = pointers.PrimitiveDateTime(nil)
		}
		if *input.Status == constants.PipelineStatushold {
			newPipeline.HoldDateTime = pointers.PrimitiveDateTime(nil)
		}
	}

	newPipeline.DueDateTime = dueDateTime
	if strings.ToUpper(input.OrderType) == constants.PipelineOrderTypeInterior {
		//newPipeline.Status = stringUtls.ToObject(constants.PipelineStatusStandBy)
		newPipeline.DueDateTime = nil
		newPipeline.ActivationDateTime = nil
	}

	// Todo, - if order type = data entry  ===do not run timer unless comparables->mls added
	if strings.ToUpper(input.OrderType) == constants.PipelineOrderTypeDataEntry {
		newPipeline.DueDateTime = nil
		newPipeline.ActivationDateTime = nil
		//input.Status = stringUtls.ToObject(constants.PipelineStatusStandBy)
	}
	// 	if any order type
	//   - if company = ALtisrouce mkt
	//   - and has no photo loaded, dont run the timer
	if input.CompanyID != nil {
		if *input.CompanyID == constants.ALtisrouceMktID {
			newPipeline.DueDateTime = nil
			//input.Status = stringUtls.ToObject(constants.PipelineStatusStandBy)
		}
	}
	if input.PremiumCompany != nil {
		newPipeline.PremiumCompany = input.PremiumCompany
	}
	if input.PremiumCompanyID != nil {
		newPipeline.PremiumCompanyID = input.PremiumCompanyID
	}

	//Saving to database
	res, err := DbCollections.Pipelines.InsertOne(ctx, newPipeline)
	if err != nil {
		return "", err
	}

	//Todo, added auto add doc from instruction module
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	newPipelineId := stringUtls.TrimObjectChar(lastInsertedIDStr)
	_ = autoCreatePipelineDoc(ctx, newPipelineId, &userId, input.CompanyID, "")

	//send email to client of new order successfully added

	//SendEmailToClientNewOrder(ctx, userId, input.OrderNumber)

	//Send email notification to client
	clientInfo, err := getPipelineAuthor(ctx, userId)
	if err != nil && clientInfo == nil {
		log.Debug("error %v", err)
		return "", err
	}

	emailMessage := `This email is to confirm that your order has been received by our system and will processed within 24 hours. If additional information is needed to complete the following order,
	you will be notified via email. Time will resume once we received the necessary information. Thank you and we appreciate your business`

	emailInput := email.EmailPipelineNotification{
		Action:      "New Order", // hold, unhold, rush, super rush, new order
		Recipient:   clientInfo.Email,
		Address:     input.Address,
		Body:        emailMessage,
		OrderNumber: input.OrderNumber,
		OrderType:   input.OrderType,
		Company:     input.Company,
	}
	email.PipelineNotification(emailInput)

	if input.IsSuperRush != nil {
		if *input.IsSuperRush {
			SendEmailIfSuperRushorder(ctx, input.OrderNumber)
		}
	}
	return newPipelineId, nil
}

func autoCreatePipelineDoc(ctx context.Context, pipelineId string, clientID, companyId *string, createdBy string) error {

	instructionsRaw, err := SearchInstructions(ctx, 0, 1000, nil, clientID, companyId)
	if err != nil {
		return err
	}

	_, err = AutoAddPipelineDoc(ctx, pipelineId, instructionsRaw, createdBy)
	if err != nil {
		return err
	}
	return nil
}

func GetPipelines(ctx context.Context, userId string) ([]*models.Pipeline, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["userId"] = userId

	cur, err := DbCollections.Pipelines.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query pipeline: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Pipeline, 0)
	for cur.Next(ctx) {
		a := &Pipeline{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil

}

func GetPipelineById(ctx context.Context, id string) (*models.Pipeline, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	a := &Pipeline{}
	err := DbCollections.Pipelines.FindOne(ctx, filter).Decode(a)
	if err != nil {
		log.Error("Failed to query pipeline: %v", err)
		return nil, errs.DbError
	}
	return a.ToModels(), nil
}

func GetPipelineByIdDataStore(ctx context.Context, id string) (*Pipeline, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}
	a := &Pipeline{}
	err := DbCollections.Pipelines.FindOne(ctx, filter).Decode(a)
	if err != nil {
		log.Error("Failed to query pipeline: %v", err)
		return nil, errs.DbError
	}
	return a, nil
}

func GetPipelineByOrderNumber(ctx context.Context, orderNumber string) ([]*models.Pipeline, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["orderNumber"] = orderNumber

	cur, err := DbCollections.Pipelines.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query pipeline: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*models.Pipeline, 0)
	for cur.Next(ctx) {
		a := &Pipeline{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline: %v", err)
			return nil, err
		}
		list = append(list, a.ToModels())
	}
	return list, nil

}

func GetPipelineByOrderNumberDataStore(ctx context.Context, orderNumber string) ([]*Pipeline, error) {
	filter := bson.M{}
	filter["status"] = bson.M{"$ne": "DELETED"}
	filter["orderNumber"] = orderNumber

	cur, err := DbCollections.Pipelines.Find(ctx, filter)
	if err != nil {
		log.Error("Failed to query pipeline: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*Pipeline, 0)
	for cur.Next(ctx) {
		a := &Pipeline{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil

}

func UpdatePipeline(ctx context.Context, id string, input models.UpdatePipelineInput, myFullname, myID string, isActivePipelineAndRunTimer *bool) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	updateDoc := bson.M{}

	update := bson.M{
		"lastUpdateTime": pointers.PrimitiveDateTime(nil),
	}

	if input.OrderNumber != nil {
		update["orderNumber"] = *input.OrderNumber
	}
	if input.Status != nil {
		update["status"] = *input.Status
	}
	if input.Address != nil {
		update["address"] = *input.Address
	}
	if input.Country != nil {
		update["country"] = *input.Country
	}
	if input.County != nil {
		update["county"] = *input.County
	}
	if input.Location != nil {
		update["location"] = *input.Location
	}
	if input.Company != nil {
		update["company"] = *input.Company
	}
	if input.CompanyID != nil {
		update["companyID"] = *input.CompanyID
	}
	if input.ZipCode != nil {
		update["zipCode"] = *input.ZipCode
	}
	if input.PremiumCompany != nil {
		update["premiumCompany"] = *input.PremiumCompany
	}
	if input.PremiumCompanyID != nil {
		update["premiumCompanyId"] = *input.PremiumCompanyID
	}
	if input.OtherCompany != nil {
		update["otherCompany"] = *input.OtherCompany
	}
	if input.Type != nil {
		update["type"] = *input.Type
	}
	if input.OrderType != nil {
		update["orderType"] = *input.OrderType
	}
	if input.Objective != nil {
		update["objective"] = *input.Objective
	}
	if input.Assign != nil {
		update["assign"] = *input.Assign
	}
	if input.AssignID != nil {
		update["assignId"] = *input.AssignID
	}
	if input.Mls != nil {
		update["mls"] = *input.Mls
	}
	if input.IsRushOrder != nil {
		update["isRushOrder"] = *input.IsRushOrder
	}
	if input.IsSuperRush != nil {
		update["isSuperRush"] = *input.IsSuperRush
	}
	if input.IsInspection != nil {
		update["isInspection"] = *input.IsInspection
	}
	if input.IsInitialBpo != nil {
		update["isInitialBpo"] = *input.IsInitialBpo
	}
	if input.OrderFee != nil {
		update["orderFee"] = *input.OrderFee
	}
	if input.TotalFee != nil {
		update["totalFee"] = *input.TotalFee
	}
	if input.IsSyncedToTurboBpo != nil {
		update["isSyncedToTurboBpo"] = *input.IsSyncedToTurboBpo
	}
	if input.RatingOverAll != nil {
		update["ratingOverAll"] = *input.RatingOverAll
	}
	if input.RatingTimeliness != nil {
		update["ratingTimeliness"] = *input.RatingTimeliness
	}
	if input.RatingQuality != nil {
		update["ratingQuality"] = *input.RatingQuality
	}
	if input.RatingFeedback != nil {
		update["ratingFeedback"] = *input.RatingFeedback
	}
	if input.IsForQa != nil {
		update["isForQa"] = *input.IsForQa
	}
	if input.IsHold != nil {
		update["ishold"] = *input.IsHold
	}
	if input.HoldRemarks != nil {
		update["holdRemarks"] = *input.HoldRemarks
	}
	if input.UnHoldRemarks != nil {
		update["unHoldRemarks"] = *input.UnHoldRemarks
	}
	if input.CancelRemarks != nil {
		update["cancelRemarks"] = *input.CancelRemarks
	}

	// dueDatetime calculation

	currentPipeline, err := GetPipelineByIdDataStore(ctx, id)
	if err != nil {
		return false, err
	}

	pipelineState, err := GetPipelineStateDataStore(ctx)

	if err != nil || pipelineState == nil {
		log.Error("%s", "something wrong on pipeline state")
		return false, errs.PipelineStateError
	}

	//dueDatetime read turnAroundtime
	turnAroundtime, err := pipelineState.TurnAoundTime(ctx)
	if err != nil || turnAroundtime == nil {
		return false, errs.InvalidTurnAroundTime
	}

	if input.IsRushOrder != nil {
		if *input.IsRushOrder {
			if pipelineState.OPRush != nil {
				update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(orderIsRush)

			}
		} else {
			if currentPipeline.OrderType != nil {
				update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*currentPipeline.OrderType, float64(*turnAroundtime)))
				//update.DueDateTime = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*currentPipeline.OrderType, float64(*turnAroundtime)))
			}
		}
	}
	if input.IsSuperRush != nil {
		if *input.IsSuperRush {
			if pipelineState.OPSuperRush != nil {
				//update.DueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsSuperRush)
				update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(orderIsSuperRush)
			}
		} else {

			if currentPipeline.OrderType != nil {
				update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*currentPipeline.OrderType, float64(*turnAroundtime)))
				//	update.DueDateTime = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*currentPipeline.OrderType, float64(*turnAroundtime)))
			}
		}
	}

	if input.OrderType != nil && stringUtls.ObjectTOString(input.OrderType) != stringUtls.ObjectTOString(currentPipeline.OrderType) {
		//update.DueDateTime = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*input.OrderType, float64(*turnAroundtime)))
		update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(stringUtls.ObjectTOString(input.OrderType), float64(*turnAroundtime)))
		log.Debug("@debug this was read %s", update["dueDateTime"])
		if strings.ToUpper(stringUtls.ObjectTOString(input.OrderType)) == constants.PipelineOrderTypeInterior {
			//update.DueDateTime = nil

			update["dueDateTime"] = nil
			update["activationDateTime"] = nil
			//update.ActivationDateTime = nil
		}

		if strings.ToUpper(stringUtls.ObjectTOString(input.OrderType)) == constants.PipelineOrderTypeDataEntry {
			update["dueDateTime"] = nil
			update["activationDateTime"] = nil
		}

	}

	// Logic to activate pipeline and run the timer
	if pointers.ObjectTOBool(isActivePipelineAndRunTimer) {
		update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*currentPipeline.OrderType, float64(*turnAroundtime)))

		log.Debug("@debug duedate time %s", PrettyPrint(update["dueDateTime"]))
	}

	//ALtisrouceMktID
	if (stringUtls.ObjectTOString(input.CompanyID) != stringUtls.ObjectTOString(currentPipeline.CompanyID)) && stringUtls.ObjectTOString(input.CompanyID) == constants.ALtisrouceMktID {
		//update.DueDateTime = nil
		update["dueDateTime"] = nil
	}

	///end duedate calculation

	if input.AssignID != nil && input.Assign != nil {
		//update.AssignDateTime = pointers.PrimitiveDateTime(nil)
		update["assignDateTime"] = pointers.PrimitiveDateTime(nil)
		updateDoc["$addToSet"] = AddSetPipelineAssignedHistory("Assigned", *input.Assign, *input.AssignID, myFullname, myID)
		//Send email to contractor assigned on this order
		SendEmailToContractor(ctx, currentPipeline, *input.AssignID)
		//Send email to client owner on this order
		SendEmailToClientOrderWasAssigned(ctx, currentPipeline, *input.AssignID)
	}

	if input.Status != nil {
		//Todo, refactor use case statement
		if *input.Status == constants.PipelineStatusActive {
			update["activationDateTime"] = primitive.DateTime(millis.NowInMillis())
			//get duedattime and add new duedatetime

			if currentPipeline.Status != nil {
				if *currentPipeline.Status == constants.PipelineStatushold {
					// holdDate - duedate
					if currentPipeline.HoldDateTime != nil {
						holdDatetime := pointers.PrimativeToDateTime(*currentPipeline.HoldDateTime)
						dueDateTime := pointers.PrimativeToDateTime(*currentPipeline.DueDateTime)
						dueVariance := dueDateTime.Sub(holdDatetime)
						//update.DueDateTime = pointers.PrimitiveDateTimeAddHr(dueVariance.Hours())
						update["dueDateTime"] = pointers.PrimitiveDateTimeAddHr(dueVariance.Hours())
					}
				}
			}
		}
		if *input.Status == constants.PipelineStatushold {
			//update.HoldDateTime = pointers.PrimitiveDateTime(nil)
			update["holdDateTime"] = pointers.PrimitiveDateTime(nil)
			//Send to client of order is onhold

			SendEmailToClientOrderOnHold(ctx, currentPipeline)
		}

		if *input.Status == constants.PipelineStatusUnhold {
			//Send to client of order is unhold

			SendEmailToClientOrderUnHold(ctx, currentPipeline)
		}

		if *input.Status == constants.PipelineStatusStandBy {
			//Send to client of order is standby

			//SendEmailToClientOrderStandBy(ctx, currentPipeline)
		}

		if *input.Status == constants.PipelineStatusComplete {
			//Todo, add status complete validation
			// if countTrue(currentPipeline.IsProcessIform, currentPipeline.IsProcessIfill, currentPipeline.IsProcessReview) <= 1 { // update to complete validation
			// 	return false, errs.CannotProcessPipeline
			// }
			// end validation
			UpdateIformStatus(ctx, id, constants.PipelineStatusComplete)
			//Send to client of order is unhold

			//Todo
			//	SendEmailToClientOrderComplete(ctx, currentPipeline)
		}
		if *input.Status == constants.PipelineStatusCancelled {
			//send email if order was cancelled

			SendEmailToClienIfOrderCancel(ctx, currentPipeline)

		}
	}

	if input.AuthorID != nil {
		authorUserRaw, err := getPipelineAuthor(ctx, *input.AuthorID)
		if err != nil {
			return false, err
		}
		update["userId"] = input.AuthorID
		update["userName"] = stringUtls.ToObject(authorUserRaw.FullName())

	}

	if input.IsProcessIform != nil {
		update["isProcessIform"] = input.IsProcessIform
		update["processIformModifiedDate"] = pointers.PrimitiveDateTime(nil)
	}

	if input.IsProcessIfill != nil {
		update["isProcessIfill"] = input.IsProcessIfill
		update["ifillProcessModifiedDate"] = pointers.PrimitiveDateTime(nil)
	}

	if input.IsProcessReview != nil {
		update["isProcessReview"] = input.IsProcessReview
		update["processReviewModifiedDate"] = pointers.PrimitiveDateTime(nil)
	}

	// 	if any order type
	//   - if company = ALtisrouce mkt
	//   - and has no photo loaded, dont run the timer
	// if input.CompanyID != nil {
	// 	if *input.CompanyID == "5dd6317d778617cbe9f32926" {
	// 		update.Status = stringUtls.ToObject(constants.PipelineStatusStandBy)
	// 		update.DueDateTime = nil
	// 	}
	// }
	log.Debug("@debug update payload %s", PrettyPrint(update))
	updateDoc["$set"] = update
	_, err = DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}

	// updateDoc = bson.M{
	// 	"$set": setDoc,
	// }
	// _, err = DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	// if err != nil {
	// 	return false, err
	// }

	//update userActive Counters
	// update assign counter on user collection
	if input.AssignID != nil {
		if currentPipeline.AssignId == nil {
			UpdateUserOrderCount(ctx, stringUtls.ObjectTOString(input.AssignID), pointers.Int(1), nil, nil)
		}
		if stringUtls.ObjectTOString(currentPipeline.AssignId) != stringUtls.ObjectTOString(input.AssignID) {
			UpdateUserOrderCount(ctx, stringUtls.ObjectTOString(input.AssignID), pointers.Int(1), nil, nil)
			//minus to current contractor
			UpdateUserOrderCount(ctx, stringUtls.ObjectTOString(currentPipeline.AssignId), pointers.Int(-1), nil, nil)
		}
	}

	// Create invoice request if status completed
	if stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusComplete || stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusStandBy {

		invoiceInput := models.SaveInvoiceInput{}
		if stringUtls.ObjectTOString(currentPipeline.Status) == constants.PipelineStatusStandBy && *input.Status == constants.PipelineStatusComplete {
			invoiceInput = models.SaveInvoiceInput{
				Type:        "submit",
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
			}

		} else if *input.Status == string(constants.PipelineStatusComplete) || *input.Status == string(constants.PipelineStatusStandBy) {

			invoiceInput = models.SaveInvoiceInput{
				Type:        "order",
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
			}

		}
		_, err := SaveInvoice(ctx, invoiceInput, "System Automation", currentPipeline, nil, nil)
		if err != nil {
			return false, err
		}

		// Add entry on request
		pipelinePhotosTotals, err := GetPipelinePhotosCountGroupByPipelineId(ctx, []string{currentPipeline.ID.Hex()})

		hasPhotos := false
		if len(pipelinePhotosTotals) > 0 {
			hasPhotos = true
		}

		AddRequest(ctx, currentPipeline.ID.Hex(), myFullname, myID, currentPipeline, hasPhotos)

		//Add invoice request history!
		requestHistory := RequestHistory{
			PipelineId:  currentPipeline.ID.Hex(),
			Status:      "PENDING",
			ClientId:    stringUtls.ObjectTOString(currentPipeline.UserId),
			EmployeeId:  myID,
			Type:        stringUtls.ToObject("INVOICE"),
			OrderNumber: currentPipeline.OrderNumber,
			Address:     currentPipeline.Address,
			Company:     currentPipeline.Company,
			//Remarks:         input.Remarks,
			CreatedDateTime: currentPipeline.CreatedDateTime,
		}
		AddRequestHistory(ctx, requestHistory)
	}

	if input.RatingOverAll != nil || input.RatingFeedback != nil {
		fmt.Println("send email")
		messageBody := fmt.Sprintf("rating notfication overall : %v, feedback:  %s", pointers.ToInt(input.RatingOverAll), stringUtls.ObjectTOString(input.RatingFeedback))
		err := email.Init("daryljamesbalangigue@gmail.com", "rating notfication", messageBody)
		if err != nil {
			//	return false, err
		}
		err = email.Init("marlonpamisa@gmail.com", "rating notfication", messageBody)
		if err != nil {
			//	return false, err
		}
	}

	//Send Email notification
	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		log.Debug("error %v", err)
		return false, err
	}

	if input.IsHold != nil {

		if *input.IsHold {
			emailMessage := `ORDER ON HOLD - Please Respond to the message to reactivate. Thank you.
			"%s"`

			emailInput := email.EmailPipelineNotification{
				Action:      "Hold", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        fmt.Sprintf(emailMessage, stringUtls.ObjectTOString(input.HoldRemarks)),
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(currentPipeline.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)
		} else {
			emailMessage := `ORDER UNHOLD .
			"%s"`

			emailInput := email.EmailPipelineNotification{
				Action:      "UnHold", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        fmt.Sprintf(emailMessage, stringUtls.ObjectTOString(input.UnHoldRemarks)),
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(currentPipeline.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)

		}
		if stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusActive && stringUtls.ObjectTOString(currentPipeline.Status) != constants.PipelineStatusActive {
			emailMessage := `ORDER WAS REACTIVATED - This order will receive priority status and be completed as soon as possible. Thank you.
			"please complete"`

			emailInput := email.EmailPipelineNotification{
				Action:      "Re-Activated", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        emailMessage,
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(input.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)
		}

		if stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusStandBy && stringUtls.ObjectTOString(currentPipeline.Status) != constants.PipelineStatusStandBy {
			emailMessage := `Hello Thank you for your report. It is being return for the following Item(s): 1.REQUIRED
			(***please read the attached directions which begin on page 4 for a successful upload). a.
			Upload the tax record for the subject use the drop down menun to select Taxt Record (PDF only) 2.
			Add the subject's Apn # to the report`

			emailInput := email.EmailPipelineNotification{
				Action:      "Cancelled", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        emailMessage,
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(input.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)
		}

		if stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusCancelled && stringUtls.ObjectTOString(currentPipeline.Status) != constants.PipelineStatusCancelled {
			emailMessage := `ORDER WAS CANCELLED - This order was cancelled either due to a duplicate, client request refund, or not falling within our company policy. If you received this notification in error, please contact us.
			"%s"`

			emailInput := email.EmailPipelineNotification{
				Action:      "Cancelled", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        fmt.Sprintf(emailMessage, stringUtls.ObjectTOString(input.CancelRemarks)),
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(input.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)
		}

		if stringUtls.ObjectTOString(input.Status) == constants.PipelineStatusComplete && stringUtls.ObjectTOString(currentPipeline.Status) != constants.PipelineStatusComplete {
			emailMessage := `Order is now complete. Report not yet submitted for your review. Thank you`

			emailInput := email.EmailPipelineNotification{
				Action:      "Complete", // hold, unhold, rush, super rush, new order
				Recipient:   clientInfo.Email,
				Address:     stringUtls.ObjectTOString(currentPipeline.Address),
				Body:        emailMessage,
				OrderNumber: stringUtls.ObjectTOString(currentPipeline.OrderNumber),
				OrderType:   stringUtls.ObjectTOString(input.OrderType),
				Company:     stringUtls.ObjectTOString(currentPipeline.Company),
			}
			email.PipelineNotification(emailInput)
		}
	}
	return true, nil
}

func SendEmailIfSuperRushorder(ctx context.Context, orderNumber string) error {

	messageBody := fmt.Sprintf("Super rush has created with Order Number :%s", orderNumber)
	//todo, configurable order@turbo.com
	err := email.Init("marlonpamisa@gmail.com", "Super Rush order", messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClienIfOrderCancel(ctx context.Context, currentPipeline *Pipeline) error {

	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		return err
	}
	messageBody := fmt.Sprintf("your order was cancelled with the Order Number :", *currentPipeline.OrderNumber)
	err = email.Init(clientInfo.Email, "Order cancelled", messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SearchFilterForClient(userId string, filter *models.PipelineFilterInput) bson.M {
	daoFilter := SearchFilterDefault(filter)
	daoFilter["userId"] = userId
	return daoFilter
}

func SearchFilterForContractor(contactorId string, filter *models.PipelineFilterInput) bson.M {
	daoFilter := SearchFilterDefault(filter)
	daoFilter["assignId"] = contactorId
	return daoFilter
}

func SearchFilterForAdmin(filter *models.PipelineFilterInput) bson.M {
	daoFilter := SearchFilterDefault(filter)
	return daoFilter
}

func EmptyFilter() bson.M {
	return bson.M{}
}

func FilterQA() bson.M {
	filter := bson.M{}
	statuses := []string{
		constants.QualityControlStatusActive,
		constants.QualityControlStatusHold,
	}
	filter["status"] = bson.M{"$in": statuses}
	return filter
}

func SearchFilterByStatuses(statuses []string) bson.M {
	filter := bson.M{"status": bson.M{"$in": statuses}}
	//filter["activationDateTime"] =  pointers.PrimitiveDateTime(&varDateFrom)
	return filter
}

func PipelineIdsFilter(ids []string) bson.M {
	var objIds []primitive.ObjectID
	for _, id := range ids {
		obj, _ := primitive.ObjectIDFromHex(id)
		objIds = append(objIds, obj)
	}
	return bson.M{"_id": bson.M{"$in": objIds}}

}

func SearchFilterDefault(filter *models.PipelineFilterInput) bson.M {
	daoFilter := bson.M{
		//	"isForQa": false,
	}
	if contains(filter.Status, constants.PipelineStatusActive) {

		if !pointers.ObjectTOBool(filter.IsProcessIfill) && !pointers.ObjectTOBool(filter.IsProcessIform) { //Apply patch display for QA pipeline on billing
			daoFilter["$or"] = []bson.M{
				{"isForQa": false},
				{"isForQa": bson.M{
					"$exists": false,
				}},
			}
		}
	}

	if filter.AuthorID != nil {
		daoFilter["userId"] = *filter.AuthorID
	}
	defaultStatusFilter := []string{ // for pipelinse status active
		constants.PipelineStatusActive,
		constants.PipelineStatusSubmit,
		constants.PipelineStatusLate,
		constants.PipelineStatushold,
		constants.PipelineStatusUnhold,
		constants.PipelineStatusActivePhotos,
	}

	// if !contains(filter.Status, constants.PipelineStatusQC) {
	// 	daoFilter["status"] = bson.M{"$in": defaultStatusFilter}
	// }
	if len(filter.Status) == 1 {
		if contains(filter.Status, constants.PipelineStatusActive) {
			//filter by pipeline status active
			daoFilter["status"] = bson.M{"$in": defaultStatusFilter}

		} else if contains(filter.Status, constants.PipelineStatusActivePhotos) {
			// daoFilter["$or"] = []bson.M{
			// 	{"status": constants.PipelineStatusActive},
			// 	{"photosCount": bson.M{"$gte": 1}},
			// }
			daoFilter["status"] = constants.PipelineStatusActive
			daoFilter["photosCount"] = bson.M{"$gte": 1}

		} else if contains(filter.Status, constants.PipelineStatusSubmit) { // if type submit  has 2 or more process and has one or more photos
			daoFilter["$or"] = []bson.M{
				{"isProcessIform": true},
				{"isProcessIfill": true},
				{"isProcessReview": true},
				{"photosCount": bson.M{"$gte": 1}},
			}
		} else if contains(filter.Status, constants.PipelineStatusQC) {
			daoFilter["isForQa"] = true
			//	daoFilter["status"] = nil
		} else {
			if contains(filter.Status, "CANCELLED") {
				filter.Status = append(filter.Status, stringUtls.ToObject(constants.PipelineStatusCancelled))
			}
			daoFilter["status"] = bson.M{"$in": filter.Status}
		}
	}
	// temporary patch for Billing error
	// Todo, refactor, make it readable
	if len(filter.Status) > 1 {
		daoFilter["status"] = bson.M{"$in": filter.Status}
	}

	// if contains(filter.Status, constants.PipelineStatusQC) {
	// 	daoFilter["isForQa"] = true
	// }

	if filter.OrderNumber != nil {
		daoFilter["orderNumber"] = bson.M{
			"$regex":   filter.OrderNumber,
			"$options": "i",
		}
	}
	if filter.Address != nil {
		daoFilter["address"] = bson.M{
			"$regex":   filter.Address,
			"$options": "i",
		}
	}
	if filter.Country != nil {
		daoFilter["country"] = bson.M{
			"$regex":   filter.Country,
			"$options": "i",
		}
	}
	if filter.AssignID != nil || filter.QcUserID != nil || filter.ReviewerUserID != nil {
		//Combine all ids into array of strings
		userIDs := stringUtls.ObjectTOArrString(filter.AssignID, filter.QcUserID, filter.ReviewerUserID)
		//daoFilter["assignId"] = filter.AssignID
		daoFilter["assignId"] = bson.M{"$in": userIDs}

	}
	if filter.Company != nil {
		daoFilter["company"] = filter.Company
	}
	if filter.OrderType != nil {
		daoFilter["orderType"] = filter.OrderType
	}

	daoFilter2 := bson.M{}

	if filter.DateFrom != nil {
		premDateFrom, err := pointers.StringTimeToPrimitive(*filter.DateFrom)
		if err != nil {
			log.Error("failed time parse error : %v", err)
		}
		daoFilter2["$gte"] = premDateFrom
		daoFilter["createdDateTime"] = daoFilter2
	}

	if filter.DateTo != nil {
		premDateTo, err := pointers.StringTimeToPrimitiveAddHr(*filter.DateTo, 24)
		if err != nil {
			log.Error("failed time to parse error : %v", err)
		}
		daoFilter2["$lte"] = premDateTo
		daoFilter["createdDateTime"] = daoFilter2
	}

	if filter.IsProcessIfill != nil {
		daoFilter["isProcessIfill"] = *filter.IsProcessIfill
	}

	if filter.IsProcessIform != nil {
		daoFilter["isProcessIform"] = *filter.IsProcessIform
	}

	if filter.IsProcessReview != nil {
		daoFilter["isProcessReview"] = *filter.IsProcessReview
	}

	if filter.IsBilled != nil {
		daoFilter["isBilled"] = *filter.IsBilled
	}

	return daoFilter
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func UpdatePipelineStatuses(ctx context.Context, filter *bson.M, newStatus string, newOrderFee *float64) (bool, error) {

	update := bson.M{
		"status": newStatus,
	}
	if newOrderFee != nil {
		update["orderFee"] = *newOrderFee
		update["totalFee"] = *newOrderFee
	}
	updateDoc := bson.M{
		"$set": update,
	}
	res, err := DbCollections.Pipelines.UpdateMany(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func SearchPipelines(ctx context.Context, filter bson.M, offset, limit int) ([]*Pipeline, error) {
	log.Debug("@debug pipeline filter: %s", PrettyPrint(filter))
	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawPipelines := make([]*Pipeline, 0)
	for cur.Next(ctx) {
		a := &Pipeline{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		rawPipelines = append(rawPipelines, a)
	}

	return rawPipelines, nil
}

func GetPipelinesCount(ctx context.Context, filter bson.M) (*int64, error) {

	count, err := DbCollections.Pipelines.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query pipeline: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetPipelineCompleteOrderFilterByMonthYear(ctx context.Context, filter models.SalesAnalyticsFilterInput) ([]*models.SalesAnalytics, error) {

	project01 := bson.M{
		"_id":    0,
		"day":    bson.M{"$dayOfMonth": "$createdDateTime"},
		"month":  bson.M{"$month": "$createdDateTime"},
		"year":   bson.M{"$year": "$createdDateTime"},
		"status": "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"day":   "$day",
			"month": "$month",
			"year":  "$year",
		},
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{
		"month": filter.Month,
		"year":  filter.Year,
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawSalesAnalytics := make([]*models.SalesAnalytics, 0)
	for cur.Next(ctx) {
		a := &SalesAnalytics{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		rawSalesAnalytics = append(rawSalesAnalytics, a.ToModels())
	}
	return rawSalesAnalytics, nil

}

func GetPipelineSumByPaymentStatusPerMonth(ctx context.Context, filter models.OrderAnalyticsFilterInput) ([]*models.OrderAnalytics, error) {

	project01 := bson.M{
		"_id":    0,
		"userId": "$userId",
		"month":  bson.M{"$month": "$createdDateTime"},
		"year":   bson.M{"$year": "$createdDateTime"},
		"status": "$status",
		"paid": bson.M{
			"$cond": bson.M{
				"if":   bson.M{"$eq": []string{"paid", "$paymentStatus"}},
				"then": "$totalFee",
				"else": 0,
			},
		},
		"unpaid": bson.M{
			"$cond": bson.M{
				"if":   bson.M{"$eq": []string{"unpaid", "$paymentStatus"}},
				"then": "$totalFee",
				"else": 0,
			},
		},
	}

	group := bson.M{
		"_id": bson.M{
			"month":  "$month",
			"year":   "$year",
			"userId": "$userId",
		},
		"paid":   bson.M{"$sum": "$paid"},
		"unpaid": bson.M{"$sum": "$unpaid"},
	}

	Colfilter := bson.M{
		"year": filter.Year,
	}

	if filter.Client != nil {
		Colfilter["userId"] = *filter.Client
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawOrderAnalytics := make([]*models.OrderAnalytics, 0)
	for cur.Next(ctx) {
		a := &OrderAnalytics{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		rawOrderAnalytics = append(rawOrderAnalytics, a.ToModels())
	}

	return rawOrderAnalytics, nil

}

func GetPipelineCountOrderPerMonthAndYearFilterByUserIdsGroupByUserId(ctx context.Context, month int, year int, userIds []string) (map[string]*int, error) {

	project01 := bson.M{
		"_id":    0,
		"userId": "$userId",
		"month":  bson.M{"$month": "$createdDateTime"},
		"year":   bson.M{"$year": "$createdDateTime"},
		"status": "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"month":  "$month",
			"year":   "$year",
			"userId": "$userId",
		},
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{
		"month":  month,
		"year":   year,
		"userId": bson.M{"$in": userIds},
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	clientOrderTotals := make(map[string]*int)
	for cur.Next(ctx) {
		a := &ClientOrder{}

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		clientOrderTotals[a.ID.ClientId] = a.Count
	}

	return clientOrderTotals, nil

}

func GetPipelineOrderByUserIdSumPaidUnpaid(ctx context.Context, userIds []string) (map[string]*ClientBalance, error) {

	project01 := bson.M{
		"_id":    0,
		"userId": "$userId",
		"status": "$status",
		"paid": bson.M{
			"$cond": bson.M{
				"if":   bson.M{"$eq": []string{"paid", "$paymentStatus"}},
				"then": "$totalFee",
				"else": 0,
			},
		},
		"unpaid": bson.M{
			"$cond": bson.M{
				"if":   bson.M{"$eq": []string{"unpaid", "$paymentStatus"}},
				"then": "$totalFee",
				"else": 0,
			},
		},
	}

	group := bson.M{
		"_id": bson.M{
			"userId": "$userId",
		},
		"paid":   bson.M{"$sum": "$paid"},
		"unpaid": bson.M{"$sum": "$unpaid"},
	}

	Colfilter := bson.M{
		"userId": bson.M{"$in": userIds},
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	clientBalanceMap := make(map[string]*ClientBalance)

	for cur.Next(ctx) {
		a := &ClientBalance{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		clientBalanceMap[a.ClientId.ClientId] = a
	}

	return clientBalanceMap, nil

}

func GetPipelineCountOrderPerMonthAndYearFilterByUserIdsAndyearGroupByUserId(ctx context.Context, year int, userIds []string) (map[string]*ContractorAssignOrder, error) {
	project01 := bson.M{
		"_id":      0,
		"assignId": "$assignId",
		"month":    bson.M{"$month": "$createdDateTime"},
		"year":     bson.M{"$year": "$createdDateTime"},
		"status":   "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"month":    "$month",
			"year":     "$year",
			"assignId": "$assignId",
		},
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{
		"year": year,
	}

	if len(userIds) > 0 {
		Colfilter["assignId"] = bson.M{"$in": userIds}
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	clientOrderTotals := make(map[string]*ContractorAssignOrder)
	for cur.Next(ctx) {
		a := &ContractorAssignOrder{}

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		clientOrderTotals[a.ID.ContactorId] = a
	}

	return clientOrderTotals, nil

}

func GetOrderSubmitted(ctx context.Context, year int) ([]*OrderSubmitted, error) {
	project01 := bson.M{
		"_id":         0,
		"coordinator": "$coordinator",
		"month":       bson.M{"$month": "$createdDateTime"},
		"year":        bson.M{"$year": "$createdDateTime"},
		"status":      "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"month":       "$month",
			"year":        "$year",
			"coordinator": "$coordinator",
		},
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{
		"year": year,
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	orderSubmits := make([]*OrderSubmitted, 0)
	for cur.Next(ctx) {
		a := &OrderSubmitted{}

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}

		orderSubmits = append(orderSubmits, a)
	}

	return orderSubmits, nil

}

func GetCreatedOrderPerday(ctx context.Context, assignedID *string) (*int, error) {

	group := bson.M{
		"_id":   0,
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{}

	if assignedID != nil {
		Colfilter["userId"] = *assignedID
	}

	layout := "2006-01-02"
	daoFilter2 := bson.M{}
	currentTime := time.Now()
	varDateFrom, err := time.Parse(layout, currentTime.Format("2006-01-02"))
	log.Debug("current date %v", PrettyPrint(varDateFrom))
	if err != nil {
		log.Debug("failed time parse error : %v", err)
	}
	daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
	//add 1 day
	currentTime = currentTime.AddDate(0, 0, 1)

	varDateTo, err := time.Parse(layout, currentTime.Format("2006-01-02"))
	if err != nil {
		log.Debug("failed time parse error : %v", err)
	}
	daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
	Colfilter["lastUpdateTime"] = daoFilter2

	pipe := []bson.M{
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	a := &AssignedCount{}

	for cur.Next(ctx) {

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		fmt.Println(PrettyPrint(a))
	}

	return a.Count, nil

}

func GetTodayOrderCount(ctx context.Context) (*int, error) {

	group := bson.M{
		"_id":   0,
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{}

	layout := "2006-01-02"
	daoFilter2 := bson.M{}
	currentTime := time.Now()
	varDateFrom, err := time.Parse(layout, currentTime.Format("2006-01-02"))
	log.Debug("current date %v", PrettyPrint(varDateFrom))
	if err != nil {
		log.Debug("failed time parse error : %v", err)
	}
	daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
	//add 1 day
	currentTime = currentTime.AddDate(0, 0, 1)

	varDateTo, err := time.Parse(layout, currentTime.Format("2006-01-02"))
	if err != nil {
		log.Debug("failed time parse error : %v", err)
	}
	daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
	Colfilter["createdDateTime"] = daoFilter2

	pipe := []bson.M{
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.Pipelines.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query pipeline:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	a := &AssignedCount{}

	for cur.Next(ctx) {

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode pipeline entry: %v", err)
			return nil, err
		}
		fmt.Println(PrettyPrint(a))
	}

	return a.Count, nil

}

func AddSetPipelineHistory(ctx context.Context, pipelineID string, action string, value string, modifiedBy string) error {

	historyRaw := &PipelineHistory{
		LogDateTime: primitive.DateTime(millis.NowInMillis()),
		Action:      stringUtls.ToObject(action),
		Value:       stringUtls.ToObject(value),
		ModifiedBy:  stringUtls.ToObject(modifiedBy),
	}

	objId, _ := primitive.ObjectIDFromHex(pipelineID)
	filter := bson.D{{"_id", objId}}

	updateStatus := bson.M{}

	incrementValue := 1
	if action == "ADD_PHOTO" {
		incrementValue = 1
	}

	if action == "DELETE_PHOTO" {
		incrementValue = -1
	}

	updateDoc := bson.M{
		"$inc": bson.M{"photosCount": incrementValue},
		"$addToSet": bson.M{
			"history": historyRaw,
		},
	}

	if action == "ADD_PHOTO" {
		pipelineRaw, err := GetPipelineByIdDataStore(ctx, pipelineID)
		if err != nil {
			return err
		}
		if pipelineRaw.OrderType != nil {
			log.Debug("@debig pipelineRaw.PhotosCount %s", PrettyPrint(pipelineRaw.PhotosCount))
			log.Debug("@debig pipelineRaw.PhotosCount %s", *pipelineRaw.OrderType)
			log.Debug("@debig pipelineRaw.PhotosCount %s", stringUtls.ObjectTOString(pipelineRaw.CompanyID))
			if pipelineRaw.PhotosCount == nil &&
				(strings.ToUpper(*pipelineRaw.OrderType) == constants.PipelineOrderTypeInterior ||
					stringUtls.ObjectTOString(pipelineRaw.CompanyID) == constants.ALtisrouceMktID) {

				pipelineState, err := GetPipelineStateDataStore(ctx)
				if err != nil || pipelineState == nil {
					log.Error("%s", "something wrong on pipeline state")
					return errs.PipelineStateError
				}
				//dueDatetime read turnAroundtime
				turnAroundtime, err := pipelineState.TurnAoundTime(ctx)
				if err != nil || turnAroundtime == nil {
					return errs.InvalidTurnAroundTime
				}
				dueDateTime := pointers.PrimitiveDateTimeAddHr(CalculateDueDate(*pipelineRaw.OrderType, float64(*turnAroundtime)))

				if pipelineRaw.IsRushOrder != nil {
					if *pipelineRaw.IsRushOrder {
						dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsRush)
					}
				}
				if pipelineRaw.IsSuperRush != nil {
					if *pipelineRaw.IsSuperRush {
						dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsSuperRush)
					}
				}
				updateStatus["activationDateTime"] = pointers.PrimitiveDateTime(nil)
				updateStatus["dueDateTime"] = dueDateTime
				log.Debug("update data %s", PrettyPrint(updateStatus))

				updateDoc["$set"] = updateStatus
				//updateStatus["status"] = constants.PipelineStatusActive
			}
			// Patch, to set order to active if client upload a photo from standby
			// Todo, refactor this code, make it readable
			if pointers.ToInt(pipelineRaw.PhotosCount) == 0 && stringUtls.ObjectTOString(pipelineRaw.Status) == constants.PipelineStatusStandBy {
				updateStatus["status"] = constants.PipelineStatusActive
				updateDoc["$set"] = updateStatus
			}
		}
		updateDoc["$inc"] = bson.M{"photosCount": incrementValue}
		updateDoc["$addToSet"] = bson.M{"history": historyRaw}
		// updateDoc = bson.M{
		// 	"$inc": bson.M{"photosCount": incrementValue},
		// 	//"$set": updateStatus,
		// 	"$addToSet": bson.M{
		// 		"history": historyRaw,
		// 	},
		// }
	}

	_, err := DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePipelineDueDateTimeAndAddDocsCount(ctx context.Context, pipelineID string, orderType string, isRushOrder, isSuperRush *bool) error {

	objId, _ := primitive.ObjectIDFromHex(pipelineID)
	filter := bson.D{{"_id", objId}}

	pipelineState, err := GetPipelineStateDataStore(ctx)
	if err != nil || pipelineState == nil {
		log.Error("%s", "something wrong on pipeline state")
		return errs.PipelineStateError
	}
	//dueDatetime read turnAroundtime
	turnAroundtime, err := pipelineState.TurnAoundTime(ctx)
	if err != nil || turnAroundtime == nil {
		return errs.InvalidTurnAroundTime
	}

	dueDateTime := pointers.PrimitiveDateTimeAddHr(CalculateDueDate(orderType, float64(*turnAroundtime)))
	update := bson.M{}

	if isRushOrder != nil {
		if *isRushOrder {
			dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsRush)
		}
	}
	if isSuperRush != nil {
		if *isSuperRush {
			dueDateTime = pointers.PrimitiveDateTimeAddHr(orderIsSuperRush)
		}
	}

	update["activationDateTime"] = pointers.PrimitiveDateTime(nil)
	update["dueDateTime"] = dueDateTime

	updateDoc := bson.M{
		"$set": update,
		"$inc": bson.M{"docsCount": 1},
	}

	_, err = DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePipelineDocsCount(ctx context.Context, pipelineID string, ints int) error {

	objId, _ := primitive.ObjectIDFromHex(pipelineID)
	filter := bson.D{{"_id", objId}}

	updateDoc := bson.M{
		"$inc": bson.M{"docsCount": ints},
	}

	_, err := DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	return nil
}

func AddSetPipelineAssignedHistory(action, assignee, assigneeID, assignedBy, assignedByID string) bson.M {

	historyRaw := &PipelineAssignedHistory{
		LogDateTime:  primitive.DateTime(millis.NowInMillis()),
		Action:       stringUtls.ToObject(action),
		Assignee:     stringUtls.ToObject(assignee),
		AssigneeID:   stringUtls.ToObject(assigneeID),
		AssignedBy:   stringUtls.ToObject(assignedBy),
		AssignedByID: stringUtls.ToObject(assignedByID),
	}

	//objId, _ := primitive.ObjectIDFromHex(pipelineID)
	//filter := bson.D{{"_id", objId}}

	assignedHistory := bson.M{
		"assignedHistory": historyRaw,
	}

	// updateDoc := bson.M{
	// 	"$addToSet": bson.M{
	// 		"assignedHistory": historyRaw,
	// 	},
	// }

	// _, err := DbCollections.Pipelines.UpdateOne(ctx, filter, updateDoc)
	// if err != nil {
	// 	return err
	// }
	return assignedHistory
}

//Todo, refactor, migrate to utils package
func countTrue(vals ...*bool) int {
	trueCount := 0
	for _, v := range vals {
		if v != nil {
			if *v {
				trueCount++
			}
		}
	}
	return trueCount
}

//For API uses
func GetPipelineSearchbyOrderNumberOrAddress(ctx context.Context, orderNumber, orderAddress string, offset, limit int) ([]*Pipeline, error) {
	filter := bson.M{}
	if orderNumber != "" {
		filter["orderNumber"] = bson.M{
			"$regex":   orderNumber,
			"$options": "i",
		}
	}
	if orderAddress != "" {
		filter["address"] = bson.M{
			"$regex":   orderAddress,
			"$options": "i",
		}
	}

	return SearchPipelines(ctx, filter, 0, limit)

}
