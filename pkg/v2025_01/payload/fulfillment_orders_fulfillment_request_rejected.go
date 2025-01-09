package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersFulfillmentRequestRejected(webhook webhook.ValidWebhook) (*FulfillmentOrdersFulfillmentRequestRejected, error) {
	payload := &FulfillmentOrdersFulfillmentRequestRejectedPayload{}
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
	return &FulfillmentOrdersFulfillmentRequestRejected{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersFulfillmentRequestRejected struct {
	info webhook.Info
	data FulfillmentOrdersFulfillmentRequestRejectedPayload
}

func (webhook *FulfillmentOrdersFulfillmentRequestRejected) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersFulfillmentRequestRejected) GetData() (FulfillmentOrdersFulfillmentRequestRejectedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersFulfillmentRequestRejectedPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
	Message          string                   `json:"message"`
}
