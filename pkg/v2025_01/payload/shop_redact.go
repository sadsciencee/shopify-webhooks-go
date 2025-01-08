package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewShopRedact(webhook webhook.ValidWebhook) (*ShopRedact, error) {
	payload := &ShopRedactPayload{}
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
	return &ShopRedact{
		info: info,
		data: *payload,
	}, nil
}

type ShopRedact struct {
	info webhook.Info
	data ShopRedactPayload
}

func (webhook *ShopRedact) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ShopRedact) GetData() (ShopRedactPayload, error) {
	return webhook.data, nil
}

type ShopRedactPayload struct {
	ShopDomain string `json:"shop_domain"`
	ShopID     int64  `json:"shop_id"`
}
