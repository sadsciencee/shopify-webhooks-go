package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewInventoryItemsDelete(webhook webhook.ValidWebhook) (*InventoryItemsDelete, error) {
	payload := &InventoryItemsDeletePayload{}
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
	return &InventoryItemsDelete{
		info: info,
		data: *payload,
	}, nil
}

type InventoryItemsDelete struct {
	info webhook.Info
	data InventoryItemsDeletePayload
}

func (webhook *InventoryItemsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *InventoryItemsDelete) GetData() (InventoryItemsDeletePayload, error) {
	return webhook.data, nil
}

type InventoryItemsDeletePayload struct {
	AdminGraphqlAPIID            shopify.ID    `json:"admin_graphql_api_id"`
	CountryCodeOfOrigin          interface{}   `json:"country_code_of_origin"`
	CountryHarmonizedSystemCodes []interface{} `json:"country_harmonized_system_codes"`
	HarmonizedSystemCode         interface{}   `json:"harmonized_system_code"`
	ID                           int64         `json:"id"`
	ProvinceCodeOfOrigin         interface{}   `json:"province_code_of_origin"`
}
