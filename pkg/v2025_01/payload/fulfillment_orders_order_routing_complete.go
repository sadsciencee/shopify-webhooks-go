package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersOrderRoutingComplete(webhook webhook.ValidWebhook) (*FulfillmentOrdersOrderRoutingComplete, error) {
	payload := &FulfillmentOrdersOrderRoutingCompletePayload{}
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
	return &FulfillmentOrdersOrderRoutingComplete{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersOrderRoutingComplete struct {
	info webhook.Info
	data FulfillmentOrdersOrderRoutingCompletePayload
}

func (webhook *FulfillmentOrdersOrderRoutingComplete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersOrderRoutingComplete) GetData() (FulfillmentOrdersOrderRoutingCompletePayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersOrderRoutingCompletePayload struct {
	FulfillmentOrder struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"fulfillment_order"`
}
