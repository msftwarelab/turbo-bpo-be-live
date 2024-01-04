package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func SaveHeader(ctx context.Context, name string) (string, error) {
	headersCodeRaws, _ := datastore.GetHeadersByCode(ctx, []*string{strings.ToObject(name)}, pointers.Bool(true))
	if len(headersCodeRaws) > 0 {
		return "", errs.HeaderAlreadExist
	}
	return datastore.SaveHeader(ctx, name, nil)
}

func SaveHeaderDetail(ctx context.Context, parentID string, name string) (string, error) {
	headersCodeRaws, _ := datastore.GetHeadersByCode(ctx, []*string{strings.ToObject(name)}, pointers.Bool(false))
	if len(headersCodeRaws) > 0 {
		return "", errs.HeaderDetailAlreadExist
	}
	return datastore.SaveHeader(ctx, name, &parentID)
}

func UpdateHeader(ctx context.Context, id string, name string) (bool, error) {
	headersCodeRaws, _ := datastore.GetHeadersByCode(ctx, []*string{strings.ToObject(name)}, nil)
	if len(headersCodeRaws) > 0 {
		return false, errs.HeaderAlreadExist
	}
	return datastore.UpdateHeader(ctx, id, name)
}

func DeleteHeader(ctx context.Context, id string) (bool, error) {
	return datastore.DeleteHeader(ctx, id)
}
func HeaderCode(ctx context.Context, codes []*string) ([]*models.HeaderCode, error) {
	headersCodeRaws, err := datastore.GetHeadersByCode(ctx, codes, nil)
	if err != nil {
		return nil, err
	}

	headerCodeDatas := make([]*models.HeaderCode, 0)
	mapCode := make(map[string]*string)
	for _, v := range codes {
		mapCode[*v] = nil
	}
	var ids []string
	for _, v := range headersCodeRaws {
		ids = append(ids, v.ParentId)
		mapCode[v.Name] = strings.ToObject(v.ParentId)
	}
	headerNames, err := datastore.GetHeadersByID(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, v := range codes {

		for _, headerVal := range headerNames {
			if mapCode[*v] != nil {
				if *mapCode[*v] == headerVal.ID.Hex() {
					mapCode[*v] = strings.ToObject(headerVal.Name)
				}
			}
		}
		//mapCode[*v] = nil
	}

	for key, val := range mapCode {
		headerCode := &models.HeaderCode{
			Code:  strings.ToObject(key),
			Value: val,
		}
		headerCodeDatas = append(headerCodeDatas, headerCode)

	}

	return headerCodeDatas, nil
}

func AllHeader(ctx context.Context, filter *models.HeaderFilterInput, parentId *string) (*models.HeaderResult, error) {

	defaultOffsetValue := int(0)
	defaultLimitValue := int(10)
	if filter == nil {

		filter = &models.HeaderFilterInput{Offset: &defaultOffsetValue, Limit: &defaultLimitValue}
	}
	if filter.Offset == nil {
		filter.Offset = &defaultOffsetValue
	}
	if filter.Limit == nil {
		filter.Limit = &defaultLimitValue
	}

	rawHeaders, err := datastore.SearchHeaders(ctx, *filter.Offset, *filter.Limit, filter.Name, parentId)
	if err != nil {
		log.Debug("Failed to retrieve header: %v", err)
		return nil, err
	}

	count, err := datastore.GetHeadersCount(ctx, filter.Name, parentId)
	if err != nil {
		log.Debug("Failed to retrieve count of header: %v", err)
		return nil, err
	}

	headers := make([]*models.Header, 0)
	for _, u := range rawHeaders {
		headers = append(headers, u.ToModels())
	}
	toInt := int(*count)
	return &models.HeaderResult{
		TotalCount: &toInt,
		Results:    headers,
	}, nil
}
