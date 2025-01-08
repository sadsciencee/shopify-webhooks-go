package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCheckoutsDelete(webhook webhook.ValidWebhook) (*CheckoutsDelete, error) {
	payload := &CheckoutsDeletePayload{}
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
	return &CheckoutsDelete{
		info: info,
		data: *payload,
	}, nil
}

type CheckoutsDelete struct {
	info webhook.Info
	data CheckoutsDeletePayload
}

func (webhook *CheckoutsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CheckoutsDelete) GetData() (CheckoutsDeletePayload, error) {
	return webhook.data, nil
}

type CheckoutsDeletePayload struct {
	BuyerAcceptsSmsMarketing bool        `json:"buyer_accepts_sms_marketing"`
	CartToken                string      `json:"cart_token"`
	ID                       int64       `json:"id"`
	PresentmentCurrency      string      `json:"presentment_currency"`
	ReservationToken         interface{} `json:"reservation_token"`
	SmsMarketingPhone        interface{} `json:"sms_marketing_phone"`
	SubtotalPrice            string      `json:"subtotal_price"`
	TotalDiscounts           string      `json:"total_discounts"`
	TotalDuties              interface{} `json:"total_duties"`
	TotalLineItemsPrice      string      `json:"total_line_items_price"`
	TotalPrice               string      `json:"total_price"`
	TotalTax                 string      `json:"total_tax"`
}
