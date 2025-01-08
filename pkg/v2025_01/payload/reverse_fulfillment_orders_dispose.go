package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReverseFulfillmentOrdersDispose(webhook webhook.ValidWebhook) (*ReverseFulfillmentOrdersDispose, error) {
	payload := &ReverseFulfillmentOrdersDisposePayload{}
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
	return &ReverseFulfillmentOrdersDispose{
		info: info,
		data: *payload,
	}, nil
}

type ReverseFulfillmentOrdersDispose struct {
	info webhook.Info
	data ReverseFulfillmentOrdersDisposePayload
}

func (webhook *ReverseFulfillmentOrdersDispose) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReverseFulfillmentOrdersDispose) GetData() (ReverseFulfillmentOrdersDisposePayload, error) {
	return webhook.data, nil
}

type ReverseFulfillmentOrdersDisposePayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	Dispositions      []struct {
		Location struct {
			AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
			ID                int64  `json:"id"`
		} `json:"location"`
		Quantity                        int64       `json:"quantity"`
		ReverseDeliveryLineItem         interface{} `json:"reverse_delivery_line_item"`
		ReverseFulfillmentOrderLineItem struct {
			AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
			ID                int64  `json:"id"`
		} `json:"reverse_fulfillment_order_line_item"`
		Type string `json:"type"`
	} `json:"dispositions"`
	ID                int64 `json:"id"`
	TotalDispositions int64 `json:"total_dispositions"`
}
