package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDiscountsRedeemcodeRemoved(webhook webhook.ValidWebhook) (*DiscountsRedeemcodeRemoved, error) {
	payload := &DiscountsRedeemcodeRemovedPayload{}
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
	return &DiscountsRedeemcodeRemoved{
		info: info,
		data: *payload,
	}, nil
}

type DiscountsRedeemcodeRemoved struct {
	info webhook.Info
	data DiscountsRedeemcodeRemovedPayload
}

func (webhook *DiscountsRedeemcodeRemoved) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DiscountsRedeemcodeRemoved) GetData() (DiscountsRedeemcodeRemovedPayload, error) {
	return webhook.data, nil
}

type DiscountsRedeemcodeRemovedPayload struct {
	AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
	RedeemCode        struct {
		Code string `json:"code"`
		ID   string `json:"id"`
	} `json:"redeem_code"`
	UpdatedAt time.Time `json:"updated_at"`
}
