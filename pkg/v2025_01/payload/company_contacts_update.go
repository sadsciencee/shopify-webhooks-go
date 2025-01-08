package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompanyContactsUpdate(webhook webhook.ValidWebhook) (*CompanyContactsUpdate, error) {
	payload := &CompanyContactsUpdatePayload{}
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
	return &CompanyContactsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CompanyContactsUpdate struct {
	info webhook.Info
	data CompanyContactsUpdatePayload
}

func (webhook *CompanyContactsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompanyContactsUpdate) GetData() (CompanyContactsUpdatePayload, error) {
	return webhook.data, nil
}

type CompanyContactsUpdatePayload struct {
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
