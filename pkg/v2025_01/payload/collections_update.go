package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCollectionsUpdate(webhook webhook.ValidWebhook) (*CollectionsUpdate, error) {
	payload := &CollectionsUpdatePayload{}
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
	return &CollectionsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CollectionsUpdate struct {
	info webhook.Info
	data CollectionsUpdatePayload
}

func (webhook *CollectionsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionsUpdate) GetData() (CollectionsUpdatePayload, error) {
	return webhook.data, nil
}

type CollectionsUpdatePayload struct {
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	BodyHTML          string      `json:"body_html"`
	Handle            string      `json:"handle"`
	ID                int64       `json:"id"`
	PublishedAt       time.Time   `json:"published_at"`
	PublishedScope    string      `json:"published_scope"`
	SortOrder         interface{} `json:"sort_order"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	Title             string      `json:"title"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
