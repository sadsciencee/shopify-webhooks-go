package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbOrdersRiskAssessmentChanged(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewOrdersRiskAssessmentChanged(w)
	if err != nil {
		t.Fatalf("failed to parse OrdersRiskAssessmentChanged: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestOrdersRiskAssessmentChanged(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.OrdersRiskAssessmentChanged, cbOrdersRiskAssessmentChanged, 59)
}
