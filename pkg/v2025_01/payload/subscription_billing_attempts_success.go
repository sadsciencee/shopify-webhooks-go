package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionBillingAttemptsSuccess(webhook webhook.ValidWebhook) (*SubscriptionBillingAttemptsSuccess, error) {
	payload := &SubscriptionBillingAttemptsSuccessPayload{}
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
	return &SubscriptionBillingAttemptsSuccess{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingAttemptsSuccess struct {
	info webhook.Info
	data SubscriptionBillingAttemptsSuccessPayload
}

func (webhook *SubscriptionBillingAttemptsSuccess) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingAttemptsSuccess) GetData() (SubscriptionBillingAttemptsSuccessPayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingAttemptsSuccessPayload struct {
	AdminGraphqlAPIID                     interface{} `json:"admin_graphql_api_id"`
	AdminGraphqlAPIOrderID                string      `json:"admin_graphql_api_order_id"`
	AdminGraphqlAPISubscriptionContractID string      `json:"admin_graphql_api_subscription_contract_id"`
	ErrorCode                             interface{} `json:"error_code"`
	ErrorMessage                          interface{} `json:"error_message"`
	ID                                    string      `json:"id"`
	IdempotencyKey                        string      `json:"idempotency_key"`
	OrderID                               int64       `json:"order_id"`
	Ready                                 bool        `json:"ready"`
	SubscriptionContractID                int64       `json:"subscription_contract_id"`
}
