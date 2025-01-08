package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewLocalesCreate(webhook webhook.ValidWebhook) (*LocalesCreate, error) {
	payload := &LocalesCreatePayload{}
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
	return &LocalesCreate{
		info: info,
		data: *payload,
	}, nil
}

type LocalesCreate struct {
	info webhook.Info
	data LocalesCreatePayload
}

func (webhook *LocalesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *LocalesCreate) GetData() (LocalesCreatePayload, error) {
	return webhook.data, nil
}

type LocalesCreatePayload struct {
	Locale    string `json:"locale"`
	Published bool   `json:"published"`
}
