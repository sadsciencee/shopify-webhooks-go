package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewAppSubscriptionsApproachingCappedAmount(webhook webhook.ValidWebhook) (*AppSubscriptionsApproachingCappedAmount, error) {
	payload := &AppSubscriptionsApproachingCappedAmountPayload{}
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
	return &AppSubscriptionsApproachingCappedAmount{
		info: info,
		data: *payload,
	}, nil
}

type AppSubscriptionsApproachingCappedAmount struct {
	info webhook.Info
	data AppSubscriptionsApproachingCappedAmountPayload
}

func (webhook *AppSubscriptionsApproachingCappedAmount) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AppSubscriptionsApproachingCappedAmount) GetData() (AppSubscriptionsApproachingCappedAmountPayload, error) {
	return webhook.data, nil
}

type AppSubscriptionsApproachingCappedAmountPayload struct {
	AppSubscription struct {
		AdminGraphqlAPIID     string    `json:"admin_graphql_api_id"`
		AdminGraphqlAPIShopID string    `json:"admin_graphql_api_shop_id"`
		BalanceUsed           int64     `json:"balance_used"`
		CappedAmount          string    `json:"capped_amount"`
		CreatedAt             time.Time `json:"created_at"`
		CurrencyCode          string    `json:"currency_code"`
		Name                  string    `json:"name"`
		UpdatedAt             time.Time `json:"updated_at"`
	} `json:"app_subscription"`
}
