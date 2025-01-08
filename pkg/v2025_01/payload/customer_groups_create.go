package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomerGroupsCreate(webhook webhook.ValidWebhook) (*CustomerGroupsCreate, error) {
	payload := &CustomerGroupsCreatePayload{}
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
	return &CustomerGroupsCreate{
		info: info,
		data: *payload,
	}, nil
}

type CustomerGroupsCreate struct {
	info webhook.Info
	data CustomerGroupsCreatePayload
}

func (webhook *CustomerGroupsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerGroupsCreate) GetData() (CustomerGroupsCreatePayload, error) {
	return webhook.data, nil
}

type CustomerGroupsCreatePayload struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Query     string    `json:"query"`
	UpdatedAt time.Time `json:"updated_at"`
}
