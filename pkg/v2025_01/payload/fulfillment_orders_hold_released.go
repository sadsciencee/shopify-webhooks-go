package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersHoldReleased(webhook webhook.ValidWebhook) (*FulfillmentOrdersHoldReleased, error) {
	payload := &FulfillmentOrdersHoldReleasedPayload{}
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
	return &FulfillmentOrdersHoldReleased{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersHoldReleased struct {
	info webhook.Info
	data FulfillmentOrdersHoldReleasedPayload
}

func (webhook *FulfillmentOrdersHoldReleased) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersHoldReleased) GetData() (FulfillmentOrdersHoldReleasedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersHoldReleasedPayload struct {
	FulfillmentOrder shopify.FulfillmentOrder `json:"fulfillment_order"`
}
