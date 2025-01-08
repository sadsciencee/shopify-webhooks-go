package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewMetafieldDefinitionsCreate(webhook webhook.ValidWebhook) (*MetafieldDefinitionsCreate, error) {
	payload := &MetafieldDefinitionsCreatePayload{}
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
	return &MetafieldDefinitionsCreate{
		info: info,
		data: *payload,
	}, nil
}

type MetafieldDefinitionsCreate struct {
	info webhook.Info
	data MetafieldDefinitionsCreatePayload
}

func (webhook *MetafieldDefinitionsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *MetafieldDefinitionsCreate) GetData() (MetafieldDefinitionsCreatePayload, error) {
	return webhook.data, nil
}

type MetafieldDefinitionsCreatePayload struct {
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
