package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewSegmentsUpdate(webhook webhook.ValidWebhook) (*SegmentsUpdate, error) {
	payload := &SegmentsUpdatePayload{}
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
	return &SegmentsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type SegmentsUpdate struct {
	info webhook.Info
	data SegmentsUpdatePayload
}

func (webhook *SegmentsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SegmentsUpdate) GetData() (SegmentsUpdatePayload, error) {
	return webhook.data, nil
}

type SegmentsUpdatePayload struct {
	CreationDate time.Time `json:"creationDate"`
	ID           int64     `json:"id"`
	LastEditDate time.Time `json:"lastEditDate"`
	Name         string    `json:"name"`
	Query        string    `json:"query"`
}
