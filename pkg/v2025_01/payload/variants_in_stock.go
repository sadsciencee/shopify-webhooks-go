package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewVariantsInStock(webhook webhook.ValidWebhook) (*VariantsInStock, error) {
	payload := &VariantsInStockPayload{}
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
	return &VariantsInStock{
		info: info,
		data: *payload,
	}, nil
}

type VariantsInStock struct {
	info webhook.Info
	data VariantsInStockPayload
}

func (webhook *VariantsInStock) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *VariantsInStock) GetData() (VariantsInStockPayload, error) {
	return webhook.data, nil
}

type VariantsInStockPayload struct {
	AdminGraphqlAPIID    shopify.ID  `json:"admin_graphql_api_id"`
	Barcode              interface{} `json:"barcode"`
	CompareAtPrice       string      `json:"compare_at_price"`
	CreatedAt            time.Time   `json:"created_at"`
	ID                   int64       `json:"id"`
	ImageID              interface{} `json:"image_id"`
	InventoryPolicy      string      `json:"inventory_policy"`
	InventoryQuantity    int64       `json:"inventory_quantity"`
	OldInventoryQuantity int64       `json:"old_inventory_quantity"`
	Option1              string      `json:"option1"`
	Option2              *string     `json:"option2"`
	Option3              *string     `json:"option3"`
	Position             int64       `json:"position"`
	Price                string      `json:"price"`
	ProductID            int64       `json:"product_id"`
	Sku                  *string     `json:"sku"`
	Taxable              bool        `json:"taxable"`
	Title                string      `json:"title"`
	UpdatedAt            time.Time   `json:"updated_at"`
}
