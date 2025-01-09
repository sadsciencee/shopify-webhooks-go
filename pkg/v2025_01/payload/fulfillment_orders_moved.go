package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersMoved(webhook webhook.ValidWebhook) (*FulfillmentOrdersMoved, error) {
	payload := &FulfillmentOrdersMovedPayload{}
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
	return &FulfillmentOrdersMoved{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersMoved struct {
	info webhook.Info
	data FulfillmentOrdersMovedPayload
}

func (webhook *FulfillmentOrdersMoved) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersMoved) GetData() (FulfillmentOrdersMovedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersMovedPayload struct {
	DestinationLocationID              shopify.ID `json:"destination_location_id"`
	FulfillmentOrderLineItemsRequested []struct {
		ID       shopify.ID `json:"id"`
		Quantity int64      `json:"quantity"`
	} `json:"fulfillment_order_line_items_requested"`
	MovedFulfillmentOrder    shopify.FulfillmentOrder `json:"moved_fulfillment_order"`
	OriginalFulfillmentOrder shopify.FulfillmentOrder `json:"original_fulfillment_order"`
	SourceLocation           struct {
		ID shopify.ID `json:"id"`
	} `json:"source_location"`
}
