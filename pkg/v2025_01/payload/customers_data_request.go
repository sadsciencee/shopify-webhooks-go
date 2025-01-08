package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersDataRequest(webhook webhook.ValidWebhook) (*CustomersDataRequest, error) {
	payload := &CustomersDataRequestPayload{}
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
	return &CustomersDataRequest{
		info: info,
		data: *payload,
	}, nil
}

type CustomersDataRequest struct {
	info webhook.Info
	data CustomersDataRequestPayload
}

func (webhook *CustomersDataRequest) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersDataRequest) GetData() (CustomersDataRequestPayload, error) {
	return webhook.data, nil
}

type CustomersDataRequestPayload struct {
	Customer struct {
		Email string `json:"email"`
		ID    int64  `json:"id"`
		Phone string `json:"phone"`
	} `json:"customer"`
	DataRequest struct {
		ID int64 `json:"id"`
	} `json:"data_request"`
	OrdersRequested []int64 `json:"orders_requested"`
	ShopDomain      string  `json:"shop_domain"`
	ShopID          int64   `json:"shop_id"`
}
