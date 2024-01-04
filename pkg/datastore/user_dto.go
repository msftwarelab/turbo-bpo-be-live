package datastore

import (
	"fmt"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                                   primitive.ObjectID  `bson:"_id,omitempty"`
	Email                                string              `bson:"email,omitempty"`
	Fullname                             string              `bson:"fullName,omitempty"`
	Password                             string              `bson:"password,omitempty"`
	FirstName                            string              `bson:"firstName,omitempty"`
	LastName                             string              `bson:"lastName,omitempty"`
	Company                              *string             `bson:"company,omitempty"`
	PhoneNumber                          *string             `bson:"phoneNumber,omitempty"`
	Address                              *string             `bson:"address,omitempty"`
	City                                 *string             `bson:"city,omitempty"`
	State                                *string             `bson:"state,omitempty"`
	Zipcode                              *string             `bson:"zipCode,omitempty"`
	Title                                *string             `bson:"title,omitempty"`
	About                                *string             `bson:"about,omitempty"`
	Roles                                []*string           `bson:"role"`
	Status                               string              `bson:"status"`
	Hdyfu                                *string             `bson:"hdyfu,omitempty"`
	PhoneConsultation                    *bool               `bson:"phoneConsultation,omitempty"`
	ImABroker                            *bool               `bson:"imABroker,omitempty"`
	Broker                               *string             `bson:"broker,omitempty"`
	BrokerLicense                        *string             `bson:"brokerLicense,omitempty"`
	Agent                                *string             `bson:"agent,omitempty"`
	AgentLicense                         *string             `bson:"agentLicense,omitempty"`
	LicenseDate                          *string             `bson:"licenseDate,omitempty"`
	LicenseExpirationDate                *string             `bson:"licenseExpirationDate,omitempty"`
	Brokerage                            *string             `bson:"brokerage,omitempty"`
	YearOfExperience                     *string             `bson:"yearOfExperience,omitempty"`
	ProfilePicture                       *string             `bson:"profilePicture,omitempty"`
	CreatedDateTime                      primitive.DateTime  `bson:"createdDateTime"`
	OrderTotal                           int                 `bson:"orderTotalomitempty"`
	AssignmentPercentage                 float64             `bson:"assignmentPercentage"`
	LastUpdateTime                       *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
	UpdateLogs                           []*UserLogRaw       `bson:"updateLogs,omitempty"`
	CompanyList                          []*string           `bson:"companyList,omitempty"`
	PermissionGroupID                    *string             `bson:"permissionGroupId,omitempty"`
	ResetPasswordToken                   *string             `bson:"resetPasswordToken,omitempty"`
	ResetPasswordTokenExpirationDateTime *primitive.DateTime `bson:"resetPasswordTokenExpirationDateTime,omitempty"`
	IsResetPasswordTokenUsed             *bool               `bson:"isResetPasswordTokenUsed,omitempty"`
	PriceModule                          *PriceModule        `bson:"priceModule,omitempty"`
	Disclaimer                           *string             `bson:"disclaimer,omitempty"`
	Theme                                *string             `bson:"theme,omitempty"`
	IsEnableEmailNotification            *bool               `json:"isEnableEmailNotification"`
	AssignDate                           *primitive.DateTime `bson:"assignDate"`
	AssignActive                         *int                `bson:"assignActive"`
	AssignHold                           *int                `bson:"assignHold"`
	AssignRush                           *int                `bson:"assignRush"`
	AssignStandby                        *int                `bson:"assignStandby"`
}

type PriceModule struct {
	Credits                  *float64 `json:"credits"`
	Orderinterior            *float64 `bson:"orderinterior,omitempty"`
	Orderexterior            *float64 `bson:"orderexterior,omitempty"`
	OrderdataEntry           *float64 `bson:"orderdataEntry,omitempty"`
	Orderrush                *float64 `bson:"orderrush,omitempty"`
	OrdersuperRush           *float64 `bson:"ordersuperRush,omitempty"`
	OrderconditionReport     *float64 `bson:"orderconditionReport,omitempty"`
	OrderrentalAddendum      *float64 `bson:"orderrentalAddendum,omitempty"`
	PhotoExterior            *float64 `bson:"photoExterior,omitempty"`
	PhotoInteriorVacantLb    *float64 `bson:"photoInteriorVacantLB,omitempty"`
	PhotoInteriorAppointment *float64 `bson:"photoInteriorAppointment,omitempty"`
}

func (u *User) ToModels() *models.User {
	var companyList []*string
	for _, v := range u.CompanyList {
		companyList = append(companyList, v)
	}
	priceModule := &models.PriceModule{}

	if u.PriceModule != nil {
		priceModule = &models.PriceModule{
			Credits:                  u.PriceModule.Credits,
			Orderinterior:            u.PriceModule.Orderinterior,
			Orderexterior:            u.PriceModule.Orderexterior,
			OrderdataEntry:           u.PriceModule.OrderdataEntry,
			Orderrush:                u.PriceModule.Orderrush,
			OrdersuperRush:           u.PriceModule.OrdersuperRush,
			OrderconditionReport:     u.PriceModule.OrderconditionReport,
			OrderrentalAddendum:      u.PriceModule.OrderrentalAddendum,
			PhotoExterior:            u.PriceModule.PhotoExterior,
			PhotoInteriorVacantLb:    u.PriceModule.PhotoInteriorVacantLb,
			PhotoInteriorAppointment: u.PriceModule.PhotoInteriorAppointment,
		}
	}
	return &models.User{
		ID:                        u.ID.Hex(),
		Email:                     u.Email,
		FirstName:                 &u.FirstName,
		LastName:                  &u.LastName,
		Company:                   u.Company,
		PhoneNumber:               u.PhoneNumber,
		Address:                   u.Address,
		State:                     u.State,
		ZipCode:                   u.Zipcode,
		Title:                     u.Title,
		About:                     u.About,
		Hdyfu:                     u.Hdyfu,
		PhoneConsultation:         u.PhoneConsultation,
		ImABroker:                 u.ImABroker,
		Broker:                    u.Broker,
		BrokerLicense:             u.BrokerLicense,
		Agent:                     u.Agent,
		AgentLicense:              u.AgentLicense,
		LicenseDate:               u.LicenseDate,
		LicenseExpirationDate:     u.LicenseExpirationDate,
		Brokerage:                 u.Brokerage,
		YearOfExperience:          u.YearOfExperience,
		ProfilePicture:            u.ProfilePicture,
		Roles:                     u.Roles,
		Status:                    &u.Status,
		City:                      u.City,
		CreatedDateTime:           strings.ToObject(TimeConversion(&u.CreatedDateTime)),
		LastUpdateTime:            strings.ToObject(TimeConversion(u.LastUpdateTime)),
		CompanyList:               companyList,
		PermissionGroupID:         u.PermissionGroupID,
		PriceModule:               priceModule,
		Disclaimer:                u.Disclaimer,
		Theme:                     u.Theme,
		IsEnableEmailNotification: u.IsEnableEmailNotification,
		AssignDate:                strings.ToObject(TimeConversion(u.AssignDate)),
		AssignActive:              u.AssignActive,
		AssignHold:                u.AssignHold,
		AssignRush:                u.AssignRush,
		AssignStandby:             u.AssignStandby,
	}
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type UserLogRaw struct {
	Datetime primitive.DateTime `bson:"datetime"`
	Value    string             `bson:"value"`
}
