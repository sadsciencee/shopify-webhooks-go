package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewAppUninstalled(webhook webhook.ValidWebhook) (*AppUninstalled, error) {
	payload := &AppUninstalledPayload{}
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
	return &AppUninstalled{
		info: info,
		data: *payload,
	}, nil
}

type AppUninstalled struct {
	info webhook.Info
	data AppUninstalledPayload
}

func (webhook *AppUninstalled) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *AppUninstalled) GetData() (AppUninstalledPayload, error) {
	return webhook.data, nil
}

type AppUninstalledPayload struct {
	Address1                             string      `json:"address1"`
	Address2                             *string     `json:"address2"`
	AutoConfigureTaxInclusivity          interface{} `json:"auto_configure_tax_inclusivity"`
	CheckoutAPISupported                 bool        `json:"checkout_api_supported"`
	City                                 string      `json:"city"`
	Country                              string      `json:"country"`
	CountryCode                          string      `json:"country_code"`
	CountryName                          string      `json:"country_name"`
	CountyTaxes                          interface{} `json:"county_taxes"`
	CreatedAt                            time.Time   `json:"created_at"`
	Currency                             string      `json:"currency"`
	CustomerEmail                        string      `json:"customer_email"`
	Domain                               interface{} `json:"domain"`
	EligibleForPayments                  bool        `json:"eligible_for_payments"`
	Email                                string      `json:"email"`
	EnabledPresentmentCurrencies         []string    `json:"enabled_presentment_currencies"`
	Finances                             bool        `json:"finances"`
	GoogleAppsDomain                     interface{} `json:"google_apps_domain"`
	GoogleAppsLoginEnabled               interface{} `json:"google_apps_login_enabled"`
	HasDiscounts                         bool        `json:"has_discounts"`
	HasGiftCards                         bool        `json:"has_gift_cards"`
	HasStorefront                        bool        `json:"has_storefront"`
	IanaTimezone                         interface{} `json:"iana_timezone"`
	ID                                   int64       `json:"id"`
	Latitude                             interface{} `json:"latitude"`
	Longitude                            interface{} `json:"longitude"`
	MarketingSmsConsentEnabledAtCheckout bool        `json:"marketing_sms_consent_enabled_at_checkout"`
	MoneyFormat                          string      `json:"money_format"`
	MoneyInEmailsFormat                  string      `json:"money_in_emails_format"`
	MoneyWithCurrencyFormat              string      `json:"money_with_currency_format"`
	MoneyWithCurrencyInEmailsFormat      string      `json:"money_with_currency_in_emails_format"`
	MultiLocationEnabled                 bool        `json:"multi_location_enabled"`
	MyshopifyDomain                      interface{} `json:"myshopify_domain"`
	Name                                 string      `json:"name"`
	PasswordEnabled                      interface{} `json:"password_enabled"`
	Phone                                string      `json:"phone"`
	PlanDisplayName                      string      `json:"plan_display_name"`
	PlanName                             string      `json:"plan_name"`
	PreLaunchEnabled                     bool        `json:"pre_launch_enabled"`
	PrimaryLocale                        string      `json:"primary_locale"`
	PrimaryLocationID                    int64       `json:"primary_location_id"`
	Province                             string      `json:"province"`
	ProvinceCode                         string      `json:"province_code"`
	RequiresExtraPaymentsAgreement       bool        `json:"requires_extra_payments_agreement"`
	SetupRequired                        bool        `json:"setup_required"`
	ShopOwner                            string      `json:"shop_owner"`
	Source                               interface{} `json:"source"`
	TaxShipping                          interface{} `json:"tax_shipping"`
	TaxesIncluded                        interface{} `json:"taxes_included"`
	Timezone                             string      `json:"timezone"`
	TransactionalSmsDisabled             bool        `json:"transactional_sms_disabled"`
	UpdatedAt                            time.Time   `json:"updated_at"`
	WeightUnit                           string      `json:"weight_unit"`
	Zip                                  string      `json:"zip"`
}
