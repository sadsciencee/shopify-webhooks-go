package shopify

import "time"

type MoneyV2 struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
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
