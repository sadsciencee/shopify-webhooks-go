package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProductFeedsUpdate(webhook webhook.ValidWebhook) (*ProductFeedsUpdate, error) {
	payload := &ProductFeedsUpdatePayload{}
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
	return &ProductFeedsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ProductFeedsUpdate struct {
	info webhook.Info
	data ProductFeedsUpdatePayload
}

func (webhook *ProductFeedsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductFeedsUpdate) GetData() (ProductFeedsUpdatePayload, error) {
	return webhook.data, nil
}

type ProductFeedsUpdatePayload struct {
	Country  string `json:"country"`
	ID       string `json:"id"`
	Language string `json:"language"`
	Status   string `json:"status"`
}
