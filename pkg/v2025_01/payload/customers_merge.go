package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewCustomersMerge(webhook webhook.ValidWebhook) (*CustomersMerge, error) {
	payload := &CustomersMergePayload{}
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
	return &CustomersMerge{
		info: info,
		data: *payload,
	}, nil
}

type CustomersMerge struct {
	info webhook.Info
	data CustomersMergePayload
}

func (webhook *CustomersMerge) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *CustomersMerge) GetData() (CustomersMergePayload, error) {
	return webhook.data, nil
}

type CustomersMergePayload struct {
	AdminGraphqlAPICustomerDeletedID shopify.ID  `json:"admin_graphql_api_customer_deleted_id"`
	AdminGraphqlAPICustomerKeptID    shopify.ID  `json:"admin_graphql_api_customer_kept_id"`
	AdminGraphqlAPIJobID             interface{} `json:"admin_graphql_api_job_id"`
	Errors                           []struct {
		CustomerIds []int64 `json:"customer_ids"`
		Field       string  `json:"field"`
		Message     string  `json:"message"`
	} `json:"errors"`
	Status string `json:"status"`
}
