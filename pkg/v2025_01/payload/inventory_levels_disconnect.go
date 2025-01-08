package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewInventoryLevelsDisconnect(webhook webhook.ValidWebhook) (*InventoryLevelsDisconnect, error) {
	payload := &InventoryLevelsDisconnectPayload{}
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
	return &InventoryLevelsDisconnect{
		info: info,
		data: *payload,
	}, nil
}

type InventoryLevelsDisconnect struct {
	info webhook.Info
	data InventoryLevelsDisconnectPayload
}

func (webhook *InventoryLevelsDisconnect) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryLevelsDisconnect) GetData() (InventoryLevelsDisconnectPayload, error) {
	return webhook.data, nil
}

type InventoryLevelsDisconnectPayload struct {
	InventoryItemID int64 `json:"inventory_item_id"`
	LocationID      int64 `json:"location_id"`
}
