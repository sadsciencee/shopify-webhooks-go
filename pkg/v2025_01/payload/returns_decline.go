package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)

func NewReturnsDecline(webhook webhook.ValidWebhook) (*ReturnsDecline, error) {
	payload := &ReturnsDeclinePayload{}
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
	return &ReturnsDecline{
		info: info,
		data: *payload,
	}, nil
}

type ReturnsDecline struct {
	info webhook.Info
	data ReturnsDeclinePayload
}

func (webhook *ReturnsDecline) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ReturnsDecline) GetData() (ReturnsDeclinePayload, error) {
	return webhook.data, nil
}

type ReturnsDeclinePayload struct {
	AdminGraphqlAPIID shopify.ID `json:"admin_graphql_api_id"`
	Decline           struct {
		Note   string `json:"note"`
		Reason string `json:"reason"`
	} `json:"decline"`
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
