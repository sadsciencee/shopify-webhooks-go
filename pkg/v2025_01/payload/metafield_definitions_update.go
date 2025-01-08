package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewMetafieldDefinitionsUpdate(webhook webhook.ValidWebhook) (*MetafieldDefinitionsUpdate, error) {
	payload := &MetafieldDefinitionsUpdatePayload{}
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
	return &MetafieldDefinitionsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type MetafieldDefinitionsUpdate struct {
	info webhook.Info
	data MetafieldDefinitionsUpdatePayload
}

func (webhook *MetafieldDefinitionsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetafieldDefinitionsUpdate) GetData() (MetafieldDefinitionsUpdatePayload, error) {
	return webhook.data, nil
}

type MetafieldDefinitionsUpdatePayload struct {
	Access                   string        `json:"access"`
	AdminFilterStatus        string        `json:"admin_filter_status"`
	APIClientID              interface{}   `json:"api_client_id"`
	CreatedAt                time.Time     `json:"created_at"`
	Deleting                 bool          `json:"deleting"`
	Description              interface{}   `json:"description"`
	ID                       string        `json:"id"`
	Key                      string        `json:"key"`
	Name                     string        `json:"name"`
	Namespace                string        `json:"namespace"`
	Options                  []interface{} `json:"options"`
	OwnerType                string        `json:"owner_type"`
	PinnedPosition           int64         `json:"pinned_position"`
	ShopID                   int64         `json:"shop_id"`
	StandardTemplateID       interface{}   `json:"standard_template_id"`
	TypeName                 string        `json:"type_name"`
	UpdatedAt                time.Time     `json:"updated_at"`
	UseAsCollectionCondition bool          `json:"use_as_collection_condition"`
	ValidationStatus         string        `json:"validation_status"`
}
