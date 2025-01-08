package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewMarketsCreate(webhook webhook.ValidWebhook) (*MarketsCreate, error) {
	payload := &MarketsCreatePayload{}
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
	return &MarketsCreate{
		info: info,
		data: *payload,
	}, nil
}

type MarketsCreate struct {
	info webhook.Info
	data MarketsCreatePayload
}

func (webhook *MarketsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MarketsCreate) GetData() (MarketsCreatePayload, error) {
	return webhook.data, nil
}

type MarketsCreatePayload struct {
	Enabled bool   `json:"enabled"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Regions []struct {
		CountryCode string `json:"country_code"`
	} `json:"regions"`
}
