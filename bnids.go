package bnids

import (
	"context"
	"encoding/base64"
	"strings"
)

// Conf is instance of Configuration struct.
var Conf = Configuration{
	BaseURL: "https://apidev.bni.co.id:8066",
}

// Configuration is struct for API Call config.
type Configuration struct {
	BaseURL       string
	ClientID      string
	ClientSecret  string
	DealerID      string
	Username      string
	Password      string
	KodeMitra     string
	KodeLoket     string
	KodeCabang    string
	PinTransaksi  string
	AccountNumber string
}

// OAuthToken is struct contains response from GetToken.
type OAuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

// GetToken function gets token from API.
func GetToken() (OAuthToken, error) {
	return GetTokenWithContext(context.Background())
}

// GetTokenWithContext function gets token from API with context.
func GetTokenWithContext(ctx context.Context) (OAuthToken, error) {
	var endpoint = Conf.BaseURL + "/api/oauth/token"
	var additionalHeaders = map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(Conf.ClientID+":"+Conf.ClientSecret)),
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	var additionalBody = strings.NewReader("grant_type=client_credentials")
	var responseStruct = OAuthToken{}

	SendRequest(ctx, "POST", endpoint, additionalHeaders, additionalBody, &responseStruct)

	return responseStruct, nil
}
