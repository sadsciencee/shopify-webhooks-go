package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsRequest(webhook webhook.ValidWebhook) (*ReturnsRequest, error) {
	payload := &ReturnsRequestPayload{}
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
	return &ReturnsRequest{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsRequest struct {
	info webhook.Info
	data ReturnsRequestPayload
}

func (webhook *ReturnsRequest) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsRequest) GetData() (ReturnsRequestPayload, error) {
	return webhook.data, nil
}

type ReturnsRequestPayload struct {
	AdminGraphqlAPIID shopify.ID    `json:"admin_graphql_api_id"`
	ExchangeLineItems []interface{} `json:"exchange_line_items"`
	ID                int64         `json:"id"`
	Name              interface{}   `json:"name"`
	Order             struct {
		AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
		ID                int64      `json:"id"`
	} `json:"order"`
	ReturnLineItems        []interface{} `json:"return_line_items"`
	ReturnShippingFees     []interface{} `json:"return_shipping_fees"`
	Status                 string        `json:"status"`
	TotalExchangeLineItems int64         `json:"total_exchange_line_items"`
	TotalReturnLineItems   int64         `json:"total_return_line_items"`
}
