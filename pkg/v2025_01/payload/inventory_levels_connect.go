package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewInventoryLevelsConnect(webhook webhook.ValidWebhook) (*InventoryLevelsConnect, error) {
	payload := &InventoryLevelsConnectPayload{}
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
	return &InventoryLevelsConnect{
		info: info,
		data: *payload,
	}, nil
}

type InventoryLevelsConnect struct {
	info webhook.Info
	data InventoryLevelsConnectPayload
}

func (webhook *InventoryLevelsConnect) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryLevelsConnect) GetData() (InventoryLevelsConnectPayload, error) {
	return webhook.data, nil
}

type InventoryLevelsConnectPayload struct {
	AdminGraphqlAPIID shopify.ID  `json:"admin_graphql_api_id"`
	Available         interface{} `json:"available"`
	InventoryItemID   int64       `json:"inventory_item_id"`
	LocationID        int64       `json:"location_id"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
