package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbOrdersCancelled(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewOrdersCancelled(w)
	if err != nil {
		t.Fatalf("failed to parse OrdersCancelled: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestOrdersCancelled(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.OrdersCancelled, cbOrdersCancelled, 44)
}
