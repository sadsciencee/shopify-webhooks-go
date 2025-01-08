package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCollectionPublicationsDelete(webhook webhook.ValidWebhook) (*CollectionPublicationsDelete, error) {
	payload := &CollectionPublicationsDeletePayload{}
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
	return &CollectionPublicationsDelete{
		info: info,
		data: *payload,
	}, nil
}

type CollectionPublicationsDelete struct {
	info webhook.Info
	data CollectionPublicationsDeletePayload
}

func (webhook *CollectionPublicationsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionPublicationsDelete) GetData() (CollectionPublicationsDeletePayload, error) {
	return webhook.data, nil
}

type CollectionPublicationsDeletePayload struct {
	ID *int64 `json:"id"`
}
