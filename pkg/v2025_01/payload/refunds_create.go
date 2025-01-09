package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewRefundsCreate(webhook webhook.ValidWebhook) (*RefundsCreate, error) {
	payload := &RefundsCreatePayload{}
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
	return &RefundsCreate{
		info: info,
		data: *payload,
	}, nil
}

type RefundsCreate struct {
	info webhook.Info
	data RefundsCreatePayload
}

func (webhook *RefundsCreate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *RefundsCreate) GetData() (RefundsCreatePayload, error) {
	return webhook.data, nil
}

type RefundsCreatePayload struct {
	AdminGraphqlAPIID shopify.ID    `json:"admin_graphql_api_id"`
	CreatedAt         time.Time     `json:"created_at"`
	Duties            []interface{} `json:"duties"`
	ID                int64         `json:"id"`
	Note              string        `json:"note"`
	OrderAdjustments  []interface{} `json:"order_adjustments"`
	OrderID           int64         `json:"order_id"`
	ProcessedAt       time.Time     `json:"processed_at"`
	RefundLineItems   []struct {
		ID       int64 `json:"id"`
		LineItem struct {
			AdminGraphqlAPIID          string                       `json:"admin_graphql_api_id"`
			DiscountAllocations        []shopify.DiscountAllocation `json:"discount_allocations"`
			Duties                     []interface{}                `json:"duties"`
			FulfillableQuantity        int64                        `json:"fulfillable_quantity"`
			FulfillmentService         string                       `json:"fulfillment_service"`
			FulfillmentStatus          interface{}                  `json:"fulfillment_status"`
			GiftCard                   bool                         `json:"gift_card"`
			Grams                      int64                        `json:"grams"`
			ID                         int64                        `json:"id"`
			Name                       string                       `json:"name"`
			Price                      string                       `json:"price"`
			PriceSet                   shopify.MoneyBag             `json:"price_set"`
			ProductExists              bool                         `json:"product_exists"`
			ProductID                  int64                        `json:"product_id"`
			Properties                 []shopify.CustomAttribute    `json:"properties"`
			Quantity                   int64                        `json:"quantity"`
			RequiresShipping           bool                         `json:"requires_shipping"`
			Sku                        string                       `json:"sku"`
			TaxLines                   []shopify.LineItemTaxLine    `json:"tax_lines"`
			Taxable                    bool                         `json:"taxable"`
			Title                      string                       `json:"title"`
			TotalDiscount              string                       `json:"total_discount"`
			TotalDiscountSet           shopify.MoneyBag             `json:"total_discount_set"`
			VariantID                  int64                        `json:"variant_id"`
			VariantInventoryManagement string                       `json:"variant_inventory_management"`
			VariantTitle               *string                      `json:"variant_title"`
			Vendor                     *string                      `json:"vendor"`
		} `json:"line_item"`
		LineItemID  int64            `json:"line_item_id"`
		LocationID  interface{}      `json:"location_id"`
		Quantity    int64            `json:"quantity"`
		RestockType string           `json:"restock_type"`
		Subtotal    float64          `json:"subtotal"`
		SubtotalSet shopify.MoneyBag `json:"subtotal_set"`
		TotalTax    float64          `json:"total_tax"`
		TotalTaxSet shopify.MoneyBag `json:"total_tax_set"`
	} `json:"refund_line_items"`
	RefundShippingLines []interface{}    `json:"refund_shipping_lines"`
	Restock             bool             `json:"restock"`
	Return              interface{}      `json:"return"`
	TotalDutiesSet      shopify.MoneyBag `json:"total_duties_set"`
	Transactions        []struct {
		AdminGraphqlAPIID    string           `json:"admin_graphql_api_id"`
		Amount               string           `json:"amount"`
		AmountRounding       interface{}      `json:"amount_rounding"`
		Authorization        interface{}      `json:"authorization"`
		CreatedAt            time.Time        `json:"created_at"`
		Currency             interface{}      `json:"currency"`
		DeviceID             interface{}      `json:"device_id"`
		ErrorCode            interface{}      `json:"error_code"`
		Gateway              string           `json:"gateway"`
		ID                   int64            `json:"id"`
		Kind                 string           `json:"kind"`
		LocationID           interface{}      `json:"location_id"`
		ManualPaymentGateway bool             `json:"manual_payment_gateway"`
		Message              string           `json:"message"`
		OrderID              int64            `json:"order_id"`
		ParentID             interface{}      `json:"parent_id"`
		PaymentID            string           `json:"payment_id"`
		ProcessedAt          interface{}      `json:"processed_at"`
		Receipt              struct{}         `json:"receipt"`
		SourceName           string           `json:"source_name"`
		Status               string           `json:"status"`
		Test                 bool             `json:"test"`
		TotalUnsettledSet    shopify.MoneyBag `json:"total_unsettled_set"`
		UserID               interface{}      `json:"user_id"`
	} `json:"transactions"`
	UserID int64 `json:"user_id"`
}
