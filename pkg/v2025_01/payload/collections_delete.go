package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCollectionsDelete(webhook webhook.ValidWebhook) (*CollectionsDelete, error) {
	payload := &CollectionsDeletePayload{}
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
	return &CollectionsDelete{
		info: info,
		data: *payload,
	}, nil
}

type CollectionsDelete struct {
	info webhook.Info
	data CollectionsDeletePayload
}

func (webhook *CollectionsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionsDelete) GetData() (CollectionsDeletePayload, error) {
	return webhook.data, nil
}

type CollectionsDeletePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	ID                int64      `json:"id"`
	PublishedScope    string     `json:"published_scope"`
}
