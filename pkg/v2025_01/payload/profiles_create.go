package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProfilesCreate(webhook webhook.ValidWebhook) (*ProfilesCreate, error) {
	payload := &ProfilesCreatePayload{}
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
	return &ProfilesCreate{
		info: info,
		data: *payload,
	}, nil
}

type ProfilesCreate struct {
	info webhook.Info
	data ProfilesCreatePayload
}

func (webhook *ProfilesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProfilesCreate) GetData() (ProfilesCreatePayload, error) {
	return webhook.data, nil
}

type ProfilesCreatePayload struct {
	ID int64 `json:"id"`
}
