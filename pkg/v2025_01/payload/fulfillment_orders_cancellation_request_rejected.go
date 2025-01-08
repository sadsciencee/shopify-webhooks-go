package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersCancellationRequestRejected(webhook webhook.ValidWebhook) (*FulfillmentOrdersCancellationRequestRejected, error) {
	payload := &FulfillmentOrdersCancellationRequestRejectedPayload{}
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
	return &FulfillmentOrdersCancellationRequestRejected{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersCancellationRequestRejected struct {
	info webhook.Info
	data FulfillmentOrdersCancellationRequestRejectedPayload
}

func (webhook *FulfillmentOrdersCancellationRequestRejected) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersCancellationRequestRejected) GetData() (FulfillmentOrdersCancellationRequestRejectedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersCancellationRequestRejectedPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
	Message          string                   `json:"message"`
}
