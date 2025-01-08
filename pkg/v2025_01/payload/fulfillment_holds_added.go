package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentHoldsAdded(webhook webhook.ValidWebhook) (*FulfillmentHoldsAdded, error) {
	payload := &FulfillmentHoldsAddedPayload{}
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
	return &FulfillmentHoldsAdded{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentHoldsAdded struct {
	info webhook.Info
	data FulfillmentHoldsAddedPayload
}

func (webhook *FulfillmentHoldsAdded) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentHoldsAdded) GetData() (FulfillmentHoldsAddedPayload, error) {
	return webhook.data, nil
}

type FulfillmentHoldsAddedPayload struct {
	FulfillmentHold struct {
		Handle    string `json:"handle"`
		HeldByApp struct {
			ID string `json:"id"`
		} `json:"held_by_app"`
		HeldByRequestingApp bool   `json:"held_by_requesting_app"`
		ID                  string `json:"id"`
		Reason              string `json:"reason"`
		ReasonNotes         string `json:"reason_notes"`
	} `json:"fulfillment_hold"`
	FulfillmentOrder struct {
		ID string `json:"id"`
	} `json:"fulfillment_order"`
}
