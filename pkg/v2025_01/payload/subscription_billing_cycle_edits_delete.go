package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewSubscriptionBillingCycleEditsDelete(webhook webhook.ValidWebhook) (*SubscriptionBillingCycleEditsDelete, error) {
	payload := &SubscriptionBillingCycleEditsDeletePayload{}
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
	return &SubscriptionBillingCycleEditsDelete{
		info: info,
		data: *payload,
	}, nil
}

type SubscriptionBillingCycleEditsDelete struct {
	info webhook.Info
	data SubscriptionBillingCycleEditsDeletePayload
}

func (webhook *SubscriptionBillingCycleEditsDelete) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *SubscriptionBillingCycleEditsDelete) GetData() (SubscriptionBillingCycleEditsDeletePayload, error) {
	return webhook.data, nil
}

type SubscriptionBillingCycleEditsDeletePayload struct {
	BillingAttemptExpectedDate time.Time   `json:"billing_attempt_expected_date"`
	ContractEdit               interface{} `json:"contract_edit"`
	CycleEndAt                 time.Time   `json:"cycle_end_at"`
	CycleIndex                 int64       `json:"cycle_index"`
	CycleStartAt               time.Time   `json:"cycle_start_at"`
	Edited                     bool        `json:"edited"`
	Skipped                    bool        `json:"skipped"`
	SubscriptionContractID     int64       `json:"subscription_contract_id"`
}
