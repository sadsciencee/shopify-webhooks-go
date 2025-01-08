package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewOrdersDelete(webhook webhook.ValidWebhook) (*OrdersDelete, error) {
	payload := &OrdersDeletePayload{}
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
	return &OrdersDelete{
		info: info,
		data: *payload,
	}, nil
}

type OrdersDelete struct {
	info webhook.Info
	data OrdersDeletePayload
}

func (webhook *OrdersDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrdersDelete) GetData() (OrdersDeletePayload, error) {
	return webhook.data, nil
}

type OrdersDeletePayload struct {
	ID int64 `json:"id"`
}
