package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewSubscriptionBillingCyclesSkip(webhook webhook.ValidWebhook) (*SubscriptionBillingCyclesSkip, error) {
	payload := &SubscriptionBillingCyclesSkipPayload{}
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
	return &SubscriptionBillingCyclesSkip{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingCyclesSkip struct {
	info webhook.Info
	data SubscriptionBillingCyclesSkipPayload
}

func (webhook *SubscriptionBillingCyclesSkip) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingCyclesSkip) GetData() (SubscriptionBillingCyclesSkipPayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingCyclesSkipPayload struct {
	BillingAttemptExpectedDate time.Time   `json:"billing_attempt_expected_date"`
	ContractEdit               interface{} `json:"contract_edit"`
	CycleEndAt                 time.Time   `json:"cycle_end_at"`
	CycleIndex                 int64       `json:"cycle_index"`
	CycleStartAt               time.Time   `json:"cycle_start_at"`
	Edited                     bool        `json:"edited"`
	Skipped                    bool        `json:"skipped"`
	SubscriptionContractID     int64       `json:"subscription_contract_id"`
}
