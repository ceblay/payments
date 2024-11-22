package adapters

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

var payDunyaBaseURL = "https://app.paydunya.com/sandbox-api"

const MASTER_KEY = "53Dlks95-Mc6y-eSnN-6j1F-OvDOOgXhGcGJ"
const PRIVATE_KEY = "test_private_HnAb25VYDooJ9pY87Xt9W3bYUG1"
const TOKEN = "o8FCTdANsYKsdagZhu4I"

type PayDunyaPaymentService struct {
	Url     string
	Country string
}
type payDunyaInvoice struct {
	Amount      float32 `json:"total_amount"`
	Description string  `json:"description"`
}
type payDunyaShop struct {
	Name string `json:"name"`
}

type payDunyaPaymentPayload struct {
	Invoice payDunyaInvoice `json:"invoice"`
	Shop    payDunyaShop    `json:"store"`
}

func NewPayDunyaPaymentService() *PayDunyaPaymentService {
	// if countryCode == "" || countryCode != "SN" {
	// 	countryCode = "SN"
	// }
	return &PayDunyaPaymentService{
		Url:     payDunyaBaseURL,
		Country: "SN", // countryCode,
	}
}

func (p *PayDunyaPaymentService) InitiatePayment(amount float32) (string, error) {
	payload := payDunyaPaymentPayload{
		Invoice: payDunyaInvoice{
			Amount:      amount,
			Description: "A quick payment",
		},
		Shop: payDunyaShop{
			Name: "Funky Store",
		},
	}

	_, body, errs := gorequest.New().Post(fmt.Sprintf("%s/v1/checkout-invoice/create", p.Url)).
		Set("PAYDUNYA-MASTER-KEY", MASTER_KEY).
		Set("PAYDUNYA-PRIVATE-KEY", PRIVATE_KEY).
		Set("PAYDUNYA-TOKEN", TOKEN).
		SendStruct(payload).
		End()

	for _, err := range errs {
		if err != nil {
			return "", fmt.Errorf("Error initiating payment on PayDunya: %v", err)
		}
	}

	return body, nil
}
