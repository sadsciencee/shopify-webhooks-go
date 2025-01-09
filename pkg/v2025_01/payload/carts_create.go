package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCartsCreate(webhook webhook.ValidWebhook) (*CartsCreate, error) {
	payload := &CartsCreatePayload{}
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
	return &CartsCreate{
		info: info,
		data: *payload,
	}, nil
}

type CartsCreate struct {
	info webhook.Info
	data CartsCreatePayload
}

func (webhook *CartsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CartsCreate) GetData() (CartsCreatePayload, error) {
	return webhook.data, nil
}

type CartsCreatePayload struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	LineItems []struct {
		DiscountedPrice      string            `json:"discounted_price"`
		DiscountedPriceSet   shopify.MoneyBag  `json:"discounted_price_set"`
		Discounts            []interface{}     `json:"discounts"`
		GiftCard             bool              `json:"gift_card"`
		Grams                int64             `json:"grams"`
		ID                   int64             `json:"id"`
		Key                  string            `json:"key"`
		LinePrice            string            `json:"line_price"`
		LinePriceSet         shopify.MoneyBag  `json:"line_price_set"`
		OriginalLinePrice    string            `json:"original_line_price"`
		OriginalLinePriceSet shopify.MoneyBag  `json:"original_line_price_set"`
		OriginalPrice        string            `json:"original_price"`
		Price                string            `json:"price"`
		PriceSet             shopify.MoneyBag  `json:"price_set"`
		ProductID            int64             `json:"product_id"`
		Properties           map[string]string `json:"properties"`
		Quantity             int64             `json:"quantity"`
		Sku                  string            `json:"sku"`
		Taxable              bool              `json:"taxable"`
		Title                string            `json:"title"`
		TotalDiscount        string            `json:"total_discount"`
		TotalDiscountSet     shopify.MoneyBag  `json:"total_discount_set"`
		VariantID            int64             `json:"variant_id"`
		Vendor               string            `json:"vendor"`
	} `json:"line_items"`
	Note      interface{} `json:"note"`
	Token     string      `json:"token"`
	UpdatedAt time.Time   `json:"updated_at"`
}
