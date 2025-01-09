package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCollectionPublicationsUpdate(webhook webhook.ValidWebhook) (*CollectionPublicationsUpdate, error) {
	payload := &CollectionPublicationsUpdatePayload{}
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
	return &CollectionPublicationsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CollectionPublicationsUpdate struct {
	info webhook.Info
	data CollectionPublicationsUpdatePayload
}

func (webhook *CollectionPublicationsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionPublicationsUpdate) GetData() (CollectionPublicationsUpdatePayload, error) {
	return webhook.data, nil
}

type CollectionPublicationsUpdatePayload struct {
	CollectionID  int64       `json:"collection_id"`
	CreatedAt     interface{} `json:"created_at"`
	ID            *int64      `json:"id"`
	PublicationID interface{} `json:"publication_id"`
	Published     bool        `json:"published"`
	PublishedAt   time.Time   `json:"published_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
