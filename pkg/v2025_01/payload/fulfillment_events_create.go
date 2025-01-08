package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewFulfillmentEventsCreate(webhook webhook.ValidWebhook) (*FulfillmentEventsCreate, error) {
	payload := &FulfillmentEventsCreatePayload{}
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
	return &FulfillmentEventsCreate{
		info: info,
		data: *payload,
	}, nil
}

type FulfillmentEventsCreate struct {
	info webhook.Info
	data FulfillmentEventsCreatePayload
}

func (webhook *FulfillmentEventsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *FulfillmentEventsCreate) GetData() (FulfillmentEventsCreatePayload, error) {
	return webhook.data, nil
}

type FulfillmentEventsCreatePayload struct {
	Address1            interface{} `json:"address1"`
	AdminGraphqlAPIID   string      `json:"admin_graphql_api_id"`
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
