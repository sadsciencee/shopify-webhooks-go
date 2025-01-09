package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionContractsPause(webhook webhook.ValidWebhook) (*SubscriptionContractsPause, error) {
	payload := &SubscriptionContractsPausePayload{}
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
	return &SubscriptionContractsPause{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionContractsPause struct {
	info webhook.Info
	data SubscriptionContractsPausePayload
}

func (webhook *SubscriptionContractsPause) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionContractsPause) GetData() (SubscriptionContractsPausePayload, error) {
	return webhook.data, nil
}

type SubscriptionContractsPausePayload struct {
	AdminGraphqlAPICustomerID    string                            `json:"admin_graphql_api_customer_id"`
	AdminGraphqlAPIID            shopify.ID                        `json:"admin_graphql_api_id"`
	AdminGraphqlAPIOriginOrderID string                            `json:"admin_graphql_api_origin_order_id"`
	BillingPolicy                shopify.SellingPlanBillingPolicy  `json:"billing_policy"`
	CurrencyCode                 string                            `json:"currency_code"`
	CustomerID                   int64                             `json:"customer_id"`
	DeliveryPolicy               shopify.SellingPlanDeliveryPolicy `json:"delivery_policy"`
	ID                           int64                             `json:"id"`
	OriginOrderID                int64                             `json:"origin_order_id"`
	RevisionID                   string                            `json:"revision_id"`
	Status                       string                            `json:"status"`
}
