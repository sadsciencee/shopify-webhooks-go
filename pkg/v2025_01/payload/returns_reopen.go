package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsReopen(webhook webhook.ValidWebhook) (*ReturnsReopen, error) {
	payload := &ReturnsReopenPayload{}
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
	return &ReturnsReopen{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsReopen struct {
	info webhook.Info
	data ReturnsReopenPayload
}

func (webhook *ReturnsReopen) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsReopen) GetData() (ReturnsReopenPayload, error) {
	return webhook.data, nil
}

type ReturnsReopenPayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	ID                int64      `json:"id"`
	OrderID           int64      `json:"order_id"`
	Status            string     `json:"status"`
}
