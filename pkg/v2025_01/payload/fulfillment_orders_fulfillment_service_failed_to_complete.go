package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersFulfillmentServiceFailedToComplete(webhook webhook.ValidWebhook) (*FulfillmentOrdersFulfillmentServiceFailedToComplete, error) {
	payload := &FulfillmentOrdersFulfillmentServiceFailedToCompletePayload{}
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
	return &FulfillmentOrdersFulfillmentServiceFailedToComplete{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersFulfillmentServiceFailedToComplete struct {
	info webhook.Info
	data FulfillmentOrdersFulfillmentServiceFailedToCompletePayload
}

func (webhook *FulfillmentOrdersFulfillmentServiceFailedToComplete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersFulfillmentServiceFailedToComplete) GetData() (FulfillmentOrdersFulfillmentServiceFailedToCompletePayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersFulfillmentServiceFailedToCompletePayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
	Message          string                   `json:"message"`
}
