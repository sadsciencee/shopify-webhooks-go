package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCollectionsCreate(webhook webhook.ValidWebhook) (*CollectionsCreate, error) {
	payload := &CollectionsCreatePayload{}
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
	return &CollectionsCreate{
		info: info,
		data: *payload,
	}, nil
}

type CollectionsCreate struct {
	info webhook.Info
	data CollectionsCreatePayload
}

func (webhook *CollectionsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionsCreate) GetData() (CollectionsCreatePayload, error) {
	return webhook.data, nil
}

type CollectionsCreatePayload struct {
	AdminGraphqlAPIID shopify.ID  `json:"admin_graphql_api_id"`
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
