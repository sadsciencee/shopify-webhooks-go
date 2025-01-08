package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewDeliveryPromiseSettingsUpdate(webhook webhook.ValidWebhook) (*DeliveryPromiseSettingsUpdate, error) {
	payload := &DeliveryPromiseSettingsUpdatePayload{}
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
	return &DeliveryPromiseSettingsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type DeliveryPromiseSettingsUpdate struct {
	info webhook.Info
	data DeliveryPromiseSettingsUpdatePayload
}

func (webhook *DeliveryPromiseSettingsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DeliveryPromiseSettingsUpdate) GetData() (DeliveryPromiseSettingsUpdatePayload, error) {
	return webhook.data, nil
}

type DeliveryPromiseSettingsUpdatePayload struct {
	DeliveryDatesEnabled bool   `json:"delivery_dates_enabled"`
	ProcessingTime       string `json:"processing_time"`
	ShopID               string `json:"shop_id"`
}
