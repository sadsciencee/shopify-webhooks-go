package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCustomerGroupsUpdate(webhook webhook.ValidWebhook) (*CustomerGroupsUpdate, error) {
	payload := &CustomerGroupsUpdatePayload{}
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
	return &CustomerGroupsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CustomerGroupsUpdate struct {
	info webhook.Info
	data CustomerGroupsUpdatePayload
}

func (webhook *CustomerGroupsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerGroupsUpdate) GetData() (CustomerGroupsUpdatePayload, error) {
	return webhook.data, nil
}

type CustomerGroupsUpdatePayload struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Query     string    `json:"query"`
	UpdatedAt time.Time `json:"updated_at"`
}
