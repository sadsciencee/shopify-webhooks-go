package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

const (
	shopifyTopicHeader       = "X-Shopify-Topic" // orders/create
	shopifyChecksumHeader    = "X-Shopify-Hmac-Sha256"
	shopifyShopDomainHeader  = "X-Shopify-Shop-Domain" // {shop}.myshopify.com
	shopifyApiVersionHeader  = "X-Shopify-API-Version" // 2025-01
	shopifyWebhookIdHeader   = "X-Shopify-Webhook-Id"
	shopifyTriggeredAtHeader = "X-Shopify-Triggered-At" // 2023-03-29T18:00:27.877041743Z
	shopifyEventIdHeader     = "X-Shopify-Event-Id"
)

type UnknownWebhook struct {
	info Info
	data io.Reader
}

func (uw *UnknownWebhook) GetReader() (io.Reader, error) {
	return uw.data, nil
}

func (uw *UnknownWebhook) GetInfo() (Info, error) {
	return uw.info, nil
}

func (uw *UnknownWebhook) LogData() error {
	if uw.data == nil {
		return fmt.Errorf("webhook data is nil")
	}

	buf := new(bytes.Buffer)

	if _, err := io.Copy(buf, uw.data); err != nil {
		return fmt.Errorf("failed to read webhook data: %v", err)
	}

	uw.data = bytes.NewReader(buf.Bytes())

	fmt.Printf("Webhook data: %s\n", buf.String())

	return nil
}

type ValidationError struct {
	code    string
	message string
	shop    string
	topic   Topic
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Webhook error: %s, code: %s, shop: %s, topic: %s", e.message, e.code, e.shop, e.topic)
}

type Handler func(webhook *UnknownWebhook) error
type CallbackErrorHandler func(w http.ResponseWriter, err error)

// NewMiddleware creates a handler that validates incoming Shopify webhooks.
// It accepts an API secret for validating the webhook signature and a handler function
// for processing the webhook payload. Any errors returned by the handler will default
// to a 500 Internal Server Error response.
func NewMiddleware(apiSecret string, handler Handler) func(http.Handler) http.Handler {
	return NewMiddlewareWithErrorHandler(apiSecret, handler, nil)
}

// NewMiddlewareWithErrorHandler implements NewMiddleware with custom error handling.
// If errorHandler is nil, errors default to a 500 Internal Server Error response.
func NewMiddlewareWithErrorHandler(apiSecret string, handler Handler, errorHandler CallbackErrorHandler) func(http.Handler) http.Handler {
	if errorHandler == nil {
		errorHandler = func(w http.ResponseWriter, err error) {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				fmt.Println("Method not allowed")
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			webhook, err := validateRequest(r, apiSecret)
			if err != nil {
				fmt.Printf("Invalid request: %v\n", err.Error())
				http.Error(w, "Invalid signature", http.StatusUnauthorized)
				return
			}

			if err := handler(webhook); err != nil {
				errorHandler(w, err)
				return
			}

			w.WriteHeader(http.StatusOK)
		})
	}
}

func validateRequest(httpRequest *http.Request, shopifyApiSecret string) (*UnknownWebhook, error) {
	fmt.Println("Request received")
	if httpRequest == nil {
		return nil, newValidationError("INVALID_REQUEST", "Request is nil", &Info{})
	}

	if httpRequest.Body == nil {
		return nil, newValidationError("INVALID_BODY", "Request body is nil", &Info{})
	}

	requestBody, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, newValidationError("READ_ERROR", fmt.Sprintf("Failed to read body: %v", err), &Info{})
	}

	shopifySha256 := httpRequest.Header.Get(shopifyChecksumHeader)
	if shopifySha256 == "" {
		return nil, newValidationError("MISSING_HMAC", "Missing HMAC header", &Info{})
	}

	actualMac := []byte(shopifySha256)

	if shopifyApiSecret == "" {
		return nil, newValidationError("MISSING_SECRET", "API secret not configured", &Info{})
	}

	mac := hmac.New(sha256.New, []byte(shopifyApiSecret))
	httpRequest.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	mac.Write(requestBody)
	macSum := mac.Sum(nil)
	expectedMac := []byte(base64.StdEncoding.EncodeToString(macSum))

	info, err := parseWebhookInfo(httpRequest)
	if err != nil {
		return nil, newValidationError("INVALID_INFO", fmt.Sprintf("Failed to parse webhook info: %v", err), &Info{})
	}
	if !hmac.Equal(actualMac, expectedMac) {
		return nil, newValidationError("INVALID_HMAC", "Invalid HMAC", info)
	}

	return newUnknownWebhook(bytes.NewReader(requestBody), *info), nil
}

func newValidationError(code string, message string, info *Info) *ValidationError {
	return &ValidationError{
		code:    code,
		message: message,
		shop:    info.ShopDomain,
		topic:   info.Topic,
	}
}

func newUnknownWebhook(body io.Reader, info Info) *UnknownWebhook {
	return &UnknownWebhook{
		info: info,
		data: body,
	}
}

func parseWebhookInfo(r *http.Request) (*Info, error) {
	topic, err := ValidateTopic(r.Header.Get(shopifyTopicHeader))
	if err != nil {
		return nil, err
	}
	return &Info{
		Topic:       topic,
		ShopDomain:  r.Header.Get(shopifyShopDomainHeader),
		ApiVersion:  r.Header.Get(shopifyApiVersionHeader),
		WebhookId:   r.Header.Get(shopifyWebhookIdHeader),
		TriggeredAt: r.Header.Get(shopifyTriggeredAtHeader),
		EventId:     r.Header.Get(shopifyEventIdHeader),
	}, nil
}
