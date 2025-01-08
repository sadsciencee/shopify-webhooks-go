package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewAppPurchasesOneTimeUpdate(webhook webhook.ValidWebhook) (*AppPurchasesOneTimeUpdate, error) {
	payload := &AppPurchasesOneTimeUpdatePayload{}
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
	return &AppPurchasesOneTimeUpdate{
		info: info,
		data: *payload,
	}, nil
}

type AppPurchasesOneTimeUpdate struct {
	info webhook.Info
	data AppPurchasesOneTimeUpdatePayload
}

func (webhook *AppPurchasesOneTimeUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AppPurchasesOneTimeUpdate) GetData() (AppPurchasesOneTimeUpdatePayload, error) {
	return webhook.data, nil
}

type AppPurchasesOneTimeUpdatePayload struct {
	AppPurchaseOneTime struct {
		AdminGraphqlAPIID     string    `json:"admin_graphql_api_id"`
		AdminGraphqlAPIShopID string    `json:"admin_graphql_api_shop_id"`
		CreatedAt             time.Time `json:"created_at"`
		Name                  string    `json:"name"`
		Status                string    `json:"status"`
		UpdatedAt             time.Time `json:"updated_at"`
	} `json:"app_purchase_one_time"`
}
