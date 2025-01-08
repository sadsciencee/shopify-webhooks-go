package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReverseDeliveriesAttachDeliverable(webhook webhook.ValidWebhook) (*ReverseDeliveriesAttachDeliverable, error) {
	payload := &ReverseDeliveriesAttachDeliverablePayload{}
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
	return &ReverseDeliveriesAttachDeliverable{
		info: info,
		data: *payload,
	}, nil
}

type ReverseDeliveriesAttachDeliverable struct {
	info webhook.Info
	data ReverseDeliveriesAttachDeliverablePayload
}

func (webhook *ReverseDeliveriesAttachDeliverable) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReverseDeliveriesAttachDeliverable) GetData() (ReverseDeliveriesAttachDeliverablePayload, error) {
	return webhook.data, nil
}

type ReverseDeliveriesAttachDeliverablePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	ID                int64      `json:"id"`
	Return            struct {
		AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
		ID                int64      `json:"id"`
	} `json:"return"`
	ShippingDeliverable struct {
		Label struct {
			CreatedAt     interface{} `json:"created_at"`
			PublicFileURL interface{} `json:"public_file_url"`
		} `json:"label"`
		Tracking struct {
			CarrierName    string      `json:"carrier_name"`
			TrackingNumber string      `json:"tracking_number"`
			TrackingURL    interface{} `json:"tracking_url"`
		} `json:"tracking"`
	} `json:"shipping_deliverable"`
}
