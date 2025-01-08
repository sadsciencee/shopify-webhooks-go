package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompanyContactsDelete(webhook webhook.ValidWebhook) (*CompanyContactsDelete, error) {
	payload := &CompanyContactsDeletePayload{}
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
	return &CompanyContactsDelete{
		info: info,
		data: *payload,
	}, nil
}

type CompanyContactsDelete struct {
	info webhook.Info
	data CompanyContactsDeletePayload
}

func (webhook *CompanyContactsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompanyContactsDelete) GetData() (CompanyContactsDeletePayload, error) {
	return webhook.data, nil
}

type CompanyContactsDeletePayload struct {
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
}
