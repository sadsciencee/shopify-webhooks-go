package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewMarketsUpdate(webhook webhook.ValidWebhook) (*MarketsUpdate, error) {
	payload := &MarketsUpdatePayload{}
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
	return &MarketsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type MarketsUpdate struct {
	info webhook.Info
	data MarketsUpdatePayload
}

func (webhook *MarketsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MarketsUpdate) GetData() (MarketsUpdatePayload, error) {
	return webhook.data, nil
}

type MarketsUpdatePayload struct {
	Enabled bool   `json:"enabled"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Regions []struct {
		CountryCode string `json:"country_code"`
	} `json:"regions"`
}
