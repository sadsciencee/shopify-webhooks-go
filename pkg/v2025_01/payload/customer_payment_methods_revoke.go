package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomerPaymentMethodsRevoke(webhook webhook.ValidWebhook) (*CustomerPaymentMethodsRevoke, error) {
	payload := &CustomerPaymentMethodsRevokePayload{}
	info, err := webhook.GetInfo()
	if err != nil {
		return nil, err
	}
	reader, err := webhook.GetReader()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(payload); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}
	return &CustomerPaymentMethodsRevoke{
		info: info,
		data: *payload,
	}, nil
}

type CustomerPaymentMethodsRevoke struct {
	info webhook.Info
	data CustomerPaymentMethodsRevokePayload
}

func (webhook *CustomerPaymentMethodsRevoke) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerPaymentMethodsRevoke) GetData() (CustomerPaymentMethodsRevokePayload, error) {
	return webhook.data, nil
}

type CustomerPaymentMethodsRevokePayload struct {
	AdminGraphqlAPICustomerID string `json:"admin_graphql_api_customer_id"`
	AdminGraphqlAPIID         string `json:"admin_graphql_api_id"`
	CustomerID                int64  `json:"customer_id"`
	InstrumentType            string `json:"instrument_type"`
	PaymentInstrument         struct {
		Brand      string `json:"brand"`
		LastDigits string `json:"last_digits"`
		Month      int64  `json:"month"`
		Name       string `json:"name"`
		Year       int64  `json:"year"`
	} `json:"payment_instrument"`
	Token string `json:"token"`
}
