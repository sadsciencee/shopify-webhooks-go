package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewThemesPublish(webhook webhook.ValidWebhook) (*ThemesPublish, error) {
	payload := &ThemesPublishPayload{}
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
	return &ThemesPublish{
		info: info,
		data: *payload,
	}, nil
}

type ThemesPublish struct {
	info webhook.Info
	data ThemesPublishPayload
}

func (webhook *ThemesPublish) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ThemesPublish) GetData() (ThemesPublishPayload, error) {
	return webhook.data, nil
}

type ThemesPublishPayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	CreatedAt         time.Time  `json:"created_at"`
	ID                int64      `json:"id"`
	Name              string     `json:"name"`
	Previewable       bool       `json:"previewable"`
	Processing        bool       `json:"processing"`
	Role              string     `json:"role"`
	ThemeStoreID      int64      `json:"theme_store_id"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
