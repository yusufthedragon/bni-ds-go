package bnids

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	Conf.ClientID = ""
	Conf.ClientSecret = ""
	Conf.DealerID = ""
	Conf.Username = ""
	Conf.Password = ""
	Conf.KodeMitra = ""
	Conf.KodeLoket = ""
	Conf.KodeCabang = ""
	Conf.PinTransaksi = ""
	Conf.AccountNumber = ""
}

func TestGetToken(t *testing.T) {
	oAuthToken, err := GetToken()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("OAuthToken: %+v\n", oAuthToken)
}
