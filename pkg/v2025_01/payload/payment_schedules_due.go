package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewPaymentSchedulesDue(webhook webhook.ValidWebhook) (*PaymentSchedulesDue, error) {
	payload := &PaymentSchedulesDuePayload{}
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
	return &PaymentSchedulesDue{
		info: info,
		data: *payload,
	}, nil
}

type PaymentSchedulesDue struct {
	info webhook.Info
	data PaymentSchedulesDuePayload
}

func (webhook *PaymentSchedulesDue) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *PaymentSchedulesDue) GetData() (PaymentSchedulesDuePayload, error) {
	return webhook.data, nil
}

type PaymentSchedulesDuePayload struct {
	AdminGraphqlAPIID   shopify.ID `json:"admin_graphql_api_id"`
	Amount              string     `json:"amount"`
	BalanceDue          string     `json:"balance_due"`
	CompletedAt         time.Time  `json:"completed_at"`
	CreatedAt           time.Time  `json:"created_at"`
	Currency            string     `json:"currency"`
	DueAt               time.Time  `json:"due_at"`
	ID                  int64      `json:"id"`
	IssuedAt            time.Time  `json:"issued_at"`
	PaymentTermsID      int64      `json:"payment_terms_id"`
	PresentmentCurrency string     `json:"presentment_currency"`
	TotalBalance        string     `json:"total_balance"`
	TotalPrice          string     `json:"total_price"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
