package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewPaymentTermsCreate(webhook webhook.ValidWebhook) (*PaymentTermsCreate, error) {
	payload := &PaymentTermsCreatePayload{}
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
	return &PaymentTermsCreate{
		info: info,
		data: *payload,
	}, nil
}

type PaymentTermsCreate struct {
	info webhook.Info
	data PaymentTermsCreatePayload
}

func (webhook *PaymentTermsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *PaymentTermsCreate) GetData() (PaymentTermsCreatePayload, error) {
	return webhook.data, nil
}

type PaymentTermsCreatePayload struct {
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
	CreatedAt         time.Time `json:"created_at"`
	DueInDays         int64     `json:"due_in_days"`
	ID                int64     `json:"id"`
	PaymentSchedules  []struct {
		AdminGraphqlAPIID   string    `json:"admin_graphql_api_id"`
		Amount              string    `json:"amount"`
		BalanceDue          string    `json:"balance_due"`
		CompletedAt         time.Time `json:"completed_at"`
		CreatedAt           time.Time `json:"created_at"`
		Currency            string    `json:"currency"`
		DueAt               time.Time `json:"due_at"`
		ID                  int64     `json:"id"`
		IssuedAt            time.Time `json:"issued_at"`
		PaymentTermsID      int64     `json:"payment_terms_id"`
		PresentmentCurrency string    `json:"presentment_currency"`
		TotalBalance        string    `json:"total_balance"`
		TotalPrice          string    `json:"total_price"`
		UpdatedAt           time.Time `json:"updated_at"`
	} `json:"payment_schedules"`
	PaymentTermsName string    `json:"payment_terms_name"`
	PaymentTermsType string    `json:"payment_terms_type"`
	UpdatedAt        time.Time `json:"updated_at"`
}
