package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewLocalesUpdate(webhook webhook.ValidWebhook) (*LocalesUpdate, error) {
	payload := &LocalesUpdatePayload{}
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
	return &LocalesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type LocalesUpdate struct {
	info webhook.Info
	data LocalesUpdatePayload
}

func (webhook *LocalesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *LocalesUpdate) GetData() (LocalesUpdatePayload, error) {
	return webhook.data, nil
}

type LocalesUpdatePayload struct {
	Locale    string `json:"locale"`
	Published bool   `json:"published"`
}
