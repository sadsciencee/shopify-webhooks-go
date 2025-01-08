package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewLocationsDelete(webhook webhook.ValidWebhook) (*LocationsDelete, error) {
	payload := &LocationsDeletePayload{}
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
	return &LocationsDelete{
		info: info,
		data: *payload,
	}, nil
}

type LocationsDelete struct {
	info webhook.Info
	data LocationsDeletePayload
}

func (webhook *LocationsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *LocationsDelete) GetData() (LocationsDeletePayload, error) {
	return webhook.data, nil
}

type LocationsDeletePayload struct {
	ID int64 `json:"id"`
}
