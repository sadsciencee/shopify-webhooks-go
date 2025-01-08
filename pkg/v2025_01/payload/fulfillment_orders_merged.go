package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersMerged(webhook webhook.ValidWebhook) (*FulfillmentOrdersMerged, error) {
	payload := &FulfillmentOrdersMergedPayload{}
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
	return &FulfillmentOrdersMerged{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersMerged struct {
	info webhook.Info
	data FulfillmentOrdersMergedPayload
}

func (webhook *FulfillmentOrdersMerged) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersMerged) GetData() (FulfillmentOrdersMergedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersMergedPayload struct {
	FulfillmentOrderMerges struct {
		FulfillmentOrder struct {
			ID     string `json:"id"`
			Status string `json:"status"`
		} `json:"fulfillment_order"`
	} `json:"fulfillment_order_merges"`
	MergeIntents []struct {
		FulfillmentOrderID        int64 `json:"fulfillment_order_id"`
		FulfillmentOrderLineItems []struct {
			ID       int64 `json:"id"`
			Quantity int64 `json:"quantity"`
		} `json:"fulfillment_order_line_items"`
	} `json:"merge_intents"`
}
