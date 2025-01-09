package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewInventoryLevelsUpdate(webhook webhook.ValidWebhook) (*InventoryLevelsUpdate, error) {
	payload := &InventoryLevelsUpdatePayload{}
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
	return &InventoryLevelsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type InventoryLevelsUpdate struct {
	info webhook.Info
	data InventoryLevelsUpdatePayload
}

func (webhook *InventoryLevelsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryLevelsUpdate) GetData() (InventoryLevelsUpdatePayload, error) {
	return webhook.data, nil
}

type InventoryLevelsUpdatePayload struct {
	AdminGraphqlAPIID shopify.ID  `json:"admin_graphql_api_id"`
	Available         interface{} `json:"available"`
	InventoryItemID   int64       `json:"inventory_item_id"`
	LocationID        int64       `json:"location_id"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
