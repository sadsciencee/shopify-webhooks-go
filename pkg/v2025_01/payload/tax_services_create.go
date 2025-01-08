package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewTaxServicesCreate(webhook webhook.ValidWebhook) (*TaxServicesCreate, error) {
	payload := &TaxServicesCreatePayload{}
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
	return &TaxServicesCreate{
		info: info,
		data: *payload,
	}, nil
}

type TaxServicesCreate struct {
	info webhook.Info
	data TaxServicesCreatePayload
}

func (webhook *TaxServicesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *TaxServicesCreate) GetData() (TaxServicesCreatePayload, error) {
	return webhook.data, nil
}

type TaxServicesCreatePayload struct {
	Active bool   `json:"active"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}
