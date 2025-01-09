package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProductPublicationsDelete(webhook webhook.ValidWebhook) (*ProductPublicationsDelete, error) {
	payload := &ProductPublicationsDeletePayload{}
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
	return &ProductPublicationsDelete{
		info: info,
		data: *payload,
	}, nil
}

type ProductPublicationsDelete struct {
	info webhook.Info
	data ProductPublicationsDeletePayload
}

func (webhook *ProductPublicationsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductPublicationsDelete) GetData() (ProductPublicationsDeletePayload, error) {
	return webhook.data, nil
}

type ProductPublicationsDeletePayload struct {
	ID interface{} `json:"id"`
}
