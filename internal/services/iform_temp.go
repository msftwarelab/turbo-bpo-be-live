package services

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/datastore"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func UpdateIformTemp(ctx context.Context, pipelineId string, input models.UpdateIformTempInput, updatedBy string) (bool, error) {
	return datastore.UpdateIformTemp(ctx, pipelineId, input, updatedBy)
}

func IformTemp(ctx context.Context, pipelineID, myID, myName string, myRole []*string) (*models.IformTemp, error) {
	//validation if role == client && pipeline.assign == myID
	if contains(myRole, constants.UserRoleClient) {
		pipelineRaw, err := datastore.GetPipelineByIdDataStore(ctx, pipelineID)
		if err != nil {
			return nil, err
		}
		if pipelineRaw != nil {
			if pipelineRaw.AssignId != nil {
				if *pipelineRaw.AssignId == myID {
					// validate if thi order already charge to credit
					if pipelineRaw.OrderNumber == nil {
						return nil, errs.InvalidPipelineData
					}
					creditLedgerFilter := datastore.FilterByOrderNumber(*pipelineRaw.OrderNumber)
					creditLedgerRaw, _ := datastore.GetCreditLedger(ctx, creditLedgerFilter)

					if creditLedgerRaw == nil {
						//Todo, Insert record on credit ledger the goal is charging of using iform
						input := datastore.CreditLedger{
							ClientName:      strings.ToObject(myName),
							ClientID:        strings.ToObject(myID),
							Type:            strings.ToObject("ORDER"),
							OrderNumber:     pipelineRaw.OrderNumber,
							OrderAddress:    pipelineRaw.Address,
							CreatedDateTime: pointers.PrimitiveDateTime(nil),
							IformCharge:     pointers.Float64(constants.IformProcessingFee),
							Amount:          pointers.Float64(-constants.IformProcessingFee),
						}
						_, err := datastore.AddCreditLedger(ctx, myID, input, myName)
						if err != nil {
							return nil, errs.CannotChargeCredit
						}

					}

				}
			}
		}
	}
	return datastore.GetIformTempByPipelineId(ctx, pipelineID)
}

func IformTempByOrderNumber(ctx context.Context, orderNumber string) (*models.Iform, error) {
	return datastore.GetIformTempByOrderNumber(ctx, orderNumber)
}
