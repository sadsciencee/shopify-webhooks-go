package payload

import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/v2025_01/shopify"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)

func NewDraftOrdersUpdate(webhook webhook.ValidWebhook) (*DraftOrdersUpdate, error) {
	payload := &DraftOrdersUpdatePayload{}
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
	return &DraftOrdersUpdate{
		info: info,
		data: *payload,
	}, nil
}

type DraftOrdersUpdate struct {
	info webhook.Info
	data DraftOrdersUpdatePayload
}

func (webhook *DraftOrdersUpdate) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *DraftOrdersUpdate) GetData() (DraftOrdersUpdatePayload, error) {
	return webhook.data, nil
}

type DraftOrdersUpdatePayload struct {
	AdminGraphqlAPIID            shopify.ID  `json:"admin_graphql_api_id"`
	AllowDiscountCodesInCheckout bool        `json:"allow_discount_codes_in_checkout?"`
	APIClientID                  interface{} `json:"api_client_id"`
	AppliedDiscount              struct {
		Amount      string `json:"amount"`
		Description string `json:"description"`
		Title       string `json:"title"`
		Value       string `json:"value"`
		ValueType   string `json:"value_type"`
	} `json:"applied_discount"`
	B2b                 bool             `json:"b2b?"`
	BillingAddress      *shopify.Address `json:"billing_address"`
	CompletedAt         time.Time        `json:"completed_at"`
	CreatedAt           time.Time        `json:"created_at"`
	CreatedOnAPIVersion interface{}      `json:"created_on_api_version"`
	Currency            string           `json:"currency"`
	Customer            shopify.Customer `json:"customer"`
	Email               string           `json:"email"`
	ID                  int64            `json:"id"`
	InvoiceSentAt       interface{}      `json:"invoice_sent_at"`
	InvoiceURL          string           `json:"invoice_url"`
	LineItems           []struct {
		AdminGraphqlAPIID  string                    `json:"admin_graphql_api_id"`
		AppliedDiscount    interface{}               `json:"applied_discount"`
		Custom             bool                      `json:"custom"`
		FulfillmentService string                    `json:"fulfillment_service"`
		GiftCard           bool                      `json:"gift_card"`
		Grams              int64                     `json:"grams"`
		ID                 int64                     `json:"id"`
		Name               string                    `json:"name"`
		Price              string                    `json:"price"`
		ProductID          int64                     `json:"product_id"`
		Properties         []shopify.CustomAttribute `json:"properties"`
		Quantity           int64                     `json:"quantity"`
		RequiresShipping   bool                      `json:"requires_shipping"`
		Sku                string                    `json:"sku"`
		TaxLines           []shopify.LineItemTaxLine `json:"tax_lines"`
		Taxable            bool                      `json:"taxable"`
		Title              string                    `json:"title"`
		VariantID          int64                     `json:"variant_id"`
		VariantTitle       string                    `json:"variant_title"`
		Vendor             string                    `json:"vendor"`
	} `json:"line_items"`
	Name           string        `json:"name"`
	Note           *string       `json:"note"`
	NoteAttributes []interface{} `json:"note_attributes"`
	OrderID        interface{}   `json:"order_id"`
	PaymentTerms   struct {
		CanPayEarly      bool      `json:"can_pay_early"`
		CreatedAt        time.Time `json:"created_at"`
		DueInDays        int64     `json:"due_in_days"`
		ID               int64     `json:"id"`
		PaymentSchedules []struct {
			Amount                     string    `json:"amount"`
			BalanceDue                 string    `json:"balance_due"`
			BalanceDueCurrency         string    `json:"balance_due_currency"`
			CompletedAt                time.Time `json:"completed_at"`
			CreatedAt                  time.Time `json:"created_at"`
			Currency                   string    `json:"currency"`
			DueAt                      time.Time `json:"due_at"`
			ID                         int64     `json:"id"`
			IssuedAt                   time.Time `json:"issued_at"`
			OutstandingBalance         string    `json:"outstanding_balance"`
			OutstandingBalanceCurrency string    `json:"outstanding_balance_currency"`
			PaymentTermsID             int64     `json:"payment_terms_id"`
			ReferenceID                int64     `json:"reference_id"`
			ReferenceType              string    `json:"reference_type"`
			TotalBalance               string    `json:"total_balance"`
			TotalBalanceCurrency       string    `json:"total_balance_currency"`
			TotalPrice                 string    `json:"total_price"`
			TotalPriceCurrency         string    `json:"total_price_currency"`
			UpdatedAt                  time.Time `json:"updated_at"`
		} `json:"payment_schedules"`
		PaymentTermsName string    `json:"payment_terms_name"`
		PaymentTermsType string    `json:"payment_terms_type"`
		UpdatedAt        time.Time `json:"updated_at"`
	} `json:"payment_terms"`
	ShippingAddress *shopify.Address `json:"shipping_address"`
	ShippingLine    struct {
		Custom bool        `json:"custom"`
		Handle interface{} `json:"handle"`
		Price  string      `json:"price"`
		Title  string      `json:"title"`
	} `json:"shipping_line"`
	Status        string                 `json:"status"`
	SubtotalPrice string                 `json:"subtotal_price"`
	Tags          string                 `json:"tags"`
	TaxExempt     bool                   `json:"tax_exempt"`
	TaxLines      []shopify.OrderTaxLine `json:"tax_lines"`
	TaxesIncluded bool                   `json:"taxes_included"`
	TotalPrice    string                 `json:"total_price"`
	TotalTax      string                 `json:"total_tax"`
	UpdatedAt     time.Time              `json:"updated_at"`
}
