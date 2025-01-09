package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewOrdersEdited(webhook webhook.ValidWebhook) (*OrdersEdited, error) {
	payload := &OrdersEditedPayload{}
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
	return &OrdersEdited{
		info: info,
		data: *payload,
	}, nil
}

type OrdersEdited struct {
	info webhook.Info
	data OrdersEditedPayload
}

func (webhook *OrdersEdited) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrdersEdited) GetData() (OrdersEditedPayload, error) {
	return webhook.data, nil
}

type OrdersEditedPayload struct {
	OrderEdit struct {
		AppID       *int64    `json:"app_id"`
		CommittedAt time.Time `json:"committed_at"`
		CreatedAt   time.Time `json:"created_at"`
		Discounts   struct {
			LineItem struct {
				Additions []interface{} `json:"additions"`
				Removals  []interface{} `json:"removals"`
			} `json:"line_item"`
		} `json:"discounts"`
		ID        int64 `json:"id"`
		LineItems struct {
			Additions []struct {
				Delta int64 `json:"delta"`
				ID    int64 `json:"id"`
			} `json:"additions"`
			Removals []struct {
				Delta int64 `json:"delta"`
				ID    int64 `json:"id"`
			} `json:"removals"`
		} `json:"line_items"`
		NotifyCustomer bool  `json:"notify_customer"`
		OrderID        int64 `json:"order_id"`
		ShippingLines  struct {
			Additions []interface{} `json:"additions"`
			Removals  []interface{} `json:"removals"`
		} `json:"shipping_lines"`
		StaffNote string      `json:"staff_note"`
		UserID    interface{} `json:"user_id"`
	} `json:"order_edit"`
}
