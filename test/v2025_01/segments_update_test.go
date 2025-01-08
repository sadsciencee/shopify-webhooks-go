package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbSegmentsUpdate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewSegmentsUpdate(w)
	if err != nil {
		t.Fatalf("failed to parse SegmentsUpdate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestSegmentsUpdate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.SegmentsUpdate, cbSegmentsUpdate, 152)
}
