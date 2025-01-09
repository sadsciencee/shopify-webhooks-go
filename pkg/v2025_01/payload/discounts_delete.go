package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDiscountsDelete(webhook webhook.ValidWebhook) (*DiscountsDelete, error) {
	payload := &DiscountsDeletePayload{}
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
	return &DiscountsDelete{
		info: info,
		data: *payload,
	}, nil
}

type DiscountsDelete struct {
	info webhook.Info
	data DiscountsDeletePayload
}

func (webhook *DiscountsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DiscountsDelete) GetData() (DiscountsDeletePayload, error) {
	return webhook.data, nil
}

type DiscountsDeletePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	DeletedAt         time.Time  `json:"deleted_at"`
}
