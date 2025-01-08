package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDisputesCreate(webhook webhook.ValidWebhook) (*DisputesCreate, error) {
	payload := &DisputesCreatePayload{}
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
	return &DisputesCreate{
		info: info,
		data: *payload,
	}, nil
}

type DisputesCreate struct {
	info webhook.Info
	data DisputesCreatePayload
}

func (webhook *DisputesCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DisputesCreate) GetData() (DisputesCreatePayload, error) {
	return webhook.data, nil
}

type DisputesCreatePayload struct {
	Amount            string      `json:"amount"`
	Currency          string      `json:"currency"`
	EvidenceDueBy     string      `json:"evidence_due_by"`
	EvidenceSentOn    interface{} `json:"evidence_sent_on"`
	FinalizedOn       interface{} `json:"finalized_on"`
	ID                int64       `json:"id"`
	InitiatedAt       time.Time   `json:"initiated_at"`
	NetworkReasonCode string      `json:"network_reason_code"`
	OrderID           int64       `json:"order_id"`
	Reason            string      `json:"reason"`
	Status            string      `json:"status"`
	Type              string      `json:"type"`
}
