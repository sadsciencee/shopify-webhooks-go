package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbMetaobjectsUpdate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewMetaobjectsUpdate(w)
	if err != nil {
		t.Fatalf("failed to parse MetaobjectsUpdate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestMetaobjectsUpdate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.MetaobjectsUpdate, cbMetaobjectsUpdate, 73)
}
