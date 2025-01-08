package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewInventoryItemsUpdate(webhook webhook.ValidWebhook) (*InventoryItemsUpdate, error) {
	payload := &InventoryItemsUpdatePayload{}
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
	return &InventoryItemsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type InventoryItemsUpdate struct {
	info webhook.Info
	data InventoryItemsUpdatePayload
}

func (webhook *InventoryItemsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryItemsUpdate) GetData() (InventoryItemsUpdatePayload, error) {
	return webhook.data, nil
}

type InventoryItemsUpdatePayload struct {
	AdminGraphqlAPIID            string        `json:"admin_graphql_api_id"`
	Cost                         interface{}   `json:"cost"`
	CountryCodeOfOrigin          interface{}   `json:"country_code_of_origin"`
	CountryHarmonizedSystemCodes []interface{} `json:"country_harmonized_system_codes"`
	CreatedAt                    time.Time     `json:"created_at"`
	HarmonizedSystemCode         interface{}   `json:"harmonized_system_code"`
	ID                           int64         `json:"id"`
	ProvinceCodeOfOrigin         interface{}   `json:"province_code_of_origin"`
	RequiresShipping             bool          `json:"requires_shipping"`
	Sku                          string        `json:"sku"`
	Tracked                      bool          `json:"tracked"`
	UpdatedAt                    time.Time     `json:"updated_at"`
}
