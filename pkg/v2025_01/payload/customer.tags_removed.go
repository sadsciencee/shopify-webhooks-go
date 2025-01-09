package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomerTagsRemoved(webhook webhook.ValidWebhook) (*CustomerTagsRemoved, error) {
	payload := &CustomerTagsRemovedPayload{}
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
	return &CustomerTagsRemoved{
		info: info,
		data: *payload,
	}, nil
}

type CustomerTagsRemoved struct {
	info webhook.Info
	data CustomerTagsRemovedPayload
}

func (webhook *CustomerTagsRemoved) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerTagsRemoved) GetData() (CustomerTagsRemovedPayload, error) {
	return webhook.data, nil
}

type CustomerTagsRemovedPayload struct {
	CustomerID shopify.ID `json:"customerId"`
	OccurredAt time.Time  `json:"occurredAt"`
	Tags       []string   `json:"tags"`
}
