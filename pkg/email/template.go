package email

import (
	"context"

	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

type EmailPipelineNotification struct {
	Action      string // hold, unhold, rush, super rush, new order
	Recipient   string
	Address     string
	Body        string
	OrderNumber string
	OrderType   string
	Company     string
}

func PipelineNotification(input EmailPipelineNotification) error {
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Create a new message with template
	m := mg.NewMessage("Turbo-BPO Order <no-reply@turbo-bpo.com>", input.Action, input.Action)
	m.SetTemplate("default")

	// Add recipients
	m.AddRecipient(input.Recipient)
	// Add the variables to be used by the template
	m.AddVariable("address", input.Address)
	m.AddVariable("body", input.Body)
	m.AddVariable("orderNumber", input.OrderNumber)
	m.AddVariable("orderType", input.OrderType)
	m.AddVariable("company", input.Company)

	_, id, err := mg.Send(ctx, m)
	fmt.Printf("Queued: %s", id)
	return err
}
