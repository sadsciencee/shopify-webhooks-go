package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewScheduledProductListingsUpdate(webhook webhook.ValidWebhook) (*ScheduledProductListingsUpdate, error) {
	payload := &ScheduledProductListingsUpdatePayload{}
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
	return &ScheduledProductListingsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ScheduledProductListingsUpdate struct {
	info webhook.Info
	data ScheduledProductListingsUpdatePayload
}

func (webhook *ScheduledProductListingsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ScheduledProductListingsUpdate) GetData() (ScheduledProductListingsUpdatePayload, error) {
	return webhook.data, nil
}

type ScheduledProductListingsUpdatePayload struct {
	ScheduledProductListing struct {
		Available   bool                    `json:"available"`
		BodyHTML    string                  `json:"body_html"`
		CreatedAt   time.Time               `json:"created_at"`
		Handle      string                  `json:"handle"`
		Images      []interface{}           `json:"images"`
		Options     []shopify.ProductOption `json:"options"`
		ProductID   int64                   `json:"product_id"`
		ProductType string                  `json:"product_type"`
		PublishAt   interface{}             `json:"publish_at"`
		Tags        string                  `json:"tags"`
		Title       string                  `json:"title"`
		UpdatedAt   time.Time               `json:"updated_at"`
		Variants    []struct {
			Available           bool        `json:"available"`
			Barcode             interface{} `json:"barcode"`
			CompareAtPrice      string      `json:"compare_at_price"`
			CreatedAt           time.Time   `json:"created_at"`
			FormattedPrice      string      `json:"formatted_price"`
			FulfillmentService  string      `json:"fulfillment_service"`
			Grams               int64       `json:"grams"`
			ID                  int64       `json:"id"`
			ImageID             interface{} `json:"image_id"`
			InventoryManagement interface{} `json:"inventory_management"`
			InventoryPolicy     string      `json:"inventory_policy"`
			InventoryQuantity   int64       `json:"inventory_quantity"`
			OptionValues        []struct {
				Name     string `json:"name"`
				OptionID int64  `json:"option_id"`
				Value    string `json:"value"`
			} `json:"option_values"`
			Position         int64       `json:"position"`
			Price            string      `json:"price"`
			RequiresShipping bool        `json:"requires_shipping"`
			Sku              interface{} `json:"sku"`
			Taxable          bool        `json:"taxable"`
			Title            string      `json:"title"`
			UpdatedAt        time.Time   `json:"updated_at"`
			Weight           float64     `json:"weight"`
			WeightUnit       string      `json:"weight_unit"`
		} `json:"variants"`
		Vendor string `json:"vendor"`
	} `json:"scheduled_product_listing"`
}
