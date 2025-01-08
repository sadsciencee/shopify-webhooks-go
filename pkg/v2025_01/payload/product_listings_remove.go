package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProductListingsRemove(webhook webhook.ValidWebhook) (*ProductListingsRemove, error) {
	payload := &ProductListingsRemovePayload{}
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
	return &ProductListingsRemove{
		info: info,
		data: *payload,
	}, nil
}

type ProductListingsRemove struct {
	info webhook.Info
	data ProductListingsRemovePayload
}

func (webhook *ProductListingsRemove) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductListingsRemove) GetData() (ProductListingsRemovePayload, error) {
	return webhook.data, nil
}

type ProductListingsRemovePayload struct {
	ProductListing struct {
		ProductID int64 `json:"product_id"`
	} `json:"product_listing"`
}
