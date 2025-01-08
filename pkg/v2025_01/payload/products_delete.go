package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProductsDelete(webhook webhook.ValidWebhook) (*ProductsDelete, error) {
	payload := &ProductsDeletePayload{}
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
	return &ProductsDelete{
		info: info,
		data: *payload,
	}, nil
}

type ProductsDelete struct {
	info webhook.Info
	data ProductsDeletePayload
}

func (webhook *ProductsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductsDelete) GetData() (ProductsDeletePayload, error) {
	return webhook.data, nil
}

type ProductsDeletePayload struct {
	ID int64 `json:"id"`
}
