package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionContractsCreate(webhook webhook.ValidWebhook) (*SubscriptionContractsCreate, error) {
	payload := &SubscriptionContractsCreatePayload{}
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
	return &SubscriptionContractsCreate{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionContractsCreate struct {
	info webhook.Info
	data SubscriptionContractsCreatePayload
}

func (webhook *SubscriptionContractsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionContractsCreate) GetData() (SubscriptionContractsCreatePayload, error) {
	return webhook.data, nil
}

type SubscriptionContractsCreatePayload struct {
	AdminGraphqlAPICustomerID    string `json:"admin_graphql_api_customer_id"`
	AdminGraphqlAPIID            string `json:"admin_graphql_api_id"`
	AdminGraphqlAPIOriginOrderID string `json:"admin_graphql_api_origin_order_id"`
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
