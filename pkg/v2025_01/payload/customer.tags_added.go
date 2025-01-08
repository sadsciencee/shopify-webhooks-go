package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomerTagsAdded(webhook webhook.ValidWebhook) (*CustomerTagsAdded, error) {
	payload := &CustomerTagsAddedPayload{}
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
	return &CustomerTagsAdded{
		info: info,
		data: *payload,
	}, nil
}

type CustomerTagsAdded struct {
	info webhook.Info
	data CustomerTagsAddedPayload
}

func (webhook *CustomerTagsAdded) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerTagsAdded) GetData() (CustomerTagsAddedPayload, error) {
	return webhook.data, nil
}

type CustomerTagsAddedPayload struct {
	CustomerID string    `json:"customerId"`
	OccurredAt time.Time `json:"occurredAt"`
	Tags       []string  `json:"tags"`
}
