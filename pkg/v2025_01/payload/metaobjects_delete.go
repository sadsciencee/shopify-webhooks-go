package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewMetaobjectsDelete(webhook webhook.ValidWebhook) (*MetaobjectsDelete, error) {
	payload := &MetaobjectsDeletePayload{}
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
	return &MetaobjectsDelete{
		info: info,
		data: *payload,
	}, nil
}

type MetaobjectsDelete struct {
	info webhook.Info
	data MetaobjectsDeletePayload
}

func (webhook *MetaobjectsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetaobjectsDelete) GetData() (MetaobjectsDeletePayload, error) {
	return webhook.data, nil
}

type MetaobjectsDeletePayload struct {
	CreatedByAppID *shopify.ID `json:"created_by_app_id"`
	Handle         string      `json:"handle"`
	ID             shopify.ID  `json:"id"`
	Type           string      `json:"type"`
}
