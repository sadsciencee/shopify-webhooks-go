package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewTenderTransactionsCreate(webhook webhook.ValidWebhook) (*TenderTransactionsCreate, error) {
	payload := &TenderTransactionsCreatePayload{}
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
	return &TenderTransactionsCreate{
		info: info,
		data: *payload,
	}, nil
}

type TenderTransactionsCreate struct {
	info webhook.Info
	data TenderTransactionsCreatePayload
}

func (webhook *TenderTransactionsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *TenderTransactionsCreate) GetData() (TenderTransactionsCreatePayload, error) {
	return webhook.data, nil
}

type TenderTransactionsCreatePayload struct {
	Amount          shopify.Decimal `json:"amount"`
	Currency        string          `json:"currency"`
	ID              int64           `json:"id"`
	OrderID         int64           `json:"order_id"`
	PaymentDetails  interface{}     `json:"payment_details"`
	PaymentMethod   string          `json:"payment_method"`
	ProcessedAt     *time.Time      `json:"processed_at"`
	RemoteReference string          `json:"remote_reference"`
	Test            bool            `json:"test"`
	UserID          interface{}     `json:"user_id"`
}
