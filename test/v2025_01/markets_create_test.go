package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbMarketsCreate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewMarketsCreate(w)
	if err != nil {
		t.Fatalf("failed to parse MarketsCreate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestMarketsCreate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.MarketsCreate, cbMarketsCreate, 180)
}
