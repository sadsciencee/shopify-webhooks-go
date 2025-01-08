package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProductFeedsCreate(webhook webhook.ValidWebhook) (*ProductFeedsCreate, error) {
	payload := &ProductFeedsCreatePayload{}
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
	return &ProductFeedsCreate{
		info: info,
		data: *payload,
	}, nil
}

type ProductFeedsCreate struct {
	info webhook.Info
	data ProductFeedsCreatePayload
}

func (webhook *ProductFeedsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductFeedsCreate) GetData() (ProductFeedsCreatePayload, error) {
	return webhook.data, nil
}

type ProductFeedsCreatePayload struct {
	Country  string `json:"country"`
	ID       string `json:"id"`
	Language string `json:"language"`
	Status   string `json:"status"`
}
