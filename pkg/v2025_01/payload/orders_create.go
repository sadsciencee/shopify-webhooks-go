package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewOrdersCreate(webhook webhook.ValidWebhook) (*OrdersCreate, error) {
	payload := &OrdersCreatePayload{}
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
	return &OrdersCreate{
		info: info,
		data: *payload,
	}, nil
}

type OrdersCreate struct {
	info webhook.Info
	data OrdersCreatePayload
}

func (webhook *OrdersCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrdersCreate) GetData() (OrdersCreatePayload, error) {
	return webhook.data, nil
}

type OrdersCreatePayload struct {
	AdminGraphqlAPIID             string           `json:"admin_graphql_api_id"`
	AppID                         interface{}      `json:"app_id"`
	BillingAddress                *shopify.Address `json:"billing_address"`
	BrowserIP                     interface{}      `json:"browser_ip"`
	BuyerAcceptsMarketing         bool             `json:"buyer_accepts_marketing"`
	CancelReason                  string           `json:"cancel_reason"`
	CancelledAt                   time.Time        `json:"cancelled_at"`
	CartToken                     interface{}      `json:"cart_token"`
	CheckoutID                    interface{}      `json:"checkout_id"`
	CheckoutToken                 interface{}      `json:"checkout_token"`
	ClientDetails                 interface{}      `json:"client_details"`
	ClosedAt                      interface{}      `json:"closed_at"`
	ConfirmationNumber            interface{}      `json:"confirmation_number"`
	Confirmed                     bool             `json:"confirmed"`
	ContactEmail                  string           `json:"contact_email"`
	CreatedAt                     time.Time        `json:"created_at"`
	Currency                      string           `json:"currency"`
	CurrentShippingPriceSet       shopify.MoneyBag `json:"current_shipping_price_set"`
	CurrentSubtotalPrice          string           `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet       shopify.MoneyBag `json:"current_subtotal_price_set"`
	CurrentTotalAdditionalFeesSet interface{}      `json:"current_total_additional_fees_set"`
	CurrentTotalDiscounts         string           `json:"current_total_discounts"`
	CurrentTotalDiscountsSet      shopify.MoneyBag `json:"current_total_discounts_set"`
	CurrentTotalDutiesSet         interface{}      `json:"current_total_duties_set"`
	CurrentTotalPrice             string           `json:"current_total_price"`
	CurrentTotalPriceSet          shopify.MoneyBag `json:"current_total_price_set"`
	CurrentTotalTax               string           `json:"current_total_tax"`
	CurrentTotalTaxSet            shopify.MoneyBag `json:"current_total_tax_set"`
	Customer                      shopify.Customer `json:"customer"`
	CustomerLocale                string           `json:"customer_locale"`
	DeviceID                      interface{}      `json:"device_id"`
	DiscountApplications          []interface{}    `json:"discount_applications"`
	DiscountCodes                 []interface{}    `json:"discount_codes"`
	DutiesIncluded                bool             `json:"duties_included"`
	Email                         string           `json:"email"`
	EstimatedTaxes                bool             `json:"estimated_taxes"`
	FinancialStatus               string           `json:"financial_status"`
	FulfillmentStatus             interface{}      `json:"fulfillment_status"`
	Fulfillments                  []interface{}    `json:"fulfillments"`
	ID                            int64            `json:"id"`
	LandingSite                   interface{}      `json:"landing_site"`
	LandingSiteRef                interface{}      `json:"landing_site_ref"`
	LineItems                     []struct {
		AdminGraphqlAPIID          string                    `json:"admin_graphql_api_id"`
		AttributedStaffs           []shopify.AttributedStaff `json:"attributed_staffs"`
		CurrentQuantity            int64                     `json:"current_quantity"`
		DiscountAllocations        []interface{}             `json:"discount_allocations"`
		Duties                     []interface{}             `json:"duties"`
		FulfillableQuantity        int64                     `json:"fulfillable_quantity"`
		FulfillmentService         string                    `json:"fulfillment_service"`
		FulfillmentStatus          interface{}               `json:"fulfillment_status"`
		GiftCard                   bool                      `json:"gift_card"`
		Grams                      int64                     `json:"grams"`
		ID                         int64                     `json:"id"`
		Name                       string                    `json:"name"`
		Price                      string                    `json:"price"`
		PriceSet                   shopify.MoneyBag          `json:"price_set"`
		ProductExists              bool                      `json:"product_exists"`
		ProductID                  int64                     `json:"product_id"`
		Properties                 []interface{}             `json:"properties"`
		Quantity                   int64                     `json:"quantity"`
		RequiresShipping           bool                      `json:"requires_shipping"`
		SalesLineItemGroupID       interface{}               `json:"sales_line_item_group_id"`
		Sku                        string                    `json:"sku"`
		TaxLines                   []shopify.LineItemTaxLine `json:"tax_lines"`
		Taxable                    bool                      `json:"taxable"`
		Title                      string                    `json:"title"`
		TotalDiscount              string                    `json:"total_discount"`
		TotalDiscountSet           shopify.MoneyBag          `json:"total_discount_set"`
		VariantID                  int64                     `json:"variant_id"`
		VariantInventoryManagement string                    `json:"variant_inventory_management"`
		VariantTitle               interface{}               `json:"variant_title"`
		Vendor                     interface{}               `json:"vendor"`
	} `json:"line_items"`
	LocationID                            interface{}            `json:"location_id"`
	MerchantBusinessEntityID              string                 `json:"merchant_business_entity_id"`
	MerchantOfRecordAppID                 interface{}            `json:"merchant_of_record_app_id"`
	Name                                  string                 `json:"name"`
	Note                                  interface{}            `json:"note"`
	NoteAttributes                        []interface{}          `json:"note_attributes"`
	Number                                int64                  `json:"number"`
	OrderNumber                           int64                  `json:"order_number"`
	OrderStatusURL                        string                 `json:"order_status_url"`
	OriginalTotalAdditionalFeesSet        interface{}            `json:"original_total_additional_fees_set"`
	OriginalTotalDutiesSet                interface{}            `json:"original_total_duties_set"`
	PaymentGatewayNames                   []string               `json:"payment_gateway_names"`
	PaymentTerms                          interface{}            `json:"payment_terms"`
	Phone                                 interface{}            `json:"phone"`
	PoNumber                              interface{}            `json:"po_number"`
	PresentmentCurrency                   string                 `json:"presentment_currency"`
	ProcessedAt                           time.Time              `json:"processed_at"`
	Reference                             interface{}            `json:"reference"`
	ReferringSite                         interface{}            `json:"referring_site"`
	Refunds                               []interface{}          `json:"refunds"`
	Returns                               []interface{}          `json:"returns"`
	ShippingAddress                       *shopify.Address       `json:"shipping_address"`
	ShippingLines                         []shopify.ShippingLine `json:"shipping_lines"`
	SourceIdentifier                      interface{}            `json:"source_identifier"`
	SourceName                            string                 `json:"source_name"`
	SourceURL                             interface{}            `json:"source_url"`
	SubtotalPrice                         string                 `json:"subtotal_price"`
	SubtotalPriceSet                      shopify.MoneyBag       `json:"subtotal_price_set"`
	Tags                                  string                 `json:"tags"`
	TaxExempt                             bool                   `json:"tax_exempt"`
	TaxLines                              []shopify.OrderTaxLine `json:"tax_lines"`
	TaxesIncluded                         bool                   `json:"taxes_included"`
	Test                                  bool                   `json:"test"`
	Token                                 string                 `json:"token"`
	TotalCashRoundingPaymentAdjustmentSet shopify.MoneyBag       `json:"total_cash_rounding_payment_adjustment_set"`
	TotalCashRoundingRefundAdjustmentSet  shopify.MoneyBag       `json:"total_cash_rounding_refund_adjustment_set"`
	TotalDiscounts                        string                 `json:"total_discounts"`
	TotalDiscountsSet                     shopify.MoneyBag       `json:"total_discounts_set"`
	TotalLineItemsPrice                   string                 `json:"total_line_items_price"`
	TotalLineItemsPriceSet                shopify.MoneyBag       `json:"total_line_items_price_set"`
	TotalOutstanding                      string                 `json:"total_outstanding"`
	TotalPrice                            string                 `json:"total_price"`
	TotalPriceSet                         shopify.MoneyBag       `json:"total_price_set"`
	TotalShippingPriceSet                 shopify.MoneyBag       `json:"total_shipping_price_set"`
	TotalTax                              string                 `json:"total_tax"`
	TotalTaxSet                           shopify.MoneyBag       `json:"total_tax_set"`
	TotalTipReceived                      string                 `json:"total_tip_received"`
	TotalWeight                           int64                  `json:"total_weight"`
	UpdatedAt                             time.Time              `json:"updated_at"`
	UserID                                interface{}            `json:"user_id"`
}
