package datastore

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"strings"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/constants"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/email"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/config"
	stringsUtils "github.com/lonmarsDev/bpo-golang-grahpql/pkg/utils/strings"
)

func ComposeEmailBody(body string, vals map[string]string) string {
	for key, value := range vals {
		body = strings.ReplaceAll(body, key, value)
	}
	return body
}

func SendEmailToContractor(ctx context.Context, currentPipeline *Pipeline, newContractorID string) error {
	contractorInfo, err := getPipelineAuthor(ctx, newContractorID)
	if err != nil && contractorInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderAssignment)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = *currentPipeline.OrderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(contractorInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientOrderWasAssigned(ctx context.Context, currentPipeline *Pipeline, newContractorID string) error {
	contractorInfo, err := getPipelineAuthor(ctx, newContractorID)
	if err != nil && contractorInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderAssignment)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = *currentPipeline.OrderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(contractorInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientOrderOnHold(ctx context.Context, currentPipeline *Pipeline) error {

	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		log.Debug("error %v", err)
		return err
	}
	messageBody := fmt.Sprintf("your Order: %s was on hold.", *currentPipeline.OrderNumber)
	parsedMessageBody, err := ParseTemplate(messageBody)
	if err != nil {
		log.Debug("error %v", err)
		return err
	}

	err = email.Init(clientInfo.Email, "Order onhold", parsedMessageBody)
	if err != nil {
		log.Debug("error %v", err)
		return err
	}
	return nil
}

func SendEmailToClientOrderUnHold(ctx context.Context, currentPipeline *Pipeline) error {

	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		return err
	}
	messageBody := fmt.Sprintf("Order: %s was unhold.", *currentPipeline.OrderNumber)
	err = email.Init(clientInfo.Email, "Order unhold", messageBody)
	if err != nil {
		log.Debug("error %v", err)
		return err
	}
	err = email.Init("orders@turbo.com", "Order unhold", messageBody)
	if err != nil {
		log.Debug("error %v", err)
		return err
	}
	return nil
}

func SendEmailToClientOrderStandBy(ctx context.Context, currentPipeline *Pipeline) error {

	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		return err
	}

	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderComplateForStandby)
	if err != nil {
		return err
	}
	var parsedKey map[string]string
	parsedKey[constants.EmailParseKeyOrderNumber] = stringsUtils.ObjectTOString(currentPipeline.OrderNumber)
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientOrderComplete(ctx context.Context, currentPipeline *Pipeline) error {

	clientInfo, err := getPipelineAuthor(ctx, *currentPipeline.UserId)
	if err != nil && clientInfo == nil {
		return err
	}

	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderComplete)
	if err != nil {
		return err
	}
	var parsedKey map[string]string
	parsedKey[constants.EmailParseKeyOrderNumber] = *currentPipeline.OrderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil

}

func SendEmailToClientNewOrder(ctx context.Context, clientID, orderNumber string) error {

	clientInfo, err := getPipelineAuthor(ctx, clientID)
	if err != nil && clientInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderNew)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientOrderRush(ctx context.Context, clientID, orderNumber string) error {

	clientInfo, err := getPipelineAuthor(ctx, clientID)
	if err != nil && clientInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderRush)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientOrderSuperRush(ctx context.Context, clientID, orderNumber string) error {

	clientInfo, err := getPipelineAuthor(ctx, clientID)
	if err != nil && clientInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderSuperRush)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailToClientCancelOrder(ctx context.Context, clientID, orderNumber string) error {

	clientInfo, err := getPipelineAuthor(ctx, clientID)
	if err != nil && clientInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderCancel)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

//send to client of rush order
// func SendEmailToClientRushOrder(ctx context.Context, clientID, orderNumber string) error {

// 	clientInfo, err := getPipelineAuthor(ctx, clientID)
// 	if err != nil && clientInfo == nil {
// 		return err
// 	}
// 	//get email template
// 	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderCancel)
// 	if err != nil && emailTemplate == nil {
// 		return err
// 	}
// 	var parsedKey = make(map[string]string)
// 	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
// 	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
// 	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//account email notification

func SendEmailAccountRegistration(ctx context.Context, clientID, orderNumber string) error {

	clientInfo, err := getPipelineAuthor(ctx, clientID)
	if err != nil && clientInfo == nil {
		return err
	}
	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateOrderNew)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)
	parsedKey[constants.EmailParseKeyOrderNumber] = orderNumber
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(clientInfo.Email, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailAccountForgetPassword(ctx context.Context, emailAddress, token string) error {

	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateAccountForgetPassword)
	if err != nil && emailTemplate == nil {
		return err
	}
	var parsedKey = make(map[string]string)

	resetpasswordLink := fmt.Sprintf("%s%s", config.AppConfig.GetString("appResetPasswordUrl"), token)
	parsedKey[constants.EmailParseKeyAccountResetPassword] = resetpasswordLink
	messageBody := ComposeEmailBody(emailTemplate.Message, parsedKey)
	err = email.Init(emailAddress, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailAccountChangePasswordSuccess(ctx context.Context, emailAddress string) error {

	//get email template
	emailTemplate, err := GetEmailTemplateByTemplateCode(ctx, constants.EmailTemplateAccountChangePasswordSuccess)
	if err != nil && emailTemplate == nil {
		return err
	}
	messageBody := ComposeEmailBody(emailTemplate.Message, nil)
	err = email.Init(emailAddress, emailTemplate.Subject, messageBody)
	if err != nil {
		return err
	}
	return nil
}

func ParseTemplate(data interface{}) (content string, err error) {

	// ParseFiles creates a new 	Template and parses the template definitions from
	// the named files. The returned template's name will have the (base) name and
	// (parsed) contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	tmpl, err := template.ParseFiles("../pkg/email/email_template.html")
	if err != nil {
		return "", err
	}

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	buf := new(bytes.Buffer)

	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
