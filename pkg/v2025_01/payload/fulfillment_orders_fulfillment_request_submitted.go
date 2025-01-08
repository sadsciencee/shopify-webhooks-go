package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewFulfillmentOrdersFulfillmentRequestSubmitted(webhook webhook.ValidWebhook) (*FulfillmentOrdersFulfillmentRequestSubmitted, error) {
	payload := &FulfillmentOrdersFulfillmentRequestSubmittedPayload{}
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
	return &FulfillmentOrdersFulfillmentRequestSubmitted{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentOrdersFulfillmentRequestSubmitted struct {
	info webhook.Info
	data FulfillmentOrdersFulfillmentRequestSubmittedPayload
}

func (webhook *FulfillmentOrdersFulfillmentRequestSubmitted) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentOrdersFulfillmentRequestSubmitted) GetData() (FulfillmentOrdersFulfillmentRequestSubmittedPayload, error) {
	return webhook.data, nil
}

type FulfillmentOrdersFulfillmentRequestSubmittedPayload struct {
	FulfillmentOrderMerchantRequest struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"fulfillment_order_merchant_request"`
	OriginalFulfillmentOrder struct {
		ID            string `json:"id"`
		RequestStatus string `json:"request_status"`
		Status        string `json:"status"`
	} `json:"original_fulfillment_order"`
	SubmittedFulfillmentOrder struct {
		ID            string `json:"id"`
		RequestStatus string `json:"request_status"`
		Status        string `json:"status"`
	} `json:"submitted_fulfillment_order"`
}
