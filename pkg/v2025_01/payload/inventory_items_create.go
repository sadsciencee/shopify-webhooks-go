package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewInventoryItemsCreate(webhook webhook.ValidWebhook) (*InventoryItemsCreate, error) {
	payload := &InventoryItemsCreatePayload{}
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
	return &InventoryItemsCreate{
		info: info,
		data: *payload,
	}, nil
}

type InventoryItemsCreate struct {
	info webhook.Info
	data InventoryItemsCreatePayload
}

func (webhook *InventoryItemsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryItemsCreate) GetData() (InventoryItemsCreatePayload, error) {
	return webhook.data, nil
}

type InventoryItemsCreatePayload struct {
	AdminGraphqlAPIID            shopify.ID    `json:"admin_graphql_api_id"`
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
