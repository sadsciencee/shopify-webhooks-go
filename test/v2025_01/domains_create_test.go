package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbDomainsCreate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewDomainsCreate(w)
	if err != nil {
		t.Fatalf("failed to parse DomainsCreate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestDomainsCreate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.DomainsCreate, cbDomainsCreate, 38)
}
