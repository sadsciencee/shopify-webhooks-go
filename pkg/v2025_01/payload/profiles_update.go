package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewProfilesUpdate(webhook webhook.ValidWebhook) (*ProfilesUpdate, error) {
	payload := &ProfilesUpdatePayload{}
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
	return &ProfilesUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ProfilesUpdate struct {
	info webhook.Info
	data ProfilesUpdatePayload
}

func (webhook *ProfilesUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProfilesUpdate) GetData() (ProfilesUpdatePayload, error) {
	return webhook.data, nil
}

type ProfilesUpdatePayload struct {
	ID int64 `json:"id"`
}
