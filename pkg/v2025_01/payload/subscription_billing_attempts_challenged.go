package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSubscriptionBillingAttemptsChallenged(webhook webhook.ValidWebhook) (*SubscriptionBillingAttemptsChallenged, error) {
	payload := &SubscriptionBillingAttemptsChallengedPayload{}
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
	return &SubscriptionBillingAttemptsChallenged{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingAttemptsChallenged struct {
	info webhook.Info
	data SubscriptionBillingAttemptsChallengedPayload
}

func (webhook *SubscriptionBillingAttemptsChallenged) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingAttemptsChallenged) GetData() (SubscriptionBillingAttemptsChallengedPayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingAttemptsChallengedPayload struct {
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
