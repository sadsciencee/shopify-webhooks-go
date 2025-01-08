package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewAppSubscriptionsUpdate(webhook webhook.ValidWebhook) (*AppSubscriptionsUpdate, error) {
	payload := &AppSubscriptionsUpdatePayload{}
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
	return &AppSubscriptionsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type AppSubscriptionsUpdate struct {
	info webhook.Info
	data AppSubscriptionsUpdatePayload
}

func (webhook *AppSubscriptionsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AppSubscriptionsUpdate) GetData() (AppSubscriptionsUpdatePayload, error) {
	return webhook.data, nil
}

type AppSubscriptionsUpdatePayload struct {
	AppSubscription struct {
		AdminGraphqlAPIID     string    `json:"admin_graphql_api_id"`
		AdminGraphqlAPIShopID string    `json:"admin_graphql_api_shop_id"`
		CappedAmount          string    `json:"capped_amount"`
		CreatedAt             time.Time `json:"created_at"`
		Currency              string    `json:"currency"`
		Name                  string    `json:"name"`
		Status                string    `json:"status"`
		UpdatedAt             time.Time `json:"updated_at"`
	} `json:"app_subscription"`
}
