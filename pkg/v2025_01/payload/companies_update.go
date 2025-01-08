package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompaniesUpdate(webhook webhook.ValidWebhook) (*CompaniesUpdate, error) {
	payload := &CompaniesUpdatePayload{}
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
	return &CompaniesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CompaniesUpdate struct {
	info webhook.Info
	data CompaniesUpdatePayload
}

func (webhook *CompaniesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompaniesUpdate) GetData() (CompaniesUpdatePayload, error) {
	return webhook.data, nil
}

type CompaniesUpdatePayload struct {
	AdminGraphqlAPIID            string    `json:"admin_graphql_api_id"`
	CreatedAt                    time.Time `json:"created_at"`
	CustomerSince                string    `json:"customer_since"`
	ExternalID                   string    `json:"external_id"`
	MainContactAdminGraphqlAPIID string    `json:"main_contact_admin_graphql_api_id"`
	Name                         string    `json:"name"`
	Note                         string    `json:"note"`
	UpdatedAt                    time.Time `json:"updated_at"`
}
