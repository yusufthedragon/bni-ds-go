package bnids

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// SendRequest function sends request to API.
func SendRequest(ctx context.Context, httpMethod string, endpoint string, additionalHeaders map[string]string, additionalBody io.Reader, responseStruct interface{}) {
	var req, err = http.NewRequestWithContext(ctx, httpMethod, endpoint, additionalBody)

	if err != nil {
		log.Fatalln(err)
	}

	// Set the request headers
	for key, value := range additionalHeaders {
		req.Header.Set(key, value)
	}

	var client = &http.Client{}
	var resp, errRequest = client.Do(req)

	if errRequest != nil {
		log.Fatalln(errRequest)
	}

	defer resp.Body.Close()

	// Read the response body
	var body, errResponse = ioutil.ReadAll(resp.Body)

	if errResponse != nil {
		log.Fatalln(errResponse)
	}

	// Decode the response JSON into target struct
	fmt.Println(string(body))
	json.Unmarshal(body, &responseStruct)
}

// SendRequestTransaction sends request for transaction.
func SendRequestTransaction(ctx context.Context, params interface{}, responseStruct interface{}) {
	var token, errGetToken = GetToken()

	if errGetToken != nil {
		log.Fatalln(errGetToken)
	}

	var endpoint = Conf.BaseURL + "/keagenan/transaction?access_token=" + token.AccessToken
	var additionalHeaders = map[string]string{
		"Content-Type": "application/json",
		"X-Auth":       base64.StdEncoding.EncodeToString([]byte(Conf.Username + ":" + Conf.Password)),
	}
	var additionalBody, err = json.Marshal(params)

	if err != nil {
		log.Fatalln(err)
	}

	SendRequest(ctx, "POST", endpoint, additionalHeaders, bytes.NewBuffer(additionalBody), &responseStruct)
}
