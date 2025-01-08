package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewLocationsActivate(webhook webhook.ValidWebhook) (*LocationsActivate, error) {
	payload := &LocationsActivatePayload{}
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
	return &LocationsActivate{
		info: info,
		data: *payload,
	}, nil
}

type LocationsActivate struct {
	info webhook.Info
	data LocationsActivatePayload
}

func (webhook *LocationsActivate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *LocationsActivate) GetData() (LocationsActivatePayload, error) {
	return webhook.data, nil
}

type LocationsActivatePayload struct {
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
