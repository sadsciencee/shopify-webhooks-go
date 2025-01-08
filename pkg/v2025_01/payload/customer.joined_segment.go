package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomerJoinedSegment(webhook webhook.ValidWebhook) (*CustomerJoinedSegment, error) {
	payload := &CustomerJoinedSegmentPayload{}
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
	return &CustomerJoinedSegment{
		info: info,
		data: *payload,
	}, nil
}

type CustomerJoinedSegment struct {
	info webhook.Info
	data CustomerJoinedSegmentPayload
}

func (webhook *CustomerJoinedSegment) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerJoinedSegment) GetData() (CustomerJoinedSegmentPayload, error) {
	return webhook.data, nil
}

type CustomerJoinedSegmentPayload struct {
	CustomerID string `json:"customer_id"`
	SegmentID  string `json:"segment_id"`
	ShopID     string `json:"shop_id"`
}
