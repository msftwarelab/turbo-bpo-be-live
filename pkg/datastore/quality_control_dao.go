package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/millis"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveQualityControl(ctx context.Context, pipelineId string, createdBy string) (string, error) {

	count, err := GetPipelineQualityControlCount(ctx, pipelineId)
	if err != nil {
		return "", err
	}
	pipelineFilter := FilterById(pipelineId)
	pipelineRaw, err := SearchPipelines(ctx, *pipelineFilter, 0, 1)
	if err != nil {
		return "", errs.InvalidPipelineId
	}
	if len(pipelineRaw) == 0 {
		return "", errs.InvalidPipelineId
	}
	if pipelineRaw[0] == nil {
		return "", errs.InvalidPipelineId
	}

	newQualityControl := &QualityControl{
		Status:          pointers.String(constants.QualityControlStatusActive),
		QcId:            millis.NowInMillis(),
		PipelineId:      pointers.String(pipelineId),
		ContractorId:    pipelineRaw[0].AssignId,
		OrderNumber:     pipelineRaw[0].OrderNumber,
		OrderAddress:    pipelineRaw[0].Address,
		OrderType:       pipelineRaw[0].OrderType,
		Requests:        pointers.Int(int(*count)),
		CreatedBy:       pointers.String(createdBy),
		IsAccepted:      pointers.Bool(false),
		CreatedDateTime: primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.QualityControls.InsertOne(ctx, newQualityControl)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func FilterByQualityControlID(id string) (*bson.M, error) {
	id64 := utils.StringToInt64(id)
	if id64 == nil {
		return nil, errs.InvalidQcID
	}
	return &bson.M{"qcId": id64}, nil
}
func UpdateQualityControl(ctx context.Context, id string, input models.UpdateQualityControlInput, updatedBy string) (bool, error) {

	id64 := utils.StringToInt64(id)
	if id64 == nil {
		return false, errs.InvalidQcID
	}
	filter := bson.M{"qcId": id64}

	update := &QualityControl{LastUpdateTime: pointers.PrimitiveDateTime(nil)}

	if input.Status != nil {
		update.Status = input.Status
		if *input.Status == constants.QualityControlRequestStatusApproved {
			update.RequestStatus = strings.ToObject(constants.QualityControlRequestStatusApproved)
		}
		if *input.Status == constants.QualityControlRequestStatusDD {
			update.RequestStatus = strings.ToObject(constants.QualityControlRequestStatusDD)
		}
	}

	if input.RequestType != nil {
		update.RequestType = input.RequestType
		update.RequestStatus = strings.ToObject(constants.QualityControlRequestStatusPending)
	}

	updateDoc := bson.M{
		"$set": update,
	}

	if input.AssignID != nil {

		userfilter := FilterByCollectionIds([]string{*input.AssignID})
		user, err := GetUser(ctx, userfilter)
		if err != nil {
			return false, err
		}
		qualityCotnrolRaw, err := SearchQualityControls(ctx, 0, 1, filter)
		if err != nil {
			return false, err
		}

		if len(qualityCotnrolRaw) == 0 {
			return false, errs.InvalidId
		}

		update.Assignee = input.AssignID
		if user != nil {
			update.AssigneeName = strings.ToObject(user.FullName())
		}
		if input.IsAcceptRequest != nil {
			update.IsAccepted = input.IsAcceptRequest
		}
		history := &QualityControlHistory{
			By:          &updatedBy,
			Status:      strings.ToObject("REASSIGN"),
			Date:        pointers.PrimitiveDateTime(nil),
			Reason:      input.AssignReason,
			NewAssignee: strings.ToObject(user.FullName()),
			//	CurrentAssignee: qualityCotnrolRaw[0].AssigneeName,
		}
		if qualityCotnrolRaw[0].AssigneeName != nil {
			history.CurrentAssignee = qualityCotnrolRaw[0].AssigneeName
		}
		updateDoc = bson.M{
			"$set": update,
			"$addToSet": bson.M{
				"history": history,
			},
		}
	}

	res, err := DbCollections.QualityControls.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func DeleteQualityControl(ctx context.Context, id string) (bool, error) {
	filter := bson.D{{"qcId", id}}
	_, err := DbCollections.Accounts.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetQualityControls(ctx context.Context, filter *bson.M) ([]*QualityControl, error) {

	cur, err := DbCollections.QualityControls.Find(ctx, filter)

	if err != nil {
		log.Error("Failed to query quality controls: %v", err)
		return nil, errs.DbError
	}

	defer cur.Close(ctx)

	list := make([]*QualityControl, 0)
	for cur.Next(ctx) {
		a := &QualityControl{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode quality controls entry: %v", err)
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

func SearchQualityControls(ctx context.Context, offset, limit int, filter bson.M) ([]*QualityControl, error) {

	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"createdDateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.QualityControls.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query quality controls:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawQualityControls := make([]*QualityControl, 0)
	for cur.Next(ctx) {
		a := &QualityControl{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode quality controls entry: %v", err)
			return nil, err
		}
		rawQualityControls = append(rawQualityControls, a)
	}

	return rawQualityControls, nil
}

func GetQualityControlsCount(ctx context.Context, filter bson.M) (*int64, error) {

	count, err := DbCollections.QualityControls.CountDocuments(ctx, filter)
	if err != nil {
		log.Debug("Failed to query quality controls: %v", err)
		return nil, err
	}

	return &count, nil
}

func GetQcCountOrderPerMonthAndYearFilterByUserIdsAndyearGroupByUserId(ctx context.Context, year int) (map[string]*ContractorAssignOrder, error) {
	project01 := bson.M{
		"_id":          0,
		"contractorId": "$contractorId",
		"month":        bson.M{"$month": "$createdDateTime"},
		"year":         bson.M{"$year": "$createdDateTime"},
		"status":       "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"month":        "$month",
			"year":         "$year",
			"contractorId": "$contractorId",
		},
		"count": bson.M{"$sum": 1},
	}

	Colfilter := bson.M{
		"year":        year,
		"requestType": bson.M{"$in": []string{constants.QcTypeDataDiscrepancyNqc, constants.QcTypeSubmitNqc, constants.QcTypeNormal, constants.QcTypefullRecomp}},
	}

	pipe := []bson.M{
		{"$project": project01},
		{"$match": Colfilter},
		{"$group": group},
	}

	cur, err := DbCollections.QualityControls.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query quality control:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	clientOrderTotals := make(map[string]*ContractorAssignOrder)
	for cur.Next(ctx) {
		a := &ContractorAssignOrder{}

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode quality control entry: %v", err)
			return nil, err
		}
		clientOrderTotals[a.ID.ContactorId] = a
	}

	return clientOrderTotals, nil

}

func GetCompletedQcCountOrderPerMonthAndYearFilterByUserIdsAndyearGroupByUserId(ctx context.Context, year int) ([]*QcCompleted, error) {
	project01 := bson.M{
		"_id":          0,
		"assigneeName": "$assigneeName",
		"month":        bson.M{"$month": "$createdDateTime"},
		"year":         bson.M{"$year": "$createdDateTime"},
		"status":       "$status",
	}

	group := bson.M{
		"_id": bson.M{
			"month":        "$month",
			"year":         "$year",
			"assigneeName": "$assigneeName",
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

	cur, err := DbCollections.QualityControls.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query quality control:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	qcsCompleteds := make([]*QcCompleted, 0)
	for cur.Next(ctx) {
		a := &QcCompleted{}

		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode quality control entry: %v", err)
			return nil, err
		}
		qcsCompleteds = append(qcsCompleteds, a)
	}

	return qcsCompleteds, nil

}

func UpdateQcRequest(ctx context.Context, id string, input models.UpdateQcRequestInput, myName string) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{
		{"lastUpdateTime", pointers.PrimitiveDateTime(nil)},
	}

	if input.RequestType != nil {

		setDoc = append(setDoc, bson.E{"requestType", *input.RequestType})
	}

	if input.Status != nil {
		if strings.ObjectTOString(input.RequestType) == constants.QualityControlRequestTypeDataDisCrepancy || strings.ObjectTOString(input.RequestType) == constants.QualityControlRequestTypeDataDisCrepancy {
			setDoc = append(setDoc, bson.E{"status", constants.QualityControlRequestStatusApproved})
		}
		setDoc = append(setDoc, bson.E{"status", *input.Status})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.QualityControls.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func UpdateQualityControlForBatch(ctx context.Context, id string, input QualityControl) (bool, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objId}}

	setDoc := bson.D{}

	if input.Assignee != nil {
		setDoc = append(setDoc, bson.E{"assignee", *input.Assignee})
	}

	if input.AssigneeName != nil {
		setDoc = append(setDoc, bson.E{"assigneeName", *input.AssigneeName})
	}
	if input.OrderNumber != nil {
		setDoc = append(setDoc, bson.E{"orderNumber", *input.OrderNumber})
	}
	if input.OrderAddress != nil {
		setDoc = append(setDoc, bson.E{"orderAddress", *input.OrderAddress})
	}
	if input.OrderCompany != nil {
		setDoc = append(setDoc, bson.E{"orderCompany", *input.OrderCompany})
	}

	updateDoc := bson.M{
		"$set": setDoc,
	}
	res, err := DbCollections.QualityControls.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}
