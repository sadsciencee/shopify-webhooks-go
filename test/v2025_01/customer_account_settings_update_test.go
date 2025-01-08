package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbCustomerAccountSettingsUpdate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewCustomerAccountSettingsUpdate(w)
	if err != nil {
		t.Fatalf("failed to parse CustomerAccountSettingsUpdate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestCustomerAccountSettingsUpdate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.CustomerAccountSettingsUpdate, cbCustomerAccountSettingsUpdate, 50)
}
