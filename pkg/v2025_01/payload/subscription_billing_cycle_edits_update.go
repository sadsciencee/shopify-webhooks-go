package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewSubscriptionBillingCycleEditsUpdate(webhook webhook.ValidWebhook) (*SubscriptionBillingCycleEditsUpdate, error) {
	payload := &SubscriptionBillingCycleEditsUpdatePayload{}
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
	return &SubscriptionBillingCycleEditsUpdate{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingCycleEditsUpdate struct {
	info webhook.Info
	data SubscriptionBillingCycleEditsUpdatePayload
}

func (webhook *SubscriptionBillingCycleEditsUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingCycleEditsUpdate) GetData() (SubscriptionBillingCycleEditsUpdatePayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingCycleEditsUpdatePayload struct {
	BillingAttemptExpectedDate time.Time   `json:"billing_attempt_expected_date"`
	ContractEdit               interface{} `json:"contract_edit"`
	CycleEndAt                 time.Time   `json:"cycle_end_at"`
	CycleIndex                 int64       `json:"cycle_index"`
	CycleStartAt               time.Time   `json:"cycle_start_at"`
	Edited                     bool        `json:"edited"`
	Skipped                    bool        `json:"skipped"`
	SubscriptionContractID     int64       `json:"subscription_contract_id"`
}
