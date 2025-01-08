package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewLocationsDeactivate(webhook webhook.ValidWebhook) (*LocationsDeactivate, error) {
	payload := &LocationsDeactivatePayload{}
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
	return &LocationsDeactivate{
		info: info,
		data: *payload,
	}, nil
}

type LocationsDeactivate struct {
	info webhook.Info
	data LocationsDeactivatePayload
}

func (webhook *LocationsDeactivate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *LocationsDeactivate) GetData() (LocationsDeactivatePayload, error) {
	return webhook.data, nil
}

type LocationsDeactivatePayload struct {
	Active            bool      `json:"active"`
	Address1          string    `json:"address1"`
	Address2          string    `json:"address2"`
	AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
	City              string    `json:"city"`
	Country           string    `json:"country"`
	CountryCode       string    `json:"country_code"`
	CountryName       string    `json:"country_name"`
	CreatedAt         time.Time `json:"created_at"`
	ID                int64     `json:"id"`
	Legacy            bool      `json:"legacy"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone"`
	Province          string    `json:"province"`
	ProvinceCode      string    `json:"province_code"`
	UpdatedAt         time.Time `json:"updated_at"`
	Zip               string    `json:"zip"`
}
