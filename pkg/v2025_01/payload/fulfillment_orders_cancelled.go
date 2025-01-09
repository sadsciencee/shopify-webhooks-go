package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersCancelled(webhook webhook.ValidWebhook) (*FulfillmentOrdersCancelled, error) {
	payload := &FulfillmentOrdersCancelledPayload{}
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
	return &FulfillmentOrdersCancelled{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersCancelled struct {
	info webhook.Info
	data FulfillmentOrdersCancelledPayload
}

func (webhook *FulfillmentOrdersCancelled) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersCancelled) GetData() (FulfillmentOrdersCancelledPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersCancelledPayload struct {
	FulfillmentOrder            shopify.FulfillmentOrder `json:"fulfillment_order"`
	ReplacementFulfillmentOrder shopify.FulfillmentOrder `json:"replacement_fulfillment_order"`
}
