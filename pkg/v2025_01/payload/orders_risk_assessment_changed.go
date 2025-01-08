package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewOrdersRiskAssessmentChanged(webhook webhook.ValidWebhook) (*OrdersRiskAssessmentChanged, error) {
	payload := &OrdersRiskAssessmentChangedPayload{}
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
	return &OrdersRiskAssessmentChanged{
		info: info,
		data: *payload,
	}, nil
}

type OrdersRiskAssessmentChanged struct {
	info webhook.Info
	data OrdersRiskAssessmentChangedPayload
}

func (webhook *OrdersRiskAssessmentChanged) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrdersRiskAssessmentChanged) GetData() (OrdersRiskAssessmentChangedPayload, error) {
	return webhook.data, nil
}

type OrdersRiskAssessmentChangedPayload struct {
	AdminGraphqlAPIOrderID interface{} `json:"admin_graphql_api_order_id"`
	CreatedAt              interface{} `json:"created_at"`
	OrderID                interface{} `json:"order_id"`
	ProviderID             interface{} `json:"provider_id"`
	ProviderTitle          interface{} `json:"provider_title"`
	RiskLevel              string      `json:"risk_level"`
}
