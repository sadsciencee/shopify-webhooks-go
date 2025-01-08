package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbCustomersPurchasingSummary(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewCustomersPurchasingSummary(w)
	if err != nil {
		t.Fatalf("failed to parse CustomersPurchasingSummary: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestCustomersPurchasingSummary(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.CustomersPurchasingSummary, cbCustomersPurchasingSummary, 6)
}
