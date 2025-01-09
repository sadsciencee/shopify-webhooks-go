package test

import (
	"context"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

type CallbackForTestUtil func(*testing.T, *webhook.UnknownWebhook) error

func WebhookHelper(t *testing.T, topic webhook.Topic, tc CallbackForTestUtil, index int) {
	secret := os.Getenv("SHOPIFY_CLIENT_SECRET")
	if secret == "" {
		t.Fatal("SHOPIFY_CLIENT_SECRET environment variable is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	port := 3000 + index
	addr := fmt.Sprintf(":%d", port)
	webhookURL := fmt.Sprintf("http://localhost:%d/webhook", port)

	results := make(chan error, 1)

	handler := func(webhook *webhook.UnknownWebhook) error {
		info, err := webhook.GetInfo()
		if err != nil {
			results <- fmt.Errorf("failed to get webhook info: %v", err)
			return nil
		}

		if info.Topic != topic {
			results <- fmt.Errorf("expected topic %s, got %s", topic, info.Topic)
			return nil
		}

		if err := tc(t, webhook); err != nil {
			results <- fmt.Errorf("webhook validation failed: %v", err)
			return nil
		}

		results <- nil
		close(results)
		return nil
	}

	errorHandler := func(w http.ResponseWriter, err error) {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	webhookHandler := webhook.NewHandlerWithErrorHandler(secret, handler, errorHandler)
	mux := http.NewServeMux()
	mux.Handle("/webhook", webhookHandler)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan struct{})
	defer close(done)

	go func() {
		t.Logf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			t.Errorf("Server error: %v", err)
		}
	}()

	go func() {
		select {
		case <-done:
			server.Shutdown(context.Background())
		case <-ctx.Done():
			server.Shutdown(context.Background())
		}
	}()

	time.Sleep(time.Second)

	cmd := exec.CommandContext(ctx, "shopify", "app", "webhook", "trigger",
		"--address", webhookURL,
		"--topic", string(topic),
		"--api-version", "2025-01",
		"--client-secret", secret)

	output, err := cmd.CombinedOutput()
	t.Logf("Webhook trigger command output: %s", string(output))
	if err != nil {
		t.Fatalf("Failed to execute webhook trigger command: %v\nOutput: %s", err, string(output))
	}

	select {
	case err := <-results:
		if err != nil {
			t.Error(err)
		}
		return
	case <-ctx.Done():
		t.Error("Timeout waiting for webhook")
	}
}
