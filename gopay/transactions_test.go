package gopay

import (
	"fmt"
	"log"
	"testing"

	bnids "github.com/yusufthedragon/bni-ds-go"
)

func init() {
	bnids.Conf.ClientID = ""
	bnids.Conf.ClientSecret = ""
	bnids.Conf.DealerID = ""
	bnids.Conf.Username = ""
	bnids.Conf.Password = ""
	bnids.Conf.KodeMitra = ""
	bnids.Conf.KodeLoket = ""
	bnids.Conf.KodeCabang = ""
	bnids.Conf.PinTransaksi = ""
	bnids.Conf.AccountNumber = ""
}

func TestGopayInquiry(t *testing.T) {
	data := InquiryParams{
		TypeReq:  "trans",
		DealerID: bnids.Conf.DealerID,
		FiturID:  "FTR497",
		Data: DataInquiryParams{
			KodeMitra:    bnids.Conf.KodeMitra,
			KodeLoket:    bnids.Conf.KodeLoket,
			KodeCabang:   bnids.Conf.KodeCabang,
			MCPID:        "GOPAY_CUSTOMER",
			Amount:       30000,
			BillingID1:   "081290731455",
			PinTransaksi: bnids.Conf.PinTransaksi,
			AccountNum:   bnids.Conf.AccountNumber,
		},
	}

	inquiry, err := CreateInquiry(&data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gopay Inquiry: %+v\n", inquiry)
}

func TestGopayPayment(t *testing.T) {
	data := PaymentParams{
		TypeReq:  "trans",
		DealerID: bnids.Conf.DealerID,
		FiturID:  "FTR498",
		Data: DataPaymentParams{
			KodeMitra:       bnids.Conf.KodeMitra,
			KodeLoket:       bnids.Conf.KodeLoket,
			KodeCabang:      bnids.Conf.KodeCabang,
			MCPID:           "GOPAY_CUSTOMER",
			HCustomerNumber: "081290731455",
			HCustomerName:   "KRISNALDI DUMB",
			HTotalAmount:    "0",
			HReference:      "825783320042456",
			CFee:            "0",
			JenisLayanan:    "Top Up GO-PAY Customer",
			BiayaLoket:      "1000",
			TotalBayar:      "30000",
			HAccountNumber:  bnids.Conf.AccountNumber,
			PinTransaksi:    bnids.Conf.PinTransaksi,
		},
	}

	payment, err := CreatePayment(&data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Gopay Payment: %+v\n", payment)
}
