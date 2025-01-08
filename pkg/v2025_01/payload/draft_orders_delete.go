package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewDraftOrdersDelete(webhook webhook.ValidWebhook) (*DraftOrdersDelete, error) {
	payload := &DraftOrdersDeletePayload{}
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
	return &DraftOrdersDelete{
		info: info,
		data: *payload,
	}, nil
}

type DraftOrdersDelete struct {
	info webhook.Info
	data DraftOrdersDeletePayload
}

func (webhook *DraftOrdersDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DraftOrdersDelete) GetData() (DraftOrdersDeletePayload, error) {
	return webhook.data, nil
}

type DraftOrdersDeletePayload struct {
	ID int64 `json:"id"`
}
