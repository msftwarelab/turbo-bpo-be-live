package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Iform struct {
	ID                                    primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	LastUpdateTime                        *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
	FormType                              *string             `bson:"formType,omitempty"`
	PipelineID                            *string             `bson:"pipelineId,omitempty"`
	ClientID                              *string             `bson:"clientId,omitempty"`
	TxtClient                             *string             `bson:"txtClient,omitempty"`
	TxtCompany                            *string             `bson:"txtCompany,omitempty"`
	TxtOrderNumber                        *string             `bson:"txtOrderNumber,omitempty"`
	CmbOrderType                          *string             `bson:"cmbOrderType,omitempty"`
	TxtAddress                            *string             `bson:"txtAddress,omitempty"`
	TxtLocation                           *string             `bson:"txtLocation,omitempty"`
	TxtBrokerChecker                      *string             `bson:"txtBrokerChecker,omitempty"`
	TxtPreparerInfoAgent                  *string             `bson:"txtPreparerInfoAgent,omitempty"`
	TxtPreparerInfoAgentLicense           *string             `bson:"txtPreparerInfoAgentLicense,omitempty"`
	TxtPreparerInfoBroker                 *string             `bson:"txtPreparerInfoBroker,omitempty"`
	TxtPreparerInfoBrokerLicense          *string             `bson:"txtPreparerInfoBrokerLicense,omitempty"`
	TxtPreparerInfoAddress                *string             `bson:"txtPreparerInfoAddress,omitempty"`
	TxtPreparerInfoBrokerage              *string             `bson:"txtPreparerInfoBrokerage,omitempty"`
	TxtPreparerInfoAgentCompany           *string             `bson:"txtPreparerInfoAgentCompany,omitempty"`
	TxtPreparerInfoPhone                  *string             `bson:"txtPreparerInfoPhone,omitempty"`
	TxtPreparerInfoYearsOfExperience      *string             `bson:"txtPreparerInfoYearsOfExperience,omitempty"`
	TxtPreparerInfoEmail                  *string             `bson:"txtPreparerInfoEmail,omitempty"`
	TxtSubjectAddress                     *string             `bson:"txtSubjectAddress,omitempty"`
	TxtPreparerInfoMilesAwayFromSubject   *string             `bson:"txtPreparerInfoMilesAwayFromSubject,omitempty"`
	TxtAgentZip                           *string             `bson:"txtAgentZip,omitempty"`
	TxtAgentCity                          *string             `bson:"txtAgentCity,omitempty"`
	TxtAgentState                         *string             `bson:"txtAgentState,omitempty"`
	TxtDisclaimer                         *string             `bson:"txtDisclaimer,omitempty"`
	CmbLocation                           *string             `bson:"cmbLocation,omitempty"`
	TxtCounty                             *string             `bson:"txtCounty,omitempty"`
	TxtTrullia                            *string             `bson:"txtTrullia,omitempty"`
	TxtZillow                             *string             `bson:"txtZillow,omitempty"`
	TxtFindcompsnow                       *string             `bson:"txtFindcompsnow,omitempty"`
	TxtAverage                            *string             `bson:"txtAverage,omitempty"`
	CmbForm                               *string             `bson:"cmbForm,omitempty"`
	CmbForm2                              *string             `bson:"cmbForm2,omitempty"`
	TxtSaleComp1Address                   *string             `bson:"txtSaleComp1Address,omitempty"`
	TxtSaleComp2Address                   *string             `bson:"txtSaleComp2Address,omitempty"`
	TxtSaleComp3Address                   *string             `bson:"txtSaleComp3Address,omitempty"`
	TxtListComp1Address                   *string             `bson:"txtListComp1Address,omitempty"`
	TxtListComp2Address                   *string             `bson:"txtListComp2Address,omitempty"`
	TxtListComp3Address                   *string             `bson:"txtListComp3Address,omitempty"`
	TxtSubjectState                       *string             `bson:"txtSubjectState,omitempty"`
	TxtSaleComp1State                     *string             `bson:"txtSaleComp1State,omitempty"`
	TxtSaleComp2State                     *string             `bson:"txtSaleComp2State,omitempty"`
	TxtSaleComp3State                     *string             `bson:"txtSaleComp3State,omitempty"`
	TxtListComp1State                     *string             `bson:"txtListComp1State,omitempty"`
	TxtListComp2State                     *string             `bson:"txtListComp2State,omitempty"`
	TxtListComp3State                     *string             `bson:"txtListComp3State,omitempty"`
	TxtSubjectCity                        *string             `bson:"txtSubjectCity,omitempty"`
	TxtSaleComp1City                      *string             `bson:"txtSaleComp1City,omitempty"`
	TxtSaleComp2City                      *string             `bson:"txtSaleComp2City,omitempty"`
	TxtSaleComp3City                      *string             `bson:"txtSaleComp3City,omitempty"`
	TxtListComp1City                      *string             `bson:"txtListComp1City,omitempty"`
	TxtListComp2City                      *string             `bson:"txtListComp2City,omitempty"`
	TxtListComp3City                      *string             `bson:"txtListComp3City,omitempty"`
	TxtSubjectnoUnit                      *string             `bson:"txtSubjectnoUnit,omitempty"`
	TxtSubjectUnitNo                      *string             `bson:"txtSubjectUnitNo,omitempty"`
	TxtSaleComp1noUnit                    *string             `bson:"txtSaleComp1noUnit,omitempty"`
	TxtSaleComp1UnitNo                    *string             `bson:"txtSaleComp1UnitNo,omitempty"`
	TxtSaleComp2noUnit                    *string             `bson:"txtSaleComp2noUnit,omitempty"`
	TxtSaleComp2UnitNo                    *string             `bson:"txtSaleComp2UnitNo,omitempty"`
	TxtSaleComp3noUnit                    *string             `bson:"txtSaleComp3noUnit,omitempty"`
	TxtSaleComp3UnitNo                    *string             `bson:"txtSaleComp3UnitNo,omitempty"`
	TxtListComp1noUnit                    *string             `bson:"txtListComp1noUnit,omitempty"`
	TxtListComp1UnitNo                    *string             `bson:"txtListComp1UnitNo,omitempty"`
	TxtListComp2noUnit                    *string             `bson:"txtListComp2noUnit,omitempty"`
	TxtListComp2UnitNo                    *string             `bson:"txtListComp2UnitNo,omitempty"`
	TxtListComp3noUnit                    *string             `bson:"txtListComp3noUnit,omitempty"`
	TxtListComp3UnitNo                    *string             `bson:"txtListComp3UnitNo,omitempty"`
	TxtSubjectUnits                       *string             `bson:"txtSubjectUnits,omitempty"`
	TxtSaleComp1Units                     *string             `bson:"txtSaleComp1Units,omitempty"`
	TxtSaleComp2Units                     *string             `bson:"txtSaleComp2Units,omitempty"`
	TxtSaleComp3Units                     *string             `bson:"txtSaleComp3Units,omitempty"`
	TxtListComp1Units                     *string             `bson:"txtListComp1Units,omitempty"`
	TxtListComp2Units                     *string             `bson:"txtListComp2Units,omitempty"`
	TxtListComp3Units                     *string             `bson:"txtListComp3Units,omitempty"`
	TxtSubjectZip                         *string             `bson:"txtSubjectZip,omitempty"`
	TxtSaleComp1Zip                       *string             `bson:"txtSaleComp1Zip,omitempty"`
	TxtSaleComp2Zip                       *string             `bson:"txtSaleComp2Zip,omitempty"`
	TxtSaleComp3Zip                       *string             `bson:"txtSaleComp3Zip,omitempty"`
	TxtListComp1Zip                       *string             `bson:"txtListComp1Zip,omitempty"`
	TxtListComp2Zip                       *string             `bson:"txtListComp2Zip,omitempty"`
	TxtListComp3Zip                       *string             `bson:"txtListComp3Zip,omitempty"`
	TxtSubjectProximity                   *string             `bson:"txtSubjectProximity,omitempty"`
	TxtSaleComp1Proximity                 *string             `bson:"txtSaleComp1Proximity,omitempty"`
	TxtSaleComp2Proximity                 *string             `bson:"txtSaleComp2Proximity,omitempty"`
	TxtSaleComp3Proximity                 *string             `bson:"txtSaleComp3Proximity,omitempty"`
	TxtListComp1Proximity                 *string             `bson:"txtListComp1Proximity,omitempty"`
	TxtListComp2Proximity                 *string             `bson:"txtListComp2Proximity,omitempty"`
	TxtListComp3Proximity                 *string             `bson:"txtListComp3Proximity,omitempty"`
	TxtSubjectDataSource                  *string             `bson:"txtSubjectDataSource,omitempty"`
	TxtSaleComp1DataSource                *string             `bson:"txtSaleComp1DataSource,omitempty"`
	TxtSaleComp2DataSource                *string             `bson:"txtSaleComp2DataSource,omitempty"`
	TxtSaleComp3DataSource                *string             `bson:"txtSaleComp3DataSource,omitempty"`
	TxtListComp1DataSource                *string             `bson:"txtListComp1DataSource,omitempty"`
	TxtListComp2DataSource                *string             `bson:"txtListComp2DataSource,omitempty"`
	TxtListComp3DataSource                *string             `bson:"txtListComp3DataSource,omitempty"`
	TxtSubjectMLSNumber                   *string             `bson:"txtSubjectMLSNumber,omitempty"`
	TxtSaleComp1MLSNumber                 *string             `bson:"txtSaleComp1MLSNumber,omitempty"`
	TxtSaleComp2MLSNumber                 *string             `bson:"txtSaleComp2MLSNumber,omitempty"`
	TxtSaleComp3MLSNumber                 *string             `bson:"txtSaleComp3MLSNumber,omitempty"`
	TxtListComp1MLSNumber                 *string             `bson:"txtListComp1MLSNumber,omitempty"`
	TxtListComp2MLSNumber                 *string             `bson:"txtListComp2MLSNumber,omitempty"`
	TxtListComp3MLSNumber                 *string             `bson:"txtListComp3MLSNumber,omitempty"`
	CmbSubjectSaleType                    *string             `bson:"cmbSubjectSaleType,omitempty"`
	CmbSaleComp1SaleType                  *string             `bson:"cmbSaleComp1SaleType,omitempty"`
	CmbSaleComp2SaleType                  *string             `bson:"cmbSaleComp2SaleType,omitempty"`
	CmbSaleComp3SaleType                  *string             `bson:"cmbSaleComp3SaleType,omitempty"`
	CmbListComp1SaleType                  *string             `bson:"cmbListComp1SaleType,omitempty"`
	CmbListComp2SaleType                  *string             `bson:"cmbListComp2SaleType,omitempty"`
	CmbListComp3SaleType                  *string             `bson:"cmbListComp3SaleType,omitempty"`
	CmbSubjectType                        *string             `bson:"cmbSubjectType,omitempty"`
	CmbSaleComp1Type                      *string             `bson:"cmbSaleComp1Type,omitempty"`
	CmbSaleComp2Type                      *string             `bson:"cmbSaleComp2Type,omitempty"`
	CmbSaleComp3Type                      *string             `bson:"cmbSaleComp3Type,omitempty"`
	CmbListComp1Type                      *string             `bson:"cmbListComp1Type,omitempty"`
	CmbListComp2Type                      *string             `bson:"cmbListComp2Type,omitempty"`
	CmbListComp3Type                      *string             `bson:"cmbListComp3Type,omitempty"`
	CmbSubjectStyle                       *string             `bson:"cmbSubjectStyle,omitempty"`
	CmbSaleComp1Style                     *string             `bson:"cmbSaleComp1Style,omitempty"`
	TxtSaleComp1StyleAdjBuiltIn           *string             `bson:"txtSaleComp1StyleAdjBuiltIn,omitempty"`
	CmbSaleComp2Style                     *string             `bson:"cmbSaleComp2Style,omitempty"`
	TxtSaleComp2StyleAdjBuiltIn           *string             `bson:"txtSaleComp2StyleAdjBuiltIn,omitempty"`
	CmbSaleComp3Style                     *string             `bson:"cmbSaleComp3Style,omitempty"`
	TxtSaleComp3StyleAdjBuiltIn           *string             `bson:"txtSaleComp3StyleAdjBuiltIn,omitempty"`
	CmbListComp1Style                     *string             `bson:"cmbListComp1Style,omitempty"`
	TxtListComp1StyleAdjBuiltIn           *string             `bson:"txtListComp1StyleAdjBuiltIn,omitempty"`
	CmbListComp2Style                     *string             `bson:"cmbListComp2Style,omitempty"`
	TxtListComp2StyleAdjBuiltIn           *string             `bson:"txtListComp2StyleAdjBuiltIn,omitempty"`
	CmbListComp3Style                     *string             `bson:"cmbListComp3Style,omitempty"`
	TxtListComp3StyleAdjBuiltIn           *string             `bson:"txtListComp3StyleAdjBuiltIn,omitempty"`
	CmbSubjectExtFinish                   *string             `bson:"cmbSubjectExtFinish,omitempty"`
	CmbSaleComp1ExtFinish                 *string             `bson:"cmbSaleComp1ExtFinish,omitempty"`
	TxtSaleComp1ExtFinishAdjBuiltIn       *string             `bson:"txtSaleComp1ExtFinishAdjBuiltIn,omitempty"`
	CmbSaleComp2ExtFinish                 *string             `bson:"cmbSaleComp2ExtFinish,omitempty"`
	TxtSaleComp2ExtFinishAdjBuiltIn       *string             `bson:"txtSaleComp2ExtFinishAdjBuiltIn,omitempty"`
	CmbSaleComp3ExtFinish                 *string             `bson:"cmbSaleComp3ExtFinish,omitempty"`
	TxtSaleComp3ExtFinishAdjBuiltIn       *string             `bson:"txtSaleComp3ExtFinishAdjBuiltIn,omitempty"`
	CmbListComp1ExtFinish                 *string             `bson:"cmbListComp1ExtFinish,omitempty"`
	TxtListComp1ExtFinishAdjBuiltIn       *string             `bson:"txtListComp1ExtFinishAdjBuiltIn,omitempty"`
	CmbListComp2ExtFinish                 *string             `bson:"cmbListComp2ExtFinish,omitempty"`
	TxtListComp2ExtFinishAdjBuiltIn       *string             `bson:"txtListComp2ExtFinishAdjBuiltIn,omitempty"`
	CmbListComp3ExtFinish                 *string             `bson:"cmbListComp3ExtFinish,omitempty"`
	TxtListComp3ExtFinishAdjBuiltIn       *string             `bson:"txtListComp3ExtFinishAdjBuiltIn,omitempty"`
	CmbSubjectCondition                   *string             `bson:"cmbSubjectCondition,omitempty"`
	CmbSaleComp1Condition                 *string             `bson:"cmbSaleComp1Condition,omitempty"`
	TxtSaleComp1ConditionAdjBuiltIn       *string             `bson:"txtSaleComp1ConditionAdjBuiltIn,omitempty"`
	CmbSaleComp2Condition                 *string             `bson:"cmbSaleComp2Condition,omitempty"`
	TxtSaleComp2ConditionAdjBuiltIn       *string             `bson:"txtSaleComp2ConditionAdjBuiltIn,omitempty"`
	CmbSaleComp3Condition                 *string             `bson:"cmbSaleComp3Condition,omitempty"`
	TxtSaleComp3ConditionAdjBuiltIn       *string             `bson:"txtSaleComp3ConditionAdjBuiltIn,omitempty"`
	CmbListComp1Condition                 *string             `bson:"cmbListComp1Condition,omitempty"`
	TxtListComp1ConditionAdjBuiltIn       *string             `bson:"txtListComp1ConditionAdjBuiltIn,omitempty"`
	CmbListComp2Condition                 *string             `bson:"cmbListComp2Condition,omitempty"`
	TxtListComp2ConditionAdjBuiltIn       *string             `bson:"txtListComp2ConditionAdjBuiltIn,omitempty"`
	CmbListComp3Condition                 *string             `bson:"cmbListComp3Condition,omitempty"`
	TxtListComp3ConditionAdjBuiltIn       *string             `bson:"txtListComp3ConditionAdjBuiltIn,omitempty"`
	CmbSubjectQuality                     *string             `bson:"cmbSubjectQuality,omitempty"`
	CmbSaleComp1Quality                   *string             `bson:"cmbSaleComp1Quality,omitempty"`
	TxtSaleComp1QualityAdjBuiltIn         *string             `bson:"txtSaleComp1QualityAdjBuiltIn,omitempty"`
	CmbSaleComp2Quality                   *string             `bson:"cmbSaleComp2Quality,omitempty"`
	TxtSaleComp2QualityAdjBuiltIn         *string             `bson:"txtSaleComp2QualityAdjBuiltIn,omitempty"`
	CmbSaleComp3Quality                   *string             `bson:"cmbSaleComp3Quality,omitempty"`
	TxtSaleComp3QualityAdjBuiltIn         *string             `bson:"txtSaleComp3QualityAdjBuiltIn,omitempty"`
	CmbListComp1Quality                   *string             `bson:"cmbListComp1Quality,omitempty"`
	TxtListComp1QualityAdjBuiltIn         *string             `bson:"txtListComp1QualityAdjBuiltIn,omitempty"`
	CmbListComp2Quality                   *string             `bson:"cmbListComp2Quality,omitempty"`
	TxtListComp2QualityAdjBuiltIn         *string             `bson:"txtListComp2QualityAdjBuiltIn,omitempty"`
	CmbListComp3Quality                   *string             `bson:"cmbListComp3Quality,omitempty"`
	TxtListComp3QualityAdjBuiltIn         *string             `bson:"txtListComp3QualityAdjBuiltIn,omitempty"`
	CmbSubjectView                        *string             `bson:"cmbSubjectView,omitempty"`
	CmbSaleComp1View                      *string             `bson:"cmbSaleComp1View,omitempty"`
	TxtSaleComp1ViewAdjBuiltIn            *string             `bson:"txtSaleComp1ViewAdjBuiltIn,omitempty"`
	CmbSaleComp2View                      *string             `bson:"cmbSaleComp2View,omitempty"`
	TxtSaleComp2ViewAdjBuiltIn            *string             `bson:"txtSaleComp2ViewAdjBuiltIn,omitempty"`
	CmbSaleComp3View                      *string             `bson:"cmbSaleComp3View,omitempty"`
	TxtSaleComp3ViewAdjBuiltIn            *string             `bson:"txtSaleComp3ViewAdjBuiltIn,omitempty"`
	CmbListComp1View                      *string             `bson:"cmbListComp1View,omitempty"`
	TxtListComp1ViewAdjBuiltIn            *string             `bson:"txtListComp1ViewAdjBuiltIn,omitempty"`
	CmbListComp2View                      *string             `bson:"cmbListComp2View,omitempty"`
	TxtListComp2ViewAdjBuiltIn            *string             `bson:"txtListComp2ViewAdjBuiltIn,omitempty"`
	CmbListComp3View                      *string             `bson:"cmbListComp3View,omitempty"`
	TxtListComp3ViewAdjBuiltIn            *string             `bson:"txtListComp3ViewAdjBuiltIn,omitempty"`
	TxtSubjectSubdivision                 *string             `bson:"txtSubjectSubdivision,omitempty"`
	TxtSaleComp1Subdivision               *string             `bson:"txtSaleComp1Subdivision,omitempty"`
	TxtSaleComp2Subdivision               *string             `bson:"txtSaleComp2Subdivision,omitempty"`
	TxtSaleComp3Subdivision               *string             `bson:"txtSaleComp3Subdivision,omitempty"`
	TxtListComp1Subdivision               *string             `bson:"txtListComp1Subdivision,omitempty"`
	TxtListComp2Subdivision               *string             `bson:"txtListComp2Subdivision,omitempty"`
	TxtListComp3Subdivision               *string             `bson:"txtListComp3Subdivision,omitempty"`
	TxtSubjectHOAFee                      *string             `bson:"txtSubjectHOAFee,omitempty"`
	TxtSaleComp1HOAFee                    *string             `bson:"txtSaleComp1HOAFee,omitempty"`
	TxtSaleComp2HOAFee                    *string             `bson:"txtSaleComp2HOAFee,omitempty"`
	TxtSaleComp3HOAFee                    *string             `bson:"txtSaleComp3HOAFee,omitempty"`
	TxtListComp1HOAFee                    *string             `bson:"txtListComp1HOAFee,omitempty"`
	TxtListComp2HOAFee                    *string             `bson:"txtListComp2HOAFee,omitempty"`
	TxtListComp3HOAFee                    *string             `bson:"txtListComp3HOAFee,omitempty"`
	TxtSubjectTotalRooms                  *string             `bson:"txtSubjectTotalRooms,omitempty"`
	TxtSaleComp1TotalRooms                *string             `bson:"txtSaleComp1TotalRooms,omitempty"`
	TxtSaleComp1TotalRoomsAdjBuiltIn      *string             `bson:"txtSaleComp1TotalRoomsAdjBuiltIn,omitempty"`
	TxtSaleComp2TotalRooms                *string             `bson:"txtSaleComp2TotalRooms,omitempty"`
	TxtSaleComp2TotalRoomsAdjBuiltIn      *string             `bson:"txtSaleComp2TotalRoomsAdjBuiltIn,omitempty"`
	TxtSaleComp3TotalRooms                *string             `bson:"txtSaleComp3TotalRooms,omitempty"`
	TxtSaleComp3TotalRoomsAdjBuiltIn      *string             `bson:"txtSaleComp3TotalRoomsAdjBuiltIn,omitempty"`
	TxtListComp1TotalRooms                *string             `bson:"txtListComp1TotalRooms,omitempty"`
	TxtListComp1TotalRoomsAdjBuiltIn      *string             `bson:"txtListComp1TotalRoomsAdjBuiltIn,omitempty"`
	TxtListComp2TotalRooms                *string             `bson:"txtListComp2TotalRooms,omitempty"`
	TxtListComp2TotalRoomsAdjBuiltIn      *string             `bson:"txtListComp2TotalRoomsAdjBuiltIn,omitempty"`
	TxtListComp3TotalRooms                *string             `bson:"txtListComp3TotalRooms,omitempty"`
	TxtListComp3TotalRoomsAdjBuiltIn      *string             `bson:"txtListComp3TotalRoomsAdjBuiltIn,omitempty"`
	TxtSubjectBedrooms                    *string             `bson:"txtSubjectBedrooms,omitempty"`
	TxtSaleComp1Bedrooms                  *string             `bson:"txtSaleComp1Bedrooms,omitempty"`
	TxtSaleComp1BedroomsAdjBuiltIn        *string             `bson:"txtSaleComp1BedroomsAdjBuiltIn,omitempty"`
	TxtSaleComp2Bedrooms                  *string             `bson:"txtSaleComp2Bedrooms,omitempty"`
	TxtSaleComp2BedroomsAdjBuiltIn        *string             `bson:"txtSaleComp2BedroomsAdjBuiltIn,omitempty"`
	TxtSaleComp3Bedrooms                  *string             `bson:"txtSaleComp3Bedrooms,omitempty"`
	TxtSaleComp3BedroomsAdjBuiltIn        *string             `bson:"txtSaleComp3BedroomsAdjBuiltIn,omitempty"`
	TxtListComp1Bedrooms                  *string             `bson:"txtListComp1Bedrooms,omitempty"`
	TxtListComp1BedroomsAdjBuiltIn        *string             `bson:"txtListComp1BedroomsAdjBuiltIn,omitempty"`
	TxtListComp2Bedrooms                  *string             `bson:"txtListComp2Bedrooms,omitempty"`
	TxtListComp2BedroomsAdjBuiltIn        *string             `bson:"txtListComp2BedroomsAdjBuiltIn,omitempty"`
	TxtListComp3Bedrooms                  *string             `bson:"txtListComp3Bedrooms,omitempty"`
	TxtListComp3BedroomsAdjBuiltIn        *string             `bson:"txtListComp3BedroomsAdjBuiltIn,omitempty"`
	TxtSubjectFullBaths                   *string             `bson:"txtSubjectFullBaths,omitempty"`
	TxtSaleComp1FullBaths                 *string             `bson:"txtSaleComp1FullBaths,omitempty"`
	TxtSaleComp1FullBathsAdjBuiltIn       *string             `bson:"txtSaleComp1FullBathsAdjBuiltIn,omitempty"`
	TxtSaleComp2FullBaths                 *string             `bson:"txtSaleComp2FullBaths,omitempty"`
	TxtSaleComp2FullBathsAdjBuiltIn       *string             `bson:"txtSaleComp2FullBathsAdjBuiltIn,omitempty"`
	TxtSaleComp3FullBaths                 *string             `bson:"txtSaleComp3FullBaths,omitempty"`
	TxtSaleComp3FullBathsAdjBuiltIn       *string             `bson:"txtSaleComp3FullBathsAdjBuiltIn,omitempty"`
	TxtListComp1FullBaths                 *string             `bson:"txtListComp1FullBaths,omitempty"`
	TxtListComp1FullBathsAdjBuiltIn       *string             `bson:"txtListComp1FullBathsAdjBuiltIn,omitempty"`
	TxtListComp2FullBaths                 *string             `bson:"txtListComp2FullBaths,omitempty"`
	TxtListComp2FullBathsAdjBuiltIn       *string             `bson:"txtListComp2FullBathsAdjBuiltIn,omitempty"`
	TxtListComp3FullBaths                 *string             `bson:"txtListComp3FullBaths,omitempty"`
	TxtListComp3FullBathsAdjBuiltIn       *string             `bson:"txtListComp3FullBathsAdjBuiltIn,omitempty"`
	TxtSubjectHalfBaths                   *string             `bson:"txtSubjectHalfBaths,omitempty"`
	TxtSaleComp1HalfBaths                 *string             `bson:"txtSaleComp1HalfBaths,omitempty"`
	TxtSaleComp1HalfBathsAdjBuiltIn       *string             `bson:"txtSaleComp1HalfBathsAdjBuiltIn,omitempty"`
	TxtSaleComp2HalfBaths                 *string             `bson:"txtSaleComp2HalfBaths,omitempty"`
	TxtSaleComp2HalfBathsAdjBuiltIn       *string             `bson:"txtSaleComp2HalfBathsAdjBuiltIn,omitempty"`
	TxtSaleComp3HalfBaths                 *string             `bson:"txtSaleComp3HalfBaths,omitempty"`
	TxtSaleComp3HalfBathsAdjBuiltIn       *string             `bson:"txtSaleComp3HalfBathsAdjBuiltIn,omitempty"`
	TxtListComp1HalfBaths                 *string             `bson:"txtListComp1HalfBaths,omitempty"`
	TxtListComp1HalfBathsAdjBuiltIn       *string             `bson:"txtListComp1HalfBathsAdjBuiltIn,omitempty"`
	TxtListComp2HalfBaths                 *string             `bson:"txtListComp2HalfBaths,omitempty"`
	TxtListComp2HalfBathsAdjBuiltIn       *string             `bson:"txtListComp2HalfBathsAdjBuiltIn,omitempty"`
	TxtListComp3HalfBaths                 *string             `bson:"txtListComp3HalfBaths,omitempty"`
	TxtListComp3HalfBathsAdjBuiltIn       *string             `bson:"txtListComp3HalfBathsAdjBuiltIn,omitempty"`
	TxtSubjectGla                         *string             `bson:"txtSubjectGLA,omitempty"`
	TxtSaleComp1gla                       *string             `bson:"txtSaleComp1GLA,omitempty"`
	TxtSaleComp1GLAAdjBuiltIn             *string             `bson:"txtSaleComp1GLAAdjBuiltIn,omitempty"`
	TxtSaleComp2gla                       *string             `bson:"txtSaleComp2GLA,omitempty"`
	TxtSaleComp2GLAAdjBuiltIn             *string             `bson:"txtSaleComp2GLAAdjBuiltIn,omitempty"`
	TxtSaleComp3gla                       *string             `bson:"txtSaleComp3GLA,omitempty"`
	TxtSaleComp3GLAAdjBuiltIn             *string             `bson:"txtSaleComp3GLAAdjBuiltIn,omitempty"`
	TxtListComp1gla                       *string             `bson:"txtListComp1GLA,omitempty"`
	TxtListComp1GLAAdjBuiltIn             *string             `bson:"txtListComp1GLAAdjBuiltIn,omitempty"`
	TxtListComp2gla                       *string             `bson:"txtListComp2GLA,omitempty"`
	TxtListComp2GLAAdjBuiltIn             *string             `bson:"txtListComp2GLAAdjBuiltIn,omitempty"`
	TxtListComp3gla                       *string             `bson:"txtListComp3GLA,omitempty"`
	TxtListComp3GLAAdjBuiltIn             *string             `bson:"txtListComp3GLAAdjBuiltIn,omitempty"`
	TxtSubjectYearBuilt                   *string             `bson:"txtSubjectYearBuilt,omitempty"`
	TxtSaleComp1YearBuilt                 *string             `bson:"txtSaleComp1YearBuilt,omitempty"`
	TxtSaleComp1YearBuiltAdjBuiltIn       *string             `bson:"txtSaleComp1YearBuiltAdjBuiltIn,omitempty"`
	TxtSaleComp2YearBuilt                 *string             `bson:"txtSaleComp2YearBuilt,omitempty"`
	TxtSaleComp2YearBuiltAdjBuiltIn       *string             `bson:"txtSaleComp2YearBuiltAdjBuiltIn,omitempty"`
	TxtSaleComp3YearBuilt                 *string             `bson:"txtSaleComp3YearBuilt,omitempty"`
	TxtSaleComp3YearBuiltAdjBuiltIn       *string             `bson:"txtSaleComp3YearBuiltAdjBuiltIn,omitempty"`
	TxtListComp1YearBuilt                 *string             `bson:"txtListComp1YearBuilt,omitempty"`
	TxtListComp1YearBuiltAdjBuiltIn       *string             `bson:"txtListComp1YearBuiltAdjBuiltIn,omitempty"`
	TxtListComp2YearBuilt                 *string             `bson:"txtListComp2YearBuilt,omitempty"`
	TxtListComp2YearBuiltAdjBuiltIn       *string             `bson:"txtListComp2YearBuiltAdjBuiltIn,omitempty"`
	TxtListComp3YearBuilt                 *string             `bson:"txtListComp3YearBuilt,omitempty"`
	TxtListComp3YearBuiltAdjBuiltIn       *string             `bson:"txtListComp3YearBuiltAdjBuiltIn,omitempty"`
	TxtSubjectAge                         *string             `bson:"txtSubjectAge,omitempty"`
	TxtSaleComp1Age                       *string             `bson:"txtSaleComp1Age,omitempty"`
	TxtSaleComp2Age                       *string             `bson:"txtSaleComp2Age,omitempty"`
	TxtSaleComp3Age                       *string             `bson:"txtSaleComp3Age,omitempty"`
	TxtListComp1Age                       *string             `bson:"txtListComp1Age,omitempty"`
	TxtListComp2Age                       *string             `bson:"txtListComp2Age,omitempty"`
	TxtListComp3Age                       *string             `bson:"txtListComp3Age,omitempty"`
	TxtSubjectAcres                       *string             `bson:"txtSubjectAcres,omitempty"`
	TxtSaleComp1Acres                     *string             `bson:"txtSaleComp1Acres,omitempty"`
	TxtSaleComp1AcresAdjBuiltIn           *string             `bson:"txtSaleComp1AcresAdjBuiltIn,omitempty"`
	TxtSaleComp2Acres                     *string             `bson:"txtSaleComp2Acres,omitempty"`
	TxtSaleComp2AcresAdjBuiltIn           *string             `bson:"txtSaleComp2AcresAdjBuiltIn,omitempty"`
	TxtSaleComp3Acres                     *string             `bson:"txtSaleComp3Acres,omitempty"`
	TxtSaleComp3AcresAdjBuiltIn           *string             `bson:"txtSaleComp3AcresAdjBuiltIn,omitempty"`
	TxtListComp1Acres                     *string             `bson:"txtListComp1Acres,omitempty"`
	TxtListComp1AcresAdjBuiltIn           *string             `bson:"txtListComp1AcresAdjBuiltIn,omitempty"`
	TxtListComp2Acres                     *string             `bson:"txtListComp2Acres,omitempty"`
	TxtListComp2AcresAdjBuiltIn           *string             `bson:"txtListComp2AcresAdjBuiltIn,omitempty"`
	TxtListComp3Acres                     *string             `bson:"txtListComp3Acres,omitempty"`
	TxtListComp3AcresAdjBuiltIn           *string             `bson:"txtListComp3AcresAdjBuiltIn,omitempty"`
	TxtSubjectSquareFeet                  *string             `bson:"txtSubjectSquareFeet,omitempty"`
	TxtSaleComp1SquareFeet                *string             `bson:"txtSaleComp1SquareFeet,omitempty"`
	TxtSaleComp2SquareFeet                *string             `bson:"txtSaleComp2SquareFeet,omitempty"`
	TxtSaleComp3SquareFeet                *string             `bson:"txtSaleComp3SquareFeet,omitempty"`
	TxtListComp1SquareFeet                *string             `bson:"txtListComp1SquareFeet,omitempty"`
	TxtListComp2SquareFeet                *string             `bson:"txtListComp2SquareFeet,omitempty"`
	TxtListComp3SquareFeet                *string             `bson:"txtListComp3SquareFeet,omitempty"`
	CmbSubjectGarage                      *string             `bson:"cmbSubjectGarage,omitempty"`
	CmbSaleComp1Garage                    *string             `bson:"cmbSaleComp1Garage,omitempty"`
	TxtSaleComp1GarageAdjBuiltIn          *string             `bson:"txtSaleComp1GarageAdjBuiltIn,omitempty"`
	CmbSaleComp2Garage                    *string             `bson:"cmbSaleComp2Garage,omitempty"`
	TxtSaleComp2GarageAdjBuiltIn          *string             `bson:"txtSaleComp2GarageAdjBuiltIn,omitempty"`
	CmbSaleComp3Garage                    *string             `bson:"cmbSaleComp3Garage,omitempty"`
	TxtSaleComp3GarageAdjBuiltIn          *string             `bson:"txtSaleComp3GarageAdjBuiltIn,omitempty"`
	CmbListComp1Garage                    *string             `bson:"cmbListComp1Garage,omitempty"`
	TxtListComp1GarageAdjBuiltIn          *string             `bson:"txtListComp1GarageAdjBuiltIn,omitempty"`
	CmbListComp2Garage                    *string             `bson:"cmbListComp2Garage,omitempty"`
	TxtListComp2GarageAdjBuiltIn          *string             `bson:"txtListComp2GarageAdjBuiltIn,omitempty"`
	CmbListComp3Garage                    *string             `bson:"cmbListComp3Garage,omitempty"`
	TxtListComp3GarageAdjBuiltIn          *string             `bson:"txtListComp3GarageAdjBuiltIn,omitempty"`
	CmbSubjectPool                        *string             `bson:"cmbSubjectPool,omitempty"`
	CmbSaleComp1Pool                      *string             `bson:"cmbSaleComp1Pool,omitempty"`
	TxtSaleComp1PoolAdjBuiltIn            *string             `bson:"txtSaleComp1PoolAdjBuiltIn,omitempty"`
	CmbSaleComp2Pool                      *string             `bson:"cmbSaleComp2Pool,omitempty"`
	TxtSaleComp2PoolAdjBuiltIn            *string             `bson:"txtSaleComp2PoolAdjBuiltIn,omitempty"`
	CmbSaleComp3Pool                      *string             `bson:"cmbSaleComp3Pool,omitempty"`
	TxtSaleComp3PoolAdjBuiltIn            *string             `bson:"txtSaleComp3PoolAdjBuiltIn,omitempty"`
	CmbListComp1Pool                      *string             `bson:"cmbListComp1Pool,omitempty"`
	TxtListComp1PoolAdjBuiltIn            *string             `bson:"txtListComp1PoolAdjBuiltIn,omitempty"`
	CmbListComp2Pool                      *string             `bson:"cmbListComp2Pool,omitempty"`
	TxtListComp2PoolAdjBuiltIn            *string             `bson:"txtListComp2PoolAdjBuiltIn,omitempty"`
	CmbListComp3Pool                      *string             `bson:"cmbListComp3Pool,omitempty"`
	TxtListComp3PoolAdjBuiltIn            *string             `bson:"txtListComp3PoolAdjBuiltIn,omitempty"`
	CmbSubjectPorchPatioDeck              *string             `bson:"cmbSubjectPorchPatioDeck,omitempty"`
	CmbSaleComp1PorchPatioDeck            *string             `bson:"cmbSaleComp1PorchPatioDeck,omitempty"`
	TxtSaleComp1PorchPatioDeckAdjBuiltIn  *string             `bson:"txtSaleComp1PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbSaleComp2PorchPatioDeck            *string             `bson:"cmbSaleComp2PorchPatioDeck,omitempty"`
	TxtSaleComp2PorchPatioDeckAdjBuiltIn  *string             `bson:"txtSaleComp2PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbSaleComp3PorchPatioDeck            *string             `bson:"cmbSaleComp3PorchPatioDeck,omitempty"`
	TxtSaleComp3PorchPatioDeckAdjBuiltIn  *string             `bson:"txtSaleComp3PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbListComp1PorchPatioDeck            *string             `bson:"cmbListComp1PorchPatioDeck,omitempty"`
	TxtListComp1PorchPatioDeckAdjBuiltIn  *string             `bson:"txtListComp1PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbListComp2PorchPatioDeck            *string             `bson:"cmbListComp2PorchPatioDeck,omitempty"`
	TxtListComp2PorchPatioDeckAdjBuiltIn  *string             `bson:"txtListComp2PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbListComp3PorchPatioDeck            *string             `bson:"cmbListComp3PorchPatioDeck,omitempty"`
	TxtListComp3PorchPatioDeckAdjBuiltIn  *string             `bson:"txtListComp3PorchPatioDeckAdjBuiltIn,omitempty"`
	CmbSubjectFireplace                   *string             `bson:"cmbSubjectFireplace,omitempty"`
	CmbSaleComp1Fireplace                 *string             `bson:"cmbSaleComp1Fireplace,omitempty"`
	TxtSaleComp1FireplaceAdjBuiltIn       *string             `bson:"txtSaleComp1FireplaceAdjBuiltIn,omitempty"`
	CmbSaleComp2Fireplace                 *string             `bson:"cmbSaleComp2Fireplace,omitempty"`
	TxtSaleComp2FireplaceAdjBuiltIn       *string             `bson:"txtSaleComp2FireplaceAdjBuiltIn,omitempty"`
	CmbSaleComp3Fireplace                 *string             `bson:"cmbSaleComp3Fireplace,omitempty"`
	TxtSaleComp3FireplaceAdjBuiltIn       *string             `bson:"txtSaleComp3FireplaceAdjBuiltIn,omitempty"`
	CmbListComp1Fireplace                 *string             `bson:"cmbListComp1Fireplace,omitempty"`
	TxtListComp1FireplaceAdjBuiltIn       *string             `bson:"txtListComp1FireplaceAdjBuiltIn,omitempty"`
	CmbListComp2Fireplace                 *string             `bson:"cmbListComp2Fireplace,omitempty"`
	TxtListComp2FireplaceAdjBuiltIn       *string             `bson:"txtListComp2FireplaceAdjBuiltIn,omitempty"`
	CmbListComp3Fireplace                 *string             `bson:"cmbListComp3Fireplace,omitempty"`
	TxtListComp3FireplaceAdjBuiltIn       *string             `bson:"txtListComp3FireplaceAdjBuiltIn,omitempty"`
	CmbSubjectBasement                    *string             `bson:"cmbSubjectBasement,omitempty"`
	CmbSaleComp1Basement                  *string             `bson:"cmbSaleComp1Basement,omitempty"`
	TxtSaleComp1BasementAdjBuiltIn        *string             `bson:"txtSaleComp1BasementAdjBuiltIn,omitempty"`
	CmbSaleComp2Basement                  *string             `bson:"cmbSaleComp2Basement,omitempty"`
	TxtSaleComp2BasementAdjBuiltIn        *string             `bson:"txtSaleComp2BasementAdjBuiltIn,omitempty"`
	CmbSaleComp3Basement                  *string             `bson:"cmbSaleComp3Basement,omitempty"`
	TxtSaleComp3BasementAdjBuiltIn        *string             `bson:"txtSaleComp3BasementAdjBuiltIn,omitempty"`
	CmbListComp1Basement                  *string             `bson:"cmbListComp1Basement,omitempty"`
	TxtListComp1BasementAdjBuiltIn        *string             `bson:"txtListComp1BasementAdjBuiltIn,omitempty"`
	CmbListComp2Basement                  *string             `bson:"cmbListComp2Basement,omitempty"`
	TxtListComp2BasementAdjBuiltIn        *string             `bson:"txtListComp2BasementAdjBuiltIn,omitempty"`
	CmbListComp3Basement                  *string             `bson:"cmbListComp3Basement,omitempty"`
	TxtListComp3BasementAdjBuiltIn        *string             `bson:"txtListComp3BasementAdjBuiltIn,omitempty"`
	CmbSubjectIsFinished                  *string             `bson:"cmbSubjectIsFinished,omitempty"`
	CmbSaleComp1IsFinished                *string             `bson:"cmbSaleComp1IsFinished,omitempty"`
	TxtSaleComp1IsFinishedAdjBuiltIn      *string             `bson:"txtSaleComp1IsFinishedAdjBuiltIn,omitempty"`
	CmbSaleComp2IsFinished                *string             `bson:"cmbSaleComp2IsFinished,omitempty"`
	TxtSaleComp2IsFinishedAdjBuiltIn      *string             `bson:"txtSaleComp2IsFinishedAdjBuiltIn,omitempty"`
	CmbSaleComp3IsFinished                *string             `bson:"cmbSaleComp3IsFinished,omitempty"`
	TxtSaleComp3IsFinishedAdjBuiltIn      *string             `bson:"txtSaleComp3IsFinishedAdjBuiltIn,omitempty"`
	CmbListComp1IsFinished                *string             `bson:"cmbListComp1IsFinished,omitempty"`
	TxtListComp1IsFinishedAdjBuiltIn      *string             `bson:"txtListComp1IsFinishedAdjBuiltIn,omitempty"`
	CmbListComp2IsFinished                *string             `bson:"cmbListComp2IsFinished,omitempty"`
	TxtListComp2IsFinishedAdjBuiltIn      *string             `bson:"txtListComp2IsFinishedAdjBuiltIn,omitempty"`
	CmbListComp3IsFinished                *string             `bson:"cmbListComp3IsFinished,omitempty"`
	TxtListComp3IsFinishedAdjBuiltIn      *string             `bson:"txtListComp3IsFinishedAdjBuiltIn,omitempty"`
	CmbSubjectPercentFinished             *string             `bson:"cmbSubjectPercentFinished,omitempty"`
	CmbSaleComp1PercentFinished           *string             `bson:"cmbSaleComp1PercentFinished,omitempty"`
	TxtSaleComp1PercentFinishedAdjBuiltIn *string             `bson:"txtSaleComp1PercentFinishedAdjBuiltIn,omitempty"`
	CmbSaleComp2PercentFinished           *string             `bson:"cmbSaleComp2PercentFinished,omitempty"`
	TxtSaleComp2PercentFinishedAdjBuiltIn *string             `bson:"txtSaleComp2PercentFinishedAdjBuiltIn,omitempty"`
	CmbSaleComp3PercentFinished           *string             `bson:"cmbSaleComp3PercentFinished,omitempty"`
	TxtSaleComp3PercentFinishedAdjBuiltIn *string             `bson:"txtSaleComp3PercentFinishedAdjBuiltIn,omitempty"`
	CmbListComp1PercentFinished           *string             `bson:"cmbListComp1PercentFinished,omitempty"`
	TxtListComp1PercentFinishedAdjBuiltIn *string             `bson:"txtListComp1PercentFinishedAdjBuiltIn,omitempty"`
	CmbListComp2PercentFinished           *string             `bson:"cmbListComp2PercentFinished,omitempty"`
	TxtListComp2PercentFinishedAdjBuiltIn *string             `bson:"txtListComp2PercentFinishedAdjBuiltIn,omitempty"`
	CmbListComp3PercentFinished           *string             `bson:"cmbListComp3PercentFinished,omitempty"`
	TxtListComp3PercentFinishedAdjBuiltIn *string             `bson:"txtListComp3PercentFinishedAdjBuiltIn,omitempty"`
	TxtSubjectBasementSqFt                *string             `bson:"txtSubjectBasementSqFt,omitempty"`
	TxtSaleComp1BasementSqFt              *string             `bson:"txtSaleComp1BasementSqFt,omitempty"`
	TxtSaleComp1BasementSqFtAdjBuiltIn    *string             `bson:"txtSaleComp1BasementSqFtAdjBuiltIn,omitempty"`
	TxtSaleComp2BasementSqFt              *string             `bson:"txtSaleComp2BasementSqFt,omitempty"`
	TxtSaleComp2BasementSqFtAdjBuiltIn    *string             `bson:"txtSaleComp2BasementSqFtAdjBuiltIn,omitempty"`
	TxtSaleComp3BasementSqFt              *string             `bson:"txtSaleComp3BasementSqFt,omitempty"`
	TxtSaleComp3BasementSqFtAdjBuiltIn    *string             `bson:"txtSaleComp3BasementSqFtAdjBuiltIn,omitempty"`
	TxtListComp1BasementSqFt              *string             `bson:"txtListComp1BasementSqFt,omitempty"`
	TxtListComp1BasementSqFtAdjBuiltIn    *string             `bson:"txtListComp1BasementSqFtAdjBuiltIn,omitempty"`
	TxtListComp2BasementSqFt              *string             `bson:"txtListComp2BasementSqFt,omitempty"`
	TxtListComp2BasementSqFtAdjBuiltIn    *string             `bson:"txtListComp2BasementSqFtAdjBuiltIn,omitempty"`
	TxtListComp3BasementSqFt              *string             `bson:"txtListComp3BasementSqFt,omitempty"`
	TxtListComp3BasementSqFtAdjBuiltIn    *string             `bson:"txtListComp3BasementSqFtAdjBuiltIn,omitempty"`
	TxtSubjectOriginalListDate            *string             `bson:"txtSubjectOriginalListDate,omitempty"`
	TxtSaleComp1OriginalListDate          *string             `bson:"txtSaleComp1OriginalListDate,omitempty"`
	TxtSaleComp2OriginalListDate          *string             `bson:"txtSaleComp2OriginalListDate,omitempty"`
	TxtSaleComp3OriginalListDate          *string             `bson:"txtSaleComp3OriginalListDate,omitempty"`
	TxtListComp1OriginalListDate          *string             `bson:"txtListComp1OriginalListDate,omitempty"`
	TxtListComp2OriginalListDate          *string             `bson:"txtListComp2OriginalListDate,omitempty"`
	TxtListComp3OriginalListDate          *string             `bson:"txtListComp3OriginalListDate,omitempty"`
	TxtSubjectCurrentListDate             *string             `bson:"txtSubjectCurrentListDate,omitempty"`
	TxtSaleComp1CurrentListDate           *string             `bson:"txtSaleComp1CurrentListDate,omitempty"`
	TxtSaleComp2CurrentListDate           *string             `bson:"txtSaleComp2CurrentListDate,omitempty"`
	TxtSaleComp3CurrentListDate           *string             `bson:"txtSaleComp3CurrentListDate,omitempty"`
	TxtListComp1CurrentListDate           *string             `bson:"txtListComp1CurrentListDate,omitempty"`
	TxtListComp2CurrentListDate           *string             `bson:"txtListComp2CurrentListDate,omitempty"`
	TxtListComp3CurrentListDate           *string             `bson:"txtListComp3CurrentListDate,omitempty"`
	TxtSubjectOriginalListPrice           *string             `bson:"txtSubjectOriginalListPrice,omitempty"`
	TxtSaleComp1OriginalListPrice         *string             `bson:"txtSaleComp1OriginalListPrice,omitempty"`
	TxtSaleComp2OriginalListPrice         *string             `bson:"txtSaleComp2OriginalListPrice,omitempty"`
	TxtSaleComp3OriginalListPrice         *string             `bson:"txtSaleComp3OriginalListPrice,omitempty"`
	TxtListComp1OriginalListPrice         *string             `bson:"txtListComp1OriginalListPrice,omitempty"`
	TxtListComp2OriginalListPrice         *string             `bson:"txtListComp2OriginalListPrice,omitempty"`
	TxtListComp3OriginalListPrice         *string             `bson:"txtListComp3OriginalListPrice,omitempty"`
	TxtSubjectListPrice                   *string             `bson:"txtSubjectListPrice,omitempty"`
	TxtSaleComp1ListPrice                 *string             `bson:"txtSaleComp1ListPrice,omitempty"`
	TxtSaleComp2ListPrice                 *string             `bson:"txtSaleComp2ListPrice,omitempty"`
	TxtSaleComp3ListPrice                 *string             `bson:"txtSaleComp3ListPrice,omitempty"`
	TxtListComp1ListPrice                 *string             `bson:"txtListComp1ListPrice,omitempty"`
	TxtListComp2ListPrice                 *string             `bson:"txtListComp2ListPrice,omitempty"`
	TxtListComp3ListPrice                 *string             `bson:"txtListComp3ListPrice,omitempty"`
	TxtSubjectSalePrice                   *string             `bson:"txtSubjectSalePrice,omitempty"`
	TxtSaleComp1SalePrice                 *string             `bson:"txtSaleComp1SalePrice,omitempty"`
	TxtSaleComp2SalePrice                 *string             `bson:"txtSaleComp2SalePrice,omitempty"`
	TxtSaleComp3SalePrice                 *string             `bson:"txtSaleComp3SalePrice,omitempty"`
	TxtSubjectSaleDate                    *string             `bson:"txtSubjectSaleDate,omitempty"`
	TxtSaleComp1SaleDate                  *string             `bson:"txtSaleComp1SaleDate,omitempty"`
	TxtSaleComp2SaleDate                  *string             `bson:"txtSaleComp2SaleDate,omitempty"`
	TxtSaleComp3SaleDate                  *string             `bson:"txtSaleComp3SaleDate,omitempty"`
	CmbSubjectFinancing                   *string             `bson:"cmbSubjectFinancing,omitempty"`
	CmbSaleComp1Financing                 *string             `bson:"cmbSaleComp1Financing,omitempty"`
	CmbSaleComp2Financing                 *string             `bson:"cmbSaleComp2Financing,omitempty"`
	CmbSaleComp3Financing                 *string             `bson:"cmbSaleComp3Financing,omitempty"`
	CmbListComp1Financing                 *string             `bson:"cmbListComp1Financing,omitempty"`
	CmbListComp2Financing                 *string             `bson:"cmbListComp2Financing,omitempty"`
	CmbListComp3Financing                 *string             `bson:"cmbListComp3Financing,omitempty"`
	TxtSubjectDom                         *string             `bson:"txtSubjectDOM,omitempty"`
	TxtSaleComp1dom                       *string             `bson:"txtSaleComp1DOM,omitempty"`
	TxtSaleComp2dom                       *string             `bson:"txtSaleComp2DOM,omitempty"`
	TxtSaleComp3dom                       *string             `bson:"txtSaleComp3DOM,omitempty"`
	TxtListComp1dom                       *string             `bson:"txtListComp1DOM,omitempty"`
	TxtListComp2dom                       *string             `bson:"txtListComp2DOM,omitempty"`
	TxtListComp3dom                       *string             `bson:"txtListComp3DOM,omitempty"`
	TxtSubjectPricePerSqFt                *string             `bson:"txtSubjectPricePerSqFt,omitempty"`
	TxtSaleComp1PricePerSqFt              *string             `bson:"txtSaleComp1PricePerSqFt,omitempty"`
	TxtSaleComp2PricePerSqFt              *string             `bson:"txtSaleComp2PricePerSqFt,omitempty"`
	TxtSaleComp3PricePerSqFt              *string             `bson:"txtSaleComp3PricePerSqFt,omitempty"`
	TxtListComp1PricePerSqFt              *string             `bson:"txtListComp1PricePerSqFt,omitempty"`
	TxtListComp2PricePerSqFt              *string             `bson:"txtListComp2PricePerSqFt,omitempty"`
	TxtListComp3PricePerSqFt              *string             `bson:"txtListComp3PricePerSqFt,omitempty"`
	TxtSubjectAdjustments                 *string             `bson:"txtSubjectAdjustments,omitempty"`
	TxtSaleComp1Adjustments               *string             `bson:"txtSaleComp1Adjustments,omitempty"`
	TxtSaleComp2Adjustments               *string             `bson:"txtSaleComp2Adjustments,omitempty"`
	TxtSaleComp3Adjustments               *string             `bson:"txtSaleComp3Adjustments,omitempty"`
	TxtListComp1Adjustments               *string             `bson:"txtListComp1Adjustments,omitempty"`
	TxtListComp2Adjustments               *string             `bson:"txtListComp2Adjustments,omitempty"`
	TxtListComp3Adjustments               *string             `bson:"txtListComp3Adjustments,omitempty"`
	TxtSubjectCompTotals                  *string             `bson:"txtSubjectCompTotals,omitempty"`
	TxtSaleComp1CompTotals                *string             `bson:"txtSaleComp1CompTotals,omitempty"`
	TxtSaleComp2CompTotals                *string             `bson:"txtSaleComp2CompTotals,omitempty"`
	TxtSaleComp3CompTotals                *string             `bson:"txtSaleComp3CompTotals,omitempty"`
	TxtListComp1CompTotals                *string             `bson:"txtListComp1CompTotals,omitempty"`
	TxtListComp2CompTotals                *string             `bson:"txtListComp2CompTotals,omitempty"`
	TxtListComp3CompTotals                *string             `bson:"txtListComp3CompTotals,omitempty"`
	CmbListComp1CommentType               *string             `bson:"cmbListComp1CommentType,omitempty"`
	TxtListComp1ComparableComments        *string             `bson:"txtListComp1ComparableComments,omitempty"`
	TxtListComp1FormatAdjustments         *string             `bson:"txtListComp1FormatAdjustments,omitempty"`
	TxtListComp1MLSComments               *string             `bson:"txtListComp1MLSComments,omitempty"`
	CmbListComp2CommentType               *string             `bson:"cmbListComp2CommentType,omitempty"`
	TxtListComp2ComparableComments        *string             `bson:"txtListComp2ComparableComments,omitempty"`
	TxtListComp2FormatAdjustments         *string             `bson:"txtListComp2FormatAdjustments,omitempty"`
	TxtListComp2MLSComments               *string             `bson:"txtListComp2MLSComments,omitempty"`
	CmbListComp3CommentType               *string             `bson:"cmbListComp3CommentType,omitempty"`
	TxtListComp3ComparableComments        *string             `bson:"txtListComp3ComparableComments,omitempty"`
	TxtListComp3FormatAdjustments         *string             `bson:"txtListComp3FormatAdjustments,omitempty"`
	TxtListComp3MLSComments               *string             `bson:"txtListComp3MLSComments,omitempty"`
	CmbSaleComp1CommentType               *string             `bson:"cmbSaleComp1CommentType,omitempty"`
	TxtSaleComp1ComparableComments        *string             `bson:"txtSaleComp1ComparableComments,omitempty"`
	TxtSaleComp1FormatAdjustments         *string             `bson:"txtSaleComp1FormatAdjustments,omitempty"`
	TxtSaleComp1MLSComments               *string             `bson:"txtSaleComp1MLSComments,omitempty"`
	CmbSaleComp2CommentType               *string             `bson:"cmbSaleComp2CommentType,omitempty"`
	TxtSaleComp2ComparableComments        *string             `bson:"txtSaleComp2ComparableComments,omitempty"`
	TxtSaleComp2FormatAdjustments         *string             `bson:"txtSaleComp2FormatAdjustments,omitempty"`
	TxtSaleComp2MLSComments               *string             `bson:"txtSaleComp2MLSComments,omitempty"`
	CmbSaleComp3CommentType               *string             `bson:"cmbSaleComp3CommentType,omitempty"`
	TxtSaleComp3ComparableComments        *string             `bson:"txtSaleComp3ComparableComments,omitempty"`
	TxtSaleComp3FormatAdjustments         *string             `bson:"txtSaleComp3FormatAdjustments,omitempty"`
	TxtSaleComp3MLSComments               *string             `bson:"txtSaleComp3MLSComments,omitempty"`
	CmbNeighborhoodTrend                  *string             `bson:"cmbNeighborhoodTrend,omitempty"`
	TxtMonthlyPecent                      *string             `bson:"txtMonthlyPecent,omitempty"`
	TxtEstimatedRent                      *string             `bson:"txtEstimatedRent,omitempty"`
	TxtEstimatedDaysOnMarket              *string             `bson:"txtEstimatedDaysOnMarket,omitempty"`
	TxtNoBoarded                          *string             `bson:"txtNoBoarded,omitempty"`
	TxtNoOfActive                         *string             `bson:"txtNoOfActive,omitempty"`
	Txt6MonthPecent                       *string             `bson:"txt6MonthPecent,omitempty"`
	TxtAnnualPecent                       *string             `bson:"txtAnnualPecent,omitempty"`
	TxtListings                           *string             `bson:"txtListings,omitempty"`
	CmbSupply                             *string             `bson:"cmbSupply,omitempty"`
	TxtListingsMinValue                   *string             `bson:"txtListingsMinValue,omitempty"`
	TxtListingsRange1                     *string             `bson:"txtListingsRange1,omitempty"`
	TxtListingsMedValue                   *string             `bson:"txtListingsMedValue,omitempty"`
	TxtListingsMaxValue                   *string             `bson:"txtListingsMaxValue,omitempty"`
	TxtListingsRange2                     *string             `bson:"txtListingsRange2,omitempty"`
	TxtListingsDom                        *string             `bson:"txtListingsDOM,omitempty"`
	TxtListingsDOMRange1                  *string             `bson:"txtListingsDOMRange1,omitempty"`
	TxtListingsDOMRange2                  *string             `bson:"txtListingsDOMRange2,omitempty"`
	CmbREOTrend                           *string             `bson:"cmbREOTrend,omitempty"`
	TxtNoOfFm                             *string             `bson:"txtNoOfFM,omitempty"`
	TxtNoOfSs                             *string             `bson:"txtNoOfSS,omitempty"`
	TxtNoOfReo                            *string             `bson:"txtNoOfREO,omitempty"`
	TxtNoOfDistressed                     *string             `bson:"txtNoOfDistressed,omitempty"`
	TxtSales                              *string             `bson:"txtSales,omitempty"`
	CmbDemand                             *string             `bson:"cmbDemand,omitempty"`
	TxtSalesRange1                        *string             `bson:"txtSalesRange1,omitempty"`
	TxtSalesMedValue                      *string             `bson:"txtSalesMedValue,omitempty"`
	TxtSalesRange2                        *string             `bson:"txtSalesRange2,omitempty"`
	TxtSalesDom                           *string             `bson:"txtSalesDOM,omitempty"`
	TxtSalesDOMRange1                     *string             `bson:"txtSalesDOMRange1,omitempty"`
	TxtSalesDOMRange2                     *string             `bson:"txtSalesDOMRange2,omitempty"`
	TxtZillowNeighborhoodTrend            *string             `bson:"txtZillowNeighborhoodTrend,omitempty"`
	TxtNeighborhoodTrendComments          *string             `bson:"txtNeighborhoodTrendComments,omitempty"`
	TxtTotalListings                      *string             `bson:"txtTotalListings,omitempty"`
	TxtTotalSales                         *string             `bson:"txtTotalSales,omitempty"`
	TxtNoOfREOListings                    *string             `bson:"txtNoOfREOListings,omitempty"`
	TxtNoOfSSListings                     *string             `bson:"txtNoOfSSListings,omitempty"`
	TxtNoOfREOSales                       *string             `bson:"txtNoOfREOSales,omitempty"`
	TxtNoOfSSSales                        *string             `bson:"txtNoOfSSSales,omitempty"`
	TxtTaxID                              *string             `bson:"txtTaxID,omitempty"`
	TxtLastSaleDate                       *string             `bson:"txtLastSaleDate,omitempty"`
	TxtLastSalePrice                      *string             `bson:"txtLastSalePrice,omitempty"`
	CmbIsListed                           *string             `bson:"cmbIsListed,omitempty"`
	TxtOwnerOccupied                      *string             `bson:"txtOwnerOccupied,omitempty"`
	TxtRenterOccupied                     *string             `bson:"txtRenterOccupied,omitempty"`
	TxtMarketRent                         *string             `bson:"txtMarketRent,omitempty"`
	TxtNoOfRentals                        *string             `bson:"txtNoOfRentals,omitempty"`
	TxtTypicalDom                         *string             `bson:"txtTypicalDOM,omitempty"`
	TxtNoRentHomes                        *string             `bson:"txtNoRentHomes,omitempty"`
	TxtTypicalRentalRates                 *string             `bson:"txtTypicalRentalRates,omitempty"`
	AdjustmentPrice                       *string             `bson:"adjustmentPrice,omitempty"`
	TxtCalculatedGla                      *string             `bson:"txtCalculatedGLA,omitempty"`
	TxtCalculatedAge                      *string             `bson:"txtCalculatedAge,omitempty"`
	TxtCalculatedSaleDates                *string             `bson:"txtCalculatedSaleDates,omitempty"`
	TxtCalculatedProximity                *string             `bson:"txtCalculatedProximity,omitempty"`
	TxtCalculatedStyle                    *string             `bson:"txtCalculatedStyle,omitempty"`
	TxtCalculatedMonthsSupply             *string             `bson:"txtCalculatedMonthsSupply,omitempty"`
	TxtCalculatedProxim                   *string             `bson:"txtCalculatedProxim,omitempty"`
	TxtCalculatedGLAs                     *string             `bson:"txtCalculatedGLAs,omitempty"`
	TxtCalculatedAges                     *string             `bson:"txtCalculatedAges,omitempty"`
	TxtCalculatedCond                     *string             `bson:"txtCalculatedCond,omitempty"`
	TxtCalculatedView                     *string             `bson:"txtCalculatedView,omitempty"`
	TxtCalculatedStyle1                   *string             `bson:"txtCalculatedStyle1,omitempty"`
	TxtCalculatedLots                     *string             `bson:"txtCalculatedLots,omitempty"`
	TxtCalculatedBeds                     *string             `bson:"txtCalculatedBeds,omitempty"`
	TxtCalculatedBath                     *string             `bson:"txtCalculatedBath,omitempty"`
	Rdbresaletext                         *string             `bson:"rdbresaletext,omitempty"`
	Rdbmarketedtext                       *string             `bson:"rdbmarketedtext,omitempty"`
	Txtpmi                                *string             `bson:"txtpmi,omitempty"`
	TxtOtherComments                      *string             `bson:"txtOtherComments,omitempty"`
	Txtcbnew                              *string             `bson:"txtcbnew,omitempty"`
	Txtcbold                              *string             `bson:"txtcbold,omitempty"`
	Txtcbstyle                            *string             `bson:"txtcbstyle,omitempty"`
	Txtcblot                              *string             `bson:"txtcblot,omitempty"`
	Txtcbview                             *string             `bson:"txtcbview,omitempty"`
	Txtcbdamage                           *string             `bson:"txtcbdamage,omitempty"`
	Txtcbupgrade                          *string             `bson:"txtcbupgrade,omitempty"`
	Txtcbinfluence                        *string             `bson:"txtcbinfluence,omitempty"`
	TxtSubjectComments                    *string             `bson:"txtSubjectComments,omitempty"`
	TxtNeighborhoodComments               *string             `bson:"txtNeighborhoodComments,omitempty"`
	TxtNeighborhoodTrend                  *string             `bson:"txtNeighborhoodTrend,omitempty"`
	TxtValidation1                        *string             `bson:"txtValidation1,omitempty"`
	TxtUniqueComments                     *string             `bson:"txtUniqueComments,omitempty"`
	TxtMarketingStrategy                  *string             `bson:"txtMarketingStrategy,omitempty"`
	TxtDisclaimer2                        *string             `bson:"txtDisclaimer2,omitempty"`
	TxtBrokerComments                     *string             `bson:"txtBrokerComments,omitempty"`
	TxtValidation                         *string             `bson:"txtValidation,omitempty"`
	Txt30DayQuickSale                     *string             `bson:"txt30DayQuickSale,omitempty"`
	Txt60DayQuickSale                     *string             `bson:"txt60DayQuickSale,omitempty"`
	Txt90DayAsIsValue                     *string             `bson:"txt90DayAsIsValue,omitempty"`
	Txt120DayQuickSale                    *string             `bson:"txt120DayQuickSale,omitempty"`
	Txt180DayQuickSale                    *string             `bson:"txt180DayQuickSale,omitempty"`
	TxtListPriceFinalValues               *string             `bson:"txtListPriceFinalValues,omitempty"`
	Txt30DayListPriceFinalValues          *string             `bson:"txt30DayListPriceFinalValues,omitempty"`
	Txt30DayQuickSaleRepaired             *string             `bson:"txt30DayQuickSaleRepaired,omitempty"`
	Txt60DayQuickSaleRepaired             *string             `bson:"txt60DayQuickSaleRepaired,omitempty"`
	Txt90DayAsIsValueRepaired             *string             `bson:"txt90DayAsIsValueRepaired,omitempty"`
	Txt120DayQuickSaleRepaired            *string             `bson:"txt120DayQuickSaleRepaired,omitempty"`
	Txt180DayQuickSaleRepaired            *string             `bson:"txt180DayQuickSaleRepaired,omitempty"`
	TxtListPriceRepaired                  *string             `bson:"txtListPriceRepaired,omitempty"`
	Txt30DayListPriceRepaired             *string             `bson:"txt30DayListPriceRepaired,omitempty"`
	CmbHouse                              *string             `bson:"cmbHouse,omitempty"`
	CmbPositive                           *string             `bson:"cmbPositive,omitempty"`
	CmbNegative                           *string             `bson:"cmbNegative,omitempty"`
	CmbView                               *string             `bson:"cmbView,omitempty"`
	CmbMarket                             *string             `bson:"cmbMarket,omitempty"`
	CmbPricing                            *string             `bson:"cmbPricing,omitempty"`
	CmbListing                            *string             `bson:"cmbListing,omitempty"`
	CmbExtra                              *string             `bson:"cmbExtra,omitempty"`
	TxtUnique                             *string             `bson:"txtUnique,omitempty"`
	PriceComment                          *string             `bson:"priceComment,omitempty"`
	RangeComment                          *string             `bson:"rangeComment,omitempty"`
	ProxException                         *string             `bson:"proxException,omitempty"`
	GlaException                          *string             `bson:"glaException,omitempty"`
	AgeException                          *string             `bson:"ageException,omitempty"`
	CondException                         *string             `bson:"condException,omitempty"`
	ViewException                         *string             `bson:"viewException,omitempty"`
	StyleException                        *string             `bson:"styleException,omitempty"`
	LotException                          *string             `bson:"lotException,omitempty"`
	BedException                          *string             `bson:"bedException,omitempty"`
	BathException                         *string             `bson:"bathException,omitempty"`
	CompletedDate                         *primitive.DateTime `bson:"competedDate,omitempty"`
	Status                                *string             `bson:"status,omitempty"`
	Logs                                  []IformHistoryLog   `bson:"logs,omitempty"`
}

func (u *Iform) ToModels() *models.Iform {

	logList := make([]*models.IformHistory, 0)
	for _, v := range u.Logs {
		logList = append(logList, v.ToModels())
	}

	return &models.Iform{
		ID:                                    strings.ToObject(u.ID.Hex()),
		FormType:                              u.FormType,
		PipelineID:                            u.PipelineID,
		TxtClient:                             u.TxtClient,
		TxtCompany:                            u.TxtCompany,
		TxtOrderNumber:                        u.TxtOrderNumber,
		CmbOrderType:                          u.CmbOrderType,
		TxtAddress:                            u.TxtAddress,
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
		TxtSubjectAddress:                     u.TxtSubjectAddress,
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
		History:                               logList,
	}
}

func (u *Iform) ToIformTempModels() *models.IformTemp {

	logList := make([]*models.IformHistory, 0)
	for _, v := range u.Logs {
		logList = append(logList, v.ToModels())
	}

	return &models.IformTemp{
		ID:                                    strings.ToObject(u.ID.Hex()),
		FormType:                              u.FormType,
		PipelineID:                            u.PipelineID,
		TxtClient:                             u.TxtClient,
		TxtCompany:                            u.TxtCompany,
		TxtOrderNumber:                        u.TxtOrderNumber,
		CmbOrderType:                          u.CmbOrderType,
		TxtAddress:                            u.TxtAddress,
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
		TxtSubjectAddress:                     u.TxtSubjectAddress,
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
		History:                               logList,
	}
}
