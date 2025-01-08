package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionBillingAttemptsFailure(webhook webhook.ValidWebhook) (*SubscriptionBillingAttemptsFailure, error) {
	payload := &SubscriptionBillingAttemptsFailurePayload{}
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
	return &SubscriptionBillingAttemptsFailure{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingAttemptsFailure struct {
	info webhook.Info
	data SubscriptionBillingAttemptsFailurePayload
}

func (webhook *SubscriptionBillingAttemptsFailure) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingAttemptsFailure) GetData() (SubscriptionBillingAttemptsFailurePayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingAttemptsFailurePayload struct {
	AdminGraphqlAPIID                     *shopify.ID `json:"admin_graphql_api_id"`
	AdminGraphqlAPIOrderID                shopify.ID  `json:"admin_graphql_api_order_id"`
	AdminGraphqlAPISubscriptionContractID shopify.ID  `json:"admin_graphql_api_subscription_contract_id"`
	ErrorCode                             interface{} `json:"error_code"`
	ErrorMessage                          interface{} `json:"error_message"`
	ID                                    string      `json:"id"`
	IdempotencyKey                        string      `json:"idempotency_key"`
	OrderID                               int64       `json:"order_id"`
	Ready                                 bool        `json:"ready"`
	SubscriptionContractID                int64       `json:"subscription_contract_id"`
}
