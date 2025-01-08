package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewBulkOperationsFinish(webhook webhook.ValidWebhook) (*BulkOperationsFinish, error) {
	payload := &BulkOperationsFinishPayload{}
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
	return &BulkOperationsFinish{
		info: info,
		data: *payload,
	}, nil
}

type BulkOperationsFinish struct {
	info webhook.Info
	data BulkOperationsFinishPayload
}

func (webhook *BulkOperationsFinish) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *BulkOperationsFinish) GetData() (BulkOperationsFinishPayload, error) {
	return webhook.data, nil
}

type BulkOperationsFinishPayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	CompletedAt       time.Time  `json:"completed_at"`
	CreatedAt         time.Time  `json:"created_at"`
	ErrorCode         *string    `json:"error_code"`
	Status            string     `json:"status"`
	Type              string     `json:"type"`
}
