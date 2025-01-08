package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewDomainsUpdate(webhook webhook.ValidWebhook) (*DomainsUpdate, error) {
	payload := &DomainsUpdatePayload{}
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
	return &DomainsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type DomainsUpdate struct {
	info webhook.Info
	data DomainsUpdatePayload
}

func (webhook *DomainsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DomainsUpdate) GetData() (DomainsUpdatePayload, error) {
	return webhook.data, nil
}

type DomainsUpdatePayload struct {
	Host         string `json:"host"`
	ID           int64  `json:"id"`
	Localization struct {
		AlternateLocales []interface{} `json:"alternate_locales"`
		Country          interface{}   `json:"country"`
		DefaultLocale    string        `json:"default_locale"`
	} `json:"localization"`
	SslEnabled bool `json:"ssl_enabled"`
}
