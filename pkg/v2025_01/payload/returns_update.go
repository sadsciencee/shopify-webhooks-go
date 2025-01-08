package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsUpdate(webhook webhook.ValidWebhook) (*ReturnsUpdate, error) {
	payload := &ReturnsUpdatePayload{}
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
	return &ReturnsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsUpdate struct {
	info webhook.Info
	data ReturnsUpdatePayload
}

func (webhook *ReturnsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsUpdate) GetData() (ReturnsUpdatePayload, error) {
	return webhook.data, nil
}

type ReturnsUpdatePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	RestockingFees    struct {
		Removals []interface{} `json:"removals"`
		Updates  []interface{} `json:"updates"`
	} `json:"restocking_fees"`
	ReturnLineItems struct {
		Removals []struct {
			AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
			Delta             int64      `json:"delta"`
		} `json:"removals"`
	} `json:"return_line_items"`
	ReturnShippingFees struct {
		Removals []interface{} `json:"removals"`
		Updates  []interface{} `json:"updates"`
	} `json:"return_shipping_fees"`
}
