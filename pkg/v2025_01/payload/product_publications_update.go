package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewProductPublicationsUpdate(webhook webhook.ValidWebhook) (*ProductPublicationsUpdate, error) {
	payload := &ProductPublicationsUpdatePayload{}
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
	return &ProductPublicationsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ProductPublicationsUpdate struct {
	info webhook.Info
	data ProductPublicationsUpdatePayload
}

func (webhook *ProductPublicationsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductPublicationsUpdate) GetData() (ProductPublicationsUpdatePayload, error) {
	return webhook.data, nil
}

type ProductPublicationsUpdatePayload struct {
	CreatedAt     *time.Time  `json:"created_at"`
	ID            string      `json:"id"`
	ProductID     int64       `json:"product_id"`
	PublicationID interface{} `json:"publication_id"`
	Published     bool        `json:"published"`
	PublishedAt   time.Time   `json:"published_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
