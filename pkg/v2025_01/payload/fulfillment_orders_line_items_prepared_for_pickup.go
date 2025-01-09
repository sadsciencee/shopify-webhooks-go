package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersLineItemsPreparedForPickup(webhook webhook.ValidWebhook) (*FulfillmentOrdersLineItemsPreparedForPickup, error) {
	payload := &FulfillmentOrdersLineItemsPreparedForPickupPayload{}
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
	return &FulfillmentOrdersLineItemsPreparedForPickup{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersLineItemsPreparedForPickup struct {
	info webhook.Info
	data FulfillmentOrdersLineItemsPreparedForPickupPayload
}

func (webhook *FulfillmentOrdersLineItemsPreparedForPickup) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersLineItemsPreparedForPickup) GetData() (FulfillmentOrdersLineItemsPreparedForPickupPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersLineItemsPreparedForPickupPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
}
