package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersPlacedOnHold(webhook webhook.ValidWebhook) (*FulfillmentOrdersPlacedOnHold, error) {
	payload := &FulfillmentOrdersPlacedOnHoldPayload{}
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
	return &FulfillmentOrdersPlacedOnHold{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersPlacedOnHold struct {
	info webhook.Info
	data FulfillmentOrdersPlacedOnHoldPayload
}

func (webhook *FulfillmentOrdersPlacedOnHold) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersPlacedOnHold) GetData() (FulfillmentOrdersPlacedOnHoldPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersPlacedOnHoldPayload struct {
	CreatedFulfillmentHold struct {
		Handle    string `json:"handle"`
		HeldByApp struct {
			ID shopify.ID `json:"id"`
		} `json:"held_by_app"`
		HeldByRequestingApp bool       `json:"held_by_requesting_app"`
		ID                  shopify.ID `json:"id"`
		Reason              string     `json:"reason"`
		ReasonNotes         string     `json:"reason_notes"`
	} `json:"created_fulfillment_hold"`
	FulfillmentOrder struct {
		FulfillmentHolds []struct {
			Handle    string `json:"handle"`
			HeldByApp struct {
				ID shopify.ID `json:"id"`
			} `json:"held_by_app"`
			HeldByRequestingApp bool       `json:"held_by_requesting_app"`
			ID                  shopify.ID `json:"id"`
			Reason              string     `json:"reason"`
			ReasonNotes         string     `json:"reason_notes"`
		} `json:"fulfillment_holds"`
		ID     shopify.ID `json:"id"`
		Status string     `json:"status"`
	} `json:"fulfillment_order"`
	HeldFulfillmentOrderLineItems []struct {
		ID       shopify.ID `json:"id"`
		Quantity int64      `json:"quantity"`
	} `json:"held_fulfillment_order_line_items"`
	RemainingFulfillmentOrder struct {
		ID     shopify.ID `json:"id"`
		Status string     `json:"status"`
	} `json:"remaining_fulfillment_order"`
}
