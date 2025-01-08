package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewProductsUpdate(webhook webhook.ValidWebhook) (*ProductsUpdate, error) {
	payload := &ProductsUpdatePayload{}
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
	return &ProductsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ProductsUpdate struct {
	info webhook.Info
	data ProductsUpdatePayload
}

func (webhook *ProductsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductsUpdate) GetData() (ProductsUpdatePayload, error) {
	return webhook.data, nil
}

type ProductsUpdatePayload struct {
	AdminGraphqlAPIID                 string        `json:"admin_graphql_api_id"`
	BodyHTML                          string        `json:"body_html"`
	Category                          interface{}   `json:"category"`
	CreatedAt                         time.Time     `json:"created_at"`
	Handle                            string        `json:"handle"`
	HasVariantsThatRequiresComponents bool          `json:"has_variants_that_requires_components"`
	ID                                int64         `json:"id"`
	Image                             interface{}   `json:"image"`
	Images                            []interface{} `json:"images"`
	Media                             []interface{} `json:"media"`
	Options                           []interface{} `json:"options"`
	ProductType                       string        `json:"product_type"`
	PublishedAt                       time.Time     `json:"published_at"`
	PublishedScope                    string        `json:"published_scope"`
	Status                            string        `json:"status"`
	Tags                              string        `json:"tags"`
	TemplateSuffix                    interface{}   `json:"template_suffix"`
	Title                             string        `json:"title"`
	UpdatedAt                         time.Time     `json:"updated_at"`
	VariantGids                       []struct {
		AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
		UpdatedAt         time.Time `json:"updated_at"`
	} `json:"variant_gids"`
	Variants []struct {
		AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
		Barcode              interface{} `json:"barcode"`
		CompareAtPrice       string      `json:"compare_at_price"`
		CreatedAt            time.Time   `json:"created_at"`
		ID                   int64       `json:"id"`
		ImageID              interface{} `json:"image_id"`
		InventoryItemID      interface{} `json:"inventory_item_id"`
		InventoryPolicy      string      `json:"inventory_policy"`
		InventoryQuantity    int64       `json:"inventory_quantity"`
		OldInventoryQuantity int64       `json:"old_inventory_quantity"`
		Option1              string      `json:"option1"`
		Option2              interface{} `json:"option2"`
		Option3              interface{} `json:"option3"`
		Position             int64       `json:"position"`
		Price                string      `json:"price"`
		ProductID            int64       `json:"product_id"`
		Sku                  interface{} `json:"sku"`
		Taxable              bool        `json:"taxable"`
		Title                string      `json:"title"`
		UpdatedAt            time.Time   `json:"updated_at"`
	} `json:"variants"`
	Vendor string `json:"vendor"`
}
