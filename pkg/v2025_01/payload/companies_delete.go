package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompaniesDelete(webhook webhook.ValidWebhook) (*CompaniesDelete, error) {
	payload := &CompaniesDeletePayload{}
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
	return &CompaniesDelete{
		info: info,
		data: *payload,
	}, nil
}

type CompaniesDelete struct {
	info webhook.Info
	data CompaniesDeletePayload
}

func (webhook *CompaniesDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompaniesDelete) GetData() (CompaniesDeletePayload, error) {
	return webhook.data, nil
}

type CompaniesDeletePayload struct {
	AdminGraphqlAPIID            shopify.ID `json:"admin_graphql_api_id"`
	CreatedAt                    time.Time  `json:"created_at"`
	CustomerSince                string     `json:"customer_since"`
	ExternalID                   string     `json:"external_id"`
	MainContactAdminGraphqlAPIID shopify.ID `json:"main_contact_admin_graphql_api_id"`
	Name                         string     `json:"name"`
	Note                         string     `json:"note"`
	UpdatedAt                    time.Time  `json:"updated_at"`
}
