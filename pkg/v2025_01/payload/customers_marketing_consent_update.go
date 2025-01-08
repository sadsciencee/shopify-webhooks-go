package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersMarketingConsentUpdate(webhook webhook.ValidWebhook) (*CustomersMarketingConsentUpdate, error) {
	payload := &CustomersMarketingConsentUpdatePayload{}
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
	return &CustomersMarketingConsentUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CustomersMarketingConsentUpdate struct {
	info webhook.Info
	data CustomersMarketingConsentUpdatePayload
}

func (webhook *CustomersMarketingConsentUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersMarketingConsentUpdate) GetData() (CustomersMarketingConsentUpdatePayload, error) {
	return webhook.data, nil
}

type CustomersMarketingConsentUpdatePayload struct {
	ID                  int64       `json:"id"`
	Phone               interface{} `json:"phone"`
	SmsMarketingConsent struct {
		ConsentCollectedFrom string      `json:"consent_collected_from"`
		ConsentUpdatedAt     interface{} `json:"consent_updated_at"`
		OptInLevel           interface{} `json:"opt_in_level"`
		State                interface{} `json:"state"`
	} `json:"sms_marketing_consent"`
}
