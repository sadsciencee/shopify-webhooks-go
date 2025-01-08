package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersCancellationRequestSubmitted(webhook webhook.ValidWebhook) (*FulfillmentOrdersCancellationRequestSubmitted, error) {
	payload := &FulfillmentOrdersCancellationRequestSubmittedPayload{}
	info, err := webhook.GetInfo()
	if err != nil {
		return nil, err
	}
	reader, err := webhook.GetReader()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(payload); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}
	return &FulfillmentOrdersCancellationRequestSubmitted{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersCancellationRequestSubmitted struct {
	info webhook.Info
	data FulfillmentOrdersCancellationRequestSubmittedPayload
}

func (webhook *FulfillmentOrdersCancellationRequestSubmitted) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersCancellationRequestSubmitted) GetData() (FulfillmentOrdersCancellationRequestSubmittedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersCancellationRequestSubmittedPayload struct {
	FulfillmentOrder struct {
		ID            string `json:"id"`
		RequestStatus string `json:"request_status"`
		Status        string `json:"status"`
	} `json:"fulfillment_order"`
	FulfillmentOrderMerchantRequest struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"fulfillment_order_merchant_request"`
}
