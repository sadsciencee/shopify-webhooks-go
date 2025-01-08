package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsCancel(webhook webhook.ValidWebhook) (*ReturnsCancel, error) {
	payload := &ReturnsCancelPayload{}
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
	return &ReturnsCancel{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsCancel struct {
	info webhook.Info
	data ReturnsCancelPayload
}

func (webhook *ReturnsCancel) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsCancel) GetData() (ReturnsCancelPayload, error) {
	return webhook.data, nil
}

type ReturnsCancelPayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	ID                int64  `json:"id"`
	OrderID           int64  `json:"order_id"`
	Status            string `json:"status"`
}
