package v2025_01__test

import (
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"github.com/sadsciencee/shopify-webhooks-go/test"
	"testing"
)

func cbCompanyContactRolesRevoke(t *testing.T, w *webhook.UnknownWebhook) error {
	webhook, err := payload.NewCompanyContactRolesRevoke(w)
	if err != nil {
		t.Fatalf("failed to parse CompanyContactRolesRevoke: %v", err)
	}
	_, err = webhook.GetData()
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	return nil
}

func TestCompanyContactRolesRevoke(t *testing.T) {
	t.Parallel()
	test.WebhookHelper(t, webhook.CompanyContactRolesRevoke, cbCompanyContactRolesRevoke, 54)
}
