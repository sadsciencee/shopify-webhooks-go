package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewProductFeedsIncrementalSync(webhook webhook.ValidWebhook) (*ProductFeedsIncrementalSync, error) {
	payload := &ProductFeedsIncrementalSyncPayload{}
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
	return &ProductFeedsIncrementalSync{
		info: info,
		data: *payload,
	}, nil
}

type ProductFeedsIncrementalSync struct {
	info webhook.Info
	data ProductFeedsIncrementalSyncPayload
}

func (webhook *ProductFeedsIncrementalSync) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductFeedsIncrementalSync) GetData() (ProductFeedsIncrementalSyncPayload, error) {
	return webhook.data, nil
}

type ProductFeedsIncrementalSyncPayload struct {
	Metadata struct {
		Action          string        `json:"action"`
		OccuredAt       time.Time     `json:"occured_at"`
		Resource        string        `json:"resource"`
		TruncatedFields []interface{} `json:"truncatedFields"`
		Type            string        `json:"type"`
	} `json:"metadata"`
	Product struct {
		CreatedAt   time.Time  `json:"createdAt"`
		Description string     `json:"description"`
		Handle      string     `json:"handle"`
		ID          shopify.ID `json:"id"`
		Images      struct {
			Edges []struct {
				Node struct {
					Height int64      `json:"height"`
					ID     shopify.ID `json:"id"`
					URL    string     `json:"url"`
					Width  int64      `json:"width"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"images"`
		IsPublished    bool   `json:"isPublished"`
		OnlineStoreURL string `json:"onlineStoreUrl"`
		Options        []struct {
			Name   string   `json:"name"`
			Values []string `json:"values"`
		} `json:"options"`
		ProductType string    `json:"productType"`
		PublishedAt time.Time `json:"publishedAt"`
		Seo         struct {
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"seo"`
		Tags      []string  `json:"tags"`
		Title     string    `json:"title"`
		UpdatedAt time.Time `json:"updatedAt"`
		Variants  struct {
			Edges []struct {
				Node struct {
					AvailableForSale bool        `json:"availableForSale"`
					Barcode          interface{} `json:"barcode"`
					CompareAtPrice   interface{} `json:"compareAtPrice"`
					CreatedAt        time.Time   `json:"createdAt"`
					ID               shopify.ID  `json:"id"`
					Image            struct {
						Height int64      `json:"height"`
						ID     shopify.ID `json:"id"`
						URL    string     `json:"url"`
						Width  int64      `json:"width"`
					} `json:"image"`
					InventoryPolicy   string          `json:"inventoryPolicy"`
					Price             shopify.MoneyV2 `json:"price"`
					QuantityAvailable int64           `json:"quantityAvailable"`
					RequireShipping   bool            `json:"requireShipping"`
					SelectedOptions   []struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					} `json:"selectedOptions"`
					Sku        string    `json:"sku"`
					Title      string    `json:"title"`
					UpdatedAt  time.Time `json:"updatedAt"`
					Weight     float64   `json:"weight"`
					WeightUnit string    `json:"weightUnit"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"variants"`
		Vendor string `json:"vendor"`
	} `json:"product"`
	ProductFeed struct {
		Country  string     `json:"country"`
		ID       shopify.ID `json:"id"`
		Language string     `json:"language"`
		ShopID   shopify.ID `json:"shop_id"`
	} `json:"productFeed"`
	Products interface{} `json:"products"`
}
