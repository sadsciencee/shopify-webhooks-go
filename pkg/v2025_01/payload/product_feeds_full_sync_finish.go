package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"strings"
	"time"
)

func NewProductFeedsFullSyncFinish(webhook webhook.ValidWebhook) (*ProductFeedsFullSyncFinish, error) {
	payload := &ProductFeedsFullSyncFinishPayload{}
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
	return &ProductFeedsFullSyncFinish{
		info: info,
		data: *payload,
	}, nil
}

type ProductFeedsFullSyncFinish struct {
	info webhook.Info
	data ProductFeedsFullSyncFinishPayload
}

func (webhook *ProductFeedsFullSyncFinish) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *ProductFeedsFullSyncFinish) GetData() (ProductFeedsFullSyncFinishPayload, error) {
	return webhook.data, nil
}

type ProductFeedsFullSyncFinishPayload struct {
	FullSync struct {
		Count     int64       `json:"count"`
		CreatedAt CustomTime  `json:"createdAt"`
		ErrorCode interface{} `json:"errorCode"`
		Status    string      `json:"status"`
		URL       interface{} `json:"url"`
	} `json:"fullSync"`
	Metadata struct {
		Action          string        `json:"action"`
		FullSyncID      shopify.ID    `json:"fullSyncId"`
		OccurredAt      time.Time     `json:"occurred_at"`
		Resource        string        `json:"resource"`
		TruncatedFields []interface{} `json:"truncatedFields"`
		Type            string        `json:"type"`
	} `json:"metadata"`
	ProductFeed struct {
		Country  string     `json:"country"`
		ID       shopify.ID `json:"id"`
		Language string     `json:"language"`
		ShopID   shopify.ID `json:"shop_id"`
	} `json:"productFeed"`
}

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	const layout = "2006-01-02 15:04:05 -0700"
	t, err := time.Parse(layout, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}
