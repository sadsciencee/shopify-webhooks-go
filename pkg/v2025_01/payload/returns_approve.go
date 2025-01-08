package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsApprove(webhook webhook.ValidWebhook) (*ReturnsApprove, error) {
	payload := &ReturnsApprovePayload{}
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
	return &ReturnsApprove{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsApprove struct {
	info webhook.Info
	data ReturnsApprovePayload
}

func (webhook *ReturnsApprove) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsApprove) GetData() (ReturnsApprovePayload, error) {
	return webhook.data, nil
}

type ReturnsApprovePayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	ExchangeLineItems []struct {
		AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
		ID                int64       `json:"id"`
		LineItem          interface{} `json:"line_item"`
	} `json:"exchange_line_items"`
	ID    int64       `json:"id"`
	Name  interface{} `json:"name"`
	Order struct {
		AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
		ID                int64  `json:"id"`
	} `json:"order"`
	ReturnLineItems    []interface{} `json:"return_line_items"`
	ReturnShippingFees []struct {
		AdminGraphqlAPIID string           `json:"admin_graphql_api_id"`
		ID                int64            `json:"id"`
		Price             shopify.MoneyBag `json:"price"`
	} `json:"return_shipping_fees"`
	Status                 string `json:"status"`
	TotalExchangeLineItems int64  `json:"total_exchange_line_items"`
	TotalReturnLineItems   int64  `json:"total_return_line_items"`
}
