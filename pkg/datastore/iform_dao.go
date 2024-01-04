package datastore

import (
	"context"
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/awsS3"
	errs "github.com/lonmarsDev/bpo-golang-grahpql/pkg/error"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/pdf"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/pointers"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveIform(ctx context.Context, pipelineId string, u models.UpdateIformInput, pipelinesRaw []*Pipeline, createdBy string) (string, error) {

	newCompany := &Iform{

		FormType:                              u.FormType,
		LastUpdateTime:                        pointers.PrimitiveDateTime(nil),
		PipelineID:                            strings.ToObject(pipelineId),
		ClientID:                              pipelinesRaw[0].UserId,
		TxtClient:                             u.TxtClient,
		TxtCompany:                            pipelinesRaw[0].Company,
		TxtOrderNumber:                        pipelinesRaw[0].OrderNumber,
		CmbOrderType:                          u.CmbOrderType,
		TxtAddress:                            pipelinesRaw[0].Address,
		TxtLocation:                           u.TxtLocation,
		TxtBrokerChecker:                      u.TxtBrokerChecker,
		TxtPreparerInfoAgent:                  u.TxtPreparerInfoAgent,
		TxtPreparerInfoAgentLicense:           u.TxtPreparerInfoAgentLicense,
		TxtPreparerInfoBroker:                 u.TxtPreparerInfoBroker,
		TxtPreparerInfoBrokerLicense:          u.TxtPreparerInfoBrokerLicense,
		TxtPreparerInfoAddress:                u.TxtPreparerInfoAddress,
		TxtPreparerInfoBrokerage:              u.TxtPreparerInfoBrokerage,
		TxtPreparerInfoAgentCompany:           u.TxtPreparerInfoAgentCompany,
		TxtPreparerInfoPhone:                  u.TxtPreparerInfoPhone,
		TxtPreparerInfoYearsOfExperience:      u.TxtPreparerInfoYearsOfExperience,
		TxtPreparerInfoEmail:                  u.TxtPreparerInfoEmail,
		TxtSubjectAddress:                     pipelinesRaw[0].Address,
		TxtPreparerInfoMilesAwayFromSubject:   u.TxtPreparerInfoMilesAwayFromSubject,
		TxtAgentZip:                           u.TxtAgentZip,
		TxtAgentCity:                          u.TxtAgentCity,
		TxtAgentState:                         u.TxtAgentState,
		TxtDisclaimer:                         u.TxtDisclaimer,
		CmbLocation:                           u.CmbLocation,
		TxtCounty:                             u.TxtCounty,
		TxtTrullia:                            u.TxtTrullia,
		TxtZillow:                             u.TxtZillow,
		TxtFindcompsnow:                       u.TxtFindcompsnow,
		TxtAverage:                            u.TxtAverage,
		CmbForm:                               u.CmbForm,
		CmbForm2:                              u.CmbForm2,
		TxtSaleComp1Address:                   u.TxtSaleComp1Address,
		TxtSaleComp2Address:                   u.TxtSaleComp2Address,
		TxtSaleComp3Address:                   u.TxtSaleComp3Address,
		TxtListComp1Address:                   u.TxtListComp1Address,
		TxtListComp2Address:                   u.TxtListComp2Address,
		TxtListComp3Address:                   u.TxtListComp3Address,
		TxtSubjectState:                       u.TxtSubjectState,
		TxtSaleComp1State:                     u.TxtSaleComp1State,
		TxtSaleComp2State:                     u.TxtSaleComp2State,
		TxtSaleComp3State:                     u.TxtSaleComp3State,
		TxtListComp1State:                     u.TxtListComp1State,
		TxtListComp2State:                     u.TxtListComp2State,
		TxtListComp3State:                     u.TxtListComp3State,
		TxtSubjectCity:                        u.TxtSubjectCity,
		TxtSaleComp1City:                      u.TxtSaleComp1City,
		TxtSaleComp2City:                      u.TxtSaleComp2City,
		TxtSaleComp3City:                      u.TxtSaleComp3City,
		TxtListComp1City:                      u.TxtListComp1City,
		TxtListComp2City:                      u.TxtListComp2City,
		TxtListComp3City:                      u.TxtListComp3City,
		TxtSubjectnoUnit:                      u.TxtSubjectnoUnit,
		TxtSubjectUnitNo:                      u.TxtSubjectUnitNo,
		TxtSaleComp1noUnit:                    u.TxtSaleComp1noUnit,
		TxtSaleComp1UnitNo:                    u.TxtSaleComp1UnitNo,
		TxtSaleComp2noUnit:                    u.TxtSaleComp2noUnit,
		TxtSaleComp2UnitNo:                    u.TxtSaleComp2UnitNo,
		TxtSaleComp3noUnit:                    u.TxtSaleComp3noUnit,
		TxtSaleComp3UnitNo:                    u.TxtSaleComp3UnitNo,
		TxtListComp1noUnit:                    u.TxtListComp1noUnit,
		TxtListComp1UnitNo:                    u.TxtListComp1UnitNo,
		TxtListComp2noUnit:                    u.TxtListComp2noUnit,
		TxtListComp2UnitNo:                    u.TxtListComp2UnitNo,
		TxtListComp3noUnit:                    u.TxtListComp3noUnit,
		TxtListComp3UnitNo:                    u.TxtListComp3UnitNo,
		TxtSubjectUnits:                       u.TxtSubjectUnits,
		TxtSaleComp1Units:                     u.TxtSaleComp1Units,
		TxtSaleComp2Units:                     u.TxtSaleComp2Units,
		TxtSaleComp3Units:                     u.TxtSaleComp3Units,
		TxtListComp1Units:                     u.TxtListComp1Units,
		TxtListComp2Units:                     u.TxtListComp2Units,
		TxtListComp3Units:                     u.TxtListComp3Units,
		TxtSubjectZip:                         u.TxtSubjectZip,
		TxtSaleComp1Zip:                       u.TxtSaleComp1Zip,
		TxtSaleComp2Zip:                       u.TxtSaleComp2Zip,
		TxtSaleComp3Zip:                       u.TxtSaleComp3Zip,
		TxtListComp1Zip:                       u.TxtListComp1Zip,
		TxtListComp2Zip:                       u.TxtListComp2Zip,
		TxtListComp3Zip:                       u.TxtListComp3Zip,
		TxtSubjectProximity:                   u.TxtSubjectProximity,
		TxtSaleComp1Proximity:                 u.TxtSaleComp1Proximity,
		TxtSaleComp2Proximity:                 u.TxtSaleComp2Proximity,
		TxtSaleComp3Proximity:                 u.TxtSaleComp3Proximity,
		TxtListComp1Proximity:                 u.TxtListComp1Proximity,
		TxtListComp2Proximity:                 u.TxtListComp2Proximity,
		TxtListComp3Proximity:                 u.TxtListComp3Proximity,
		TxtSubjectDataSource:                  u.TxtSubjectDataSource,
		TxtSaleComp1DataSource:                u.TxtSaleComp1DataSource,
		TxtSaleComp2DataSource:                u.TxtSaleComp2DataSource,
		TxtSaleComp3DataSource:                u.TxtSaleComp3DataSource,
		TxtListComp1DataSource:                u.TxtListComp1DataSource,
		TxtListComp2DataSource:                u.TxtListComp2DataSource,
		TxtListComp3DataSource:                u.TxtListComp3DataSource,
		TxtSubjectMLSNumber:                   u.TxtSubjectMLSNumber,
		TxtSaleComp1MLSNumber:                 u.TxtSaleComp1MLSNumber,
		TxtSaleComp2MLSNumber:                 u.TxtSaleComp2MLSNumber,
		TxtSaleComp3MLSNumber:                 u.TxtSaleComp3MLSNumber,
		TxtListComp1MLSNumber:                 u.TxtListComp1MLSNumber,
		TxtListComp2MLSNumber:                 u.TxtListComp2MLSNumber,
		TxtListComp3MLSNumber:                 u.TxtListComp3MLSNumber,
		CmbSubjectSaleType:                    u.CmbSubjectSaleType,
		CmbSaleComp1SaleType:                  u.CmbSaleComp1SaleType,
		CmbSaleComp2SaleType:                  u.CmbSaleComp2SaleType,
		CmbSaleComp3SaleType:                  u.CmbSaleComp3SaleType,
		CmbListComp1SaleType:                  u.CmbListComp1SaleType,
		CmbListComp2SaleType:                  u.CmbListComp2SaleType,
		CmbListComp3SaleType:                  u.CmbListComp3SaleType,
		CmbSubjectType:                        u.CmbSubjectType,
		CmbSaleComp1Type:                      u.CmbSaleComp1Type,
		CmbSaleComp2Type:                      u.CmbSaleComp2Type,
		CmbSaleComp3Type:                      u.CmbSaleComp3Type,
		CmbListComp1Type:                      u.CmbListComp1Type,
		CmbListComp2Type:                      u.CmbListComp2Type,
		CmbListComp3Type:                      u.CmbListComp3Type,
		CmbSubjectStyle:                       u.CmbSubjectStyle,
		CmbSaleComp1Style:                     u.CmbSaleComp1Style,
		TxtSaleComp1StyleAdjBuiltIn:           u.TxtSaleComp1StyleAdjBuiltIn,
		CmbSaleComp2Style:                     u.CmbSaleComp2Style,
		TxtSaleComp2StyleAdjBuiltIn:           u.TxtSaleComp2StyleAdjBuiltIn,
		CmbSaleComp3Style:                     u.CmbSaleComp3Style,
		TxtSaleComp3StyleAdjBuiltIn:           u.TxtSaleComp3StyleAdjBuiltIn,
		CmbListComp1Style:                     u.CmbListComp1Style,
		TxtListComp1StyleAdjBuiltIn:           u.TxtListComp1StyleAdjBuiltIn,
		CmbListComp2Style:                     u.CmbListComp2Style,
		TxtListComp2StyleAdjBuiltIn:           u.TxtListComp2StyleAdjBuiltIn,
		CmbListComp3Style:                     u.CmbListComp3Style,
		TxtListComp3StyleAdjBuiltIn:           u.TxtListComp3StyleAdjBuiltIn,
		CmbSubjectExtFinish:                   u.CmbSubjectExtFinish,
		CmbSaleComp1ExtFinish:                 u.CmbSaleComp1ExtFinish,
		TxtSaleComp1ExtFinishAdjBuiltIn:       u.TxtSaleComp1ExtFinishAdjBuiltIn,
		CmbSaleComp2ExtFinish:                 u.CmbSaleComp2ExtFinish,
		TxtSaleComp2ExtFinishAdjBuiltIn:       u.TxtSaleComp2ExtFinishAdjBuiltIn,
		CmbSaleComp3ExtFinish:                 u.CmbSaleComp3ExtFinish,
		TxtSaleComp3ExtFinishAdjBuiltIn:       u.TxtSaleComp3ExtFinishAdjBuiltIn,
		CmbListComp1ExtFinish:                 u.CmbListComp1ExtFinish,
		TxtListComp1ExtFinishAdjBuiltIn:       u.TxtListComp1ExtFinishAdjBuiltIn,
		CmbListComp2ExtFinish:                 u.CmbListComp2ExtFinish,
		TxtListComp2ExtFinishAdjBuiltIn:       u.TxtListComp2ExtFinishAdjBuiltIn,
		CmbListComp3ExtFinish:                 u.CmbListComp3ExtFinish,
		TxtListComp3ExtFinishAdjBuiltIn:       u.TxtListComp3ExtFinishAdjBuiltIn,
		CmbSubjectCondition:                   u.CmbSubjectCondition,
		CmbSaleComp1Condition:                 u.CmbSaleComp1Condition,
		TxtSaleComp1ConditionAdjBuiltIn:       u.TxtSaleComp1ConditionAdjBuiltIn,
		CmbSaleComp2Condition:                 u.CmbSaleComp2Condition,
		TxtSaleComp2ConditionAdjBuiltIn:       u.TxtSaleComp2ConditionAdjBuiltIn,
		CmbSaleComp3Condition:                 u.CmbSaleComp3Condition,
		TxtSaleComp3ConditionAdjBuiltIn:       u.TxtSaleComp3ConditionAdjBuiltIn,
		CmbListComp1Condition:                 u.CmbListComp1Condition,
		TxtListComp1ConditionAdjBuiltIn:       u.TxtListComp1ConditionAdjBuiltIn,
		CmbListComp2Condition:                 u.CmbListComp2Condition,
		TxtListComp2ConditionAdjBuiltIn:       u.TxtListComp2ConditionAdjBuiltIn,
		CmbListComp3Condition:                 u.CmbListComp3Condition,
		TxtListComp3ConditionAdjBuiltIn:       u.TxtListComp3ConditionAdjBuiltIn,
		CmbSubjectQuality:                     u.CmbSubjectQuality,
		CmbSaleComp1Quality:                   u.CmbSaleComp1Quality,
		TxtSaleComp1QualityAdjBuiltIn:         u.TxtSaleComp1QualityAdjBuiltIn,
		CmbSaleComp2Quality:                   u.CmbSaleComp2Quality,
		TxtSaleComp2QualityAdjBuiltIn:         u.TxtSaleComp2QualityAdjBuiltIn,
		CmbSaleComp3Quality:                   u.CmbSaleComp3Quality,
		TxtSaleComp3QualityAdjBuiltIn:         u.TxtSaleComp3QualityAdjBuiltIn,
		CmbListComp1Quality:                   u.CmbListComp1Quality,
		TxtListComp1QualityAdjBuiltIn:         u.TxtListComp1QualityAdjBuiltIn,
		CmbListComp2Quality:                   u.CmbListComp2Quality,
		TxtListComp2QualityAdjBuiltIn:         u.TxtListComp2QualityAdjBuiltIn,
		CmbListComp3Quality:                   u.CmbListComp3Quality,
		TxtListComp3QualityAdjBuiltIn:         u.TxtListComp3QualityAdjBuiltIn,
		CmbSubjectView:                        u.CmbSubjectView,
		CmbSaleComp1View:                      u.CmbSaleComp1View,
		TxtSaleComp1ViewAdjBuiltIn:            u.TxtSaleComp1ViewAdjBuiltIn,
		CmbSaleComp2View:                      u.CmbSaleComp2View,
		TxtSaleComp2ViewAdjBuiltIn:            u.TxtSaleComp2ViewAdjBuiltIn,
		CmbSaleComp3View:                      u.CmbSaleComp3View,
		TxtSaleComp3ViewAdjBuiltIn:            u.TxtSaleComp3ViewAdjBuiltIn,
		CmbListComp1View:                      u.CmbListComp1View,
		TxtListComp1ViewAdjBuiltIn:            u.TxtListComp1ViewAdjBuiltIn,
		CmbListComp2View:                      u.CmbListComp2View,
		TxtListComp2ViewAdjBuiltIn:            u.TxtListComp2ViewAdjBuiltIn,
		CmbListComp3View:                      u.CmbListComp3View,
		TxtListComp3ViewAdjBuiltIn:            u.TxtListComp3ViewAdjBuiltIn,
		TxtSubjectSubdivision:                 u.TxtSubjectSubdivision,
		TxtSaleComp1Subdivision:               u.TxtSaleComp1Subdivision,
		TxtSaleComp2Subdivision:               u.TxtSaleComp2Subdivision,
		TxtSaleComp3Subdivision:               u.TxtSaleComp3Subdivision,
		TxtListComp1Subdivision:               u.TxtListComp1Subdivision,
		TxtListComp2Subdivision:               u.TxtListComp2Subdivision,
		TxtListComp3Subdivision:               u.TxtListComp3Subdivision,
		TxtSubjectHOAFee:                      u.TxtSubjectHOAFee,
		TxtSaleComp1HOAFee:                    u.TxtSaleComp1HOAFee,
		TxtSaleComp2HOAFee:                    u.TxtSaleComp2HOAFee,
		TxtSaleComp3HOAFee:                    u.TxtSaleComp3HOAFee,
		TxtListComp1HOAFee:                    u.TxtListComp1HOAFee,
		TxtListComp2HOAFee:                    u.TxtListComp2HOAFee,
		TxtListComp3HOAFee:                    u.TxtListComp3HOAFee,
		TxtSubjectTotalRooms:                  u.TxtSubjectTotalRooms,
		TxtSaleComp1TotalRooms:                u.TxtSaleComp1TotalRooms,
		TxtSaleComp1TotalRoomsAdjBuiltIn:      u.TxtSaleComp1TotalRoomsAdjBuiltIn,
		TxtSaleComp2TotalRooms:                u.TxtSaleComp2TotalRooms,
		TxtSaleComp2TotalRoomsAdjBuiltIn:      u.TxtSaleComp2TotalRoomsAdjBuiltIn,
		TxtSaleComp3TotalRooms:                u.TxtSaleComp3TotalRooms,
		TxtSaleComp3TotalRoomsAdjBuiltIn:      u.TxtSaleComp3TotalRoomsAdjBuiltIn,
		TxtListComp1TotalRooms:                u.TxtListComp1TotalRooms,
		TxtListComp1TotalRoomsAdjBuiltIn:      u.TxtListComp1TotalRoomsAdjBuiltIn,
		TxtListComp2TotalRooms:                u.TxtListComp2TotalRooms,
		TxtListComp2TotalRoomsAdjBuiltIn:      u.TxtListComp2TotalRoomsAdjBuiltIn,
		TxtListComp3TotalRooms:                u.TxtListComp3TotalRooms,
		TxtListComp3TotalRoomsAdjBuiltIn:      u.TxtListComp3TotalRoomsAdjBuiltIn,
		TxtSubjectBedrooms:                    u.TxtSubjectBedrooms,
		TxtSaleComp1Bedrooms:                  u.TxtSaleComp1Bedrooms,
		TxtSaleComp1BedroomsAdjBuiltIn:        u.TxtSaleComp1BedroomsAdjBuiltIn,
		TxtSaleComp2Bedrooms:                  u.TxtSaleComp2Bedrooms,
		TxtSaleComp2BedroomsAdjBuiltIn:        u.TxtSaleComp2BedroomsAdjBuiltIn,
		TxtSaleComp3Bedrooms:                  u.TxtSaleComp3Bedrooms,
		TxtSaleComp3BedroomsAdjBuiltIn:        u.TxtSaleComp3BedroomsAdjBuiltIn,
		TxtListComp1Bedrooms:                  u.TxtListComp1Bedrooms,
		TxtListComp1BedroomsAdjBuiltIn:        u.TxtListComp1BedroomsAdjBuiltIn,
		TxtListComp2Bedrooms:                  u.TxtListComp2Bedrooms,
		TxtListComp2BedroomsAdjBuiltIn:        u.TxtListComp2BedroomsAdjBuiltIn,
		TxtListComp3Bedrooms:                  u.TxtListComp3Bedrooms,
		TxtListComp3BedroomsAdjBuiltIn:        u.TxtListComp3BedroomsAdjBuiltIn,
		TxtSubjectFullBaths:                   u.TxtSubjectFullBaths,
		TxtSaleComp1FullBaths:                 u.TxtSaleComp1FullBaths,
		TxtSaleComp1FullBathsAdjBuiltIn:       u.TxtSaleComp1FullBathsAdjBuiltIn,
		TxtSaleComp2FullBaths:                 u.TxtSaleComp2FullBaths,
		TxtSaleComp2FullBathsAdjBuiltIn:       u.TxtSaleComp2FullBathsAdjBuiltIn,
		TxtSaleComp3FullBaths:                 u.TxtSaleComp3FullBaths,
		TxtSaleComp3FullBathsAdjBuiltIn:       u.TxtSaleComp3FullBathsAdjBuiltIn,
		TxtListComp1FullBaths:                 u.TxtListComp1FullBaths,
		TxtListComp1FullBathsAdjBuiltIn:       u.TxtListComp1FullBathsAdjBuiltIn,
		TxtListComp2FullBaths:                 u.TxtListComp2FullBaths,
		TxtListComp2FullBathsAdjBuiltIn:       u.TxtListComp2FullBathsAdjBuiltIn,
		TxtListComp3FullBaths:                 u.TxtListComp3FullBaths,
		TxtListComp3FullBathsAdjBuiltIn:       u.TxtListComp3FullBathsAdjBuiltIn,
		TxtSubjectHalfBaths:                   u.TxtSubjectHalfBaths,
		TxtSaleComp1HalfBaths:                 u.TxtSaleComp1HalfBaths,
		TxtSaleComp1HalfBathsAdjBuiltIn:       u.TxtSaleComp1HalfBathsAdjBuiltIn,
		TxtSaleComp2HalfBaths:                 u.TxtSaleComp2HalfBaths,
		TxtSaleComp2HalfBathsAdjBuiltIn:       u.TxtSaleComp2HalfBathsAdjBuiltIn,
		TxtSaleComp3HalfBaths:                 u.TxtSaleComp3HalfBaths,
		TxtSaleComp3HalfBathsAdjBuiltIn:       u.TxtSaleComp3HalfBathsAdjBuiltIn,
		TxtListComp1HalfBaths:                 u.TxtListComp1HalfBaths,
		TxtListComp1HalfBathsAdjBuiltIn:       u.TxtListComp1HalfBathsAdjBuiltIn,
		TxtListComp2HalfBaths:                 u.TxtListComp2HalfBaths,
		TxtListComp2HalfBathsAdjBuiltIn:       u.TxtListComp2HalfBathsAdjBuiltIn,
		TxtListComp3HalfBaths:                 u.TxtListComp3HalfBaths,
		TxtListComp3HalfBathsAdjBuiltIn:       u.TxtListComp3HalfBathsAdjBuiltIn,
		TxtSubjectGla:                         u.TxtSubjectGla,
		TxtSaleComp1gla:                       u.TxtSaleComp1gla,
		TxtSaleComp1GLAAdjBuiltIn:             u.TxtSaleComp1GLAAdjBuiltIn,
		TxtSaleComp2gla:                       u.TxtSaleComp2gla,
		TxtSaleComp2GLAAdjBuiltIn:             u.TxtSaleComp2GLAAdjBuiltIn,
		TxtSaleComp3gla:                       u.TxtSaleComp3gla,
		TxtSaleComp3GLAAdjBuiltIn:             u.TxtSaleComp3GLAAdjBuiltIn,
		TxtListComp1gla:                       u.TxtListComp1gla,
		TxtListComp1GLAAdjBuiltIn:             u.TxtListComp1GLAAdjBuiltIn,
		TxtListComp2gla:                       u.TxtListComp2gla,
		TxtListComp2GLAAdjBuiltIn:             u.TxtListComp2GLAAdjBuiltIn,
		TxtListComp3gla:                       u.TxtListComp3gla,
		TxtListComp3GLAAdjBuiltIn:             u.TxtListComp3GLAAdjBuiltIn,
		TxtSubjectYearBuilt:                   u.TxtSubjectYearBuilt,
		TxtSaleComp1YearBuilt:                 u.TxtSaleComp1YearBuilt,
		TxtSaleComp1YearBuiltAdjBuiltIn:       u.TxtSaleComp1YearBuiltAdjBuiltIn,
		TxtSaleComp2YearBuilt:                 u.TxtSaleComp2YearBuilt,
		TxtSaleComp2YearBuiltAdjBuiltIn:       u.TxtSaleComp2YearBuiltAdjBuiltIn,
		TxtSaleComp3YearBuilt:                 u.TxtSaleComp3YearBuilt,
		TxtSaleComp3YearBuiltAdjBuiltIn:       u.TxtSaleComp3YearBuiltAdjBuiltIn,
		TxtListComp1YearBuilt:                 u.TxtListComp1YearBuilt,
		TxtListComp1YearBuiltAdjBuiltIn:       u.TxtListComp1YearBuiltAdjBuiltIn,
		TxtListComp2YearBuilt:                 u.TxtListComp2YearBuilt,
		TxtListComp2YearBuiltAdjBuiltIn:       u.TxtListComp2YearBuiltAdjBuiltIn,
		TxtListComp3YearBuilt:                 u.TxtListComp3YearBuilt,
		TxtListComp3YearBuiltAdjBuiltIn:       u.TxtListComp3YearBuiltAdjBuiltIn,
		TxtSubjectAge:                         u.TxtSubjectAge,
		TxtSaleComp1Age:                       u.TxtSaleComp1Age,
		TxtSaleComp2Age:                       u.TxtSaleComp2Age,
		TxtSaleComp3Age:                       u.TxtSaleComp3Age,
		TxtListComp1Age:                       u.TxtListComp1Age,
		TxtListComp2Age:                       u.TxtListComp2Age,
		TxtListComp3Age:                       u.TxtListComp3Age,
		TxtSubjectAcres:                       u.TxtSubjectAcres,
		TxtSaleComp1Acres:                     u.TxtSaleComp1Acres,
		TxtSaleComp1AcresAdjBuiltIn:           u.TxtSaleComp1AcresAdjBuiltIn,
		TxtSaleComp2Acres:                     u.TxtSaleComp2Acres,
		TxtSaleComp2AcresAdjBuiltIn:           u.TxtSaleComp2AcresAdjBuiltIn,
		TxtSaleComp3Acres:                     u.TxtSaleComp3Acres,
		TxtSaleComp3AcresAdjBuiltIn:           u.TxtSaleComp3AcresAdjBuiltIn,
		TxtListComp1Acres:                     u.TxtListComp1Acres,
		TxtListComp1AcresAdjBuiltIn:           u.TxtListComp1AcresAdjBuiltIn,
		TxtListComp2Acres:                     u.TxtListComp2Acres,
		TxtListComp2AcresAdjBuiltIn:           u.TxtListComp2AcresAdjBuiltIn,
		TxtListComp3Acres:                     u.TxtListComp3Acres,
		TxtListComp3AcresAdjBuiltIn:           u.TxtListComp3AcresAdjBuiltIn,
		TxtSubjectSquareFeet:                  u.TxtSubjectSquareFeet,
		TxtSaleComp1SquareFeet:                u.TxtSaleComp1SquareFeet,
		TxtSaleComp2SquareFeet:                u.TxtSaleComp2SquareFeet,
		TxtSaleComp3SquareFeet:                u.TxtSaleComp3SquareFeet,
		TxtListComp1SquareFeet:                u.TxtListComp1SquareFeet,
		TxtListComp2SquareFeet:                u.TxtListComp2SquareFeet,
		TxtListComp3SquareFeet:                u.TxtListComp3SquareFeet,
		CmbSubjectGarage:                      u.CmbSubjectGarage,
		CmbSaleComp1Garage:                    u.CmbSaleComp1Garage,
		TxtSaleComp1GarageAdjBuiltIn:          u.TxtSaleComp1GarageAdjBuiltIn,
		CmbSaleComp2Garage:                    u.CmbSaleComp2Garage,
		TxtSaleComp2GarageAdjBuiltIn:          u.TxtSaleComp2GarageAdjBuiltIn,
		CmbSaleComp3Garage:                    u.CmbSaleComp3Garage,
		TxtSaleComp3GarageAdjBuiltIn:          u.TxtSaleComp3GarageAdjBuiltIn,
		CmbListComp1Garage:                    u.CmbListComp1Garage,
		TxtListComp1GarageAdjBuiltIn:          u.TxtListComp1GarageAdjBuiltIn,
		CmbListComp2Garage:                    u.CmbListComp2Garage,
		TxtListComp2GarageAdjBuiltIn:          u.TxtListComp2GarageAdjBuiltIn,
		CmbListComp3Garage:                    u.CmbListComp3Garage,
		TxtListComp3GarageAdjBuiltIn:          u.TxtListComp3GarageAdjBuiltIn,
		CmbSubjectPool:                        u.CmbSubjectPool,
		CmbSaleComp1Pool:                      u.CmbSaleComp1Pool,
		TxtSaleComp1PoolAdjBuiltIn:            u.TxtSaleComp1PoolAdjBuiltIn,
		CmbSaleComp2Pool:                      u.CmbSaleComp2Pool,
		TxtSaleComp2PoolAdjBuiltIn:            u.TxtSaleComp2PoolAdjBuiltIn,
		CmbSaleComp3Pool:                      u.CmbSaleComp3Pool,
		TxtSaleComp3PoolAdjBuiltIn:            u.TxtSaleComp3PoolAdjBuiltIn,
		CmbListComp1Pool:                      u.CmbListComp1Pool,
		TxtListComp1PoolAdjBuiltIn:            u.TxtListComp1PoolAdjBuiltIn,
		CmbListComp2Pool:                      u.CmbListComp2Pool,
		TxtListComp2PoolAdjBuiltIn:            u.TxtListComp2PoolAdjBuiltIn,
		CmbListComp3Pool:                      u.CmbListComp3Pool,
		TxtListComp3PoolAdjBuiltIn:            u.TxtListComp3PoolAdjBuiltIn,
		CmbSubjectPorchPatioDeck:              u.CmbSubjectPorchPatioDeck,
		CmbSaleComp1PorchPatioDeck:            u.CmbSaleComp1PorchPatioDeck,
		TxtSaleComp1PorchPatioDeckAdjBuiltIn:  u.TxtSaleComp1PorchPatioDeckAdjBuiltIn,
		CmbSaleComp2PorchPatioDeck:            u.CmbSaleComp2PorchPatioDeck,
		TxtSaleComp2PorchPatioDeckAdjBuiltIn:  u.TxtSaleComp2PorchPatioDeckAdjBuiltIn,
		CmbSaleComp3PorchPatioDeck:            u.CmbSaleComp3PorchPatioDeck,
		TxtSaleComp3PorchPatioDeckAdjBuiltIn:  u.TxtSaleComp3PorchPatioDeckAdjBuiltIn,
		CmbListComp1PorchPatioDeck:            u.CmbListComp1PorchPatioDeck,
		TxtListComp1PorchPatioDeckAdjBuiltIn:  u.TxtListComp1PorchPatioDeckAdjBuiltIn,
		CmbListComp2PorchPatioDeck:            u.CmbListComp2PorchPatioDeck,
		TxtListComp2PorchPatioDeckAdjBuiltIn:  u.TxtListComp2PorchPatioDeckAdjBuiltIn,
		CmbListComp3PorchPatioDeck:            u.CmbListComp3PorchPatioDeck,
		TxtListComp3PorchPatioDeckAdjBuiltIn:  u.TxtListComp3PorchPatioDeckAdjBuiltIn,
		CmbSubjectFireplace:                   u.CmbSubjectFireplace,
		CmbSaleComp1Fireplace:                 u.CmbSaleComp1Fireplace,
		TxtSaleComp1FireplaceAdjBuiltIn:       u.TxtSaleComp1FireplaceAdjBuiltIn,
		CmbSaleComp2Fireplace:                 u.CmbSaleComp2Fireplace,
		TxtSaleComp2FireplaceAdjBuiltIn:       u.TxtSaleComp2FireplaceAdjBuiltIn,
		CmbSaleComp3Fireplace:                 u.CmbSaleComp3Fireplace,
		TxtSaleComp3FireplaceAdjBuiltIn:       u.TxtSaleComp3FireplaceAdjBuiltIn,
		CmbListComp1Fireplace:                 u.CmbListComp1Fireplace,
		TxtListComp1FireplaceAdjBuiltIn:       u.TxtListComp1FireplaceAdjBuiltIn,
		CmbListComp2Fireplace:                 u.CmbListComp2Fireplace,
		TxtListComp2FireplaceAdjBuiltIn:       u.TxtListComp2FireplaceAdjBuiltIn,
		CmbListComp3Fireplace:                 u.CmbListComp3Fireplace,
		TxtListComp3FireplaceAdjBuiltIn:       u.TxtListComp3FireplaceAdjBuiltIn,
		CmbSubjectBasement:                    u.CmbSubjectBasement,
		CmbSaleComp1Basement:                  u.CmbSaleComp1Basement,
		TxtSaleComp1BasementAdjBuiltIn:        u.TxtSaleComp1BasementAdjBuiltIn,
		CmbSaleComp2Basement:                  u.CmbSaleComp2Basement,
		TxtSaleComp2BasementAdjBuiltIn:        u.TxtSaleComp2BasementAdjBuiltIn,
		CmbSaleComp3Basement:                  u.CmbSaleComp3Basement,
		TxtSaleComp3BasementAdjBuiltIn:        u.TxtSaleComp3BasementAdjBuiltIn,
		CmbListComp1Basement:                  u.CmbListComp1Basement,
		TxtListComp1BasementAdjBuiltIn:        u.TxtListComp1BasementAdjBuiltIn,
		CmbListComp2Basement:                  u.CmbListComp2Basement,
		TxtListComp2BasementAdjBuiltIn:        u.TxtListComp2BasementAdjBuiltIn,
		CmbListComp3Basement:                  u.CmbListComp3Basement,
		TxtListComp3BasementAdjBuiltIn:        u.TxtListComp3BasementAdjBuiltIn,
		CmbSubjectIsFinished:                  u.CmbSubjectIsFinished,
		CmbSaleComp1IsFinished:                u.CmbSaleComp1IsFinished,
		TxtSaleComp1IsFinishedAdjBuiltIn:      u.TxtSaleComp1IsFinishedAdjBuiltIn,
		CmbSaleComp2IsFinished:                u.CmbSaleComp2IsFinished,
		TxtSaleComp2IsFinishedAdjBuiltIn:      u.TxtSaleComp2IsFinishedAdjBuiltIn,
		CmbSaleComp3IsFinished:                u.CmbSaleComp3IsFinished,
		TxtSaleComp3IsFinishedAdjBuiltIn:      u.TxtSaleComp3IsFinishedAdjBuiltIn,
		CmbListComp1IsFinished:                u.CmbListComp1IsFinished,
		TxtListComp1IsFinishedAdjBuiltIn:      u.TxtListComp1IsFinishedAdjBuiltIn,
		CmbListComp2IsFinished:                u.CmbListComp2IsFinished,
		TxtListComp2IsFinishedAdjBuiltIn:      u.TxtListComp2IsFinishedAdjBuiltIn,
		CmbListComp3IsFinished:                u.CmbListComp3IsFinished,
		TxtListComp3IsFinishedAdjBuiltIn:      u.TxtListComp3IsFinishedAdjBuiltIn,
		CmbSubjectPercentFinished:             u.CmbSubjectPercentFinished,
		CmbSaleComp1PercentFinished:           u.CmbSaleComp1PercentFinished,
		TxtSaleComp1PercentFinishedAdjBuiltIn: u.TxtSaleComp1PercentFinishedAdjBuiltIn,
		CmbSaleComp2PercentFinished:           u.CmbSaleComp2PercentFinished,
		TxtSaleComp2PercentFinishedAdjBuiltIn: u.TxtSaleComp2PercentFinishedAdjBuiltIn,
		CmbSaleComp3PercentFinished:           u.CmbSaleComp3PercentFinished,
		TxtSaleComp3PercentFinishedAdjBuiltIn: u.TxtSaleComp3PercentFinishedAdjBuiltIn,
		CmbListComp1PercentFinished:           u.CmbListComp1PercentFinished,
		TxtListComp1PercentFinishedAdjBuiltIn: u.TxtListComp1PercentFinishedAdjBuiltIn,
		CmbListComp2PercentFinished:           u.CmbListComp2PercentFinished,
		TxtListComp2PercentFinishedAdjBuiltIn: u.TxtListComp2PercentFinishedAdjBuiltIn,
		CmbListComp3PercentFinished:           u.CmbListComp3PercentFinished,
		TxtListComp3PercentFinishedAdjBuiltIn: u.TxtListComp3PercentFinishedAdjBuiltIn,
		TxtSubjectBasementSqFt:                u.TxtSubjectBasementSqFt,
		TxtSaleComp1BasementSqFt:              u.TxtSaleComp1BasementSqFt,
		TxtSaleComp1BasementSqFtAdjBuiltIn:    u.TxtSaleComp1BasementSqFtAdjBuiltIn,
		TxtSaleComp2BasementSqFt:              u.TxtSaleComp2BasementSqFt,
		TxtSaleComp2BasementSqFtAdjBuiltIn:    u.TxtSaleComp2BasementSqFtAdjBuiltIn,
		TxtSaleComp3BasementSqFt:              u.TxtSaleComp3BasementSqFt,
		TxtSaleComp3BasementSqFtAdjBuiltIn:    u.TxtSaleComp3BasementSqFtAdjBuiltIn,
		TxtListComp1BasementSqFt:              u.TxtListComp1BasementSqFt,
		TxtListComp1BasementSqFtAdjBuiltIn:    u.TxtListComp1BasementSqFtAdjBuiltIn,
		TxtListComp2BasementSqFt:              u.TxtListComp2BasementSqFt,
		TxtListComp2BasementSqFtAdjBuiltIn:    u.TxtListComp2BasementSqFtAdjBuiltIn,
		TxtListComp3BasementSqFt:              u.TxtListComp3BasementSqFt,
		TxtListComp3BasementSqFtAdjBuiltIn:    u.TxtListComp3BasementSqFtAdjBuiltIn,
		TxtSubjectOriginalListDate:            u.TxtSubjectOriginalListDate,
		TxtSaleComp1OriginalListDate:          u.TxtSaleComp1OriginalListDate,
		TxtSaleComp2OriginalListDate:          u.TxtSaleComp2OriginalListDate,
		TxtSaleComp3OriginalListDate:          u.TxtSaleComp3OriginalListDate,
		TxtListComp1OriginalListDate:          u.TxtListComp1OriginalListDate,
		TxtListComp2OriginalListDate:          u.TxtListComp2OriginalListDate,
		TxtListComp3OriginalListDate:          u.TxtListComp3OriginalListDate,
		TxtSubjectCurrentListDate:             u.TxtSubjectCurrentListDate,
		TxtSaleComp1CurrentListDate:           u.TxtSaleComp1CurrentListDate,
		TxtSaleComp2CurrentListDate:           u.TxtSaleComp2CurrentListDate,
		TxtSaleComp3CurrentListDate:           u.TxtSaleComp3CurrentListDate,
		TxtListComp1CurrentListDate:           u.TxtListComp1CurrentListDate,
		TxtListComp2CurrentListDate:           u.TxtListComp2CurrentListDate,
		TxtListComp3CurrentListDate:           u.TxtListComp3CurrentListDate,
		TxtSubjectOriginalListPrice:           u.TxtSubjectOriginalListPrice,
		TxtSaleComp1OriginalListPrice:         u.TxtSaleComp1OriginalListPrice,
		TxtSaleComp2OriginalListPrice:         u.TxtSaleComp2OriginalListPrice,
		TxtSaleComp3OriginalListPrice:         u.TxtSaleComp3OriginalListPrice,
		TxtListComp1OriginalListPrice:         u.TxtListComp1OriginalListPrice,
		TxtListComp2OriginalListPrice:         u.TxtListComp2OriginalListPrice,
		TxtListComp3OriginalListPrice:         u.TxtListComp3OriginalListPrice,
		TxtSubjectListPrice:                   u.TxtSubjectListPrice,
		TxtSaleComp1ListPrice:                 u.TxtSaleComp1ListPrice,
		TxtSaleComp2ListPrice:                 u.TxtSaleComp2ListPrice,
		TxtSaleComp3ListPrice:                 u.TxtSaleComp3ListPrice,
		TxtListComp1ListPrice:                 u.TxtListComp1ListPrice,
		TxtListComp2ListPrice:                 u.TxtListComp2ListPrice,
		TxtListComp3ListPrice:                 u.TxtListComp3ListPrice,
		TxtSubjectSalePrice:                   u.TxtSubjectSalePrice,
		TxtSaleComp1SalePrice:                 u.TxtSaleComp1SalePrice,
		TxtSaleComp2SalePrice:                 u.TxtSaleComp2SalePrice,
		TxtSaleComp3SalePrice:                 u.TxtSaleComp3SalePrice,
		TxtSubjectSaleDate:                    u.TxtSubjectSaleDate,
		TxtSaleComp1SaleDate:                  u.TxtSaleComp1SaleDate,
		TxtSaleComp2SaleDate:                  u.TxtSaleComp2SaleDate,
		TxtSaleComp3SaleDate:                  u.TxtSaleComp3SaleDate,
		CmbSubjectFinancing:                   u.CmbSubjectFinancing,
		CmbSaleComp1Financing:                 u.CmbSaleComp1Financing,
		CmbSaleComp2Financing:                 u.CmbSaleComp2Financing,
		CmbSaleComp3Financing:                 u.CmbSaleComp3Financing,
		CmbListComp1Financing:                 u.CmbListComp1Financing,
		CmbListComp2Financing:                 u.CmbListComp2Financing,
		CmbListComp3Financing:                 u.CmbListComp3Financing,
		TxtSubjectDom:                         u.TxtSubjectDom,
		TxtSaleComp1dom:                       u.TxtSaleComp1dom,
		TxtSaleComp2dom:                       u.TxtSaleComp2dom,
		TxtSaleComp3dom:                       u.TxtSaleComp3dom,
		TxtListComp1dom:                       u.TxtListComp1dom,
		TxtListComp2dom:                       u.TxtListComp2dom,
		TxtListComp3dom:                       u.TxtListComp3dom,
		TxtSubjectPricePerSqFt:                u.TxtSubjectPricePerSqFt,
		TxtSaleComp1PricePerSqFt:              u.TxtSaleComp1PricePerSqFt,
		TxtSaleComp2PricePerSqFt:              u.TxtSaleComp2PricePerSqFt,
		TxtSaleComp3PricePerSqFt:              u.TxtSaleComp3PricePerSqFt,
		TxtListComp1PricePerSqFt:              u.TxtListComp1PricePerSqFt,
		TxtListComp2PricePerSqFt:              u.TxtListComp2PricePerSqFt,
		TxtListComp3PricePerSqFt:              u.TxtListComp3PricePerSqFt,
		TxtSubjectAdjustments:                 u.TxtSubjectAdjustments,
		TxtSaleComp1Adjustments:               u.TxtSaleComp1Adjustments,
		TxtSaleComp2Adjustments:               u.TxtSaleComp2Adjustments,
		TxtSaleComp3Adjustments:               u.TxtSaleComp3Adjustments,
		TxtListComp1Adjustments:               u.TxtListComp1Adjustments,
		TxtListComp2Adjustments:               u.TxtListComp2Adjustments,
		TxtListComp3Adjustments:               u.TxtListComp3Adjustments,
		TxtSubjectCompTotals:                  u.TxtSubjectCompTotals,
		TxtSaleComp1CompTotals:                u.TxtSaleComp1CompTotals,
		TxtSaleComp2CompTotals:                u.TxtSaleComp2CompTotals,
		TxtSaleComp3CompTotals:                u.TxtSaleComp3CompTotals,
		TxtListComp1CompTotals:                u.TxtListComp1CompTotals,
		TxtListComp2CompTotals:                u.TxtListComp2CompTotals,
		TxtListComp3CompTotals:                u.TxtListComp3CompTotals,
		CmbListComp1CommentType:               u.CmbListComp1CommentType,
		TxtListComp1ComparableComments:        u.TxtListComp1ComparableComments,
		TxtListComp1FormatAdjustments:         u.TxtListComp1FormatAdjustments,
		TxtListComp1MLSComments:               u.TxtListComp1MLSComments,
		CmbListComp2CommentType:               u.CmbListComp2CommentType,
		TxtListComp2ComparableComments:        u.TxtListComp2ComparableComments,
		TxtListComp2FormatAdjustments:         u.TxtListComp2FormatAdjustments,
		TxtListComp2MLSComments:               u.TxtListComp2MLSComments,
		CmbListComp3CommentType:               u.CmbListComp3CommentType,
		TxtListComp3ComparableComments:        u.TxtListComp3ComparableComments,
		TxtListComp3FormatAdjustments:         u.TxtListComp3FormatAdjustments,
		TxtListComp3MLSComments:               u.TxtListComp3MLSComments,
		CmbSaleComp1CommentType:               u.CmbSaleComp1CommentType,
		TxtSaleComp1ComparableComments:        u.TxtSaleComp1ComparableComments,
		TxtSaleComp1FormatAdjustments:         u.TxtSaleComp1FormatAdjustments,
		TxtSaleComp1MLSComments:               u.TxtSaleComp1MLSComments,
		CmbSaleComp2CommentType:               u.CmbSaleComp2CommentType,
		TxtSaleComp2ComparableComments:        u.TxtSaleComp2ComparableComments,
		TxtSaleComp2FormatAdjustments:         u.TxtSaleComp2FormatAdjustments,
		TxtSaleComp2MLSComments:               u.TxtSaleComp2MLSComments,
		CmbSaleComp3CommentType:               u.CmbSaleComp3CommentType,
		TxtSaleComp3ComparableComments:        u.TxtSaleComp3ComparableComments,
		TxtSaleComp3FormatAdjustments:         u.TxtSaleComp3FormatAdjustments,
		TxtSaleComp3MLSComments:               u.TxtSaleComp3MLSComments,
		CmbNeighborhoodTrend:                  u.CmbNeighborhoodTrend,
		TxtMonthlyPecent:                      u.TxtMonthlyPecent,
		TxtEstimatedRent:                      u.TxtEstimatedRent,
		TxtEstimatedDaysOnMarket:              u.TxtEstimatedDaysOnMarket,
		TxtNoBoarded:                          u.TxtNoBoarded,
		TxtNoOfActive:                         u.TxtNoOfActive,
		Txt6MonthPecent:                       u.Txt6MonthPecent,
		TxtAnnualPecent:                       u.TxtAnnualPecent,
		TxtListings:                           u.TxtListings,
		CmbSupply:                             u.CmbSupply,
		TxtListingsMinValue:                   u.TxtListingsMinValue,
		TxtListingsRange1:                     u.TxtListingsRange1,
		TxtListingsMedValue:                   u.TxtListingsMedValue,
		TxtListingsMaxValue:                   u.TxtListingsMaxValue,
		TxtListingsRange2:                     u.TxtListingsRange2,
		TxtListingsDom:                        u.TxtListingsDom,
		TxtListingsDOMRange1:                  u.TxtListingsDOMRange1,
		TxtListingsDOMRange2:                  u.TxtListingsDOMRange2,
		CmbREOTrend:                           u.CmbREOTrend,
		TxtNoOfFm:                             u.TxtNoOfFm,
		TxtNoOfSs:                             u.TxtNoOfSs,
		TxtNoOfReo:                            u.TxtNoOfReo,
		TxtNoOfDistressed:                     u.TxtNoOfDistressed,
		TxtSales:                              u.TxtSales,
		CmbDemand:                             u.CmbDemand,
		TxtSalesRange1:                        u.TxtSalesRange1,
		TxtSalesMedValue:                      u.TxtSalesMedValue,
		TxtSalesRange2:                        u.TxtSalesRange2,
		TxtSalesDom:                           u.TxtSalesDom,
		TxtSalesDOMRange1:                     u.TxtSalesDOMRange1,
		TxtSalesDOMRange2:                     u.TxtSalesDOMRange2,
		TxtZillowNeighborhoodTrend:            u.TxtZillowNeighborhoodTrend,
		TxtNeighborhoodTrendComments:          u.TxtNeighborhoodTrendComments,
		TxtTotalListings:                      u.TxtTotalListings,
		TxtTotalSales:                         u.TxtTotalSales,
		TxtNoOfREOListings:                    u.TxtNoOfREOListings,
		TxtNoOfSSListings:                     u.TxtNoOfSSListings,
		TxtNoOfREOSales:                       u.TxtNoOfREOSales,
		TxtNoOfSSSales:                        u.TxtNoOfSSSales,
		TxtTaxID:                              u.TxtTaxID,
		TxtLastSaleDate:                       u.TxtLastSaleDate,
		TxtLastSalePrice:                      u.TxtLastSalePrice,
		CmbIsListed:                           u.CmbIsListed,
		TxtOwnerOccupied:                      u.TxtOwnerOccupied,
		TxtRenterOccupied:                     u.TxtRenterOccupied,
		TxtMarketRent:                         u.TxtMarketRent,
		TxtNoOfRentals:                        u.TxtNoOfRentals,
		TxtTypicalDom:                         u.TxtTypicalDom,
		TxtNoRentHomes:                        u.TxtNoRentHomes,
		TxtTypicalRentalRates:                 u.TxtTypicalRentalRates,
		AdjustmentPrice:                       u.AdjustmentPrice,
		TxtCalculatedGla:                      u.TxtCalculatedGla,
		TxtCalculatedAge:                      u.TxtCalculatedAge,
		TxtCalculatedSaleDates:                u.TxtCalculatedSaleDates,
		TxtCalculatedProximity:                u.TxtCalculatedProximity,
		TxtCalculatedStyle:                    u.TxtCalculatedStyle,
		TxtCalculatedMonthsSupply:             u.TxtCalculatedMonthsSupply,
		TxtCalculatedProxim:                   u.TxtCalculatedProxim,
		TxtCalculatedGLAs:                     u.TxtCalculatedGLAs,
		TxtCalculatedAges:                     u.TxtCalculatedAges,
		TxtCalculatedCond:                     u.TxtCalculatedCond,
		TxtCalculatedView:                     u.TxtCalculatedView,
		TxtCalculatedStyle1:                   u.TxtCalculatedStyle1,
		TxtCalculatedLots:                     u.TxtCalculatedLots,
		TxtCalculatedBeds:                     u.TxtCalculatedBeds,
		TxtCalculatedBath:                     u.TxtCalculatedBath,
		Rdbresaletext:                         u.Rdbresaletext,
		Rdbmarketedtext:                       u.Rdbmarketedtext,
		Txtpmi:                                u.Txtpmi,
		TxtOtherComments:                      u.TxtOtherComments,
		Txtcbnew:                              u.Txtcbnew,
		Txtcbold:                              u.Txtcbold,
		Txtcbstyle:                            u.Txtcbstyle,
		Txtcblot:                              u.Txtcblot,
		Txtcbview:                             u.Txtcbview,
		Txtcbdamage:                           u.Txtcbdamage,
		Txtcbupgrade:                          u.Txtcbupgrade,
		Txtcbinfluence:                        u.Txtcbinfluence,
		TxtSubjectComments:                    u.TxtSubjectComments,
		TxtNeighborhoodComments:               u.TxtNeighborhoodComments,
		TxtNeighborhoodTrend:                  u.TxtNeighborhoodTrend,
		TxtValidation1:                        u.TxtValidation1,
		TxtUniqueComments:                     u.TxtUniqueComments,
		TxtMarketingStrategy:                  u.TxtMarketingStrategy,
		TxtDisclaimer2:                        u.TxtDisclaimer2,
		TxtBrokerComments:                     u.TxtBrokerComments,
		TxtValidation:                         u.TxtValidation,
		Txt30DayQuickSale:                     u.Txt30DayQuickSale,
		Txt60DayQuickSale:                     u.Txt60DayQuickSale,
		Txt90DayAsIsValue:                     u.Txt90DayAsIsValue,
		Txt120DayQuickSale:                    u.Txt120DayQuickSale,
		Txt180DayQuickSale:                    u.Txt180DayQuickSale,
		TxtListPriceFinalValues:               u.TxtListPriceFinalValues,
		Txt30DayListPriceFinalValues:          u.Txt30DayListPriceFinalValues,
		Txt30DayQuickSaleRepaired:             u.Txt30DayQuickSaleRepaired,
		Txt60DayQuickSaleRepaired:             u.Txt60DayQuickSaleRepaired,
		Txt90DayAsIsValueRepaired:             u.Txt90DayAsIsValueRepaired,
		Txt120DayQuickSaleRepaired:            u.Txt120DayQuickSaleRepaired,
		Txt180DayQuickSaleRepaired:            u.Txt180DayQuickSaleRepaired,
		TxtListPriceRepaired:                  u.TxtListPriceRepaired,
		Txt30DayListPriceRepaired:             u.Txt30DayListPriceRepaired,
		CmbHouse:                              u.CmbHouse,
		CmbPositive:                           u.CmbPositive,
		CmbNegative:                           u.CmbNegative,
		CmbView:                               u.CmbView,
		CmbMarket:                             u.CmbMarket,
		CmbPricing:                            u.CmbPricing,
		CmbListing:                            u.CmbListing,
		CmbExtra:                              u.CmbExtra,
		TxtUnique:                             u.TxtUnique,
		PriceComment:                          u.PriceComment,
		RangeComment:                          u.RangeComment,
		ProxException:                         u.ProxException,
		GlaException:                          u.GlaException,
		AgeException:                          u.AgeException,
		CondException:                         u.CondException,
		ViewException:                         u.ViewException,
		StyleException:                        u.StyleException,
		LotException:                          u.LotException,
		BedException:                          u.BedException,
		BathException:                         u.BathException,
		//	CreatedBy:                             createdBy,
		//	CreatedDateTime:                       primitive.DateTime(millis.NowInMillis()),
	}

	res, err := DbCollections.Iforms.InsertOne(ctx, newCompany)
	if err != nil {
		return "", err
	}
	lastInsertedIDStr := fmt.Sprintf("%v", res.InsertedID)
	return strings.TrimObjectChar(lastInsertedIDStr), nil
}

func UpdateIformStatus(ctx context.Context, pipelineId string, newStatus string) (bool, error) {

	filter := bson.D{{"_id", pipelineId}}

	setDoc := bson.D{
		{"status", newStatus},
		{"competedDate", pointers.PrimitiveDateTime(nil)},
	}

	update := bson.M{
		"$set": setDoc,
	}

	res, err := DbCollections.Iforms.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func UpdateIform(ctx context.Context, pipelineId string, input models.UpdateIformInput, updatedBy string) (bool, error) {

	pipelineFilter := FilterByIds([]string{pipelineId})
	pipelinesRaw, err := SearchPipelines(ctx, *pipelineFilter, 0, 1)
	if len(pipelinesRaw) == 0 || err != nil {
		return false, errs.InvalidPipelineId
	}

	iformRaw, err := GetIformByPipelineId(ctx, pipelineId)
	if err != nil || iformRaw == nil {
		_, err := SaveIform(ctx, pipelineId, input, pipelinesRaw, updatedBy)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	filter := bson.M{"pipelineId": pipelineId}

	updateDoc := Iform{
		LastUpdateTime: pointers.PrimitiveDateTime(nil),
	}

	if input.FormType != nil {
		updateDoc.FormType = input.FormType
	}
	if pipelinesRaw[0].UserId != nil {
		updateDoc.ClientID = pipelinesRaw[0].UserId
	}
	if input.TxtClient != nil {
		updateDoc.TxtClient = input.TxtClient
	}
	if input.TxtCompany != nil {
		updateDoc.TxtCompany = input.TxtCompany
	}
	if input.TxtOrderNumber != nil {
		updateDoc.TxtOrderNumber = input.TxtOrderNumber
	}
	if input.CmbOrderType != nil {
		updateDoc.CmbOrderType = input.CmbOrderType
	}
	if input.TxtAddress != nil {
		updateDoc.TxtAddress = input.TxtAddress
	}
	if input.TxtLocation != nil {
		updateDoc.TxtLocation = input.TxtLocation
	}
	if input.TxtBrokerChecker != nil {
		updateDoc.TxtBrokerChecker = input.TxtBrokerChecker
	}
	if input.TxtPreparerInfoAgent != nil {
		updateDoc.TxtPreparerInfoAgent = input.TxtPreparerInfoAgent
	}
	if input.TxtPreparerInfoAgentLicense != nil {
		updateDoc.TxtPreparerInfoAgentLicense = input.TxtPreparerInfoAgentLicense
	}
	if input.TxtPreparerInfoBroker != nil {
		updateDoc.TxtPreparerInfoBroker = input.TxtPreparerInfoBroker
	}
	if input.TxtPreparerInfoBrokerLicense != nil {
		updateDoc.TxtPreparerInfoBrokerLicense = input.TxtPreparerInfoBrokerLicense
	}
	if input.TxtPreparerInfoAddress != nil {
		updateDoc.TxtPreparerInfoAddress = input.TxtPreparerInfoAddress
	}
	if input.TxtPreparerInfoBrokerage != nil {
		updateDoc.TxtPreparerInfoBrokerage = input.TxtPreparerInfoBrokerage
	}
	if input.TxtPreparerInfoAgentCompany != nil {
		updateDoc.TxtPreparerInfoAgentCompany = input.TxtPreparerInfoAgentCompany
	}
	if input.TxtPreparerInfoPhone != nil {
		updateDoc.TxtPreparerInfoPhone = input.TxtPreparerInfoPhone
	}
	if input.TxtPreparerInfoYearsOfExperience != nil {
		updateDoc.TxtPreparerInfoYearsOfExperience = input.TxtPreparerInfoYearsOfExperience
	}
	if input.TxtPreparerInfoEmail != nil {
		updateDoc.TxtPreparerInfoEmail = input.TxtPreparerInfoEmail
	}
	if input.TxtSubjectAddress != nil {
		updateDoc.TxtSubjectAddress = input.TxtSubjectAddress
	}
	if input.TxtPreparerInfoMilesAwayFromSubject != nil {
		updateDoc.TxtPreparerInfoMilesAwayFromSubject = input.TxtPreparerInfoMilesAwayFromSubject
	}
	if input.TxtAgentZip != nil {
		updateDoc.TxtAgentZip = input.TxtAgentZip
	}
	if input.TxtAgentCity != nil {
		updateDoc.TxtAgentCity = input.TxtAgentCity
	}
	if input.TxtAgentState != nil {
		updateDoc.TxtAgentState = input.TxtAgentState
	}
	if input.TxtDisclaimer != nil {
		updateDoc.TxtDisclaimer = input.TxtDisclaimer
	}
	if input.CmbLocation != nil {
		updateDoc.CmbLocation = input.CmbLocation
	}
	if input.TxtCounty != nil {
		updateDoc.TxtCounty = input.TxtCounty
	}
	if input.TxtTrullia != nil {
		updateDoc.TxtTrullia = input.TxtTrullia
	}
	if input.TxtZillow != nil {
		updateDoc.TxtZillow = input.TxtZillow
	}
	if input.TxtFindcompsnow != nil {
		updateDoc.TxtFindcompsnow = input.TxtFindcompsnow
	}
	if input.TxtAverage != nil {
		updateDoc.TxtAverage = input.TxtAverage
	}
	if input.CmbForm != nil {
		updateDoc.CmbForm = input.CmbForm
	}
	if input.CmbForm2 != nil {
		updateDoc.CmbForm2 = input.CmbForm2
	}
	if input.TxtSaleComp1Address != nil {
		updateDoc.TxtSaleComp1Address = input.TxtSaleComp1Address
	}
	if input.TxtSaleComp2Address != nil {
		updateDoc.TxtSaleComp2Address = input.TxtSaleComp2Address
	}
	if input.TxtSaleComp3Address != nil {
		updateDoc.TxtSaleComp3Address = input.TxtSaleComp3Address
	}
	if input.TxtListComp1Address != nil {
		updateDoc.TxtListComp1Address = input.TxtListComp1Address
	}
	if input.TxtListComp2Address != nil {
		updateDoc.TxtListComp2Address = input.TxtListComp2Address
	}
	if input.TxtListComp3Address != nil {
		updateDoc.TxtListComp3Address = input.TxtListComp3Address
	}
	if input.TxtSubjectState != nil {
		updateDoc.TxtSubjectState = input.TxtSubjectState
	}
	if input.TxtSaleComp1State != nil {
		updateDoc.TxtSaleComp1State = input.TxtSaleComp1State
	}
	if input.TxtSaleComp2State != nil {
		updateDoc.TxtSaleComp2State = input.TxtSaleComp2State
	}
	if input.TxtSaleComp3State != nil {
		updateDoc.TxtSaleComp3State = input.TxtSaleComp3State
	}
	if input.TxtListComp1State != nil {
		updateDoc.TxtListComp1State = input.TxtListComp1State
	}
	if input.TxtListComp2State != nil {
		updateDoc.TxtListComp2State = input.TxtListComp2State
	}
	if input.TxtListComp3State != nil {
		updateDoc.TxtListComp3State = input.TxtListComp3State
	}
	if input.TxtSubjectCity != nil {
		updateDoc.TxtSubjectCity = input.TxtSubjectCity
	}
	if input.TxtSaleComp1City != nil {
		updateDoc.TxtSaleComp1City = input.TxtSaleComp1City
	}
	if input.TxtSaleComp2City != nil {
		updateDoc.TxtSaleComp2City = input.TxtSaleComp2City
	}
	if input.TxtSaleComp3City != nil {
		updateDoc.TxtSaleComp3City = input.TxtSaleComp3City
	}
	if input.TxtListComp1City != nil {
		updateDoc.TxtListComp1City = input.TxtListComp1City
	}
	if input.TxtListComp2City != nil {
		updateDoc.TxtListComp2City = input.TxtListComp2City
	}
	if input.TxtListComp3City != nil {
		updateDoc.TxtListComp3City = input.TxtListComp3City
	}
	if input.TxtSubjectnoUnit != nil {
		updateDoc.TxtSubjectnoUnit = input.TxtSubjectnoUnit
	}
	if input.TxtSubjectUnitNo != nil {
		updateDoc.TxtSubjectUnitNo = input.TxtSubjectUnitNo
	}
	if input.TxtSaleComp1noUnit != nil {
		updateDoc.TxtSaleComp1noUnit = input.TxtSaleComp1noUnit
	}
	if input.TxtSaleComp1UnitNo != nil {
		updateDoc.TxtSaleComp1UnitNo = input.TxtSaleComp1UnitNo
	}
	if input.TxtSaleComp2noUnit != nil {
		updateDoc.TxtSaleComp2noUnit = input.TxtSaleComp2noUnit
	}
	if input.TxtSaleComp2UnitNo != nil {
		updateDoc.TxtSaleComp2UnitNo = input.TxtSaleComp2UnitNo
	}
	if input.TxtSaleComp3noUnit != nil {
		updateDoc.TxtSaleComp3noUnit = input.TxtSaleComp3noUnit
	}
	if input.TxtSaleComp3UnitNo != nil {
		updateDoc.TxtSaleComp3UnitNo = input.TxtSaleComp3UnitNo
	}
	if input.TxtListComp1noUnit != nil {
		updateDoc.TxtListComp1noUnit = input.TxtListComp1noUnit
	}
	if input.TxtListComp1UnitNo != nil {
		updateDoc.TxtListComp1UnitNo = input.TxtListComp1UnitNo
	}
	if input.TxtListComp2noUnit != nil {
		updateDoc.TxtListComp2noUnit = input.TxtListComp2noUnit
	}
	if input.TxtListComp2UnitNo != nil {
		updateDoc.TxtListComp2UnitNo = input.TxtListComp2UnitNo
	}
	if input.TxtListComp3noUnit != nil {
		updateDoc.TxtListComp3noUnit = input.TxtListComp3noUnit
	}
	if input.TxtListComp3UnitNo != nil {
		updateDoc.TxtListComp3UnitNo = input.TxtListComp3UnitNo
	}
	if input.TxtSubjectUnits != nil {
		updateDoc.TxtSubjectUnits = input.TxtSubjectUnits
	}
	if input.TxtSaleComp1Units != nil {
		updateDoc.TxtSaleComp1Units = input.TxtSaleComp1Units
	}
	if input.TxtSaleComp2Units != nil {
		updateDoc.TxtSaleComp2Units = input.TxtSaleComp2Units
	}
	if input.TxtSaleComp3Units != nil {
		updateDoc.TxtSaleComp3Units = input.TxtSaleComp3Units
	}
	if input.TxtListComp1Units != nil {
		updateDoc.TxtListComp1Units = input.TxtListComp1Units
	}
	if input.TxtListComp2Units != nil {
		updateDoc.TxtListComp2Units = input.TxtListComp2Units
	}
	if input.TxtListComp3Units != nil {
		updateDoc.TxtListComp3Units = input.TxtListComp3Units
	}
	if input.TxtSubjectZip != nil {
		updateDoc.TxtSubjectZip = input.TxtSubjectZip
	}
	if input.TxtSaleComp1Zip != nil {
		updateDoc.TxtSaleComp1Zip = input.TxtSaleComp1Zip
	}
	if input.TxtSaleComp2Zip != nil {
		updateDoc.TxtSaleComp2Zip = input.TxtSaleComp2Zip
	}
	if input.TxtSaleComp3Zip != nil {
		updateDoc.TxtSaleComp3Zip = input.TxtSaleComp3Zip
	}
	if input.TxtListComp1Zip != nil {
		updateDoc.TxtListComp1Zip = input.TxtListComp1Zip
	}
	if input.TxtListComp2Zip != nil {
		updateDoc.TxtListComp2Zip = input.TxtListComp2Zip
	}
	if input.TxtListComp3Zip != nil {
		updateDoc.TxtListComp3Zip = input.TxtListComp3Zip
	}
	if input.TxtSubjectProximity != nil {
		updateDoc.TxtSubjectProximity = input.TxtSubjectProximity
	}
	if input.TxtSaleComp1Proximity != nil {
		updateDoc.TxtSaleComp1Proximity = input.TxtSaleComp1Proximity
	}
	if input.TxtSaleComp2Proximity != nil {
		updateDoc.TxtSaleComp2Proximity = input.TxtSaleComp2Proximity
	}
	if input.TxtSaleComp3Proximity != nil {
		updateDoc.TxtSaleComp3Proximity = input.TxtSaleComp3Proximity
	}
	if input.TxtListComp1Proximity != nil {
		updateDoc.TxtListComp1Proximity = input.TxtListComp1Proximity
	}
	if input.TxtListComp2Proximity != nil {
		updateDoc.TxtListComp2Proximity = input.TxtListComp2Proximity
	}
	if input.TxtListComp3Proximity != nil {
		updateDoc.TxtListComp3Proximity = input.TxtListComp3Proximity
	}
	if input.TxtSubjectDataSource != nil {
		updateDoc.TxtSubjectDataSource = input.TxtSubjectDataSource
	}
	if input.TxtSaleComp1DataSource != nil {
		updateDoc.TxtSaleComp1DataSource = input.TxtSaleComp1DataSource
	}
	if input.TxtSaleComp2DataSource != nil {
		updateDoc.TxtSaleComp2DataSource = input.TxtSaleComp2DataSource
	}
	if input.TxtSaleComp3DataSource != nil {
		updateDoc.TxtSaleComp3DataSource = input.TxtSaleComp3DataSource
	}
	if input.TxtListComp1DataSource != nil {
		updateDoc.TxtListComp1DataSource = input.TxtListComp1DataSource
	}
	if input.TxtListComp2DataSource != nil {
		updateDoc.TxtListComp2DataSource = input.TxtListComp2DataSource
	}
	if input.TxtListComp3DataSource != nil {
		updateDoc.TxtListComp3DataSource = input.TxtListComp3DataSource
	}
	if input.TxtSubjectMLSNumber != nil {
		updateDoc.TxtSubjectMLSNumber = input.TxtSubjectMLSNumber
	}
	if input.TxtSaleComp1MLSNumber != nil {
		updateDoc.TxtSaleComp1MLSNumber = input.TxtSaleComp1MLSNumber
	}
	if input.TxtSaleComp2MLSNumber != nil {
		updateDoc.TxtSaleComp2MLSNumber = input.TxtSaleComp2MLSNumber
	}
	if input.TxtSaleComp3MLSNumber != nil {
		updateDoc.TxtSaleComp3MLSNumber = input.TxtSaleComp3MLSNumber
	}
	if input.TxtListComp1MLSNumber != nil {
		updateDoc.TxtListComp1MLSNumber = input.TxtListComp1MLSNumber
	}
	if input.TxtListComp2MLSNumber != nil {
		updateDoc.TxtListComp2MLSNumber = input.TxtListComp2MLSNumber
	}
	if input.TxtListComp3MLSNumber != nil {
		updateDoc.TxtListComp3MLSNumber = input.TxtListComp3MLSNumber
	}
	if input.CmbSubjectSaleType != nil {
		updateDoc.CmbSubjectSaleType = input.CmbSubjectSaleType
	}
	if input.CmbSaleComp1SaleType != nil {
		updateDoc.CmbSaleComp1SaleType = input.CmbSaleComp1SaleType
	}
	if input.CmbSaleComp2SaleType != nil {
		updateDoc.CmbSaleComp2SaleType = input.CmbSaleComp2SaleType
	}
	if input.CmbSaleComp3SaleType != nil {
		updateDoc.CmbSaleComp3SaleType = input.CmbSaleComp3SaleType
	}
	if input.CmbListComp1SaleType != nil {
		updateDoc.CmbListComp1SaleType = input.CmbListComp1SaleType
	}
	if input.CmbListComp2SaleType != nil {
		updateDoc.CmbListComp2SaleType = input.CmbListComp2SaleType
	}
	if input.CmbListComp3SaleType != nil {
		updateDoc.CmbListComp3SaleType = input.CmbListComp3SaleType
	}
	if input.CmbSubjectType != nil {
		updateDoc.CmbSubjectType = input.CmbSubjectType
	}
	if input.CmbSaleComp1Type != nil {
		updateDoc.CmbSaleComp1Type = input.CmbSaleComp1Type
	}
	if input.CmbSaleComp2Type != nil {
		updateDoc.CmbSaleComp2Type = input.CmbSaleComp2Type
	}
	if input.CmbSaleComp3Type != nil {
		updateDoc.CmbSaleComp3Type = input.CmbSaleComp3Type
	}
	if input.CmbListComp1Type != nil {
		updateDoc.CmbListComp1Type = input.CmbListComp1Type
	}
	if input.CmbListComp2Type != nil {
		updateDoc.CmbListComp2Type = input.CmbListComp2Type
	}
	if input.CmbListComp3Type != nil {
		updateDoc.CmbListComp3Type = input.CmbListComp3Type
	}
	if input.CmbSubjectStyle != nil {
		updateDoc.CmbSubjectStyle = input.CmbSubjectStyle
	}
	if input.CmbSaleComp1Style != nil {
		updateDoc.CmbSaleComp1Style = input.CmbSaleComp1Style
	}
	if input.TxtSaleComp1StyleAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1StyleAdjBuiltIn = input.TxtSaleComp1StyleAdjBuiltIn
	}
	if input.CmbSaleComp2Style != nil {
		updateDoc.CmbSaleComp2Style = input.CmbSaleComp2Style
	}
	if input.TxtSaleComp2StyleAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2StyleAdjBuiltIn = input.TxtSaleComp2StyleAdjBuiltIn
	}
	if input.CmbSaleComp3Style != nil {
		updateDoc.CmbSaleComp3Style = input.CmbSaleComp3Style
	}
	if input.TxtSaleComp3StyleAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3StyleAdjBuiltIn = input.TxtSaleComp3StyleAdjBuiltIn
	}
	if input.CmbListComp1Style != nil {
		updateDoc.CmbListComp1Style = input.CmbListComp1Style
	}
	if input.TxtListComp1StyleAdjBuiltIn != nil {
		updateDoc.TxtListComp1StyleAdjBuiltIn = input.TxtListComp1StyleAdjBuiltIn
	}
	if input.CmbListComp2Style != nil {
		updateDoc.CmbListComp2Style = input.CmbListComp2Style
	}
	if input.TxtListComp2StyleAdjBuiltIn != nil {
		updateDoc.TxtListComp2StyleAdjBuiltIn = input.TxtListComp2StyleAdjBuiltIn
	}
	if input.CmbListComp3Style != nil {
		updateDoc.CmbListComp3Style = input.CmbListComp3Style
	}
	if input.TxtListComp3StyleAdjBuiltIn != nil {
		updateDoc.TxtListComp3StyleAdjBuiltIn = input.TxtListComp3StyleAdjBuiltIn
	}
	if input.CmbSubjectExtFinish != nil {
		updateDoc.CmbSubjectExtFinish = input.CmbSubjectExtFinish
	}
	if input.CmbSaleComp1ExtFinish != nil {
		updateDoc.CmbSaleComp1ExtFinish = input.CmbSaleComp1ExtFinish
	}
	if input.TxtSaleComp1ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1ExtFinishAdjBuiltIn = input.TxtSaleComp1ExtFinishAdjBuiltIn
	}
	if input.CmbSaleComp2ExtFinish != nil {
		updateDoc.CmbSaleComp2ExtFinish = input.CmbSaleComp2ExtFinish
	}
	if input.TxtSaleComp2ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2ExtFinishAdjBuiltIn = input.TxtSaleComp2ExtFinishAdjBuiltIn
	}
	if input.CmbSaleComp3ExtFinish != nil {
		updateDoc.CmbSaleComp3ExtFinish = input.CmbSaleComp3ExtFinish
	}
	if input.TxtSaleComp3ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3ExtFinishAdjBuiltIn = input.TxtSaleComp3ExtFinishAdjBuiltIn
	}
	if input.CmbListComp1ExtFinish != nil {
		updateDoc.CmbListComp1ExtFinish = input.CmbListComp1ExtFinish
	}
	if input.TxtListComp1ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtListComp1ExtFinishAdjBuiltIn = input.TxtListComp1ExtFinishAdjBuiltIn
	}
	if input.CmbListComp2ExtFinish != nil {
		updateDoc.CmbListComp2ExtFinish = input.CmbListComp2ExtFinish
	}
	if input.TxtListComp2ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtListComp2ExtFinishAdjBuiltIn = input.TxtListComp2ExtFinishAdjBuiltIn
	}
	if input.CmbListComp3ExtFinish != nil {
		updateDoc.CmbListComp3ExtFinish = input.CmbListComp3ExtFinish
	}
	if input.TxtListComp3ExtFinishAdjBuiltIn != nil {
		updateDoc.TxtListComp3ExtFinishAdjBuiltIn = input.TxtListComp3ExtFinishAdjBuiltIn
	}
	if input.CmbSubjectCondition != nil {
		updateDoc.CmbSubjectCondition = input.CmbSubjectCondition
	}
	if input.CmbSaleComp1Condition != nil {
		updateDoc.CmbSaleComp1Condition = input.CmbSaleComp1Condition
	}
	if input.TxtSaleComp1ConditionAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1ConditionAdjBuiltIn = input.TxtSaleComp1ConditionAdjBuiltIn
	}
	if input.CmbSaleComp2Condition != nil {
		updateDoc.CmbSaleComp2Condition = input.CmbSaleComp2Condition
	}
	if input.TxtSaleComp2ConditionAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2ConditionAdjBuiltIn = input.TxtSaleComp2ConditionAdjBuiltIn
	}
	if input.CmbSaleComp3Condition != nil {
		updateDoc.CmbSaleComp3Condition = input.CmbSaleComp3Condition
	}
	if input.TxtSaleComp3ConditionAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3ConditionAdjBuiltIn = input.TxtSaleComp3ConditionAdjBuiltIn
	}
	if input.CmbListComp1Condition != nil {
		updateDoc.CmbListComp1Condition = input.CmbListComp1Condition
	}
	if input.TxtListComp1ConditionAdjBuiltIn != nil {
		updateDoc.TxtListComp1ConditionAdjBuiltIn = input.TxtListComp1ConditionAdjBuiltIn
	}
	if input.CmbListComp2Condition != nil {
		updateDoc.CmbListComp2Condition = input.CmbListComp2Condition
	}
	if input.TxtListComp2ConditionAdjBuiltIn != nil {
		updateDoc.TxtListComp2ConditionAdjBuiltIn = input.TxtListComp2ConditionAdjBuiltIn
	}
	if input.CmbListComp3Condition != nil {
		updateDoc.CmbListComp3Condition = input.CmbListComp3Condition
	}
	if input.TxtListComp3ConditionAdjBuiltIn != nil {
		updateDoc.TxtListComp3ConditionAdjBuiltIn = input.TxtListComp3ConditionAdjBuiltIn
	}
	if input.CmbSubjectQuality != nil {
		updateDoc.CmbSubjectQuality = input.CmbSubjectQuality
	}
	if input.CmbSaleComp1Quality != nil {
		updateDoc.CmbSaleComp1Quality = input.CmbSaleComp1Quality
	}
	if input.TxtSaleComp1QualityAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1QualityAdjBuiltIn = input.TxtSaleComp1QualityAdjBuiltIn
	}
	if input.CmbSaleComp2Quality != nil {
		updateDoc.CmbSaleComp2Quality = input.CmbSaleComp2Quality
	}
	if input.TxtSaleComp2QualityAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2QualityAdjBuiltIn = input.TxtSaleComp2QualityAdjBuiltIn
	}
	if input.CmbSaleComp3Quality != nil {
		updateDoc.CmbSaleComp3Quality = input.CmbSaleComp3Quality
	}
	if input.TxtSaleComp3QualityAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3QualityAdjBuiltIn = input.TxtSaleComp3QualityAdjBuiltIn
	}
	if input.CmbListComp1Quality != nil {
		updateDoc.CmbListComp1Quality = input.CmbListComp1Quality
	}
	if input.TxtListComp1QualityAdjBuiltIn != nil {
		updateDoc.TxtListComp1QualityAdjBuiltIn = input.TxtListComp1QualityAdjBuiltIn
	}
	if input.CmbListComp2Quality != nil {
		updateDoc.CmbListComp2Quality = input.CmbListComp2Quality
	}
	if input.TxtListComp2QualityAdjBuiltIn != nil {
		updateDoc.TxtListComp2QualityAdjBuiltIn = input.TxtListComp2QualityAdjBuiltIn
	}
	if input.CmbListComp3Quality != nil {
		updateDoc.CmbListComp3Quality = input.CmbListComp3Quality
	}
	if input.TxtListComp3QualityAdjBuiltIn != nil {
		updateDoc.TxtListComp3QualityAdjBuiltIn = input.TxtListComp3QualityAdjBuiltIn
	}
	if input.CmbSubjectView != nil {
		updateDoc.CmbSubjectView = input.CmbSubjectView
	}
	if input.CmbSaleComp1View != nil {
		updateDoc.CmbSaleComp1View = input.CmbSaleComp1View
	}
	if input.TxtSaleComp1ViewAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1ViewAdjBuiltIn = input.TxtSaleComp1ViewAdjBuiltIn
	}
	if input.CmbSaleComp2View != nil {
		updateDoc.CmbSaleComp2View = input.CmbSaleComp2View
	}
	if input.TxtSaleComp2ViewAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2ViewAdjBuiltIn = input.TxtSaleComp2ViewAdjBuiltIn
	}
	if input.CmbSaleComp3View != nil {
		updateDoc.CmbSaleComp3View = input.CmbSaleComp3View
	}
	if input.TxtSaleComp3ViewAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3ViewAdjBuiltIn = input.TxtSaleComp3ViewAdjBuiltIn
	}
	if input.CmbListComp1View != nil {
		updateDoc.CmbListComp1View = input.CmbListComp1View
	}
	if input.TxtListComp1ViewAdjBuiltIn != nil {
		updateDoc.TxtListComp1ViewAdjBuiltIn = input.TxtListComp1ViewAdjBuiltIn
	}
	if input.CmbListComp2View != nil {
		updateDoc.CmbListComp2View = input.CmbListComp2View
	}
	if input.TxtListComp2ViewAdjBuiltIn != nil {
		updateDoc.TxtListComp2ViewAdjBuiltIn = input.TxtListComp2ViewAdjBuiltIn
	}
	if input.CmbListComp3View != nil {
		updateDoc.CmbListComp3View = input.CmbListComp3View
	}
	if input.TxtListComp3ViewAdjBuiltIn != nil {
		updateDoc.TxtListComp3ViewAdjBuiltIn = input.TxtListComp3ViewAdjBuiltIn
	}
	if input.TxtSubjectSubdivision != nil {
		updateDoc.TxtSubjectSubdivision = input.TxtSubjectSubdivision
	}
	if input.TxtSaleComp1Subdivision != nil {
		updateDoc.TxtSaleComp1Subdivision = input.TxtSaleComp1Subdivision
	}
	if input.TxtSaleComp2Subdivision != nil {
		updateDoc.TxtSaleComp2Subdivision = input.TxtSaleComp2Subdivision
	}
	if input.TxtSaleComp3Subdivision != nil {
		updateDoc.TxtSaleComp3Subdivision = input.TxtSaleComp3Subdivision
	}
	if input.TxtListComp1Subdivision != nil {
		updateDoc.TxtListComp1Subdivision = input.TxtListComp1Subdivision
	}
	if input.TxtListComp2Subdivision != nil {
		updateDoc.TxtListComp2Subdivision = input.TxtListComp2Subdivision
	}
	if input.TxtListComp3Subdivision != nil {
		updateDoc.TxtListComp3Subdivision = input.TxtListComp3Subdivision
	}
	if input.TxtSubjectHOAFee != nil {
		updateDoc.TxtSubjectHOAFee = input.TxtSubjectHOAFee
	}
	if input.TxtSaleComp1HOAFee != nil {
		updateDoc.TxtSaleComp1HOAFee = input.TxtSaleComp1HOAFee
	}
	if input.TxtSaleComp2HOAFee != nil {
		updateDoc.TxtSaleComp2HOAFee = input.TxtSaleComp2HOAFee
	}
	if input.TxtSaleComp3HOAFee != nil {
		updateDoc.TxtSaleComp3HOAFee = input.TxtSaleComp3HOAFee
	}
	if input.TxtListComp1HOAFee != nil {
		updateDoc.TxtListComp1HOAFee = input.TxtListComp1HOAFee
	}
	if input.TxtListComp2HOAFee != nil {
		updateDoc.TxtListComp2HOAFee = input.TxtListComp2HOAFee
	}
	if input.TxtListComp3HOAFee != nil {
		updateDoc.TxtListComp3HOAFee = input.TxtListComp3HOAFee
	}
	if input.TxtSubjectTotalRooms != nil {
		updateDoc.TxtSubjectTotalRooms = input.TxtSubjectTotalRooms
	}
	if input.TxtSaleComp1TotalRooms != nil {
		updateDoc.TxtSaleComp1TotalRooms = input.TxtSaleComp1TotalRooms
	}
	if input.TxtSaleComp1TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1TotalRoomsAdjBuiltIn = input.TxtSaleComp1TotalRoomsAdjBuiltIn
	}
	if input.TxtSaleComp2TotalRooms != nil {
		updateDoc.TxtSaleComp2TotalRooms = input.TxtSaleComp2TotalRooms
	}
	if input.TxtSaleComp2TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2TotalRoomsAdjBuiltIn = input.TxtSaleComp2TotalRoomsAdjBuiltIn
	}
	if input.TxtSaleComp3TotalRooms != nil {
		updateDoc.TxtSaleComp3TotalRooms = input.TxtSaleComp3TotalRooms
	}
	if input.TxtSaleComp3TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3TotalRoomsAdjBuiltIn = input.TxtSaleComp3TotalRoomsAdjBuiltIn
	}
	if input.TxtListComp1TotalRooms != nil {
		updateDoc.TxtListComp1TotalRooms = input.TxtListComp1TotalRooms
	}
	if input.TxtListComp1TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtListComp1TotalRoomsAdjBuiltIn = input.TxtListComp1TotalRoomsAdjBuiltIn
	}
	if input.TxtListComp2TotalRooms != nil {
		updateDoc.TxtListComp2TotalRooms = input.TxtListComp2TotalRooms
	}
	if input.TxtListComp2TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtListComp2TotalRoomsAdjBuiltIn = input.TxtListComp2TotalRoomsAdjBuiltIn
	}
	if input.TxtListComp3TotalRooms != nil {
		updateDoc.TxtListComp3TotalRooms = input.TxtListComp3TotalRooms
	}
	if input.TxtListComp3TotalRoomsAdjBuiltIn != nil {
		updateDoc.TxtListComp3TotalRoomsAdjBuiltIn = input.TxtListComp3TotalRoomsAdjBuiltIn
	}
	if input.TxtSubjectBedrooms != nil {
		updateDoc.TxtSubjectBedrooms = input.TxtSubjectBedrooms
	}
	if input.TxtSaleComp1Bedrooms != nil {
		updateDoc.TxtSaleComp1Bedrooms = input.TxtSaleComp1Bedrooms
	}
	if input.TxtSaleComp1BedroomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1BedroomsAdjBuiltIn = input.TxtSaleComp1BedroomsAdjBuiltIn
	}
	if input.TxtSaleComp2Bedrooms != nil {
		updateDoc.TxtSaleComp2Bedrooms = input.TxtSaleComp2Bedrooms
	}
	if input.TxtSaleComp2BedroomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2BedroomsAdjBuiltIn = input.TxtSaleComp2BedroomsAdjBuiltIn
	}
	if input.TxtSaleComp3Bedrooms != nil {
		updateDoc.TxtSaleComp3Bedrooms = input.TxtSaleComp3Bedrooms
	}
	if input.TxtSaleComp3BedroomsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3BedroomsAdjBuiltIn = input.TxtSaleComp3BedroomsAdjBuiltIn
	}
	if input.TxtListComp1Bedrooms != nil {
		updateDoc.TxtListComp1Bedrooms = input.TxtListComp1Bedrooms
	}
	if input.TxtListComp1BedroomsAdjBuiltIn != nil {
		updateDoc.TxtListComp1BedroomsAdjBuiltIn = input.TxtListComp1BedroomsAdjBuiltIn
	}
	if input.TxtListComp2Bedrooms != nil {
		updateDoc.TxtListComp2Bedrooms = input.TxtListComp2Bedrooms
	}
	if input.TxtListComp2BedroomsAdjBuiltIn != nil {
		updateDoc.TxtListComp2BedroomsAdjBuiltIn = input.TxtListComp2BedroomsAdjBuiltIn
	}
	if input.TxtListComp3Bedrooms != nil {
		updateDoc.TxtListComp3Bedrooms = input.TxtListComp3Bedrooms
	}
	if input.TxtListComp3BedroomsAdjBuiltIn != nil {
		updateDoc.TxtListComp3BedroomsAdjBuiltIn = input.TxtListComp3BedroomsAdjBuiltIn
	}
	if input.TxtSubjectFullBaths != nil {
		updateDoc.TxtSubjectFullBaths = input.TxtSubjectFullBaths
	}
	if input.TxtSaleComp1FullBaths != nil {
		updateDoc.TxtSaleComp1FullBaths = input.TxtSaleComp1FullBaths
	}
	if input.TxtSaleComp1FullBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1FullBathsAdjBuiltIn = input.TxtSaleComp1FullBathsAdjBuiltIn
	}
	if input.TxtSaleComp2FullBaths != nil {
		updateDoc.TxtSaleComp2FullBaths = input.TxtSaleComp2FullBaths
	}
	if input.TxtSaleComp2FullBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2FullBathsAdjBuiltIn = input.TxtSaleComp2FullBathsAdjBuiltIn
	}
	if input.TxtSaleComp3FullBaths != nil {
		updateDoc.TxtSaleComp3FullBaths = input.TxtSaleComp3FullBaths
	}
	if input.TxtSaleComp3FullBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3FullBathsAdjBuiltIn = input.TxtSaleComp3FullBathsAdjBuiltIn
	}
	if input.TxtListComp1FullBaths != nil {
		updateDoc.TxtListComp1FullBaths = input.TxtListComp1FullBaths
	}
	if input.TxtListComp1FullBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp1FullBathsAdjBuiltIn = input.TxtListComp1FullBathsAdjBuiltIn
	}
	if input.TxtListComp2FullBaths != nil {
		updateDoc.TxtListComp2FullBaths = input.TxtListComp2FullBaths
	}
	if input.TxtListComp2FullBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp2FullBathsAdjBuiltIn = input.TxtListComp2FullBathsAdjBuiltIn
	}
	if input.TxtListComp3FullBaths != nil {
		updateDoc.TxtListComp3FullBaths = input.TxtListComp3FullBaths
	}
	if input.TxtListComp3FullBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp3FullBathsAdjBuiltIn = input.TxtListComp3FullBathsAdjBuiltIn
	}
	if input.TxtSubjectHalfBaths != nil {
		updateDoc.TxtSubjectHalfBaths = input.TxtSubjectHalfBaths
	}
	if input.TxtSaleComp1HalfBaths != nil {
		updateDoc.TxtSaleComp1HalfBaths = input.TxtSaleComp1HalfBaths
	}
	if input.TxtSaleComp1HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1HalfBathsAdjBuiltIn = input.TxtSaleComp1HalfBathsAdjBuiltIn
	}
	if input.TxtSaleComp2HalfBaths != nil {
		updateDoc.TxtSaleComp2HalfBaths = input.TxtSaleComp2HalfBaths
	}
	if input.TxtSaleComp2HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2HalfBathsAdjBuiltIn = input.TxtSaleComp2HalfBathsAdjBuiltIn
	}
	if input.TxtSaleComp3HalfBaths != nil {
		updateDoc.TxtSaleComp3HalfBaths = input.TxtSaleComp3HalfBaths
	}
	if input.TxtSaleComp3HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3HalfBathsAdjBuiltIn = input.TxtSaleComp3HalfBathsAdjBuiltIn
	}
	if input.TxtListComp1HalfBaths != nil {
		updateDoc.TxtListComp1HalfBaths = input.TxtListComp1HalfBaths
	}
	if input.TxtListComp1HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp1HalfBathsAdjBuiltIn = input.TxtListComp1HalfBathsAdjBuiltIn
	}
	if input.TxtListComp2HalfBaths != nil {
		updateDoc.TxtListComp2HalfBaths = input.TxtListComp2HalfBaths
	}
	if input.TxtListComp2HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp2HalfBathsAdjBuiltIn = input.TxtListComp2HalfBathsAdjBuiltIn
	}
	if input.TxtListComp3HalfBaths != nil {
		updateDoc.TxtListComp3HalfBaths = input.TxtListComp3HalfBaths
	}
	if input.TxtListComp3HalfBathsAdjBuiltIn != nil {
		updateDoc.TxtListComp3HalfBathsAdjBuiltIn = input.TxtListComp3HalfBathsAdjBuiltIn
	}
	if input.TxtSubjectGla != nil {
		updateDoc.TxtSubjectGla = input.TxtSubjectGla
	}
	if input.TxtSaleComp1gla != nil {
		updateDoc.TxtSaleComp1gla = input.TxtSaleComp1gla
	}
	if input.TxtSaleComp1GLAAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1GLAAdjBuiltIn = input.TxtSaleComp1GLAAdjBuiltIn
	}
	if input.TxtSaleComp2gla != nil {
		updateDoc.TxtSaleComp2gla = input.TxtSaleComp2gla
	}
	if input.TxtSaleComp2GLAAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2GLAAdjBuiltIn = input.TxtSaleComp2GLAAdjBuiltIn
	}
	if input.TxtSaleComp3gla != nil {
		updateDoc.TxtSaleComp3gla = input.TxtSaleComp3gla
	}
	if input.TxtSaleComp3GLAAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3GLAAdjBuiltIn = input.TxtSaleComp3GLAAdjBuiltIn
	}
	if input.TxtListComp1gla != nil {
		updateDoc.TxtListComp1gla = input.TxtListComp1gla
	}
	if input.TxtListComp1GLAAdjBuiltIn != nil {
		updateDoc.TxtListComp1GLAAdjBuiltIn = input.TxtListComp1GLAAdjBuiltIn
	}
	if input.TxtListComp2gla != nil {
		updateDoc.TxtListComp2gla = input.TxtListComp2gla
	}
	if input.TxtListComp2GLAAdjBuiltIn != nil {
		updateDoc.TxtListComp2GLAAdjBuiltIn = input.TxtListComp2GLAAdjBuiltIn
	}
	if input.TxtListComp3gla != nil {
		updateDoc.TxtListComp3gla = input.TxtListComp3gla
	}
	if input.TxtListComp3GLAAdjBuiltIn != nil {
		updateDoc.TxtListComp3GLAAdjBuiltIn = input.TxtListComp3GLAAdjBuiltIn
	}
	if input.TxtSubjectYearBuilt != nil {
		updateDoc.TxtSubjectYearBuilt = input.TxtSubjectYearBuilt
	}
	if input.TxtSaleComp1YearBuilt != nil {
		updateDoc.TxtSaleComp1YearBuilt = input.TxtSaleComp1YearBuilt
	}
	if input.TxtSaleComp1YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1YearBuiltAdjBuiltIn = input.TxtSaleComp1YearBuiltAdjBuiltIn
	}
	if input.TxtSaleComp2YearBuilt != nil {
		updateDoc.TxtSaleComp2YearBuilt = input.TxtSaleComp2YearBuilt
	}
	if input.TxtSaleComp2YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2YearBuiltAdjBuiltIn = input.TxtSaleComp2YearBuiltAdjBuiltIn
	}
	if input.TxtSaleComp3YearBuilt != nil {
		updateDoc.TxtSaleComp3YearBuilt = input.TxtSaleComp3YearBuilt
	}
	if input.TxtSaleComp3YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3YearBuiltAdjBuiltIn = input.TxtSaleComp3YearBuiltAdjBuiltIn
	}
	if input.TxtListComp1YearBuilt != nil {
		updateDoc.TxtListComp1YearBuilt = input.TxtListComp1YearBuilt
	}
	if input.TxtListComp1YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtListComp1YearBuiltAdjBuiltIn = input.TxtListComp1YearBuiltAdjBuiltIn
	}
	if input.TxtListComp2YearBuilt != nil {
		updateDoc.TxtListComp2YearBuilt = input.TxtListComp2YearBuilt
	}
	if input.TxtListComp2YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtListComp2YearBuiltAdjBuiltIn = input.TxtListComp2YearBuiltAdjBuiltIn
	}
	if input.TxtListComp3YearBuilt != nil {
		updateDoc.TxtListComp3YearBuilt = input.TxtListComp3YearBuilt
	}
	if input.TxtListComp3YearBuiltAdjBuiltIn != nil {
		updateDoc.TxtListComp3YearBuiltAdjBuiltIn = input.TxtListComp3YearBuiltAdjBuiltIn
	}
	if input.TxtSubjectAge != nil {
		updateDoc.TxtSubjectAge = input.TxtSubjectAge
	}
	if input.TxtSaleComp1Age != nil {
		updateDoc.TxtSaleComp1Age = input.TxtSaleComp1Age
	}
	if input.TxtSaleComp2Age != nil {
		updateDoc.TxtSaleComp2Age = input.TxtSaleComp2Age
	}
	if input.TxtSaleComp3Age != nil {
		updateDoc.TxtSaleComp3Age = input.TxtSaleComp3Age
	}
	if input.TxtListComp1Age != nil {
		updateDoc.TxtListComp1Age = input.TxtListComp1Age
	}
	if input.TxtListComp2Age != nil {
		updateDoc.TxtListComp2Age = input.TxtListComp2Age
	}
	if input.TxtListComp3Age != nil {
		updateDoc.TxtListComp3Age = input.TxtListComp3Age
	}
	if input.TxtSubjectAcres != nil {
		updateDoc.TxtSubjectAcres = input.TxtSubjectAcres
	}
	if input.TxtSaleComp1Acres != nil {
		updateDoc.TxtSaleComp1Acres = input.TxtSaleComp1Acres
	}
	if input.TxtSaleComp1AcresAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1AcresAdjBuiltIn = input.TxtSaleComp1AcresAdjBuiltIn
	}
	if input.TxtSaleComp2Acres != nil {
		updateDoc.TxtSaleComp2Acres = input.TxtSaleComp2Acres
	}
	if input.TxtSaleComp2AcresAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2AcresAdjBuiltIn = input.TxtSaleComp2AcresAdjBuiltIn
	}
	if input.TxtSaleComp3Acres != nil {
		updateDoc.TxtSaleComp3Acres = input.TxtSaleComp3Acres
	}
	if input.TxtSaleComp3AcresAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3AcresAdjBuiltIn = input.TxtSaleComp3AcresAdjBuiltIn
	}
	if input.TxtListComp1Acres != nil {
		updateDoc.TxtListComp1Acres = input.TxtListComp1Acres
	}
	if input.TxtListComp1AcresAdjBuiltIn != nil {
		updateDoc.TxtListComp1AcresAdjBuiltIn = input.TxtListComp1AcresAdjBuiltIn
	}
	if input.TxtListComp2Acres != nil {
		updateDoc.TxtListComp2Acres = input.TxtListComp2Acres
	}
	if input.TxtListComp2AcresAdjBuiltIn != nil {
		updateDoc.TxtListComp2AcresAdjBuiltIn = input.TxtListComp2AcresAdjBuiltIn
	}
	if input.TxtListComp3Acres != nil {
		updateDoc.TxtListComp3Acres = input.TxtListComp3Acres
	}
	if input.TxtListComp3AcresAdjBuiltIn != nil {
		updateDoc.TxtListComp3AcresAdjBuiltIn = input.TxtListComp3AcresAdjBuiltIn
	}
	if input.TxtSubjectSquareFeet != nil {
		updateDoc.TxtSubjectSquareFeet = input.TxtSubjectSquareFeet
	}
	if input.TxtSaleComp1SquareFeet != nil {
		updateDoc.TxtSaleComp1SquareFeet = input.TxtSaleComp1SquareFeet
	}
	if input.TxtSaleComp2SquareFeet != nil {
		updateDoc.TxtSaleComp2SquareFeet = input.TxtSaleComp2SquareFeet
	}
	if input.TxtSaleComp3SquareFeet != nil {
		updateDoc.TxtSaleComp3SquareFeet = input.TxtSaleComp3SquareFeet
	}
	if input.TxtListComp1SquareFeet != nil {
		updateDoc.TxtListComp1SquareFeet = input.TxtListComp1SquareFeet
	}
	if input.TxtListComp2SquareFeet != nil {
		updateDoc.TxtListComp2SquareFeet = input.TxtListComp2SquareFeet
	}
	if input.TxtListComp3SquareFeet != nil {
		updateDoc.TxtListComp3SquareFeet = input.TxtListComp3SquareFeet
	}
	if input.CmbSubjectGarage != nil {
		updateDoc.CmbSubjectGarage = input.CmbSubjectGarage
	}
	if input.CmbSaleComp1Garage != nil {
		updateDoc.CmbSaleComp1Garage = input.CmbSaleComp1Garage
	}
	if input.TxtSaleComp1GarageAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1GarageAdjBuiltIn = input.TxtSaleComp1GarageAdjBuiltIn
	}
	if input.CmbSaleComp2Garage != nil {
		updateDoc.CmbSaleComp2Garage = input.CmbSaleComp2Garage
	}
	if input.TxtSaleComp2GarageAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2GarageAdjBuiltIn = input.TxtSaleComp2GarageAdjBuiltIn
	}
	if input.CmbSaleComp3Garage != nil {
		updateDoc.CmbSaleComp3Garage = input.CmbSaleComp3Garage
	}
	if input.TxtSaleComp3GarageAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3GarageAdjBuiltIn = input.TxtSaleComp3GarageAdjBuiltIn
	}
	if input.CmbListComp1Garage != nil {
		updateDoc.CmbListComp1Garage = input.CmbListComp1Garage
	}
	if input.TxtListComp1GarageAdjBuiltIn != nil {
		updateDoc.TxtListComp1GarageAdjBuiltIn = input.TxtListComp1GarageAdjBuiltIn
	}
	if input.CmbListComp2Garage != nil {
		updateDoc.CmbListComp2Garage = input.CmbListComp2Garage
	}
	if input.TxtListComp2GarageAdjBuiltIn != nil {
		updateDoc.TxtListComp2GarageAdjBuiltIn = input.TxtListComp2GarageAdjBuiltIn
	}
	if input.CmbListComp3Garage != nil {
		updateDoc.CmbListComp3Garage = input.CmbListComp3Garage
	}
	if input.TxtListComp3GarageAdjBuiltIn != nil {
		updateDoc.TxtListComp3GarageAdjBuiltIn = input.TxtListComp3GarageAdjBuiltIn
	}
	if input.CmbSubjectPool != nil {
		updateDoc.CmbSubjectPool = input.CmbSubjectPool
	}
	if input.CmbSaleComp1Pool != nil {
		updateDoc.CmbSaleComp1Pool = input.CmbSaleComp1Pool
	}
	if input.TxtSaleComp1PoolAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1PoolAdjBuiltIn = input.TxtSaleComp1PoolAdjBuiltIn
	}
	if input.CmbSaleComp2Pool != nil {
		updateDoc.CmbSaleComp2Pool = input.CmbSaleComp2Pool
	}
	if input.TxtSaleComp2PoolAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2PoolAdjBuiltIn = input.TxtSaleComp2PoolAdjBuiltIn
	}
	if input.CmbSaleComp3Pool != nil {
		updateDoc.CmbSaleComp3Pool = input.CmbSaleComp3Pool
	}
	if input.TxtSaleComp3PoolAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3PoolAdjBuiltIn = input.TxtSaleComp3PoolAdjBuiltIn
	}
	if input.CmbListComp1Pool != nil {
		updateDoc.CmbListComp1Pool = input.CmbListComp1Pool
	}
	if input.TxtListComp1PoolAdjBuiltIn != nil {
		updateDoc.TxtListComp1PoolAdjBuiltIn = input.TxtListComp1PoolAdjBuiltIn
	}
	if input.CmbListComp2Pool != nil {
		updateDoc.CmbListComp2Pool = input.CmbListComp2Pool
	}
	if input.TxtListComp2PoolAdjBuiltIn != nil {
		updateDoc.TxtListComp2PoolAdjBuiltIn = input.TxtListComp2PoolAdjBuiltIn
	}
	if input.CmbListComp3Pool != nil {
		updateDoc.CmbListComp3Pool = input.CmbListComp3Pool
	}
	if input.TxtListComp3PoolAdjBuiltIn != nil {
		updateDoc.TxtListComp3PoolAdjBuiltIn = input.TxtListComp3PoolAdjBuiltIn
	}
	if input.CmbSubjectPorchPatioDeck != nil {
		updateDoc.CmbSubjectPorchPatioDeck = input.CmbSubjectPorchPatioDeck
	}
	if input.CmbSaleComp1PorchPatioDeck != nil {
		updateDoc.CmbSaleComp1PorchPatioDeck = input.CmbSaleComp1PorchPatioDeck
	}
	if input.TxtSaleComp1PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1PorchPatioDeckAdjBuiltIn = input.TxtSaleComp1PorchPatioDeckAdjBuiltIn
	}
	if input.CmbSaleComp2PorchPatioDeck != nil {
		updateDoc.CmbSaleComp2PorchPatioDeck = input.CmbSaleComp2PorchPatioDeck
	}
	if input.TxtSaleComp2PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2PorchPatioDeckAdjBuiltIn = input.TxtSaleComp2PorchPatioDeckAdjBuiltIn
	}
	if input.CmbSaleComp3PorchPatioDeck != nil {
		updateDoc.CmbSaleComp3PorchPatioDeck = input.CmbSaleComp3PorchPatioDeck
	}
	if input.TxtSaleComp3PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3PorchPatioDeckAdjBuiltIn = input.TxtSaleComp3PorchPatioDeckAdjBuiltIn
	}
	if input.CmbListComp1PorchPatioDeck != nil {
		updateDoc.CmbListComp1PorchPatioDeck = input.CmbListComp1PorchPatioDeck
	}
	if input.TxtListComp1PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtListComp1PorchPatioDeckAdjBuiltIn = input.TxtListComp1PorchPatioDeckAdjBuiltIn
	}
	if input.CmbListComp2PorchPatioDeck != nil {
		updateDoc.CmbListComp2PorchPatioDeck = input.CmbListComp2PorchPatioDeck
	}
	if input.TxtListComp2PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtListComp2PorchPatioDeckAdjBuiltIn = input.TxtListComp2PorchPatioDeckAdjBuiltIn
	}
	if input.CmbListComp3PorchPatioDeck != nil {
		updateDoc.CmbListComp3PorchPatioDeck = input.CmbListComp3PorchPatioDeck
	}
	if input.TxtListComp3PorchPatioDeckAdjBuiltIn != nil {
		updateDoc.TxtListComp3PorchPatioDeckAdjBuiltIn = input.TxtListComp3PorchPatioDeckAdjBuiltIn
	}
	if input.CmbSubjectFireplace != nil {
		updateDoc.CmbSubjectFireplace = input.CmbSubjectFireplace
	}
	if input.CmbSaleComp1Fireplace != nil {
		updateDoc.CmbSaleComp1Fireplace = input.CmbSaleComp1Fireplace
	}
	if input.TxtSaleComp1FireplaceAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1FireplaceAdjBuiltIn = input.TxtSaleComp1FireplaceAdjBuiltIn
	}
	if input.CmbSaleComp2Fireplace != nil {
		updateDoc.CmbSaleComp2Fireplace = input.CmbSaleComp2Fireplace
	}
	if input.TxtSaleComp2FireplaceAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2FireplaceAdjBuiltIn = input.TxtSaleComp2FireplaceAdjBuiltIn
	}
	if input.CmbSaleComp3Fireplace != nil {
		updateDoc.CmbSaleComp3Fireplace = input.CmbSaleComp3Fireplace
	}
	if input.TxtSaleComp3FireplaceAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3FireplaceAdjBuiltIn = input.TxtSaleComp3FireplaceAdjBuiltIn
	}
	if input.CmbListComp1Fireplace != nil {
		updateDoc.CmbListComp1Fireplace = input.CmbListComp1Fireplace
	}
	if input.TxtListComp1FireplaceAdjBuiltIn != nil {
		updateDoc.TxtListComp1FireplaceAdjBuiltIn = input.TxtListComp1FireplaceAdjBuiltIn
	}
	if input.CmbListComp2Fireplace != nil {
		updateDoc.CmbListComp2Fireplace = input.CmbListComp2Fireplace
	}
	if input.TxtListComp2FireplaceAdjBuiltIn != nil {
		updateDoc.TxtListComp2FireplaceAdjBuiltIn = input.TxtListComp2FireplaceAdjBuiltIn
	}
	if input.CmbListComp3Fireplace != nil {
		updateDoc.CmbListComp3Fireplace = input.CmbListComp3Fireplace
	}
	if input.TxtListComp3FireplaceAdjBuiltIn != nil {
		updateDoc.TxtListComp3FireplaceAdjBuiltIn = input.TxtListComp3FireplaceAdjBuiltIn
	}
	if input.CmbSubjectBasement != nil {
		updateDoc.CmbSubjectBasement = input.CmbSubjectBasement
	}
	if input.CmbSaleComp1Basement != nil {
		updateDoc.CmbSaleComp1Basement = input.CmbSaleComp1Basement
	}
	if input.TxtSaleComp1BasementAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1BasementAdjBuiltIn = input.TxtSaleComp1BasementAdjBuiltIn
	}
	if input.CmbSaleComp2Basement != nil {
		updateDoc.CmbSaleComp2Basement = input.CmbSaleComp2Basement
	}
	if input.TxtSaleComp2BasementAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2BasementAdjBuiltIn = input.TxtSaleComp2BasementAdjBuiltIn
	}
	if input.CmbSaleComp3Basement != nil {
		updateDoc.CmbSaleComp3Basement = input.CmbSaleComp3Basement
	}
	if input.TxtSaleComp3BasementAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3BasementAdjBuiltIn = input.TxtSaleComp3BasementAdjBuiltIn
	}
	if input.CmbListComp1Basement != nil {
		updateDoc.CmbListComp1Basement = input.CmbListComp1Basement
	}
	if input.TxtListComp1BasementAdjBuiltIn != nil {
		updateDoc.TxtListComp1BasementAdjBuiltIn = input.TxtListComp1BasementAdjBuiltIn
	}
	if input.CmbListComp2Basement != nil {
		updateDoc.CmbListComp2Basement = input.CmbListComp2Basement
	}
	if input.TxtListComp2BasementAdjBuiltIn != nil {
		updateDoc.TxtListComp2BasementAdjBuiltIn = input.TxtListComp2BasementAdjBuiltIn
	}
	if input.CmbListComp3Basement != nil {
		updateDoc.CmbListComp3Basement = input.CmbListComp3Basement
	}
	if input.TxtListComp3BasementAdjBuiltIn != nil {
		updateDoc.TxtListComp3BasementAdjBuiltIn = input.TxtListComp3BasementAdjBuiltIn
	}
	if input.CmbSubjectIsFinished != nil {
		updateDoc.CmbSubjectIsFinished = input.CmbSubjectIsFinished
	}
	if input.CmbSaleComp1IsFinished != nil {
		updateDoc.CmbSaleComp1IsFinished = input.CmbSaleComp1IsFinished
	}
	if input.TxtSaleComp1IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1IsFinishedAdjBuiltIn = input.TxtSaleComp1IsFinishedAdjBuiltIn
	}
	if input.CmbSaleComp2IsFinished != nil {
		updateDoc.CmbSaleComp2IsFinished = input.CmbSaleComp2IsFinished
	}
	if input.TxtSaleComp2IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2IsFinishedAdjBuiltIn = input.TxtSaleComp2IsFinishedAdjBuiltIn
	}
	if input.CmbSaleComp3IsFinished != nil {
		updateDoc.CmbSaleComp3IsFinished = input.CmbSaleComp3IsFinished
	}
	if input.TxtSaleComp3IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3IsFinishedAdjBuiltIn = input.TxtSaleComp3IsFinishedAdjBuiltIn
	}
	if input.CmbListComp1IsFinished != nil {
		updateDoc.CmbListComp1IsFinished = input.CmbListComp1IsFinished
	}
	if input.TxtListComp1IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp1IsFinishedAdjBuiltIn = input.TxtListComp1IsFinishedAdjBuiltIn
	}
	if input.CmbListComp2IsFinished != nil {
		updateDoc.CmbListComp2IsFinished = input.CmbListComp2IsFinished
	}
	if input.TxtListComp2IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp2IsFinishedAdjBuiltIn = input.TxtListComp2IsFinishedAdjBuiltIn
	}
	if input.CmbListComp3IsFinished != nil {
		updateDoc.CmbListComp3IsFinished = input.CmbListComp3IsFinished
	}
	if input.TxtListComp3IsFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp3IsFinishedAdjBuiltIn = input.TxtListComp3IsFinishedAdjBuiltIn
	}
	if input.CmbSubjectPercentFinished != nil {
		updateDoc.CmbSubjectPercentFinished = input.CmbSubjectPercentFinished
	}
	if input.CmbSaleComp1PercentFinished != nil {
		updateDoc.CmbSaleComp1PercentFinished = input.CmbSaleComp1PercentFinished
	}
	if input.TxtSaleComp1PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1PercentFinishedAdjBuiltIn = input.TxtSaleComp1PercentFinishedAdjBuiltIn
	}
	if input.CmbSaleComp2PercentFinished != nil {
		updateDoc.CmbSaleComp2PercentFinished = input.CmbSaleComp2PercentFinished
	}
	if input.TxtSaleComp2PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2PercentFinishedAdjBuiltIn = input.TxtSaleComp2PercentFinishedAdjBuiltIn
	}
	if input.CmbSaleComp3PercentFinished != nil {
		updateDoc.CmbSaleComp3PercentFinished = input.CmbSaleComp3PercentFinished
	}
	if input.TxtSaleComp3PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3PercentFinishedAdjBuiltIn = input.TxtSaleComp3PercentFinishedAdjBuiltIn
	}
	if input.CmbListComp1PercentFinished != nil {
		updateDoc.CmbListComp1PercentFinished = input.CmbListComp1PercentFinished
	}
	if input.TxtListComp1PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp1PercentFinishedAdjBuiltIn = input.TxtListComp1PercentFinishedAdjBuiltIn
	}
	if input.CmbListComp2PercentFinished != nil {
		updateDoc.CmbListComp2PercentFinished = input.CmbListComp2PercentFinished
	}
	if input.TxtListComp2PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp2PercentFinishedAdjBuiltIn = input.TxtListComp2PercentFinishedAdjBuiltIn
	}
	if input.CmbListComp3PercentFinished != nil {
		updateDoc.CmbListComp3PercentFinished = input.CmbListComp3PercentFinished
	}
	if input.TxtListComp3PercentFinishedAdjBuiltIn != nil {
		updateDoc.TxtListComp3PercentFinishedAdjBuiltIn = input.TxtListComp3PercentFinishedAdjBuiltIn
	}
	if input.TxtSubjectBasementSqFt != nil {
		updateDoc.TxtSubjectBasementSqFt = input.TxtSubjectBasementSqFt
	}
	if input.TxtSaleComp1BasementSqFt != nil {
		updateDoc.TxtSaleComp1BasementSqFt = input.TxtSaleComp1BasementSqFt
	}
	if input.TxtSaleComp1BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtSaleComp1BasementSqFtAdjBuiltIn = input.TxtSaleComp1BasementSqFtAdjBuiltIn
	}
	if input.TxtSaleComp2BasementSqFt != nil {
		updateDoc.TxtSaleComp2BasementSqFt = input.TxtSaleComp2BasementSqFt
	}
	if input.TxtSaleComp2BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtSaleComp2BasementSqFtAdjBuiltIn = input.TxtSaleComp2BasementSqFtAdjBuiltIn
	}
	if input.TxtSaleComp3BasementSqFt != nil {
		updateDoc.TxtSaleComp3BasementSqFt = input.TxtSaleComp3BasementSqFt
	}
	if input.TxtSaleComp3BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtSaleComp3BasementSqFtAdjBuiltIn = input.TxtSaleComp3BasementSqFtAdjBuiltIn
	}
	if input.TxtListComp1BasementSqFt != nil {
		updateDoc.TxtListComp1BasementSqFt = input.TxtListComp1BasementSqFt
	}
	if input.TxtListComp1BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtListComp1BasementSqFtAdjBuiltIn = input.TxtListComp1BasementSqFtAdjBuiltIn
	}
	if input.TxtListComp2BasementSqFt != nil {
		updateDoc.TxtListComp2BasementSqFt = input.TxtListComp2BasementSqFt
	}
	if input.TxtListComp2BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtListComp2BasementSqFtAdjBuiltIn = input.TxtListComp2BasementSqFtAdjBuiltIn
	}
	if input.TxtListComp3BasementSqFt != nil {
		updateDoc.TxtListComp3BasementSqFt = input.TxtListComp3BasementSqFt
	}
	if input.TxtListComp3BasementSqFtAdjBuiltIn != nil {
		updateDoc.TxtListComp3BasementSqFtAdjBuiltIn = input.TxtListComp3BasementSqFtAdjBuiltIn
	}
	if input.TxtSubjectOriginalListDate != nil {
		updateDoc.TxtSubjectOriginalListDate = input.TxtSubjectOriginalListDate
	}
	if input.TxtSaleComp1OriginalListDate != nil {
		updateDoc.TxtSaleComp1OriginalListDate = input.TxtSaleComp1OriginalListDate
	}
	if input.TxtSaleComp2OriginalListDate != nil {
		updateDoc.TxtSaleComp2OriginalListDate = input.TxtSaleComp2OriginalListDate
	}
	if input.TxtSaleComp3OriginalListDate != nil {
		updateDoc.TxtSaleComp3OriginalListDate = input.TxtSaleComp3OriginalListDate
	}
	if input.TxtListComp1OriginalListDate != nil {
		updateDoc.TxtListComp1OriginalListDate = input.TxtListComp1OriginalListDate
	}
	if input.TxtListComp2OriginalListDate != nil {
		updateDoc.TxtListComp2OriginalListDate = input.TxtListComp2OriginalListDate
	}
	if input.TxtListComp3OriginalListDate != nil {
		updateDoc.TxtListComp3OriginalListDate = input.TxtListComp3OriginalListDate
	}
	if input.TxtSubjectCurrentListDate != nil {
		updateDoc.TxtSubjectCurrentListDate = input.TxtSubjectCurrentListDate
	}
	if input.TxtSaleComp1CurrentListDate != nil {
		updateDoc.TxtSaleComp1CurrentListDate = input.TxtSaleComp1CurrentListDate
	}
	if input.TxtSaleComp2CurrentListDate != nil {
		updateDoc.TxtSaleComp2CurrentListDate = input.TxtSaleComp2CurrentListDate
	}
	if input.TxtSaleComp3CurrentListDate != nil {
		updateDoc.TxtSaleComp3CurrentListDate = input.TxtSaleComp3CurrentListDate
	}
	if input.TxtListComp1CurrentListDate != nil {
		updateDoc.TxtListComp1CurrentListDate = input.TxtListComp1CurrentListDate
	}
	if input.TxtListComp2CurrentListDate != nil {
		updateDoc.TxtListComp2CurrentListDate = input.TxtListComp2CurrentListDate
	}
	if input.TxtListComp3CurrentListDate != nil {
		updateDoc.TxtListComp3CurrentListDate = input.TxtListComp3CurrentListDate
	}
	if input.TxtSubjectOriginalListPrice != nil {
		updateDoc.TxtSubjectOriginalListPrice = input.TxtSubjectOriginalListPrice
	}
	if input.TxtSaleComp1OriginalListPrice != nil {
		updateDoc.TxtSaleComp1OriginalListPrice = input.TxtSaleComp1OriginalListPrice
	}
	if input.TxtSaleComp2OriginalListPrice != nil {
		updateDoc.TxtSaleComp2OriginalListPrice = input.TxtSaleComp2OriginalListPrice
	}
	if input.TxtSaleComp3OriginalListPrice != nil {
		updateDoc.TxtSaleComp3OriginalListPrice = input.TxtSaleComp3OriginalListPrice
	}
	if input.TxtListComp1OriginalListPrice != nil {
		updateDoc.TxtListComp1OriginalListPrice = input.TxtListComp1OriginalListPrice
	}
	if input.TxtListComp2OriginalListPrice != nil {
		updateDoc.TxtListComp2OriginalListPrice = input.TxtListComp2OriginalListPrice
	}
	if input.TxtListComp3OriginalListPrice != nil {
		updateDoc.TxtListComp3OriginalListPrice = input.TxtListComp3OriginalListPrice
	}
	if input.TxtSubjectListPrice != nil {
		updateDoc.TxtSubjectListPrice = input.TxtSubjectListPrice
	}
	if input.TxtSaleComp1ListPrice != nil {
		updateDoc.TxtSaleComp1ListPrice = input.TxtSaleComp1ListPrice
	}
	if input.TxtSaleComp2ListPrice != nil {
		updateDoc.TxtSaleComp2ListPrice = input.TxtSaleComp2ListPrice
	}
	if input.TxtSaleComp3ListPrice != nil {
		updateDoc.TxtSaleComp3ListPrice = input.TxtSaleComp3ListPrice
	}
	if input.TxtListComp1ListPrice != nil {
		updateDoc.TxtListComp1ListPrice = input.TxtListComp1ListPrice
	}
	if input.TxtListComp2ListPrice != nil {
		updateDoc.TxtListComp2ListPrice = input.TxtListComp2ListPrice
	}
	if input.TxtListComp3ListPrice != nil {
		updateDoc.TxtListComp3ListPrice = input.TxtListComp3ListPrice
	}
	if input.TxtSubjectSalePrice != nil {
		updateDoc.TxtSubjectSalePrice = input.TxtSubjectSalePrice
	}
	if input.TxtSaleComp1SalePrice != nil {
		updateDoc.TxtSaleComp1SalePrice = input.TxtSaleComp1SalePrice
	}
	if input.TxtSaleComp2SalePrice != nil {
		updateDoc.TxtSaleComp2SalePrice = input.TxtSaleComp2SalePrice
	}
	if input.TxtSaleComp3SalePrice != nil {
		updateDoc.TxtSaleComp3SalePrice = input.TxtSaleComp3SalePrice
	}
	if input.TxtSubjectSaleDate != nil {
		updateDoc.TxtSubjectSaleDate = input.TxtSubjectSaleDate
	}
	if input.TxtSaleComp1SaleDate != nil {
		updateDoc.TxtSaleComp1SaleDate = input.TxtSaleComp1SaleDate
	}
	if input.TxtSaleComp2SaleDate != nil {
		updateDoc.TxtSaleComp2SaleDate = input.TxtSaleComp2SaleDate
	}
	if input.TxtSaleComp3SaleDate != nil {
		updateDoc.TxtSaleComp3SaleDate = input.TxtSaleComp3SaleDate
	}
	if input.CmbSubjectFinancing != nil {
		updateDoc.CmbSubjectFinancing = input.CmbSubjectFinancing
	}
	if input.CmbSaleComp1Financing != nil {
		updateDoc.CmbSaleComp1Financing = input.CmbSaleComp1Financing
	}
	if input.CmbSaleComp2Financing != nil {
		updateDoc.CmbSaleComp2Financing = input.CmbSaleComp2Financing
	}
	if input.CmbSaleComp3Financing != nil {
		updateDoc.CmbSaleComp3Financing = input.CmbSaleComp3Financing
	}
	if input.CmbListComp1Financing != nil {
		updateDoc.CmbListComp1Financing = input.CmbListComp1Financing
	}
	if input.CmbListComp2Financing != nil {
		updateDoc.CmbListComp2Financing = input.CmbListComp2Financing
	}
	if input.CmbListComp3Financing != nil {
		updateDoc.CmbListComp3Financing = input.CmbListComp3Financing
	}
	if input.TxtSubjectDom != nil {
		updateDoc.TxtSubjectDom = input.TxtSubjectDom
	}
	if input.TxtSaleComp1dom != nil {
		updateDoc.TxtSaleComp1dom = input.TxtSaleComp1dom
	}
	if input.TxtSaleComp2dom != nil {
		updateDoc.TxtSaleComp2dom = input.TxtSaleComp2dom
	}
	if input.TxtSaleComp3dom != nil {
		updateDoc.TxtSaleComp3dom = input.TxtSaleComp3dom
	}
	if input.TxtListComp1dom != nil {
		updateDoc.TxtListComp1dom = input.TxtListComp1dom
	}
	if input.TxtListComp2dom != nil {
		updateDoc.TxtListComp2dom = input.TxtListComp2dom
	}
	if input.TxtListComp3dom != nil {
		updateDoc.TxtListComp3dom = input.TxtListComp3dom
	}
	if input.TxtSubjectPricePerSqFt != nil {
		updateDoc.TxtSubjectPricePerSqFt = input.TxtSubjectPricePerSqFt
	}
	if input.TxtSaleComp1PricePerSqFt != nil {
		updateDoc.TxtSaleComp1PricePerSqFt = input.TxtSaleComp1PricePerSqFt
	}
	if input.TxtSaleComp2PricePerSqFt != nil {
		updateDoc.TxtSaleComp2PricePerSqFt = input.TxtSaleComp2PricePerSqFt
	}
	if input.TxtSaleComp3PricePerSqFt != nil {
		updateDoc.TxtSaleComp3PricePerSqFt = input.TxtSaleComp3PricePerSqFt
	}
	if input.TxtListComp1PricePerSqFt != nil {
		updateDoc.TxtListComp1PricePerSqFt = input.TxtListComp1PricePerSqFt
	}
	if input.TxtListComp2PricePerSqFt != nil {
		updateDoc.TxtListComp2PricePerSqFt = input.TxtListComp2PricePerSqFt
	}
	if input.TxtListComp3PricePerSqFt != nil {
		updateDoc.TxtListComp3PricePerSqFt = input.TxtListComp3PricePerSqFt
	}
	if input.TxtSubjectAdjustments != nil {
		updateDoc.TxtSubjectAdjustments = input.TxtSubjectAdjustments
	}
	if input.TxtSaleComp1Adjustments != nil {
		updateDoc.TxtSaleComp1Adjustments = input.TxtSaleComp1Adjustments
	}
	if input.TxtSaleComp2Adjustments != nil {
		updateDoc.TxtSaleComp2Adjustments = input.TxtSaleComp2Adjustments
	}
	if input.TxtSaleComp3Adjustments != nil {
		updateDoc.TxtSaleComp3Adjustments = input.TxtSaleComp3Adjustments
	}
	if input.TxtListComp1Adjustments != nil {
		updateDoc.TxtListComp1Adjustments = input.TxtListComp1Adjustments
	}
	if input.TxtListComp2Adjustments != nil {
		updateDoc.TxtListComp2Adjustments = input.TxtListComp2Adjustments
	}
	if input.TxtListComp3Adjustments != nil {
		updateDoc.TxtListComp3Adjustments = input.TxtListComp3Adjustments
	}
	if input.TxtSubjectCompTotals != nil {
		updateDoc.TxtSubjectCompTotals = input.TxtSubjectCompTotals
	}
	if input.TxtSaleComp1CompTotals != nil {
		updateDoc.TxtSaleComp1CompTotals = input.TxtSaleComp1CompTotals
	}
	if input.TxtSaleComp2CompTotals != nil {
		updateDoc.TxtSaleComp2CompTotals = input.TxtSaleComp2CompTotals
	}
	if input.TxtSaleComp3CompTotals != nil {
		updateDoc.TxtSaleComp3CompTotals = input.TxtSaleComp3CompTotals
	}
	if input.TxtListComp1CompTotals != nil {
		updateDoc.TxtListComp1CompTotals = input.TxtListComp1CompTotals
	}
	if input.TxtListComp2CompTotals != nil {
		updateDoc.TxtListComp2CompTotals = input.TxtListComp2CompTotals
	}
	if input.TxtListComp3CompTotals != nil {
		updateDoc.TxtListComp3CompTotals = input.TxtListComp3CompTotals
	}
	if input.CmbListComp1CommentType != nil {
		updateDoc.CmbListComp1CommentType = input.CmbListComp1CommentType
	}
	if input.TxtListComp1ComparableComments != nil {
		updateDoc.TxtListComp1ComparableComments = input.TxtListComp1ComparableComments
	}
	if input.TxtListComp1FormatAdjustments != nil {
		updateDoc.TxtListComp1FormatAdjustments = input.TxtListComp1FormatAdjustments
	}
	if input.TxtListComp1MLSComments != nil {
		updateDoc.TxtListComp1MLSComments = input.TxtListComp1MLSComments
	}
	if input.CmbListComp2CommentType != nil {
		updateDoc.CmbListComp2CommentType = input.CmbListComp2CommentType
	}
	if input.TxtListComp2ComparableComments != nil {
		updateDoc.TxtListComp2ComparableComments = input.TxtListComp2ComparableComments
	}
	if input.TxtListComp2FormatAdjustments != nil {
		updateDoc.TxtListComp2FormatAdjustments = input.TxtListComp2FormatAdjustments
	}
	if input.TxtListComp2MLSComments != nil {
		updateDoc.TxtListComp2MLSComments = input.TxtListComp2MLSComments
	}
	if input.CmbListComp3CommentType != nil {
		updateDoc.CmbListComp3CommentType = input.CmbListComp3CommentType
	}
	if input.TxtListComp3ComparableComments != nil {
		updateDoc.TxtListComp3ComparableComments = input.TxtListComp3ComparableComments
	}
	if input.TxtListComp3FormatAdjustments != nil {
		updateDoc.TxtListComp3FormatAdjustments = input.TxtListComp3FormatAdjustments
	}
	if input.TxtListComp3MLSComments != nil {
		updateDoc.TxtListComp3MLSComments = input.TxtListComp3MLSComments
	}
	if input.CmbSaleComp1CommentType != nil {
		updateDoc.CmbSaleComp1CommentType = input.CmbSaleComp1CommentType
	}
	if input.TxtSaleComp1ComparableComments != nil {
		updateDoc.TxtSaleComp1ComparableComments = input.TxtSaleComp1ComparableComments
	}
	if input.TxtSaleComp1FormatAdjustments != nil {
		updateDoc.TxtSaleComp1FormatAdjustments = input.TxtSaleComp1FormatAdjustments
	}
	if input.TxtSaleComp1MLSComments != nil {
		updateDoc.TxtSaleComp1MLSComments = input.TxtSaleComp1MLSComments
	}
	if input.CmbSaleComp2CommentType != nil {
		updateDoc.CmbSaleComp2CommentType = input.CmbSaleComp2CommentType
	}
	if input.TxtSaleComp2ComparableComments != nil {
		updateDoc.TxtSaleComp2ComparableComments = input.TxtSaleComp2ComparableComments
	}
	if input.TxtSaleComp2FormatAdjustments != nil {
		updateDoc.TxtSaleComp2FormatAdjustments = input.TxtSaleComp2FormatAdjustments
	}
	if input.TxtSaleComp2MLSComments != nil {
		updateDoc.TxtSaleComp2MLSComments = input.TxtSaleComp2MLSComments
	}
	if input.CmbSaleComp3CommentType != nil {
		updateDoc.CmbSaleComp3CommentType = input.CmbSaleComp3CommentType
	}
	if input.TxtSaleComp3ComparableComments != nil {
		updateDoc.TxtSaleComp3ComparableComments = input.TxtSaleComp3ComparableComments
	}
	if input.TxtSaleComp3FormatAdjustments != nil {
		updateDoc.TxtSaleComp3FormatAdjustments = input.TxtSaleComp3FormatAdjustments
	}
	if input.TxtSaleComp3MLSComments != nil {
		updateDoc.TxtSaleComp3MLSComments = input.TxtSaleComp3MLSComments
	}
	if input.CmbNeighborhoodTrend != nil {
		updateDoc.CmbNeighborhoodTrend = input.CmbNeighborhoodTrend
	}
	if input.TxtMonthlyPecent != nil {
		updateDoc.TxtMonthlyPecent = input.TxtMonthlyPecent
	}
	if input.TxtEstimatedRent != nil {
		updateDoc.TxtEstimatedRent = input.TxtEstimatedRent
	}
	if input.TxtEstimatedDaysOnMarket != nil {
		updateDoc.TxtEstimatedDaysOnMarket = input.TxtEstimatedDaysOnMarket
	}
	if input.TxtNoBoarded != nil {
		updateDoc.TxtNoBoarded = input.TxtNoBoarded
	}
	if input.TxtNoOfActive != nil {
		updateDoc.TxtNoOfActive = input.TxtNoOfActive
	}
	if input.Txt6MonthPecent != nil {
		updateDoc.Txt6MonthPecent = input.Txt6MonthPecent
	}
	if input.TxtAnnualPecent != nil {
		updateDoc.TxtAnnualPecent = input.TxtAnnualPecent
	}
	if input.TxtListings != nil {
		updateDoc.TxtListings = input.TxtListings
	}
	if input.CmbSupply != nil {
		updateDoc.CmbSupply = input.CmbSupply
	}
	if input.TxtListingsMinValue != nil {
		updateDoc.TxtListingsMinValue = input.TxtListingsMinValue
	}
	if input.TxtListingsRange1 != nil {
		updateDoc.TxtListingsRange1 = input.TxtListingsRange1
	}
	if input.TxtListingsMedValue != nil {
		updateDoc.TxtListingsMedValue = input.TxtListingsMedValue
	}
	if input.TxtListingsMaxValue != nil {
		updateDoc.TxtListingsMaxValue = input.TxtListingsMaxValue
	}
	if input.TxtListingsRange2 != nil {
		updateDoc.TxtListingsRange2 = input.TxtListingsRange2
	}
	if input.TxtListingsDom != nil {
		updateDoc.TxtListingsDom = input.TxtListingsDom
	}
	if input.TxtListingsDOMRange1 != nil {
		updateDoc.TxtListingsDOMRange1 = input.TxtListingsDOMRange1
	}
	if input.TxtListingsDOMRange2 != nil {
		updateDoc.TxtListingsDOMRange2 = input.TxtListingsDOMRange2
	}
	if input.CmbREOTrend != nil {
		updateDoc.CmbREOTrend = input.CmbREOTrend
	}
	if input.TxtNoOfFm != nil {
		updateDoc.TxtNoOfFm = input.TxtNoOfFm
	}
	if input.TxtNoOfSs != nil {
		updateDoc.TxtNoOfSs = input.TxtNoOfSs
	}
	if input.TxtNoOfReo != nil {
		updateDoc.TxtNoOfReo = input.TxtNoOfReo
	}
	if input.TxtNoOfDistressed != nil {
		updateDoc.TxtNoOfDistressed = input.TxtNoOfDistressed
	}
	if input.TxtSales != nil {
		updateDoc.TxtSales = input.TxtSales
	}
	if input.CmbDemand != nil {
		updateDoc.CmbDemand = input.CmbDemand
	}
	if input.TxtSalesRange1 != nil {
		updateDoc.TxtSalesRange1 = input.TxtSalesRange1
	}
	if input.TxtSalesMedValue != nil {
		updateDoc.TxtSalesMedValue = input.TxtSalesMedValue
	}
	if input.TxtSalesRange2 != nil {
		updateDoc.TxtSalesRange2 = input.TxtSalesRange2
	}
	if input.TxtSalesDom != nil {
		updateDoc.TxtSalesDom = input.TxtSalesDom
	}
	if input.TxtSalesDOMRange1 != nil {
		updateDoc.TxtSalesDOMRange1 = input.TxtSalesDOMRange1
	}
	if input.TxtSalesDOMRange2 != nil {
		updateDoc.TxtSalesDOMRange2 = input.TxtSalesDOMRange2
	}
	if input.TxtZillowNeighborhoodTrend != nil {
		updateDoc.TxtZillowNeighborhoodTrend = input.TxtZillowNeighborhoodTrend
	}
	if input.TxtNeighborhoodTrendComments != nil {
		updateDoc.TxtNeighborhoodTrendComments = input.TxtNeighborhoodTrendComments
	}
	if input.TxtTotalListings != nil {
		updateDoc.TxtTotalListings = input.TxtTotalListings
	}
	if input.TxtTotalSales != nil {
		updateDoc.TxtTotalSales = input.TxtTotalSales
	}
	if input.TxtNoOfREOListings != nil {
		updateDoc.TxtNoOfREOListings = input.TxtNoOfREOListings
	}
	if input.TxtNoOfSSListings != nil {
		updateDoc.TxtNoOfSSListings = input.TxtNoOfSSListings
	}
	if input.TxtNoOfREOSales != nil {
		updateDoc.TxtNoOfREOSales = input.TxtNoOfREOSales
	}
	if input.TxtNoOfSSSales != nil {
		updateDoc.TxtNoOfSSSales = input.TxtNoOfSSSales
	}
	if input.TxtTaxID != nil {
		updateDoc.TxtTaxID = input.TxtTaxID
	}
	if input.TxtLastSaleDate != nil {
		updateDoc.TxtLastSaleDate = input.TxtLastSaleDate
	}
	if input.TxtLastSalePrice != nil {
		updateDoc.TxtLastSalePrice = input.TxtLastSalePrice
	}
	if input.CmbIsListed != nil {
		updateDoc.CmbIsListed = input.CmbIsListed
	}
	if input.TxtOwnerOccupied != nil {
		updateDoc.TxtOwnerOccupied = input.TxtOwnerOccupied
	}
	if input.TxtRenterOccupied != nil {
		updateDoc.TxtRenterOccupied = input.TxtRenterOccupied
	}
	if input.TxtMarketRent != nil {
		updateDoc.TxtMarketRent = input.TxtMarketRent
	}
	if input.TxtNoOfRentals != nil {
		updateDoc.TxtNoOfRentals = input.TxtNoOfRentals
	}
	if input.TxtTypicalDom != nil {
		updateDoc.TxtTypicalDom = input.TxtTypicalDom
	}
	if input.TxtNoRentHomes != nil {
		updateDoc.TxtNoRentHomes = input.TxtNoRentHomes
	}
	if input.TxtTypicalRentalRates != nil {
		updateDoc.TxtTypicalRentalRates = input.TxtTypicalRentalRates
	}
	if input.AdjustmentPrice != nil {
		updateDoc.AdjustmentPrice = input.AdjustmentPrice
	}
	if input.TxtCalculatedGla != nil {
		updateDoc.TxtCalculatedGla = input.TxtCalculatedGla
	}
	if input.TxtCalculatedAge != nil {
		updateDoc.TxtCalculatedAge = input.TxtCalculatedAge
	}
	if input.TxtCalculatedSaleDates != nil {
		updateDoc.TxtCalculatedSaleDates = input.TxtCalculatedSaleDates
	}
	if input.TxtCalculatedProximity != nil {
		updateDoc.TxtCalculatedProximity = input.TxtCalculatedProximity
	}
	if input.TxtCalculatedStyle != nil {
		updateDoc.TxtCalculatedStyle = input.TxtCalculatedStyle
	}
	if input.TxtCalculatedMonthsSupply != nil {
		updateDoc.TxtCalculatedMonthsSupply = input.TxtCalculatedMonthsSupply
	}
	if input.TxtCalculatedProxim != nil {
		updateDoc.TxtCalculatedProxim = input.TxtCalculatedProxim
	}
	if input.TxtCalculatedGLAs != nil {
		updateDoc.TxtCalculatedGLAs = input.TxtCalculatedGLAs
	}
	if input.TxtCalculatedAges != nil {
		updateDoc.TxtCalculatedAges = input.TxtCalculatedAges
	}
	if input.TxtCalculatedCond != nil {
		updateDoc.TxtCalculatedCond = input.TxtCalculatedCond
	}
	if input.TxtCalculatedView != nil {
		updateDoc.TxtCalculatedView = input.TxtCalculatedView
	}
	if input.TxtCalculatedStyle1 != nil {
		updateDoc.TxtCalculatedStyle1 = input.TxtCalculatedStyle1
	}
	if input.TxtCalculatedLots != nil {
		updateDoc.TxtCalculatedLots = input.TxtCalculatedLots
	}
	if input.TxtCalculatedBeds != nil {
		updateDoc.TxtCalculatedBeds = input.TxtCalculatedBeds
	}
	if input.TxtCalculatedBath != nil {
		updateDoc.TxtCalculatedBath = input.TxtCalculatedBath
	}
	if input.Rdbresaletext != nil {
		updateDoc.Rdbresaletext = input.Rdbresaletext
	}
	if input.Rdbmarketedtext != nil {
		updateDoc.Rdbmarketedtext = input.Rdbmarketedtext
	}
	if input.Txtpmi != nil {
		updateDoc.Txtpmi = input.Txtpmi
	}
	if input.TxtOtherComments != nil {
		updateDoc.TxtOtherComments = input.TxtOtherComments
	}
	if input.Txtcbnew != nil {
		updateDoc.Txtcbnew = input.Txtcbnew
	}
	if input.Txtcbold != nil {
		updateDoc.Txtcbold = input.Txtcbold
	}
	if input.Txtcbstyle != nil {
		updateDoc.Txtcbstyle = input.Txtcbstyle
	}
	if input.Txtcblot != nil {
		updateDoc.Txtcblot = input.Txtcblot
	}
	if input.Txtcbview != nil {
		updateDoc.Txtcbview = input.Txtcbview
	}
	if input.Txtcbdamage != nil {
		updateDoc.Txtcbdamage = input.Txtcbdamage
	}
	if input.Txtcbupgrade != nil {
		updateDoc.Txtcbupgrade = input.Txtcbupgrade
	}
	if input.Txtcbinfluence != nil {
		updateDoc.Txtcbinfluence = input.Txtcbinfluence
	}
	if input.TxtSubjectComments != nil {
		updateDoc.TxtSubjectComments = input.TxtSubjectComments
	}
	if input.TxtNeighborhoodComments != nil {
		updateDoc.TxtNeighborhoodComments = input.TxtNeighborhoodComments
	}
	if input.TxtNeighborhoodTrend != nil {
		updateDoc.TxtNeighborhoodTrend = input.TxtNeighborhoodTrend
	}
	if input.TxtValidation1 != nil {
		updateDoc.TxtValidation1 = input.TxtValidation1
	}
	if input.TxtUniqueComments != nil {
		updateDoc.TxtUniqueComments = input.TxtUniqueComments
	}
	if input.TxtMarketingStrategy != nil {
		updateDoc.TxtMarketingStrategy = input.TxtMarketingStrategy
	}
	if input.TxtDisclaimer2 != nil {
		updateDoc.TxtDisclaimer2 = input.TxtDisclaimer2
	}
	if input.TxtBrokerComments != nil {
		updateDoc.TxtBrokerComments = input.TxtBrokerComments
	}
	if input.TxtValidation != nil {
		updateDoc.TxtValidation = input.TxtValidation
	}
	if input.Txt30DayQuickSale != nil {
		updateDoc.Txt30DayQuickSale = input.Txt30DayQuickSale
	}
	if input.Txt60DayQuickSale != nil {
		updateDoc.Txt60DayQuickSale = input.Txt60DayQuickSale
	}
	if input.Txt90DayAsIsValue != nil {
		updateDoc.Txt90DayAsIsValue = input.Txt90DayAsIsValue
	}
	if input.Txt120DayQuickSale != nil {
		updateDoc.Txt120DayQuickSale = input.Txt120DayQuickSale
	}
	if input.Txt180DayQuickSale != nil {
		updateDoc.Txt180DayQuickSale = input.Txt180DayQuickSale
	}
	if input.TxtListPriceFinalValues != nil {
		updateDoc.TxtListPriceFinalValues = input.TxtListPriceFinalValues
	}
	if input.Txt30DayListPriceFinalValues != nil {
		updateDoc.Txt30DayListPriceFinalValues = input.Txt30DayListPriceFinalValues
	}
	if input.Txt30DayQuickSaleRepaired != nil {
		updateDoc.Txt30DayQuickSaleRepaired = input.Txt30DayQuickSaleRepaired
	}
	if input.Txt60DayQuickSaleRepaired != nil {
		updateDoc.Txt60DayQuickSaleRepaired = input.Txt60DayQuickSaleRepaired
	}
	if input.Txt90DayAsIsValueRepaired != nil {
		updateDoc.Txt90DayAsIsValueRepaired = input.Txt90DayAsIsValueRepaired
	}
	if input.Txt120DayQuickSaleRepaired != nil {
		updateDoc.Txt120DayQuickSaleRepaired = input.Txt120DayQuickSaleRepaired
	}
	if input.Txt180DayQuickSaleRepaired != nil {
		updateDoc.Txt180DayQuickSaleRepaired = input.Txt180DayQuickSaleRepaired
	}
	if input.TxtListPriceRepaired != nil {
		updateDoc.TxtListPriceRepaired = input.TxtListPriceRepaired
	}
	if input.Txt30DayListPriceRepaired != nil {
		updateDoc.Txt30DayListPriceRepaired = input.Txt30DayListPriceRepaired
	}
	if input.CmbHouse != nil {
		updateDoc.CmbHouse = input.CmbHouse
	}
	if input.CmbPositive != nil {
		updateDoc.CmbPositive = input.CmbPositive
	}
	if input.CmbNegative != nil {
		updateDoc.CmbNegative = input.CmbNegative
	}
	if input.CmbView != nil {
		updateDoc.CmbView = input.CmbView
	}
	if input.CmbMarket != nil {
		updateDoc.CmbMarket = input.CmbMarket
	}
	if input.CmbPricing != nil {
		updateDoc.CmbPricing = input.CmbPricing
	}
	if input.CmbListing != nil {
		updateDoc.CmbListing = input.CmbListing
	}
	if input.CmbExtra != nil {
		updateDoc.CmbExtra = input.CmbExtra
	}
	if input.TxtUnique != nil {
		updateDoc.TxtUnique = input.TxtUnique
	}
	if input.PriceComment != nil {
		updateDoc.PriceComment = input.PriceComment
	}
	if input.RangeComment != nil {
		updateDoc.RangeComment = input.RangeComment
	}
	if input.ProxException != nil {
		updateDoc.ProxException = input.ProxException
	}
	if input.GlaException != nil {
		updateDoc.GlaException = input.GlaException
	}
	if input.AgeException != nil {
		updateDoc.AgeException = input.AgeException
	}
	if input.CondException != nil {
		updateDoc.CondException = input.CondException
	}
	if input.ViewException != nil {
		updateDoc.ViewException = input.ViewException
	}
	if input.StyleException != nil {
		updateDoc.StyleException = input.StyleException
	}
	if input.LotException != nil {
		updateDoc.LotException = input.LotException
	}
	if input.BedException != nil {
		updateDoc.BedException = input.BedException
	}
	if input.BathException != nil {
		updateDoc.BathException = input.BathException
	}

	//Todo, Create History
	fileDir, err := pdf.Create(iformRaw)
	if err != nil {
		return false, err
	}
	//upload to s3

	s3FileUrl, err := awsS3.S3UploaderFromFile(*fileDir)
	if err != nil {
		return false, err
	}
	logRaw := makeHistoryLogRaw("pdffile.pdf", *s3FileUrl, updatedBy)

	update := bson.M{
		"$set": updateDoc,
		"$addToSet": bson.M{
			"logs": logRaw,
		},
	}
	res, err := DbCollections.Iforms.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	if res.ModifiedCount < 1 {
		return false, errs.NoRecordUpdate
	}
	return true, nil
}

