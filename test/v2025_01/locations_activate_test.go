package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbLocationsActivate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewLocationsActivate(w)
	if err != nil {
		t.Fatalf("failed to parse LocationsActivate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestLocationsActivate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.LocationsActivate, cbLocationsActivate, 168)
}
