package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbSubscriptionContractsCancel(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewSubscriptionContractsCancel(w)
	if err != nil {
		t.Fatalf("failed to parse SubscriptionContractsCancel: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestSubscriptionContractsCancel(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.SubscriptionContractsCancel, cbSubscriptionContractsCancel, 97)
}
