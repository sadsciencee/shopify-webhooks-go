# Usage
## Example
```go
package main

import (
	"context"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/payload"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// do not commit to source - use environment variables
	secret := "your_api_secret"

	// initialize a webhook handler. if you just want to return 500 for errors, use webhook.NewHandler(secret, yourHandler) instead
	webhookHandler := webhook.NewHandlerWithErrorHandler(secret, yourHandler, yourErrorHandler)
	mux := http.NewServeMux()
	// handler will validate the request and process webhooks
	mux.Handle("/webhook", webhookHandler)

	// from here on it's just a generic http server
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}

// used to process validated webhook requests
func yourHandler(wh *webhook.UnknownWebhook) error {
	info, err := wh.GetInfo()
	if err != nil {
		return err
	}
	switch info.Topic {
	// handle each topic you are interested in. the payload package has dedicated structs for each topic to unmarshall the data
	case webhook.OrdersCreate:
		// note - future versions will have a payload.NewOrdersUnwrap method to make this less verbose
		raw, err := payload.NewOrdersCreate(wh)
		if err != nil {
			return err
		}
		data, err := raw.GetData()
		if err != nil {
			return err
		}
		fmt.Printf("Order received with %s line items", len(data.LineItems))
		// most likely you will want to put these in a queue, shopify will be upset if you take longer than 300ms to respond with 200
	}
	if err != nil {
		return err
	}
	return nil
}

// optionally handle errors from your handler function. if not provided, a default 500 response will be used
// this specifically runs on *your* callback errors. if there is an error validating the webhook an error will always be returned to the http request
// you probably only need this if, for example, you're running shadow traffic on a new webhook topic and don't want shopify to flag a million failed webhooks if something goes wrong
func yourErrorHandler(w http.ResponseWriter, err error) {
	fmt.Printf("Error: %v\n", err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

```

## Versioning
Library versions are independent of Shopify's quarterly versions. As new Shopify version support is added, it will be added to the corresponding directory - i.e. `pkg/v<year>_<month>`

# Contribution
## Payloads
The payloads in v0.0.1 are generated from the JSON payload examples provided by Shopify.

All payload structs will unmarshall without errors when the default test webhook is triggered. If you find that a payload struct is invalid in production, PRs are welcomed and encouraged.

You may also utilize the custom payloads functionality described above for temporary fixes to payloads.

## Codegen
Initial payloads were generated with scripts in the cmd/ directory. This got me probably 80% of the way but I still manually cleaned things up to remove redundant types.

I'd love to double down on the codegen and get this completely automated. This would open the door to auto gen for partial webhooks and remove the need for custom payload callbacks.

Run these scripts in sequence to get payloads for 2025-01

- `go run github.com/sadsciencee/shopify-webhooks-go@latest payloads`
- `go run github.com/sadsciencee/shopify-webhooks-go@latest models`


# Shopify Reference Docs
- [Shopify Webhook Types](https://shopify.dev/docs/api/webhooks?reference=toml#list-of-topics-app)
- [Shopify Webhook Validation](https://shopify.dev/docs/apps/build/webhooks/subscribe/https#step-2-validate-the-origin-of-your-webhook-to-ensure-its-coming-from-shopify)


