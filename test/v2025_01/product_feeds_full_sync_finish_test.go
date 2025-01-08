package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbProductFeedsFullSyncFinish(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewProductFeedsFullSyncFinish(w)
	if err != nil {
		t.Fatalf("failed to parse ProductFeedsFullSyncFinish: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestProductFeedsFullSyncFinish(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.ProductFeedsFullSyncFinish, cbProductFeedsFullSyncFinish, 68)
}
