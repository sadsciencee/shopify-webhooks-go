package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionContractsCancel(webhook webhook.ValidWebhook) (*SubscriptionContractsCancel, error) {
	payload := &SubscriptionContractsCancelPayload{}
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
	return &SubscriptionContractsCancel{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionContractsCancel struct {
	info webhook.Info
	data SubscriptionContractsCancelPayload
}

func (webhook *SubscriptionContractsCancel) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionContractsCancel) GetData() (SubscriptionContractsCancelPayload, error) {
	return webhook.data, nil
}

type SubscriptionContractsCancelPayload struct {
	AdminGraphqlAPICustomerID    string     `json:"admin_graphql_api_customer_id"`
	AdminGraphqlAPIID            shopify.ID `json:"admin_graphql_api_id"`
	AdminGraphqlAPIOriginOrderID string     `json:"admin_graphql_api_origin_order_id"`
	BillingPolicy                struct {
		Interval      string `json:"interval"`
		IntervalCount int64  `json:"interval_count"`
		MaxCycles     int64  `json:"max_cycles"`
		MinCycles     int64  `json:"min_cycles"`
	} `json:"billing_policy"`
	CurrencyCode   string `json:"currency_code"`
	CustomerID     int64  `json:"customer_id"`
	DeliveryPolicy struct {
		Interval      string `json:"interval"`
		IntervalCount int64  `json:"interval_count"`
	} `json:"delivery_policy"`
	ID            int64  `json:"id"`
	OriginOrderID int64  `json:"origin_order_id"`
	RevisionID    string `json:"revision_id"`
	Status        string `json:"status"`
}
