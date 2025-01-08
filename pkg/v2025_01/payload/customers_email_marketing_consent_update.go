package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersEmailMarketingConsentUpdate(webhook webhook.ValidWebhook) (*CustomersEmailMarketingConsentUpdate, error) {
	payload := &CustomersEmailMarketingConsentUpdatePayload{}
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
	return &CustomersEmailMarketingConsentUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CustomersEmailMarketingConsentUpdate struct {
	info webhook.Info
	data CustomersEmailMarketingConsentUpdatePayload
}

func (webhook *CustomersEmailMarketingConsentUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersEmailMarketingConsentUpdate) GetData() (CustomersEmailMarketingConsentUpdatePayload, error) {
	return webhook.data, nil
}

type CustomersEmailMarketingConsentUpdatePayload struct {
	CustomerID            int64       `json:"customer_id"`
	EmailAddress          interface{} `json:"email_address"`
	EmailMarketingConsent struct {
		ConsentUpdatedAt interface{} `json:"consent_updated_at"`
		OptInLevel       interface{} `json:"opt_in_level"`
		State            interface{} `json:"state"`
	} `json:"email_marketing_consent"`
}
