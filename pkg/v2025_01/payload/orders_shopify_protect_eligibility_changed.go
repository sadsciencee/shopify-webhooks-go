package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewOrdersShopifyProtectEligibilityChanged(webhook webhook.ValidWebhook) (*OrdersShopifyProtectEligibilityChanged, error) {
	payload := &OrdersShopifyProtectEligibilityChangedPayload{}
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
	return &OrdersShopifyProtectEligibilityChanged{
		info: info,
		data: *payload,
	}, nil
}

type OrdersShopifyProtectEligibilityChanged struct {
	info webhook.Info
	data OrdersShopifyProtectEligibilityChangedPayload
}

func (webhook *OrdersShopifyProtectEligibilityChanged) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrdersShopifyProtectEligibilityChanged) GetData() (OrdersShopifyProtectEligibilityChangedPayload, error) {
	return webhook.data, nil
}

type OrdersShopifyProtectEligibilityChangedPayload struct {
	Eligibility struct {
		Status string `json:"status"`
	} `json:"eligibility"`
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}
