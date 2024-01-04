package datastore

import (
	"context"

	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Billing struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	InvoiceNumber   *string             `bson:"invoiceNumber,omitempty"`
	Status          *string             `bson:"status,omitempty"`
	Date            *primitive.DateTime `bson:"date,omitempty"`
	DateFrom        *primitive.DateTime `bson:"dateFrom,omitempty"`
	DateTo          *primitive.DateTime `bson:"dateTo,omitempty"`
	DueDate         *primitive.DateTime `bson:"dueDate,omitempty"`
	UserID          *string             `bson:"userId,omitempty"`
	UserName        *string             `bson:"userName,omitempty"`
	Entries         []*BillingEntry     `bson:"entries,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

type BillingEntry struct {
	OrderNumber *string  `bson:"orderNumber,omitempty"`
	Description *string  `bson:"description,omitempty"`
	Amount      *float64 `bson:"amount,omitempty"`
	Type        *string  `bson:"type,omitempty"`
}

func (u *Billing) ToModels() *models.Billing {

	entries := []*models.BillingEntry{}
	if len(u.Entries) > 0 {
		for _, v := range u.Entries {
			entry := &models.BillingEntry{
				OrderNumber: v.OrderNumber,
				Description: v.Description,
				Amount:      v.Amount,
				Type:        v.Type,
			}
			entries = append(entries, entry)
		}
	}

	return &models.Billing{
		ID:            strings.ToObject(u.ID.Hex()),
		InvoiceNumber: u.InvoiceNumber,
		Status:        u.Status,
		Date:          strings.ToObject(TimeConversion(u.Date)),
		DateFrom:      strings.ToObject(TimeConversion(u.DateFrom)),
		DateTo:        strings.ToObject(TimeConversion(u.DateTo)),
		DueDate:       strings.ToObject(TimeConversion(u.DueDate)),
		UserID:        u.UserID,
		UserName:      u.UserName,
		Entries:       entries,
	}
}

func (u *Billing) GetOrdernumbers() []string {

	var orderNumbers []string
	for _, v := range u.Entries {
		orderNumbers = append(orderNumbers, strings.ObjectTOString(v.OrderNumber))
	}
	return orderNumbers
}

func (u *Billing) ClientInfo(ctx context.Context) *User {
	if u.UserID == nil {
		return nil
	}
	filter := FilterById(*u.UserID)
	userRaw, err := GetUser(ctx, filter)
	if err != nil {
		return nil
	}
	return userRaw
}

func (u *Billing) FormatedDate() string {
	return TimeConversion(u.Date)
}
