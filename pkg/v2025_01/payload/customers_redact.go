package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersRedact(webhook webhook.ValidWebhook) (*CustomersRedact, error) {
	payload := &CustomersRedactPayload{}
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
	return &CustomersRedact{
		info: info,
		data: *payload,
	}, nil
}

type CustomersRedact struct {
	info webhook.Info
	data CustomersRedactPayload
}

func (webhook *CustomersRedact) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersRedact) GetData() (CustomersRedactPayload, error) {
	return webhook.data, nil
}

type CustomersRedactPayload struct {
	Customer struct {
		Email string `json:"email"`
		ID    int64  `json:"id"`
		Phone string `json:"phone"`
	} `json:"customer"`
	OrdersToRedact []int64 `json:"orders_to_redact"`
	ShopDomain     string  `json:"shop_domain"`
	ShopID         int64   `json:"shop_id"`
}
