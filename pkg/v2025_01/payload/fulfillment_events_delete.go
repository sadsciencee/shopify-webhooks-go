package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewFulfillmentEventsDelete(webhook webhook.ValidWebhook) (*FulfillmentEventsDelete, error) {
	payload := &FulfillmentEventsDeletePayload{}
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
	return &FulfillmentEventsDelete{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentEventsDelete struct {
	info webhook.Info
	data FulfillmentEventsDeletePayload
}

func (webhook *FulfillmentEventsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentEventsDelete) GetData() (FulfillmentEventsDeletePayload, error) {
	return webhook.data, nil
}

type FulfillmentEventsDeletePayload struct {
	Address1            interface{} `json:"address1"`
	AdminGraphqlAPIID   shopify.ID  `json:"admin_graphql_api_id"`
	City                interface{} `json:"city"`
	Country             string      `json:"country"`
	CreatedAt           time.Time   `json:"created_at"`
	EstimatedDeliveryAt interface{} `json:"estimated_delivery_at"`
	FulfillmentID       int64       `json:"fulfillment_id"`
	HappenedAt          time.Time   `json:"happened_at"`
	ID                  int64       `json:"id"`
	Latitude            interface{} `json:"latitude"`
	Longitude           interface{} `json:"longitude"`
	Message             string      `json:"message"`
	OrderID             int64       `json:"order_id"`
	Province            interface{} `json:"province"`
	ShopID              int64       `json:"shop_id"`
	Status              string      `json:"status"`
	UpdatedAt           time.Time   `json:"updated_at"`
	Zip                 interface{} `json:"zip"`
}
