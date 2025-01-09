package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewMetaobjectsUpdate(webhook webhook.ValidWebhook) (*MetaobjectsUpdate, error) {
	payload := &MetaobjectsUpdatePayload{}
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
	return &MetaobjectsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type MetaobjectsUpdate struct {
	info webhook.Info
	data MetaobjectsUpdatePayload
}

func (webhook *MetaobjectsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetaobjectsUpdate) GetData() (MetaobjectsUpdatePayload, error) {
	return webhook.data, nil
}

type MetaobjectsUpdatePayload struct {
	Capabilities struct {
		Publishable struct {
			Status string `json:"status"`
		} `json:"publishable"`
	} `json:"capabilities"`
	CreatedAt        time.Time         `json:"created_at"`
	CreatedByAppID   *shopify.ID       `json:"created_by_app_id"`
	CreatedByStaffID *shopify.ID       `json:"created_by_staff_id"`
	DefinitionID     shopify.ID        `json:"definition_id"`
	DisplayName      string            `json:"display_name"`
	Fields           map[string]string `json:"fields"`
	Handle           string            `json:"handle"`
	ID               shopify.ID        `json:"id"`
	Type             string            `json:"type"`
	UpdatedAt        time.Time         `json:"updated_at"`
}
