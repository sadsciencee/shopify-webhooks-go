package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewCollectionListingsAdd(webhook webhook.ValidWebhook) (*CollectionListingsAdd, error) {
	payload := &CollectionListingsAddPayload{}
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
	return &CollectionListingsAdd{
		info: info,
		data: *payload,
	}, nil
}

type CollectionListingsAdd struct {
	info webhook.Info
	data CollectionListingsAddPayload
}

func (webhook *CollectionListingsAdd) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CollectionListingsAdd) GetData() (CollectionListingsAddPayload, error) {
	return webhook.data, nil
}

type CollectionListingsAddPayload struct {
	CollectionListing struct {
		BodyHTML            string      `json:"body_html"`
		CollectionID        int64       `json:"collection_id"`
		DefaultProductImage interface{} `json:"default_product_image"`
		Handle              string      `json:"handle"`
		Image               interface{} `json:"image"`
		PublishedAt         time.Time   `json:"published_at"`
		SortOrder           interface{} `json:"sort_order"`
		Title               string      `json:"title"`
		UpdatedAt           time.Time   `json:"updated_at"`
	} `json:"collection_listing"`
}
