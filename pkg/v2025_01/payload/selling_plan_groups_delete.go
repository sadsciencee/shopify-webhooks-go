package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSellingPlanGroupsDelete(webhook webhook.ValidWebhook) (*SellingPlanGroupsDelete, error) {
	payload := &SellingPlanGroupsDeletePayload{}
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
	return &SellingPlanGroupsDelete{
		info: info,
		data: *payload,
	}, nil
}

type SellingPlanGroupsDelete struct {
	info webhook.Info
	data SellingPlanGroupsDeletePayload
}

func (webhook *SellingPlanGroupsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SellingPlanGroupsDelete) GetData() (SellingPlanGroupsDeletePayload, error) {
	return webhook.data, nil
}

type SellingPlanGroupsDeletePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	ID                int64      `json:"id"`
}
