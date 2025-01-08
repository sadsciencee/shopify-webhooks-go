package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersSplit(webhook webhook.ValidWebhook) (*FulfillmentOrdersSplit, error) {
	payload := &FulfillmentOrdersSplitPayload{}
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
	return &FulfillmentOrdersSplit{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersSplit struct {
	info webhook.Info
	data FulfillmentOrdersSplitPayload
}

func (webhook *FulfillmentOrdersSplit) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersSplit) GetData() (FulfillmentOrdersSplitPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersSplitPayload struct {
	FulfillmentOrder            shopify.FulfillmentOrder `json:"fulfillment_order"`
	RemainingFulfillmentOrder   shopify.FulfillmentOrder `json:"remaining_fulfillment_order"`
	ReplacementFulfillmentOrder shopify.FulfillmentOrder `json:"replacement_fulfillment_order"`
	SplitLineItems              []struct {
		ID       string `json:"id"`
		Quantity int64  `json:"quantity"`
	} `json:"split_line_items"`
}
