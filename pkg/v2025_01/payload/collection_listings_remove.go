package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCollectionListingsRemove(webhook webhook.ValidWebhook) (*CollectionListingsRemove, error) {
	payload := &CollectionListingsRemovePayload{}
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
	return &CollectionListingsRemove{
		info: info,
		data: *payload,
	}, nil
}

type CollectionListingsRemove struct {
	info webhook.Info
	data CollectionListingsRemovePayload
}

func (webhook *CollectionListingsRemove) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionListingsRemove) GetData() (CollectionListingsRemovePayload, error) {
	return webhook.data, nil
}

type CollectionListingsRemovePayload struct {
	CollectionListing struct {
		CollectionID int64 `json:"collection_id"`
	} `json:"collection_listing"`
}
