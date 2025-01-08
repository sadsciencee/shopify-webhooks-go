package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsClose(webhook webhook.ValidWebhook) (*ReturnsClose, error) {
	payload := &ReturnsClosePayload{}
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
	return &ReturnsClose{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsClose struct {
	info webhook.Info
	data ReturnsClosePayload
}

func (webhook *ReturnsClose) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsClose) GetData() (ReturnsClosePayload, error) {
	return webhook.data, nil
}

type ReturnsClosePayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	ID                int64  `json:"id"`
	OrderID           int64  `json:"order_id"`
	Status            string `json:"status"`
}
