package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDiscountsUpdate(webhook webhook.ValidWebhook) (*DiscountsUpdate, error) {
	payload := &DiscountsUpdatePayload{}
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
	return &DiscountsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type DiscountsUpdate struct {
	info webhook.Info
	data DiscountsUpdatePayload
}

func (webhook *DiscountsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DiscountsUpdate) GetData() (DiscountsUpdatePayload, error) {
	return webhook.data, nil
}

type DiscountsUpdatePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	CreatedAt         time.Time  `json:"created_at"`
	Status            string     `json:"status"`
	Title             string     `json:"title"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
