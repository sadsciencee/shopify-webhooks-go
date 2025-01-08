package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbSellingPlanGroupsUpdate(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewSellingPlanGroupsUpdate(w)
	if err != nil {
		t.Fatalf("failed to parse SellingPlanGroupsUpdate: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestSellingPlanGroupsUpdate(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.SellingPlanGroupsUpdate, cbSellingPlanGroupsUpdate, 175)
}
