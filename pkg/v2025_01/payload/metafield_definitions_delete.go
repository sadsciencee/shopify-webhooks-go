package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewMetafieldDefinitionsDelete(webhook webhook.ValidWebhook) (*MetafieldDefinitionsDelete, error) {
	payload := &MetafieldDefinitionsDeletePayload{}
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
	return &MetafieldDefinitionsDelete{
		info: info,
		data: *payload,
	}, nil
}

type MetafieldDefinitionsDelete struct {
	info webhook.Info
	data MetafieldDefinitionsDeletePayload
}

func (webhook *MetafieldDefinitionsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetafieldDefinitionsDelete) GetData() (MetafieldDefinitionsDeletePayload, error) {
	return webhook.data, nil
}

type MetafieldDefinitionsDeletePayload struct {
	CreatedByAppID interface{} `json:"created_by_app_id"`
	ID             string      `json:"id"`
	Type           string      `json:"type"`
}
