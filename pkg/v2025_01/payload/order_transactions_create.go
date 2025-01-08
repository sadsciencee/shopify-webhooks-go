package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewOrderTransactionsCreate(webhook webhook.ValidWebhook) (*OrderTransactionsCreate, error) {
	payload := &OrderTransactionsCreatePayload{}
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
	return &OrderTransactionsCreate{
		info: info,
		data: *payload,
	}, nil
}

type OrderTransactionsCreate struct {
	info webhook.Info
	data OrderTransactionsCreatePayload
}

func (webhook *OrderTransactionsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *OrderTransactionsCreate) GetData() (OrderTransactionsCreatePayload, error) {
	return webhook.data, nil
}

type OrderTransactionsCreatePayload struct {
	AdminGraphqlAPIID    shopify.ID      `json:"admin_graphql_api_id"`
	Amount               shopify.Decimal `json:"amount"`
	AmountRounding       interface{}     `json:"amount_rounding"`
	Authorization        string          `json:"authorization"`
	CreatedAt            time.Time       `json:"created_at"`
	Currency             string          `json:"currency"`
	DeviceID             interface{}     `json:"device_id"`
	ErrorCode            interface{}     `json:"error_code"`
	Gateway              string          `json:"gateway"`
	ID                   int64           `json:"id"`
	Kind                 string          `json:"kind"`
	LocationID           interface{}     `json:"location_id"`
	ManualPaymentGateway bool            `json:"manual_payment_gateway"`
	Message              interface{}     `json:"message"`
	OrderID              int64           `json:"order_id"`
	ParentID             interface{}     `json:"parent_id"`
	PaymentDetails       struct {
		AvsResultCode             interface{} `json:"avs_result_code"`
		BuyerActionInfo           interface{} `json:"buyer_action_info"`
		CreditCardBin             interface{} `json:"credit_card_bin"`
		CreditCardCompany         string      `json:"credit_card_company"`
		CreditCardExpirationMonth interface{} `json:"credit_card_expiration_month"`
		CreditCardExpirationYear  interface{} `json:"credit_card_expiration_year"`
		CreditCardName            interface{} `json:"credit_card_name"`
		CreditCardNumber          string      `json:"credit_card_number"`
		CreditCardWallet          interface{} `json:"credit_card_wallet"`
		CvvResultCode             interface{} `json:"cvv_result_code"`
		PaymentMethodName         string      `json:"payment_method_name"`
	} `json:"payment_details"`
	PaymentID         string           `json:"payment_id"`
	ProcessedAt       interface{}      `json:"processed_at"`
	Receipt           struct{}         `json:"receipt"`
	SourceName        string           `json:"source_name"`
	Status            string           `json:"status"`
	Test              bool             `json:"test"`
	TotalUnsettledSet shopify.MoneyBag `json:"total_unsettled_set"`
	UserID            interface{}      `json:"user_id"`
}
