package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewSegmentsCreate(webhook webhook.ValidWebhook) (*SegmentsCreate, error) {
	payload := &SegmentsCreatePayload{}
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
	return &SegmentsCreate{
		info: info,
		data: *payload,
	}, nil
}

type SegmentsCreate struct {
	info webhook.Info
	data SegmentsCreatePayload
}

func (webhook *SegmentsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SegmentsCreate) GetData() (SegmentsCreatePayload, error) {
	return webhook.data, nil
}

type SegmentsCreatePayload struct {
	CreationDate time.Time `json:"creationDate"`
	ID           int64     `json:"id"`
	LastEditDate time.Time `json:"lastEditDate"`
	Name         string    `json:"name"`
	Query        string    `json:"query"`
}
