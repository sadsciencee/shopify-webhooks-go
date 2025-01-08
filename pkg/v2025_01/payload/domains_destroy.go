package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewDomainsDestroy(webhook webhook.ValidWebhook) (*DomainsDestroy, error) {
	payload := &DomainsDestroyPayload{}
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
	return &DomainsDestroy{
		info: info,
		data: *payload,
	}, nil
}

type DomainsDestroy struct {
	info webhook.Info
	data DomainsDestroyPayload
}

func (webhook *DomainsDestroy) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DomainsDestroy) GetData() (DomainsDestroyPayload, error) {
	return webhook.data, nil
}

type DomainsDestroyPayload struct {
	Host         string               `json:"host"`
	ID           int64                `json:"id"`
	Localization shopify.Localization `json:"localization"`
	SslEnabled   bool                 `json:"ssl_enabled"`
}
