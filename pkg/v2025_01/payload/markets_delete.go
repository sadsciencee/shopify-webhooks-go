package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewMarketsDelete(webhook webhook.ValidWebhook) (*MarketsDelete, error) {
	payload := &MarketsDeletePayload{}
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
	return &MarketsDelete{
		info: info,
		data: *payload,
	}, nil
}

type MarketsDelete struct {
	info webhook.Info
	data MarketsDeletePayload
}

func (webhook *MarketsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MarketsDelete) GetData() (MarketsDeletePayload, error) {
	return webhook.data, nil
}

type MarketsDeletePayload struct {
	ID int64 `json:"id"`
}
