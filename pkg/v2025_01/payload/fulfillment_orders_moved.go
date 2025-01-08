package payload

import (
	"encoding/json"
	"fmt"
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
	DestinationLocationID              string `json:"destination_location_id"`
	FulfillmentOrderLineItemsRequested []struct {
		ID       string `json:"id"`
		Quantity int64  `json:"quantity"`
	} `json:"fulfillment_order_line_items_requested"`
	MovedFulfillmentOrder    FulfillmentOrder `json:"moved_fulfillment_order"`
	OriginalFulfillmentOrder FulfillmentOrder `json:"original_fulfillment_order"`
	SourceLocation           struct {
		ID string `json:"id"`
	} `json:"source_location"`
}

type FulfillmentOrder struct {
	AssignedLocationID string `json:"assigned_location_id"`
	ID                 string `json:"id"`
	Status             string `json:"status"`
}
