package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompanyContactsCreate(webhook webhook.ValidWebhook) (*CompanyContactsCreate, error) {
	payload := &CompanyContactsCreatePayload{}
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
	return &CompanyContactsCreate{
		info: info,
		data: *payload,
	}, nil
}

type CompanyContactsCreate struct {
	info webhook.Info
	data CompanyContactsCreatePayload
}

func (webhook *CompanyContactsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompanyContactsCreate) GetData() (CompanyContactsCreatePayload, error) {
	return webhook.data, nil
}

type CompanyContactsCreatePayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	Company           struct {
		AdminGraphqlAPIID            string    `json:"admin_graphql_api_id"`
		CreatedAt                    time.Time `json:"created_at"`
		CustomerSince                string    `json:"customer_since"`
		ExternalID                   string    `json:"external_id"`
		MainContactAdminGraphqlAPIID string    `json:"main_contact_admin_graphql_api_id"`
		Name                         string    `json:"name"`
		Note                         string    `json:"note"`
		UpdatedAt                    time.Time `json:"updated_at"`
	} `json:"company"`
	CreatedAt                 time.Time `json:"created_at"`
	CustomerAdminGraphqlAPIID string    `json:"customer_admin_graphql_api_id"`
	Locale                    string    `json:"locale"`
	Title                     string    `json:"title"`
	UpdatedAt                 time.Time `json:"updated_at"`
}
