package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewSellingPlanGroupsUpdate(webhook webhook.ValidWebhook) (*SellingPlanGroupsUpdate, error) {
	payload := &SellingPlanGroupsUpdatePayload{}
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
	return &SellingPlanGroupsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type SellingPlanGroupsUpdate struct {
	info webhook.Info
	data SellingPlanGroupsUpdatePayload
}

func (webhook *SellingPlanGroupsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SellingPlanGroupsUpdate) GetData() (SellingPlanGroupsUpdatePayload, error) {
	return webhook.data, nil
}

type SellingPlanGroupsUpdatePayload struct {
	AdminGraphqlAPIApp string                `json:"admin_graphql_api_app"`
	AdminGraphqlAPIID  string                `json:"admin_graphql_api_id"`
	AppID              interface{}           `json:"app_id"`
	Description        interface{}           `json:"description"`
	ID                 int64                 `json:"id"`
	MerchantCode       string                `json:"merchant_code"`
	Name               string                `json:"name"`
	Options            []string              `json:"options"`
	Position           interface{}           `json:"position"`
	ProductVariants    []interface{}         `json:"product_variants"`
	Products           []interface{}         `json:"products"`
	SellingPlans       []shopify.SellingPlan `json:"selling_plans"`
	Summary            string                `json:"summary"`
}
