package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewFulfillmentOrdersRescheduled(webhook webhook.ValidWebhook) (*FulfillmentOrdersRescheduled, error) {
	payload := &FulfillmentOrdersRescheduledPayload{}
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
	return &FulfillmentOrdersRescheduled{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersRescheduled struct {
	info webhook.Info
	data FulfillmentOrdersRescheduledPayload
}

func (webhook *FulfillmentOrdersRescheduled) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersRescheduled) GetData() (FulfillmentOrdersRescheduledPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersRescheduledPayload struct {
	FulfillmentOrder struct {
		FulfillAt time.Time  `json:"fulfill_at"`
		ID        shopify.ID `json:"id"`
		Status    string     `json:"status"`
	} `json:"fulfillment_order"`
}
