package shopify

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ID string

func (g *ID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	prefix := "gid://shopify/"
	if !strings.HasPrefix(s, prefix) {
		return fmt.Errorf("invalid ID format: must start with %s", prefix)
	}

	parts := strings.Split(s[len(prefix):], "/")
	if len(parts) != 2 {
		return fmt.Errorf("invalid ID format: must be %s{resource}/{id}", prefix)
	}

	*g = ID(s)
	return nil
}

func (g *ID) ResourceName() string {
	parts := strings.Split(string(*g)[13:], "/")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

func (g *ID) CleanId() string {
	parts := strings.Split(string(*g)[13:], "/")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func (g *ID) IntId() (int64, error) {
	id := g.CleanId()
	return strconv.ParseInt(id, 10, 64)
}

type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Decimal string

func (d *Decimal) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		sv := strings.ReplaceAll(s, ",", ".")
		if _, err := strconv.ParseFloat(sv, 64); err != nil {
			return fmt.Errorf("decimal: invalid decimal string: %s", s)
		}

		*d = Decimal(s)
		return nil
	}

	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*d = Decimal(strconv.FormatFloat(f, 'f', -1, 64))
		return nil
	}

	return fmt.Errorf("decimal: cannot unmarshal %s", string(data))
}

func (d *Decimal) ToFloat() (float64, error) {
	s := strings.ReplaceAll(string(*d), ",", ".")
	return strconv.ParseFloat(s, 64)
}

type MoneyV2 struct {
	Amount       Decimal `json:"amount"`
	CurrencyCode string  `json:"currency_code"`
}

type MoneyBag struct {
	PresentmentMoney MoneyV2 `json:"presentment_money"`
	ShopMoney        MoneyV2 `json:"shop_money"`
}

type Address struct {
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
}

type Customer struct {
	AdminGraphqlAPIID   string                  `json:"admin_graphql_api_id"`
	CreatedAt           time.Time               `json:"created_at"`
	Currency            string                  `json:"currency"`
	DefaultAddress      *CustomerDefaultAddress `json:"default_address"`
	Email               *string                 `json:"email"`
	FirstName           *string                 `json:"first_name"`
	ID                  int64                   `json:"id"`
	LastName            *string                 `json:"last_name"`
	MultipassIdentifier interface{}             `json:"multipass_identifier"`
	Note                *string                 `json:"note"`
	Phone               *string                 `json:"phone"`
	State               *string                 `json:"state"`
	TaxExempt           bool                    `json:"tax_exempt"`
	TaxExemptions       []interface{}           `json:"tax_exemptions"`
	UpdatedAt           time.Time               `json:"updated_at"`
	VerifiedEmail       bool                    `json:"verified_email"`
}

type CustomerDefaultAddress struct {
	Address1     string  `json:"address1"`
	Address2     *string `json:"address2"`
	City         string  `json:"city"`
	Company      *string `json:"company"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"country_code"`
	CountryName  string  `json:"country_name"`
	CustomerID   int64   `json:"customer_id"`
	Default      bool    `json:"default"`
	FirstName    string  `json:"first_name"`
	ID           int64   `json:"id"`
	LastName     string  `json:"last_name"`
	Name         string  `json:"name"`
	Phone        string  `json:"phone"`
	Province     string  `json:"province"`
	ProvinceCode string  `json:"province_code"`
	Zip          string  `json:"zip"`
}

type SellingPlanBillingPolicy struct {
	Interval      string `json:"interval"`
	IntervalCount int64  `json:"interval_count"`
	MaxCycles     *int64 `json:"max_cycles"`
	MinCycles     *int64 `json:"min_cycles"`
}

type SellingPlanDeliveryPolicy struct {
	Anchors           []SellingPlanAnchor `json:"anchors"`
	Cutoff            interface{}         `json:"cutoff"`
	Interval          string              `json:"interval"`
	IntervalCount     int64               `json:"interval_count"`
	PreAnchorBehavior *string             `json:"pre_anchor_behavior"`
}

type SellingPlanAnchor struct {
}

type SellingPlanPricingPolicy struct {
}

type SellingPlan struct {
	BillingPolicy   SellingPlanBillingPolicy   `json:"billing_policy"`
	DeliveryPolicy  SellingPlanDeliveryPolicy  `json:"delivery_policy"`
	Description     *string                    `json:"description"`
	Name            string                     `json:"name"`
	Options         []string                   `json:"options"`
	Position        interface{}                `json:"position"`
	PricingPolicies []SellingPlanPricingPolicy `json:"pricing_policies"`
}

type AttributedStaff struct {
	ID       string `json:"id"`
	Quantity int64  `json:"quantity"`
}

