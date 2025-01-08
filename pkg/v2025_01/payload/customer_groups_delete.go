package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomerGroupsDelete(webhook webhook.ValidWebhook) (*CustomerGroupsDelete, error) {
	payload := &CustomerGroupsDeletePayload{}
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
	return &CustomerGroupsDelete{
		info: info,
		data: *payload,
	}, nil
}

type CustomerGroupsDelete struct {
	info webhook.Info
	data CustomerGroupsDeletePayload
}

func (webhook *CustomerGroupsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerGroupsDelete) GetData() (CustomerGroupsDeletePayload, error) {
	return webhook.data, nil
}

type CustomerGroupsDeletePayload struct {
	ID int64 `json:"id"`
}
