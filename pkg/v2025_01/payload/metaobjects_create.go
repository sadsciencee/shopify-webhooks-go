package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewMetaobjectsCreate(webhook webhook.ValidWebhook) (*MetaobjectsCreate, error) {
	payload := &MetaobjectsCreatePayload{}
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
	return &MetaobjectsCreate{
		info: info,
		data: *payload,
	}, nil
}

type MetaobjectsCreate struct {
	info webhook.Info
	data MetaobjectsCreatePayload
}

func (webhook *MetaobjectsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetaobjectsCreate) GetData() (MetaobjectsCreatePayload, error) {
	return webhook.data, nil
}

type MetaobjectsCreatePayload struct {
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
