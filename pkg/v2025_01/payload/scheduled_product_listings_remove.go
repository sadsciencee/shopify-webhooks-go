package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewScheduledProductListingsRemove(webhook webhook.ValidWebhook) (*ScheduledProductListingsRemove, error) {
	payload := &ScheduledProductListingsRemovePayload{}
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
	return &ScheduledProductListingsRemove{
		info: info,
		data: *payload,
	}, nil
}

type ScheduledProductListingsRemove struct {
	info webhook.Info
	data ScheduledProductListingsRemovePayload
}

func (webhook *ScheduledProductListingsRemove) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ScheduledProductListingsRemove) GetData() (ScheduledProductListingsRemovePayload, error) {
	return webhook.data, nil
}

type ScheduledProductListingsRemovePayload struct {
	ScheduledProductListing struct {
		ProductID int64 `json:"product_id"`
	} `json:"scheduled_product_listing"`
}
