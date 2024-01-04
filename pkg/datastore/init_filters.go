package datastore

import (
	"time"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"go.mongodb.org/mongo-driver/bson"
)

func FilterByUserIds(userIds []string) *bson.M {
	return &bson.M{"userId": bson.M{"$in": userIds}}
}

func FilterByPipelineIds(ids []string) *bson.M {
	return &bson.M{"pipelineId": bson.M{"$in": ids}}
}

func FilterByAssignId(ids []string) bson.M {
	return bson.M{"assignId": bson.M{"$in": ids}}
}

func FilterByDateFromDateto(dateFrom *string, dateTo *string, userId string, userRoleList []*string) *bson.M {
	layout := "2006-01-02"
	daoFilter := bson.M{}
	daoFilter2 := bson.M{}
	if dateFrom != nil {
		varDateFrom, err := time.Parse(layout, *dateFrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		daoFilter["createdDateTime"] = daoFilter2
	}

	if dateTo != nil {
		varDateTo, err := time.Parse(layout, *dateTo)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		daoFilter["createdDateTime"] = daoFilter2
	}

	if !contains(userRoleList, constants.UserRoleAdmin) {
		daoFilter["assignee"] = userId
	}

	return &daoFilter
}

func FilterQCByDateFromDatetoAndStatuses(dateFrom *string, dateTo *string, statuses []*string, orderAssignee, orderNumber, qcAssigneeID *string, blocklistedRequestType []string) *bson.M {
	layout := "2006-01-02"
	daoFilter := bson.M{}
	//daoFilter["isAccepted"] = false
	daoFilter["$or"] = []bson.M{
		{"assignee": true},
		{"assignee": bson.M{
			"$exists": true,
		}},
	}
	daoFilter2 := bson.M{}
	if dateFrom != nil {
		varDateFrom, err := time.Parse(layout, *dateFrom)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$gte"] = pointers.PrimitiveDateTime(&varDateFrom)
		daoFilter["createdDateTime"] = daoFilter2
	}

	if dateTo != nil {
		varDateTo, err := time.Parse(layout, *dateTo)
		if err != nil {
			log.Debug("failed time parse error : %v", err)
		}
		daoFilter2["$lte"] = pointers.PrimitiveDateTime(&varDateTo)
		daoFilter["createdDateTime"] = daoFilter2
	}

	if len(statuses) > 0 {
		var statusesStr []string
		for _, v := range statuses {
			statusesStr = append(statusesStr, *v)
		}
		daoFilter["status"] = bson.M{"$in": statusesStr}
	}
	if orderAssignee != nil {
		daoFilter["orderAssignee"] = *orderAssignee
	}
	if orderNumber != nil {
		daoFilter["orderNumber"] = *orderNumber
	}

	if qcAssigneeID != nil {
		daoFilter["assignee"] = *qcAssigneeID
	}

	if len(blocklistedRequestType) > 0 {
		daoFilter["requestType"] = bson.M{"$ne": blocklistedRequestType}
	}
	return &daoFilter
}
