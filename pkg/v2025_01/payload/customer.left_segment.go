package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomerLeftSegment(webhook webhook.ValidWebhook) (*CustomerLeftSegment, error) {
	payload := &CustomerLeftSegmentPayload{}
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
	return &CustomerLeftSegment{
		info: info,
		data: *payload,
	}, nil
}

type CustomerLeftSegment struct {
	info webhook.Info
	data CustomerLeftSegmentPayload
}

func (webhook *CustomerLeftSegment) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerLeftSegment) GetData() (CustomerLeftSegmentPayload, error) {
	return webhook.data, nil
}

type CustomerLeftSegmentPayload struct {
	CustomerID string `json:"customer_id"`
	SegmentID  string `json:"segment_id"`
	ShopID     string `json:"shop_id"`
}
