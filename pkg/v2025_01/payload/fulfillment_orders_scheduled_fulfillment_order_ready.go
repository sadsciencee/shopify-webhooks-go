package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersScheduledFulfillmentOrderReady(webhook webhook.ValidWebhook) (*FulfillmentOrdersScheduledFulfillmentOrderReady, error) {
	payload := &FulfillmentOrdersScheduledFulfillmentOrderReadyPayload{}
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
	return &FulfillmentOrdersScheduledFulfillmentOrderReady{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersScheduledFulfillmentOrderReady struct {
	info webhook.Info
	data FulfillmentOrdersScheduledFulfillmentOrderReadyPayload
}

func (webhook *FulfillmentOrdersScheduledFulfillmentOrderReady) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersScheduledFulfillmentOrderReady) GetData() (FulfillmentOrdersScheduledFulfillmentOrderReadyPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersScheduledFulfillmentOrderReadyPayload struct {
	FulfillmentOrder struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"fulfillment_order"`
}
