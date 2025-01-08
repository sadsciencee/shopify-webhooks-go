package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbDisputesUpdate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewDisputesUpdate(w)
	if err != nil {
		t.Fatalf("failed to parse DisputesUpdate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestDisputesUpdate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.DisputesUpdate, cbDisputesUpdate, 11)
}
