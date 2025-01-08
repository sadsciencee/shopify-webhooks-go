package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomersCreate(webhook webhook.ValidWebhook) (*CustomersCreate, error) {
	payload := &CustomersCreatePayload{}
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
	return &CustomersCreate{
		info: info,
		data: *payload,
	}, nil
}

type CustomersCreate struct {
	info webhook.Info
	data CustomersCreatePayload
}

func (webhook *CustomersCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersCreate) GetData() (CustomersCreatePayload, error) {
	return webhook.data, nil
}

type CustomersCreatePayload struct {
	Addresses           []interface{}                   `json:"addresses"`
	AdminGraphqlAPIID   string                          `json:"admin_graphql_api_id"`
	CreatedAt           time.Time                       `json:"created_at"`
	Currency            string                          `json:"currency"`
	DefaultAddress      *shopify.CustomerDefaultAddress `json:"default_address"`
	Email               string                          `json:"email"`
	FirstName           string                          `json:"first_name"`
	ID                  int64                           `json:"id"`
	LastName            string                          `json:"last_name"`
	MultipassIdentifier interface{}                     `json:"multipass_identifier"`
	Note                string                          `json:"note"`
	Phone               interface{}                     `json:"phone"`
	State               string                          `json:"state"`
	TaxExempt           bool                            `json:"tax_exempt"`
	TaxExemptions       []interface{}                   `json:"tax_exemptions"`
	UpdatedAt           time.Time                       `json:"updated_at"`
	VerifiedEmail       bool                            `json:"verified_email"`
}
