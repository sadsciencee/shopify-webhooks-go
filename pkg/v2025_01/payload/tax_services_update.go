package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewTaxServicesUpdate(webhook webhook.ValidWebhook) (*TaxServicesUpdate, error) {
	payload := &TaxServicesUpdatePayload{}
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
	return &TaxServicesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type TaxServicesUpdate struct {
	info webhook.Info
	data TaxServicesUpdatePayload
}

func (webhook *TaxServicesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *TaxServicesUpdate) GetData() (TaxServicesUpdatePayload, error) {
	return webhook.data, nil
}

type TaxServicesUpdatePayload struct {
	Active bool    `json:"active"`
	ID     *string `json:"id"`
	Name   string  `json:"name"`
	URL    string  `json:"url"`
}
