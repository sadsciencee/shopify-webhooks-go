package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomerAccountSettingsUpdate(webhook webhook.ValidWebhook) (*CustomerAccountSettingsUpdate, error) {
	payload := &CustomerAccountSettingsUpdatePayload{}
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
	return &CustomerAccountSettingsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type CustomerAccountSettingsUpdate struct {
	info webhook.Info
	data CustomerAccountSettingsUpdatePayload
}

func (webhook *CustomerAccountSettingsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomerAccountSettingsUpdate) GetData() (CustomerAccountSettingsUpdatePayload, error) {
	return webhook.data, nil
}

type CustomerAccountSettingsUpdatePayload struct {
	CustomerAccountsVersion                  string  `json:"customer_accounts_version"`
	LoginLinksVisibleOnStorefrontAndCheckout bool    `json:"login_links_visible_on_storefront_and_checkout"`
	LoginRequiredAtCheckout                  bool    `json:"login_required_at_checkout"`
	URL                                      *string `json:"url"`
}
