package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewThemesDelete(webhook webhook.ValidWebhook) (*ThemesDelete, error) {
	payload := &ThemesDeletePayload{}
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
	return &ThemesDelete{
		info: info,
		data: *payload,
	}, nil
}

type ThemesDelete struct {
	info webhook.Info
	data ThemesDeletePayload
}

func (webhook *ThemesDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ThemesDelete) GetData() (ThemesDeletePayload, error) {
	return webhook.data, nil
}

type ThemesDeletePayload struct {
	ID int64 `json:"id"`
}
