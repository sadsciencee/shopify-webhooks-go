package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDiscountsCreate(webhook webhook.ValidWebhook) (*DiscountsCreate, error) {
	payload := &DiscountsCreatePayload{}
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
	return &DiscountsCreate{
		info: info,
		data: *payload,
	}, nil
}

type DiscountsCreate struct {
	info webhook.Info
	data DiscountsCreatePayload
}

func (webhook *DiscountsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DiscountsCreate) GetData() (DiscountsCreatePayload, error) {
	return webhook.data, nil
}

type DiscountsCreatePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	CreatedAt         time.Time  `json:"created_at"`
	Status            string     `json:"status"`
	Title             string     `json:"title"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
