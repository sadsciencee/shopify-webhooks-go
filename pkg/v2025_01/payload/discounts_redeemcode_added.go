package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDiscountsRedeemcodeAdded(webhook webhook.ValidWebhook) (*DiscountsRedeemcodeAdded, error) {
	payload := &DiscountsRedeemcodeAddedPayload{}
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
	return &DiscountsRedeemcodeAdded{
		info: info,
		data: *payload,
	}, nil
}

type DiscountsRedeemcodeAdded struct {
	info webhook.Info
	data DiscountsRedeemcodeAddedPayload
}

func (webhook *DiscountsRedeemcodeAdded) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DiscountsRedeemcodeAdded) GetData() (DiscountsRedeemcodeAddedPayload, error) {
	return webhook.data, nil
}

type DiscountsRedeemcodeAddedPayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	RedeemCode        struct {
		Code string `json:"code"`
		ID   string `json:"id"`
	} `json:"redeem_code"`
	UpdatedAt time.Time `json:"updated_at"`
}
