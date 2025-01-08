package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewDomainsCreate(webhook webhook.ValidWebhook) (*DomainsCreate, error) {
	payload := &DomainsCreatePayload{}
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
	return &DomainsCreate{
		info: info,
		data: *payload,
	}, nil
}

type DomainsCreate struct {
	info webhook.Info
	data DomainsCreatePayload
}

func (webhook *DomainsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DomainsCreate) GetData() (DomainsCreatePayload, error) {
	return webhook.data, nil
}

type DomainsCreatePayload struct {
	Host         string `json:"host"`
	ID           int64  `json:"id"`
	Localization struct {
		AlternateLocales []interface{} `json:"alternate_locales"`
		Country          interface{}   `json:"country"`
		DefaultLocale    string        `json:"default_locale"`
	} `json:"localization"`
	SslEnabled bool `json:"ssl_enabled"`
}
