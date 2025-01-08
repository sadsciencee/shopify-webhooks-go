package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewThemesUpdate(webhook webhook.ValidWebhook) (*ThemesUpdate, error) {
	payload := &ThemesUpdatePayload{}
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
	return &ThemesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ThemesUpdate struct {
	info webhook.Info
	data ThemesUpdatePayload
}

func (webhook *ThemesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ThemesUpdate) GetData() (ThemesUpdatePayload, error) {
	return webhook.data, nil
}

type ThemesUpdatePayload struct {
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
