package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompanyContactRolesAssign(webhook webhook.ValidWebhook) (*CompanyContactRolesAssign, error) {
	payload := &CompanyContactRolesAssignPayload{}
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
	return &CompanyContactRolesAssign{
		info: info,
		data: *payload,
	}, nil
}

type CompanyContactRolesAssign struct {
	info webhook.Info
	data CompanyContactRolesAssignPayload
}

func (webhook *CompanyContactRolesAssign) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompanyContactRolesAssign) GetData() (CompanyContactRolesAssignPayload, error) {
	return webhook.data, nil
}

type CompanyContactRolesAssignPayload struct {
	CompanyContact struct {
		AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
		Company           struct {
			AdminGraphqlAPIID            shopify.ID `json:"admin_graphql_api_id"`
			CreatedAt                    time.Time  `json:"created_at"`
			CustomerSince                string     `json:"customer_since"`
			ExternalID                   string     `json:"external_id"`
			MainContactAdminGraphqlAPIID shopify.ID `json:"main_contact_admin_graphql_api_id"`
			Name                         string     `json:"name"`
			Note                         string     `json:"note"`
			UpdatedAt                    time.Time  `json:"updated_at"`
		} `json:"company"`
		CreatedAt                 time.Time  `json:"created_at"`
		CustomerAdminGraphqlAPIID shopify.ID `json:"customer_admin_graphql_api_id"`
		Locale                    string     `json:"locale"`
		Title                     string     `json:"title"`
		UpdatedAt                 time.Time  `json:"updated_at"`
	} `json:"company_contact"`
	CompanyContactRole struct {
		Name string `json:"name"`
	} `json:"company_contact_role"`
	CompanyLocation struct {
		AdminGraphqlAPIID            shopify.ID       `json:"admin_graphql_api_id"`
		BillingAddress               *shopify.Address `json:"billing_address"`
		BuyerExperienceConfiguration interface{}      `json:"buyer_experience_configuration"`
		Company                      struct {
			AdminGraphqlAPIID            shopify.ID `json:"admin_graphql_api_id"`
			CreatedAt                    time.Time  `json:"created_at"`
			CustomerSince                string     `json:"customer_since"`
			ExternalID                   string     `json:"external_id"`
			MainContactAdminGraphqlAPIID shopify.ID `json:"main_contact_admin_graphql_api_id"`
			Name                         string     `json:"name"`
			Note                         string     `json:"note"`
			UpdatedAt                    time.Time  `json:"updated_at"`
		} `json:"company"`
		CreatedAt       time.Time        `json:"created_at"`
		ExternalID      string           `json:"external_id"`
		Locale          string           `json:"locale"`
		Name            string           `json:"name"`
		Note            string           `json:"note"`
		Phone           string           `json:"phone"`
		ShippingAddress *shopify.Address `json:"shipping_address"`
		TaxExemptions   []string         `json:"tax_exemptions"`
		TaxRegistration struct {
			TaxID string `json:"tax_id"`
		} `json:"tax_registration"`
		TaxSettings struct {
			TaxExempt         interface{} `json:"tax_exempt"`
			TaxExemptions     []string    `json:"tax_exemptions"`
			TaxRegistrationID string      `json:"tax_registration_id"`
		} `json:"tax_settings"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"company_location"`
}
