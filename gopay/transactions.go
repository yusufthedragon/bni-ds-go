package gopay

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	bnids "github.com/yusufthedragon/bni-ds-go"
)

// CreateInquiry function creates Gopay Inquiry.
func CreateInquiry(params *InquiryParams) (Inquiry, error) {
	return CreateInquiryWithContext(context.Background(), params)
}

// CreateInquiryWithContext function creates Gopay Inquiry with context.
func CreateInquiryWithContext(ctx context.Context, params *InquiryParams) (Inquiry, error) {
	var responseStruct = Inquiry{}

	bnids.SendRequestTransaction(ctx, &params, &responseStruct)

	return responseStruct, nil
}

// CreatePayment function creates Gopay Payment.
func CreatePayment(params *PaymentParams) (Payment, error) {
	return CreatePaymentWithContext(context.Background(), params)
}

// CreatePaymentWithContext function creates Gopay Payment with context.
func CreatePaymentWithContext(ctx context.Context, params *PaymentParams) (Payment, error) {
	var token, errGetToken = bnids.GetToken()

	if errGetToken != nil {
		log.Fatalln(errGetToken)
	}

	var endpoint = bnids.Conf.BaseURL + "/keagenan/transaction?access_token=" + token.AccessToken
	var additionalHeaders = map[string]string{
		"Content-Type": "application/json",
		"X-Auth":       base64.StdEncoding.EncodeToString([]byte(bnids.Conf.Username + ":" + bnids.Conf.Password)),
	}
	var responseStruct = Payment{}
	var additionalBody, err = json.Marshal(params)

	if err != nil {
		log.Fatalln(err)
	}

	bnids.SendRequest(ctx, "POST", endpoint, additionalHeaders, bytes.NewBuffer(additionalBody), &responseStruct)

	return responseStruct, nil
}
