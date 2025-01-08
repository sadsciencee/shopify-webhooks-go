package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSegmentsDelete(webhook webhook.ValidWebhook) (*SegmentsDelete, error) {
	payload := &SegmentsDeletePayload{}
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
	return &SegmentsDelete{
		info: info,
		data: *payload,
	}, nil
}

type SegmentsDelete struct {
	info webhook.Info
	data SegmentsDeletePayload
}

func (webhook *SegmentsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SegmentsDelete) GetData() (SegmentsDeletePayload, error) {
	return webhook.data, nil
}

type SegmentsDeletePayload struct {
	ID int64 `json:"id"`
}
