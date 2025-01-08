package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCompaniesCreate(webhook webhook.ValidWebhook) (*CompaniesCreate, error) {
	payload := &CompaniesCreatePayload{}
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
	return &CompaniesCreate{
		info: info,
		data: *payload,
	}, nil
}

type CompaniesCreate struct {
	info webhook.Info
	data CompaniesCreatePayload
}

func (webhook *CompaniesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CompaniesCreate) GetData() (CompaniesCreatePayload, error) {
	return webhook.data, nil
}

type CompaniesCreatePayload struct {
	AdminGraphqlAPIID            string    `json:"admin_graphql_api_id"`
	CreatedAt                    time.Time `json:"created_at"`
	CustomerSince                string    `json:"customer_since"`
	ExternalID                   string    `json:"external_id"`
	MainContactAdminGraphqlAPIID string    `json:"main_contact_admin_graphql_api_id"`
	Name                         string    `json:"name"`
	Note                         string    `json:"note"`
	UpdatedAt                    time.Time `json:"updated_at"`
}
