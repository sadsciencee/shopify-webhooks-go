package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewChannelsDelete(webhook webhook.ValidWebhook) (*ChannelsDelete, error) {
	payload := &ChannelsDeletePayload{}
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
	return &ChannelsDelete{
		info: info,
		data: *payload,
	}, nil
}

type ChannelsDelete struct {
	info webhook.Info
	data ChannelsDeletePayload
}

func (webhook *ChannelsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ChannelsDelete) GetData() (ChannelsDeletePayload, error) {
	return webhook.data, nil
}

type ChannelsDeletePayload struct {
	ID string `json:"id"`
}
