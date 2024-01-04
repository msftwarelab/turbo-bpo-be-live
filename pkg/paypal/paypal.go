package paypal

import (
	"encoding/json"
	"errors"

	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
	"github.com/plutov/paypal"
)

func ValidateTransaction(paypalOrderID string) error {
	//const clientID = "Aex4EBhqwxIC1rn-FFYnAsesx95xJKW0Rm7ajHWlb6_ScaN015eAD_8YFTRjpW5Mg8mfNOO6qx2jVbTb" //test
	//const secret = "EGJ3cA7JScMmob-IOs9CrB6rqQRRZVOzf-KtmJHpLy4j7bRfNg23QhPx2uHI2lwBU3NLPm3vI4Qfrg9H"   //test
	const clientID = "AaNhJKIMAAzkrzSNHLRQemf3tGQu23BhD6xA0WnMBtaCMCylJduFmAwl7rWt88JkPm-hPR2goZfNriKN" //live
	const secret = "EILIt4UeT0Kws1wHSlghjeX6qhZfGCmoi07YaLIxzAh_6tx2KVOnNIuYs4v_ZuR3m9wi3mdcRTpTSAnB"   //live
	//client, err := paypal.NewClient(clientID, secret, paypal.APIBaseSandBox)  /testing
	client, err := paypal.NewClient(clientID, secret, paypal.APIBaseLive)
	if err != nil {
		return err
	}
	accessToken, err := client.GetAccessToken()
	log.Debug("accessToken : %s", accessToken)

	//	capture, err := c.CaptureOrder("asdasdasd", paypal.CaptureOrderRequest{})
	//{paypalOrderId: "7BU303056U427781Y", billingId: "5de8cbe16c3cb8d9d8a5a9c3"}
	//c.SetAccessToken()
	oderInfo, err := client.GetOrder(paypalOrderID)
	if err != nil {
		return err
	}
	log.Debug("orderInfo %s", PrettyPrint(oderInfo))

	if oderInfo.Status != "COMPLETED" {
		return errors.New("paypal transaction failed")
	}
	return nil

}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
