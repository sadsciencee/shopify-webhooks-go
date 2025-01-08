package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomersPurchasingSummary(webhook webhook.ValidWebhook) (*CustomersPurchasingSummary, error) {
	payload := &CustomersPurchasingSummaryPayload{}
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
	return &CustomersPurchasingSummary{
		info: info,
		data: *payload,
	}, nil
}

type CustomersPurchasingSummary struct {
	info webhook.Info
	data CustomersPurchasingSummaryPayload
}

func (webhook *CustomersPurchasingSummary) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersPurchasingSummary) GetData() (CustomersPurchasingSummaryPayload, error) {
	return webhook.data, nil
}

type CustomersPurchasingSummaryPayload struct {
	AmountSpent struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"amountSpent"`
	CustomerID     string    `json:"customerId"`
	LastOrderID    string    `json:"lastOrderId"`
	NumberOfOrders int64     `json:"numberOfOrders"`
	OccurredAt     time.Time `json:"occurredAt"`
}
