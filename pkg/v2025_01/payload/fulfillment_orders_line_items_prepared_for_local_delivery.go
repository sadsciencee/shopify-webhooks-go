package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersLineItemsPreparedForLocalDelivery(webhook webhook.ValidWebhook) (*FulfillmentOrdersLineItemsPreparedForLocalDelivery, error) {
	payload := &FulfillmentOrdersLineItemsPreparedForLocalDeliveryPayload{}
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
	return &FulfillmentOrdersLineItemsPreparedForLocalDelivery{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersLineItemsPreparedForLocalDelivery struct {
	info webhook.Info
	data FulfillmentOrdersLineItemsPreparedForLocalDeliveryPayload
}

func (webhook *FulfillmentOrdersLineItemsPreparedForLocalDelivery) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersLineItemsPreparedForLocalDelivery) GetData() (FulfillmentOrdersLineItemsPreparedForLocalDeliveryPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersLineItemsPreparedForLocalDeliveryPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
}
