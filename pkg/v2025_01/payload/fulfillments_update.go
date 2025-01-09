package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewFulfillmentsUpdate(webhook webhook.ValidWebhook) (*FulfillmentsUpdate, error) {
	payload := &FulfillmentsUpdatePayload{}
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
	return &FulfillmentsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentsUpdate struct {
	info webhook.Info
	data FulfillmentsUpdatePayload
}

func (webhook *FulfillmentsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentsUpdate) GetData() (FulfillmentsUpdatePayload, error) {
	return webhook.data, nil
}

type FulfillmentsUpdatePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	CreatedAt         time.Time  `json:"created_at"`
	Destination       struct {
		Address1     string      `json:"address1"`
		Address2     *string     `json:"address2"`
		City         string      `json:"city"`
		Company      string      `json:"company"`
		Country      string      `json:"country"`
		CountryCode  string      `json:"country_code"`
		FirstName    string      `json:"first_name"`
		LastName     string      `json:"last_name"`
		Latitude     interface{} `json:"latitude"`
		Longitude    interface{} `json:"longitude"`
		Name         string      `json:"name"`
		Phone        string      `json:"phone"`
		Province     string      `json:"province"`
		ProvinceCode string      `json:"province_code"`
		Zip          string      `json:"zip"`
	} `json:"destination"`
	Email     string `json:"email"`
	ID        int64  `json:"id"`
	LineItems []struct {
		AdminGraphqlAPIID          string                       `json:"admin_graphql_api_id"`
		DiscountAllocations        []shopify.DiscountAllocation `json:"discount_allocations"`
		Duties                     []interface{}                `json:"duties"`
		FulfillableQuantity        int64                        `json:"fulfillable_quantity"`
		FulfillmentService         string                       `json:"fulfillment_service"`
		FulfillmentStatus          interface{}                  `json:"fulfillment_status"`
		GiftCard                   bool                         `json:"gift_card"`
		Grams                      int64                        `json:"grams"`
		ID                         int64                        `json:"id"`
		Name                       string                       `json:"name"`
		Price                      string                       `json:"price"`
		PriceSet                   shopify.MoneyBag             `json:"price_set"`
		ProductExists              bool                         `json:"product_exists"`
		ProductID                  int64                        `json:"product_id"`
		Properties                 []shopify.CustomAttribute    `json:"properties"`
		Quantity                   int64                        `json:"quantity"`
		RequiresShipping           bool                         `json:"requires_shipping"`
		Sku                        string                       `json:"sku"`
		TaxLines                   []shopify.LineItemTaxLine    `json:"tax_lines"`
		Taxable                    bool                         `json:"taxable"`
		Title                      string                       `json:"title"`
		TotalDiscount              string                       `json:"total_discount"`
		TotalDiscountSet           shopify.MoneyBag             `json:"total_discount_set"`
		VariantID                  int64                        `json:"variant_id"`
		VariantInventoryManagement string                       `json:"variant_inventory_management"`
		VariantTitle               interface{}                  `json:"variant_title"`
		Vendor                     interface{}                  `json:"vendor"`
	} `json:"line_items"`
	LocationID      interface{} `json:"location_id"`
	Name            string      `json:"name"`
	OrderID         int64       `json:"order_id"`
	OriginAddress   interface{} `json:"origin_address"`
	Receipt         struct{}    `json:"receipt"`
	Service         interface{} `json:"service"`
	ShipmentStatus  interface{} `json:"shipment_status"`
	Status          string      `json:"status"`
	TrackingCompany string      `json:"tracking_company"`
	TrackingNumber  string      `json:"tracking_number"`
	TrackingNumbers []string    `json:"tracking_numbers"`
	TrackingURL     string      `json:"tracking_url"`
	TrackingUrls    []string    `json:"tracking_urls"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
