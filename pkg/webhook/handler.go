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
	shopifyTopicHeader       = "X-Shopify-Topic"
	shopifyChecksumHeader    = "X-Shopify-Hmac-Sha256"
	shopifyShopDomainHeader  = "X-Shopify-Shop-Domain"
	shopifyApiVersionHeader  = "X-Shopify-API-Version"
	shopifyWebhookIdHeader   = "X-Shopify-Webhook-Id"
	shopifyTriggeredAtHeader = "X-Shopify-Triggered-At"
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

type CallbackHandler func(webhook *UnknownWebhook) error
type CallbackErrorHandler func(w http.ResponseWriter, err error)

// Handler represents the main handler for Shopify webhooks
type Handler struct {
	apiSecret    string
	handler      CallbackHandler
	errorHandler CallbackErrorHandler
}

// NewHandler creates a new webhook handler with default error handling
func NewHandler(apiSecret string, handler CallbackHandler) *Handler {
	return NewHandlerWithErrorHandler(apiSecret, handler, nil)
}

// NewHandlerWithErrorHandler creates a new webhook handler with custom error handling
func NewHandlerWithErrorHandler(apiSecret string, handler CallbackHandler, errorHandler CallbackErrorHandler) *Handler {
	if errorHandler == nil {
		errorHandler = func(w http.ResponseWriter, err error) {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}

	return &Handler{
		apiSecret:    apiSecret,
		handler:      handler,
		errorHandler: errorHandler,
	}
}

// ServeHTTP implements the http.Handler interface
func (wh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	webhook, err := validateRequest(r, wh.apiSecret)
	if err != nil {
		fmt.Printf("Invalid request: %v\n", err.Error())
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	if err := wh.handler(webhook); err != nil {
		wh.errorHandler(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// validateRequest validates the incoming webhook request using the provided Shopify API secret against the HMAC signature
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

// parseWebhookInfo creates a standardized Info struct from the incoming webhook request and validates the topic
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
