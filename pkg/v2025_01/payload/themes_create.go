package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewThemesCreate(webhook webhook.ValidWebhook) (*ThemesCreate, error) {
	payload := &ThemesCreatePayload{}
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
	return &ThemesCreate{
		info: info,
		data: *payload,
	}, nil
}

type ThemesCreate struct {
	info webhook.Info
	data ThemesCreatePayload
}

func (webhook *ThemesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ThemesCreate) GetData() (ThemesCreatePayload, error) {
	return webhook.data, nil
}

type ThemesCreatePayload struct {
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
	CreatedAt         time.Time `json:"created_at"`
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	Previewable       bool      `json:"previewable"`
	Processing        bool      `json:"processing"`
	Role              string    `json:"role"`
	ThemeStoreID      int64     `json:"theme_store_id"`
	UpdatedAt         time.Time `json:"updated_at"`
}
