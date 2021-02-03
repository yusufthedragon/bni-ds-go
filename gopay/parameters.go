package gopay

// DataInquiryParams is struct contains parameter data for create Inquiry.
type DataInquiryParams struct {
	KodeMitra    string `json:"kode_mitra"`
	KodeLoket    string `json:"kode_loket"`
	KodeCabang   string `json:"kode_cabang"`
	MCPID        string `json:"MCP_ID"`
	Amount       int    `json:"amount"`
	BillingID1   string `json:"billingId1"`
	PinTransaksi string `json:"pin_transaksi"`
	AccountNum   string `json:"account_num"`
}

// DataPaymentParams is struct contains parameter data for create Payment.
type DataPaymentParams struct {
	KodeMitra       string `json:"kode_mitra"`
	KodeLoket       string `json:"kode_loket"`
	KodeCabang      string `json:"kode_cabang"`
	MCPID           string `json:"MCP_ID"`
	HCustomerNumber string `json:"h_customerNumber"`
	HCustomerName   string `json:"h_customerName"`
	HTotalAmount    string `json:"h_totalAmount"`
	HReference      string `json:"h_reference"`
	CFee            string `json:"c_fee"`
	JenisLayanan    string `json:"jenis_layanan"`
	BiayaLoket      string `json:"biaya_loket"`
	TotalBayar      string `json:"total_bayar"`
	HAccountNumber  string `json:"h_accountNumber"`
	PinTransaksi    string `json:"pin_transaksi"`
}

// InquiryParams is struct contains parameters for create Inquiry.
type InquiryParams struct {
	TypeReq  string            `json:"typereq"`
	DealerID string            `json:"dealerId"`
	FiturID  string            `json:"fiturId"`
	Data     DataInquiryParams `json:"data"`
}

// Inquiry is struct contains Inquiry data from API response.
type Inquiry struct {
	Error           bool   `json:"error"`
	PaymentType     string `json:"PAYMENT_TYPE"`
	Channel         string `json:"CHANNEL"`
	MCPID           string `json:"MCP_ID"`
	HCustomerName   string `json:"h_customerName"`
	HCurrency       string `json:"h_currency"`
	HTotalAmount    string `json:"h_totalAmount"`
	CFee            string `json:"c_fee"`
	HCompanyCode    string `json:"h_companyCode"`
	HCustomerNumber string `json:"h_customerNumber"`
	HRequestID      string `json:"h_requestId"`
	HInquiryStatus  string `json:"h_inquiryStatus"`
	HReference      string `json:"h_reference"`
	BiayaLoket      string `json:"biaya_loket"`
}

// PaymentParams is struct contains parameters for create Payment.
type PaymentParams struct {
	TypeReq  string            `json:"typereq"`
	DealerID string            `json:"dealerId"`
	FiturID  string            `json:"fiturId"`
	Data     DataPaymentParams `json:"data"`
}

// Payment is struct contains Payment data from API response.
type Payment struct {
	TrxType string `json:"trxType"`
	Status  string `json:"status"`
	Data    struct {
		Error              bool   `json:"error"`
		PaymentType        string `json:"PAYMENT_TYPE"`
		Channel            string `json:"CHANNEL"`
		MCPID              string `json:"MCP_ID"`
		HCustomerName      string `json:"h_customerName"`
		HCurrency          string `json:"h_currency"`
		HPaidAmount        string `json:"h_paidAmount"`
		HTotalAmount       string `json:"h_totalAmount"`
		HTransactionDate   string `json:"h_transactionDate"`
		CFee               string `json:"c_fee"`
		HAccountNumber     string `json:"h_accountNumber"`
		CPayMethod         string `json:"c_payMethod"`
		CDescription       string `json:"c_description"`
		HCustomerNumber    string `json:"h_customerNumber"`
		HPaymentFlagStatus string `json:"h_paymentFlagStatus"`
		Journal            string `json:"journal"`
		HReference         string `json:"h_reference"`
		BiayaLoket         string `json:"biaya_loket"`
		TotalBayar         string `json:"total_bayar"`
		JenisLayanan       string `json:"jenis_layanan"`
		ErrorNum           string `json:"errorNum"`
	} `json:"data"`
	Result struct {
		KodeLoket  string `json:"kode_loket"`
		KdLkt      string `json:"kd_lkt"`
		Nama       string `json:"nama"`
		KodeCabang string `json:"kode_cabang"`
		KodeMitra  string `json:"kode_mitra"`
		Alamat     string `json:"alamat"`
	} `json:"result"`
	CustomerData struct {
		NoRek    string `json:"no_rek"`
		BiayaAdm string `json:"biaya_adm"`
		NoTel    string `json:"no_tel"`
		Nominal  string `json:"nominal"`
		Time     string `json:"time"`
	} `json:"CustomerData"`
}