func GetIformByPipelineId(ctx context.Context, pipelineId string) (*models.Iform, error) {
	filter := bson.D{{"pipelineId", pipelineId}}
	a := &Iform{}
	//err := DbCollections.Iforms.FindOne(ctx, filter).Decode(a)
	resRaw := DbCollections.Iforms.FindOne(ctx, filter)
	if resRaw.Err() != nil {
		log.Error("Failed to query iform: %v", resRaw.Err())
		return nil, resRaw.Err()
	}
	err := resRaw.Decode(a)
	if err != nil {
		log.Error("Failed to query iform: %v", err)
		return nil, errs.NoRecordUpdate
	}
	return a.ToModels(), nil
}

func SearchIforms(ctx context.Context, offset, limit int, clientID, address *string) ([]*Iform, error) {

	filter := bson.M{}
	if clientID != nil {
		filter["clientId"] = *clientID
	}
	if address != nil {
		filter["txtSubjectAddress"] = *address
	}
	pipe := []bson.M{
		{"$match": filter},
		{"$sort": bson.M{"lastUpdateTime": -1}},
		{"$skip": offset},
		{"$limit": limit},
	}

	cur, err := DbCollections.Iforms.Aggregate(ctx, pipe)
	if err != nil {
		log.Debug("Failed to query iform:  %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	rawIforms := make([]*Iform, 0)
	for cur.Next(ctx) {
		a := &Iform{}
		if err := cur.Decode(a); err != nil {
			log.Debug("Failed to decode iform entry: %v", err)
			return nil, err
		}
		rawIforms = append(rawIforms, a)
	}

	return rawIforms, nil
}

func GetIformByOrderNumber(ctx context.Context, orderNumber string) (*models.Iform, error) {
	filter := bson.D{{"txtOrderNumber", orderNumber}}
	a := &Iform{}
	//err := DbCollections.Iforms.FindOne(ctx, filter).Decode(a)
	resRaw := DbCollections.Iforms.FindOne(ctx, filter)
	if resRaw.Err() != nil {
		log.Error("Failed to query iform: %v", resRaw.Err())
		return nil, resRaw.Err()
	}
	err := resRaw.Decode(a)
	if err != nil {
		log.Error("Failed to query iform: %v", err)
		return nil, errs.NoRecordUpdate
	}
	return a.ToModels(), nil
}

