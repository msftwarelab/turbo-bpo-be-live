package datastore

import (
	"github.com/lonmarsDev/bpo-golang-grahpql/graphql/models"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credit struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	UserId          string              `bson:"userId,omitempty"`
	PaypalOrderID   *string             `bson:"paypalOrderId,omitempty"`
	PaypalToken     *string             `bson:"paypalToken,omitempty"`
	Amount          *float64            `bson:"amount,omitempty"`
	CreatedDateTime primitive.DateTime  `bson:"createdDateTime"`
	LastUpdateTime  *primitive.DateTime `bson:"lastUpdateTime,omitempty"`
}

type CreditTotal struct {
	TotalAmount float64 `bson:"totalAmount"`
}

func (u *Credit) ToModels() *models.Credit {
	return &models.Credit{
		ID:              strings.ToObject(u.ID.Hex()),
		PaypalOrderID:   u.PaypalOrderID,
		PaypalToken:     u.PaypalToken,
		Amount:          u.Amount,
		CreatedDateTime: strings.ToObject(TimeConversion(&u.CreatedDateTime)),
	}
}

type CreditLedger struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	ClientName      *string             `bson:"clientName,omitempty"`
	ClientID        *string             `bson:"clientId,omitempty"`
	PaypalOrderID   *string             `bson:"paypalOrderId,omitempty"`
	Type            *string             `bson:"type,omitempty"`
	OrderNumber     *string             `bson:"orderNumber,omitempty"`
	OrderAddress    *string             `bson:"orderAddress,omitempty"`
	CreatedDateTime *primitive.DateTime `bson:"createdDateTime,omitempty"`
	Balance         *float64            `bson:"balance,omitempty"`
	IformCharge     *float64            `bson:"iformCharge,omitempty"`
	Amount          *float64            `bson:"amount,omitempty"`
}

func (u *CreditLedger) ToModels() *models.CreditLedger {
	return &models.CreditLedger{
		ID:              strings.ToObject(u.ID.Hex()),
		ClientName:      u.ClientName,
		ClientID:        u.ClientID,
		PaypalOrderID:   u.PaypalOrderID,
		Type:            u.Type,
		OrderNumber:     u.OrderNumber,
		OrderAddress:    u.OrderAddress,
		CreatedDateTime: strings.ToObject(TimeConversion(u.CreatedDateTime)),
		Balance:         u.Balance,
		IformCharge:     u.IformCharge,
		Amount:          u.Amount,
	}
}
