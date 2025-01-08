package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersFulfillmentRequestAccepted(webhook webhook.ValidWebhook) (*FulfillmentOrdersFulfillmentRequestAccepted, error) {
	payload := &FulfillmentOrdersFulfillmentRequestAcceptedPayload{}
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
	return &FulfillmentOrdersFulfillmentRequestAccepted{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersFulfillmentRequestAccepted struct {
	info webhook.Info
	data FulfillmentOrdersFulfillmentRequestAcceptedPayload
}

func (webhook *FulfillmentOrdersFulfillmentRequestAccepted) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersFulfillmentRequestAccepted) GetData() (FulfillmentOrdersFulfillmentRequestAcceptedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersFulfillmentRequestAcceptedPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
	Message          string                   `json:"message"`
}
