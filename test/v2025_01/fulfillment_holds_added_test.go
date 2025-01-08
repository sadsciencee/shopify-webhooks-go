package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbFulfillmentHoldsAdded(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewFulfillmentHoldsAdded(w)
	if err != nil {
		t.Fatalf("failed to parse FulfillmentHoldsAdded: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestFulfillmentHoldsAdded(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.FulfillmentHoldsAdded, cbFulfillmentHoldsAdded, 154)
}
