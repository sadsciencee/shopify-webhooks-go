package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProfilesDelete(webhook webhook.ValidWebhook) (*ProfilesDelete, error) {
	payload := &ProfilesDeletePayload{}
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
	return &ProfilesDelete{
		info: info,
		data: *payload,
	}, nil
}

type ProfilesDelete struct {
	info webhook.Info
	data ProfilesDeletePayload
}

func (webhook *ProfilesDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProfilesDelete) GetData() (ProfilesDeletePayload, error) {
	return webhook.data, nil
}

type ProfilesDeletePayload struct {
	ID int64 `json:"id"`
}
