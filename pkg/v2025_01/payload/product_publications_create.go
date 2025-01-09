package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewProductPublicationsCreate(webhook webhook.ValidWebhook) (*ProductPublicationsCreate, error) {
	payload := &ProductPublicationsCreatePayload{}
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
	return &ProductPublicationsCreate{
		info: info,
		data: *payload,
	}, nil
}

type ProductPublicationsCreate struct {
	info webhook.Info
	data ProductPublicationsCreatePayload
}

func (webhook *ProductPublicationsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductPublicationsCreate) GetData() (ProductPublicationsCreatePayload, error) {
	return webhook.data, nil
}

type ProductPublicationsCreatePayload struct {
	CreatedAt     *time.Time  `json:"created_at"`
	ID            *string     `json:"id"`
	ProductID     int64       `json:"product_id"`
	PublicationID interface{} `json:"publication_id"`
	Published     bool        `json:"published"`
	PublishedAt   *time.Time  `json:"published_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
}
