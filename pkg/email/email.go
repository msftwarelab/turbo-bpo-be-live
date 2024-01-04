package email

import (
	"context"
	"encoding/json"

	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

var yourDomain string = "turbopipeline.com" // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = "4885edc9007912509e376a65f94232b2-4879ff27-b15fcaba"

func Init(recipient, subject, messageBody string) error {

	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "do.not.reply@turbo-bpo.com"
	body := messageBody

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}

func ToObject(i string) *string {
	return &i
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
