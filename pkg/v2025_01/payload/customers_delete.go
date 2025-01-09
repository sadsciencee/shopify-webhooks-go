package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersDelete(webhook webhook.ValidWebhook) (*CustomersDelete, error) {
	payload := &CustomersDeletePayload{}
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
	return &CustomersDelete{
		info: info,
		data: *payload,
	}, nil
}

type CustomersDelete struct {
	info webhook.Info
	data CustomersDeletePayload
}

func (webhook *CustomersDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersDelete) GetData() (CustomersDeletePayload, error) {
	return webhook.data, nil
}

type CustomersDeletePayload struct {
	Addresses         []interface{} `json:"addresses"`
	AdminGraphqlAPIID shopify.ID    `json:"admin_graphql_api_id"`
	ID                int64         `json:"id"`
	Phone             interface{}   `json:"phone"`
	TaxExemptions     []interface{} `json:"tax_exemptions"`
}
