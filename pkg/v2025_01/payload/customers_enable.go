package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomersEnable(webhook webhook.ValidWebhook) (*CustomersEnable, error) {
	payload := &CustomersEnablePayload{}
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
	return &CustomersEnable{
		info: info,
		data: *payload,
	}, nil
}

type CustomersEnable struct {
	info webhook.Info
	data CustomersEnablePayload
}

func (webhook *CustomersEnable) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersEnable) GetData() (CustomersEnablePayload, error) {
	return webhook.data, nil
}

type CustomersEnablePayload struct {
	Addresses           []interface{}                   `json:"addresses"`
	AdminGraphqlAPIID   shopify.ID                      `json:"admin_graphql_api_id"`
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