type LineItem struct {
	AdminGraphqlAPIID          string            `json:"admin_graphql_api_id"`
	AttributedStaffs           []AttributedStaff `json:"attributed_staffs"`
	CurrentQuantity            int64             `json:"current_quantity"`
	DiscountAllocations        []interface{}     `json:"discount_allocations"`
	Duties                     []interface{}     `json:"duties"`
	FulfillableQuantity        int64             `json:"fulfillable_quantity"`
	FulfillmentService         string            `json:"fulfillment_service"`
	FulfillmentStatus          interface{}       `json:"fulfillment_status"`
	GiftCard                   bool              `json:"gift_card"`
	Grams                      int64             `json:"grams"`
	ID                         int64             `json:"id"`
	Name                       string            `json:"name"`
	Price                      string            `json:"price"`
	PriceSet                   MoneyBag          `json:"price_set"`
	ProductExists              bool              `json:"product_exists"`
	ProductID                  int64             `json:"product_id"`
	Properties                 []CustomAttribute `json:"properties"`
	Quantity                   int64             `json:"quantity"`
	RequiresShipping           bool              `json:"requires_shipping"`
	SalesLineItemGroupID       interface{}       `json:"sales_line_item_group_id"`
	Sku                        string            `json:"sku"`
	TaxLines                   []LineItemTaxLine `json:"tax_lines"`
	Taxable                    bool              `json:"taxable"`
	Title                      string            `json:"title"`
	TotalDiscount              string            `json:"total_discount"`
	TotalDiscountSet           MoneyBag          `json:"total_discount_set"`
	VariantID                  int64             `json:"variant_id"`
	VariantInventoryManagement string            `json:"variant_inventory_management"`
	VariantTitle               interface{}       `json:"variant_title"`
	Vendor                     interface{}       `json:"vendor"`
}

type LineItemTaxLine struct {
	ChannelLiable             bool        `json:"channel_liable"`
	CompareAt                 float64     `json:"compare_at"`
	JurisdictionID            interface{} `json:"jurisdiction_id"`
	JurisdictionSource        interface{} `json:"jurisdiction_source"`
	JurisdictionType          interface{} `json:"jurisdiction_type"`
	Position                  int64       `json:"position"`
	Price                     string      `json:"price"`
	Rate                      float64     `json:"rate"`
	ReportingExemptAmount     interface{} `json:"reporting_exempt_amount"`
	ReportingJurisdictionCode interface{} `json:"reporting_jurisdiction_code"`
	ReportingJurisdictionName interface{} `json:"reporting_jurisdiction_name"`
	ReportingJurisdictionType interface{} `json:"reporting_jurisdiction_type"`
	ReportingNonTaxableAmount interface{} `json:"reporting_non_taxable_amount"`
	ReportingTaxableAmount    interface{} `json:"reporting_taxable_amount"`
	Source                    string      `json:"source"`
	TaxAPIClientID            interface{} `json:"tax_api_client_id"`
	TaxCalculationPrice       string      `json:"tax_calculation_price"`
	TaxRegistrationID         interface{} `json:"tax_registration_id"`
	Title                     string      `json:"title"`
	Zone                      string      `json:"zone"`
}

type OrderTaxLine struct {
	ChannelLiable bool    `json:"channel_liable"`
	Price         string  `json:"price"`
	Rate          float64 `json:"rate"`
	Title         string  `json:"title"`
}

type UnitPriceMeasurement struct {
	MeasuredType   *string `json:"measured_type"`
	QuantityUnit   *string `json:"quantity_unit"`
	QuantityValue  *string `json:"quantity_value"`
	ReferenceUnit  *string `json:"reference_unit"`
	ReferenceValue *string `json:"reference_value"`
}

type ShippingTaxLine struct{}

type ShippingLine struct {
	CarrierIdentifier             interface{}       `json:"carrier_identifier"`
	Code                          interface{}       `json:"code"`
	CurrentDiscountedPriceSet     MoneyBag          `json:"current_discounted_price_set"`
	DiscountAllocations           []interface{}     `json:"discount_allocations"`
	DiscountedPrice               string            `json:"discounted_price"`
	DiscountedPriceSet            MoneyBag          `json:"discounted_price_set"`
	ID                            int64             `json:"id"`
	IsRemoved                     bool              `json:"is_removed"`
	Phone                         *string           `json:"phone"`
	Price                         string            `json:"price"`
	PriceSet                      MoneyBag          `json:"price_set"`
	RequestedFulfillmentServiceID interface{}       `json:"requested_fulfillment_service_id"`
	Source                        string            `json:"source"`
	TaxLines                      []ShippingTaxLine `json:"tax_lines"`
	Title                         string            `json:"title"`
}

type DiscountAllocation struct {
	Amount                   string   `json:"amount"`
	AmountSet                MoneyBag `json:"amount_set"`
	DiscountApplicationIndex int64    `json:"discount_application_index"`
}

type Localization struct {
	AlternateLocales []string `json:"alternate_locales"`
	Country          *string  `json:"country"`
	DefaultLocale    string   `json:"default_locale"`
}

type DeliveryMethod struct {
	MethodType string `json:"method_type"`
}

type FulfillmentOrder struct {
	ID                 ID              `json:"id"`
	Status             *string         `json:"status"`
	RequestStatus      *string         `json:"request_status"`
	DeliveryMethod     *DeliveryMethod `json:"delivery_method"`
	Preparable         *bool           `json:"preparable"`
	AssignedLocationID *string         `json:"assigned_location_id"`
}

type ProductOption struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Position  int64    `json:"position"`
	ProductID int64    `json:"product_id"`
	Values    []string `json:"values"`
}
