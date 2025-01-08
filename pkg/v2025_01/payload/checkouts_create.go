package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCheckoutsCreate(webhook webhook.ValidWebhook) (*CheckoutsCreate, error) {
	payload := &CheckoutsCreatePayload{}
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
	return &CheckoutsCreate{
		info: info,
		data: *payload,
	}, nil
}

type CheckoutsCreate struct {
	info webhook.Info
	data CheckoutsCreatePayload
}

func (webhook *CheckoutsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CheckoutsCreate) GetData() (CheckoutsCreatePayload, error) {
	return webhook.data, nil
}

type CheckoutsCreatePayload struct {
	AbandonedCheckoutURL     string           `json:"abandoned_checkout_url"`
	BillingAddress           *shopify.Address `json:"billing_address"`
	BuyerAcceptsMarketing    bool             `json:"buyer_accepts_marketing"`
	BuyerAcceptsSmsMarketing bool             `json:"buyer_accepts_sms_marketing"`
	CartToken                string           `json:"cart_token"`
	ClosedAt                 interface{}      `json:"closed_at"`
	CompletedAt              interface{}      `json:"completed_at"`
	CreatedAt                time.Time        `json:"created_at"`
	Currency                 string           `json:"currency"`
	Customer                 shopify.Customer `json:"customer"`
	CustomerLocale           interface{}      `json:"customer_locale"`
	DeviceID                 interface{}      `json:"device_id"`
	DiscountCodes            []interface{}    `json:"discount_codes"`
	Email                    string           `json:"email"`
	Gateway                  interface{}      `json:"gateway"`
	ID                       int64            `json:"id"`
	LandingSite              interface{}      `json:"landing_site"`
	LineItems                []struct {
		AppliedDiscounts        []interface{}                `json:"applied_discounts"`
		CompareAtPrice          interface{}                  `json:"compare_at_price"`
		DestinationLocationID   int64                        `json:"destination_location_id"`
		DiscountAllocations     []interface{}                `json:"discount_allocations"`
		FulfillmentService      string                       `json:"fulfillment_service"`
		GiftCard                bool                         `json:"gift_card"`
		Grams                   int64                        `json:"grams"`
		Key                     string                       `json:"key"`
		LinePrice               string                       `json:"line_price"`
		OriginLocationID        int64                        `json:"origin_location_id"`
		PresentmentTitle        string                       `json:"presentment_title"`
		PresentmentVariantTitle string                       `json:"presentment_variant_title"`
		Price                   string                       `json:"price"`
		ProductID               int64                        `json:"product_id"`
		Properties              interface{}                  `json:"properties"`
		Quantity                int64                        `json:"quantity"`
		Rank                    interface{}                  `json:"rank"`
		RequiresShipping        bool                         `json:"requires_shipping"`
		Sku                     string                       `json:"sku"`
		TaxLines                []shopify.LineItemTaxLine    `json:"tax_lines"`
		Taxable                 bool                         `json:"taxable"`
		Title                   string                       `json:"title"`
		UnitPriceMeasurement    shopify.UnitPriceMeasurement `json:"unit_price_measurement"`
		UserID                  interface{}                  `json:"user_id"`
		VariantID               int64                        `json:"variant_id"`
		VariantPrice            string                       `json:"variant_price"`
		VariantTitle            string                       `json:"variant_title"`
		Vendor                  string                       `json:"vendor"`
	} `json:"line_items"`
	LocationID          interface{}            `json:"location_id"`
	Name                string                 `json:"name"`
	Note                interface{}            `json:"note"`
	NoteAttributes      []interface{}          `json:"note_attributes"`
	Phone               interface{}            `json:"phone"`
	PresentmentCurrency string                 `json:"presentment_currency"`
	ReferringSite       interface{}            `json:"referring_site"`
	ReservationToken    interface{}            `json:"reservation_token"`
	ShippingAddress     *shopify.Address       `json:"shipping_address"`
	ShippingLines       []interface{}          `json:"shipping_lines"`
	SmsMarketingPhone   interface{}            `json:"sms_marketing_phone"`
	Source              interface{}            `json:"source"`
	SourceIdentifier    interface{}            `json:"source_identifier"`
	SourceName          string                 `json:"source_name"`
	SourceURL           interface{}            `json:"source_url"`
	SubtotalPrice       string                 `json:"subtotal_price"`
	TaxLines            []shopify.OrderTaxLine `json:"tax_lines"`
	TaxesIncluded       bool                   `json:"taxes_included"`
	Token               string                 `json:"token"`
	TotalDiscounts      string                 `json:"total_discounts"`
	TotalDuties         interface{}            `json:"total_duties"`
	TotalLineItemsPrice string                 `json:"total_line_items_price"`
	TotalPrice          string                 `json:"total_price"`
	TotalTax            string                 `json:"total_tax"`
	TotalWeight         int64                  `json:"total_weight"`
	UpdatedAt           time.Time              `json:"updated_at"`
	UserID              interface{}            `json:"user_id"`
}
