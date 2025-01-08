package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewAppScopesUpdate(webhook webhook.ValidWebhook) (*AppScopesUpdate, error) {
	payload := &AppScopesUpdatePayload{}
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
	return &AppScopesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type AppScopesUpdate struct {
	info webhook.Info
	data AppScopesUpdatePayload
}

func (webhook *AppScopesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AppScopesUpdate) GetData() (AppScopesUpdatePayload, error) {
	return webhook.data, nil
}

type AppScopesUpdatePayload struct {
	Current   []string  `json:"current"`
	ID        int64     `json:"id"`
	Previous  []string  `json:"previous"`
	UpdatedAt time.Time `json:"updated_at"`
}
