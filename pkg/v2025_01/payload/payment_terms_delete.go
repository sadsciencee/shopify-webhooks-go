package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewPaymentTermsDelete(webhook webhook.ValidWebhook) (*PaymentTermsDelete, error) {
	payload := &PaymentTermsDeletePayload{}
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
	return &PaymentTermsDelete{
		info: info,
		data: *payload,
	}, nil
}

type PaymentTermsDelete struct {
	info webhook.Info
	data PaymentTermsDeletePayload
}

func (webhook *PaymentTermsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *PaymentTermsDelete) GetData() (PaymentTermsDeletePayload, error) {
	return webhook.data, nil
}

type PaymentTermsDeletePayload struct {
	ID int64 `json:"id"`
}
