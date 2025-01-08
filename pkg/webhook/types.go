package webhook

import (
	"fmt"
	"io"
)

type ValidWebhook interface {
	GetInfo() (Info, error)
	GetReader() (io.Reader, error)
}

type Info struct {
	Topic       Topic
	ShopDomain  string
	ApiVersion  string
	WebhookId   string
	TriggeredAt string
	EventId     string
}

type Topic string

const (
	AppScopesUpdate                                     Topic = "app/scopes_update"
	AppUninstalled                                      Topic = "app/uninstalled"
	AppPurchasesOneTimeUpdate                           Topic = "app_purchases_one_time/update"
	AppSubscriptionsApproachingCappedAmount             Topic = "app_subscriptions/approaching_capped_amount"
	AppSubscriptionsUpdate                              Topic = "app_subscriptions/update"
	AuditEventsAdminApiActivity                         Topic = "audit_events/admin_api_activity"
	BulkOperationsFinish                                Topic = "bulk_operations/finish"
	CartsCreate                                         Topic = "carts/create"
	CartsUpdate                                         Topic = "carts/update"
	ChannelsDelete                                      Topic = "channels/delete"
	CheckoutsCreate                                     Topic = "checkouts/create"
	CheckoutsDelete                                     Topic = "checkouts/delete"
	CheckoutsUpdate                                     Topic = "checkouts/update"
	CollectionListingsAdd                               Topic = "collection_listings/add"
	CollectionListingsRemove                            Topic = "collection_listings/remove"
	CollectionListingsUpdate                            Topic = "collection_listings/update"
	CollectionPublicationsCreate                        Topic = "collection_publications/create"
	CollectionPublicationsDelete                        Topic = "collection_publications/delete"
	CollectionPublicationsUpdate                        Topic = "collection_publications/update"
	CollectionsCreate                                   Topic = "collections/create"
	CollectionsDelete                                   Topic = "collections/delete"
	CollectionsUpdate                                   Topic = "collections/update"
	CompaniesCreate                                     Topic = "companies/create"
	CompaniesDelete                                     Topic = "companies/delete"
	CompaniesUpdate                                     Topic = "companies/update"
	CompanyContactRolesAssign                           Topic = "company_contact_roles/assign"
	CompanyContactRolesRevoke                           Topic = "company_contact_roles/revoke"
	CompanyContactsCreate                               Topic = "company_contacts/create"
	CompanyContactsDelete                               Topic = "company_contacts/delete"
	CompanyContactsUpdate                               Topic = "company_contacts/update"
	CompanyLocationsCreate                              Topic = "company_locations/create"
	CompanyLocationsDelete                              Topic = "company_locations/delete"
	CompanyLocationsUpdate                              Topic = "company_locations/update"
	CustomerJoinedSegment                               Topic = "customer.joined_segment"
	CustomerLeftSegment                                 Topic = "customer.left_segment"
	CustomerTagsAdded                                   Topic = "customer.tags_added"
	CustomerTagsRemoved                                 Topic = "customer.tags_removed"
	CustomerAccountSettingsUpdate                       Topic = "customer_account_settings/update"
	CustomerGroupsCreate                                Topic = "customer_groups/create"
	CustomerGroupsDelete                                Topic = "customer_groups/delete"
	CustomerGroupsUpdate                                Topic = "customer_groups/update"
	CustomerPaymentMethodsCreate                        Topic = "customer_payment_methods/create"
	CustomerPaymentMethodsRevoke                        Topic = "customer_payment_methods/revoke"
	CustomerPaymentMethodsUpdate                        Topic = "customer_payment_methods/update"
	CustomersCreate                                     Topic = "customers/create"
	CustomersDataRequest                                Topic = "customers/data_request"
	CustomersDelete                                     Topic = "customers/delete"
	CustomersDisable                                    Topic = "customers/disable"
	CustomersEnable                                     Topic = "customers/enable"
	CustomersMerge                                      Topic = "customers/merge"
	CustomersPurchasingSummary                          Topic = "customers/purchasing_summary"
	CustomersRedact                                     Topic = "customers/redact"
	CustomersUpdate                                     Topic = "customers/update"
	CustomersEmailMarketingConsentUpdate                Topic = "customers_email_marketing_consent/update"
	CustomersMarketingConsentUpdate                     Topic = "customers_marketing_consent/update"
	DeliveryPromiseSettingsUpdate                       Topic = "delivery_promise_settings/update"
	DiscountsCreate                                     Topic = "discounts/create"
	DiscountsDelete                                     Topic = "discounts/delete"
	DiscountsRedeemcodeAdded                            Topic = "discounts/redeemcode_added"
	DiscountsRedeemcodeRemoved                          Topic = "discounts/redeemcode_removed"
	DiscountsUpdate                                     Topic = "discounts/update"
	DisputesCreate                                      Topic = "disputes/create"
	DisputesUpdate                                      Topic = "disputes/update"
	DomainsCreate                                       Topic = "domains/create"
	DomainsDestroy                                      Topic = "domains/destroy"
	DomainsUpdate                                       Topic = "domains/update"
	DraftOrdersCreate                                   Topic = "draft_orders/create"
	DraftOrdersDelete                                   Topic = "draft_orders/delete"
	DraftOrdersUpdate                                   Topic = "draft_orders/update"
	FulfillmentEventsCreate                             Topic = "fulfillment_events/create"
	FulfillmentEventsDelete                             Topic = "fulfillment_events/delete"
	FulfillmentHoldsAdded                               Topic = "fulfillment_holds/added"
	FulfillmentHoldsReleased                            Topic = "fulfillment_holds/released"
	FulfillmentOrdersCancellationRequestAccepted        Topic = "fulfillment_orders/cancellation_request_accepted"
	FulfillmentOrdersCancellationRequestRejected        Topic = "fulfillment_orders/cancellation_request_rejected"
	FulfillmentOrdersCancellationRequestSubmitted       Topic = "fulfillment_orders/cancellation_request_submitted"
	FulfillmentOrdersCancelled                          Topic = "fulfillment_orders/cancelled"
	FulfillmentOrdersFulfillmentRequestAccepted         Topic = "fulfillment_orders/fulfillment_request_accepted"
	FulfillmentOrdersFulfillmentRequestRejected         Topic = "fulfillment_orders/fulfillment_request_rejected"
	FulfillmentOrdersFulfillmentRequestSubmitted        Topic = "fulfillment_orders/fulfillment_request_submitted"
	FulfillmentOrdersFulfillmentServiceFailedToComplete Topic = "fulfillment_orders/fulfillment_service_failed_to_complete"
	FulfillmentOrdersHoldReleased                       Topic = "fulfillment_orders/hold_released"
	FulfillmentOrdersLineItemsPreparedForLocalDelivery  Topic = "fulfillment_orders/line_items_prepared_for_local_delivery"
	FulfillmentOrdersLineItemsPreparedForPickup         Topic = "fulfillment_orders/line_items_prepared_for_pickup"
	FulfillmentOrdersMerged                             Topic = "fulfillment_orders/merged"
	FulfillmentOrdersMoved                              Topic = "fulfillment_orders/moved"
	FulfillmentOrdersOrderRoutingComplete               Topic = "fulfillment_orders/order_routing_complete"
	FulfillmentOrdersPlacedOnHold                       Topic = "fulfillment_orders/placed_on_hold"
	FulfillmentOrdersRescheduled                        Topic = "fulfillment_orders/rescheduled"
	FulfillmentOrdersScheduledFulfillmentOrderReady     Topic = "fulfillment_orders/scheduled_fulfillment_order_ready"
	FulfillmentOrdersSplit                              Topic = "fulfillment_orders/split"
	FulfillmentsCreate                                  Topic = "fulfillments/create"
	FulfillmentsUpdate                                  Topic = "fulfillments/update"
	InventoryItemsCreate                                Topic = "inventory_items/create"
	InventoryItemsDelete                                Topic = "inventory_items/delete"
	InventoryItemsUpdate                                Topic = "inventory_items/update"
	InventoryLevelsConnect                              Topic = "inventory_levels/connect"
	InventoryLevelsDisconnect                           Topic = "inventory_levels/disconnect"
	InventoryLevelsUpdate                               Topic = "inventory_levels/update"
	LocalesCreate                                       Topic = "locales/create"
	LocalesUpdate                                       Topic = "locales/update"
	LocationsActivate                                   Topic = "locations/activate"
	LocationsCreate                                     Topic = "locations/create"
	LocationsDeactivate                                 Topic = "locations/deactivate"
	LocationsDelete                                     Topic = "locations/delete"
	LocationsUpdate                                     Topic = "locations/update"
	MarketsCreate                                       Topic = "markets/create"
	MarketsDelete                                       Topic = "markets/delete"
	MarketsUpdate                                       Topic = "markets/update"
	MetafieldDefinitionsCreate                          Topic = "metafield_definitions/create"
	MetafieldDefinitionsDelete                          Topic = "metafield_definitions/delete"
	MetafieldDefinitionsUpdate                          Topic = "metafield_definitions/update"
	MetaobjectsCreate                                   Topic = "metaobjects/create"
	MetaobjectsDelete                                   Topic = "metaobjects/delete"
	MetaobjectsUpdate                                   Topic = "metaobjects/update"
	OrderTransactionsCreate                             Topic = "order_transactions/create"
	OrdersCancelled                                     Topic = "orders/cancelled"
	OrdersCreate                                        Topic = "orders/create"
	OrdersDelete                                        Topic = "orders/delete"
	OrdersEdited                                        Topic = "orders/edited"
	OrdersFulfilled                                     Topic = "orders/fulfilled"
	OrdersPaid                                          Topic = "orders/paid"
	OrdersPartiallyFulfilled                            Topic = "orders/partially_fulfilled"
	OrdersRiskAssessmentChanged                         Topic = "orders/risk_assessment_changed"
	OrdersShopifyProtectEligibilityChanged              Topic = "orders/shopify_protect_eligibility_changed"
	OrdersUpdated                                       Topic = "orders/updated"
	PaymentSchedulesDue                                 Topic = "payment_schedules/due"
	PaymentTermsCreate                                  Topic = "payment_terms/create"
	PaymentTermsDelete                                  Topic = "payment_terms/delete"
	PaymentTermsUpdate                                  Topic = "payment_terms/update"
	ProductFeedsCreate                                  Topic = "product_feeds/create"
	ProductFeedsFullSync                                Topic = "product_feeds/full_sync"
	ProductFeedsFullSyncFinish                          Topic = "product_feeds/full_sync_finish"
	ProductFeedsIncrementalSync                         Topic = "product_feeds/incremental_sync"
	ProductFeedsUpdate                                  Topic = "product_feeds/update"
	ProductListingsAdd                                  Topic = "product_listings/add"
	ProductListingsRemove                               Topic = "product_listings/remove"
	ProductListingsUpdate                               Topic = "product_listings/update"
	ProductPublicationsCreate                           Topic = "product_publications/create"
	ProductPublicationsDelete                           Topic = "product_publications/delete"
	ProductPublicationsUpdate                           Topic = "product_publications/update"
	ProductsCreate                                      Topic = "products/create"
	ProductsDelete                                      Topic = "products/delete"
	ProductsUpdate                                      Topic = "products/update"
	ProfilesCreate                                      Topic = "profiles/create"
	ProfilesDelete                                      Topic = "profiles/delete"
	ProfilesUpdate                                      Topic = "profiles/update"
	RefundsCreate                                       Topic = "refunds/create"
	ReturnsApprove                                      Topic = "returns/approve"
	ReturnsCancel                                       Topic = "returns/cancel"
	ReturnsClose                                        Topic = "returns/close"
	ReturnsDecline                                      Topic = "returns/decline"
	ReturnsReopen                                       Topic = "returns/reopen"
	ReturnsRequest                                      Topic = "returns/request"
	ReturnsUpdate                                       Topic = "returns/update"
	ReverseDeliveriesAttachDeliverable                  Topic = "reverse_deliveries/attach_deliverable"
	ReverseFulfillmentOrdersDispose                     Topic = "reverse_fulfillment_orders/dispose"
	ScheduledProductListingsAdd                         Topic = "scheduled_product_listings/add"
	ScheduledProductListingsRemove                      Topic = "scheduled_product_listings/remove"
	ScheduledProductListingsUpdate                      Topic = "scheduled_product_listings/update"
	SegmentsCreate                                      Topic = "segments/create"
	SegmentsDelete                                      Topic = "segments/delete"
	SegmentsUpdate                                      Topic = "segments/update"
	SellingPlanGroupsCreate                             Topic = "selling_plan_groups/create"
	SellingPlanGroupsDelete                             Topic = "selling_plan_groups/delete"
	SellingPlanGroupsUpdate                             Topic = "selling_plan_groups/update"
	ShippingAddressesCreate                             Topic = "shipping_addresses/create"
	ShippingAddressesUpdate                             Topic = "shipping_addresses/update"
	ShopRedact                                          Topic = "shop/redact"
	ShopUpdate                                          Topic = "shop/update"
	SubscriptionBillingAttemptsChallenged               Topic = "subscription_billing_attempts/challenged"
	SubscriptionBillingAttemptsFailure                  Topic = "subscription_billing_attempts/failure"
	SubscriptionBillingAttemptsSuccess                  Topic = "subscription_billing_attempts/success"
	SubscriptionBillingCycleEditsCreate                 Topic = "subscription_billing_cycle_edits/create"
	SubscriptionBillingCycleEditsDelete                 Topic = "subscription_billing_cycle_edits/delete"
	SubscriptionBillingCycleEditsUpdate                 Topic = "subscription_billing_cycle_edits/update"
	SubscriptionBillingCyclesSkip                       Topic = "subscription_billing_cycles/skip"
	SubscriptionBillingCyclesUnskip                     Topic = "subscription_billing_cycles/unskip"
	SubscriptionContractsActivate                       Topic = "subscription_contracts/activate"
	SubscriptionContractsCancel                         Topic = "subscription_contracts/cancel"
	SubscriptionContractsCreate                         Topic = "subscription_contracts/create"
	SubscriptionContractsExpire                         Topic = "subscription_contracts/expire"
	SubscriptionContractsFail                           Topic = "subscription_contracts/fail"
	SubscriptionContractsPause                          Topic = "subscription_contracts/pause"
	SubscriptionContractsUpdate                         Topic = "subscription_contracts/update"
	TaxServicesCreate                                   Topic = "tax_services/create"
	TaxServicesUpdate                                   Topic = "tax_services/update"
	TenderTransactionsCreate                            Topic = "tender_transactions/create"
	ThemesCreate                                        Topic = "themes/create"
	ThemesDelete                                        Topic = "themes/delete"
	ThemesPublish                                       Topic = "themes/publish"
	ThemesUpdate                                        Topic = "themes/update"
	VariantsInStock                                     Topic = "variants/in_stock"
	VariantsOutOfStock                                  Topic = "variants/out_of_stock"
)

var valids = map[Topic]bool{
	AppScopesUpdate:                                     true,
	AppUninstalled:                                      true,
	AppPurchasesOneTimeUpdate:                           true,
	AppSubscriptionsApproachingCappedAmount:             true,
	AppSubscriptionsUpdate:                              true,
	AuditEventsAdminApiActivity:                         true,
	BulkOperationsFinish:                                true,
	CartsCreate:                                         true,
	CartsUpdate:                                         true,
	ChannelsDelete:                                      true,
	CheckoutsCreate:                                     true,
	CheckoutsDelete:                                     true,
	CheckoutsUpdate:                                     true,
	CollectionListingsAdd:                               true,
	CollectionListingsRemove:                            true,
	CollectionListingsUpdate:                            true,
	CollectionPublicationsCreate:                        true,
	CollectionPublicationsDelete:                        true,
	CollectionPublicationsUpdate:                        true,
	CollectionsCreate:                                   true,
	CollectionsDelete:                                   true,
	CollectionsUpdate:                                   true,
	CompaniesCreate:                                     true,
	CompaniesDelete:                                     true,
	CompaniesUpdate:                                     true,
	CompanyContactRolesAssign:                           true,
	CompanyContactRolesRevoke:                           true,
	CompanyContactsCreate:                               true,
	CompanyContactsDelete:                               true,
	CompanyContactsUpdate:                               true,
	CompanyLocationsCreate:                              true,
	CompanyLocationsDelete:                              true,
	CompanyLocationsUpdate:                              true,
	CustomerJoinedSegment:                               true,
	CustomerLeftSegment:                                 true,
	CustomerTagsAdded:                                   true,
	CustomerTagsRemoved:                                 true,
	CustomerAccountSettingsUpdate:                       true,
	CustomerGroupsCreate:                                true,
	CustomerGroupsDelete:                                true,
	CustomerGroupsUpdate:                                true,
	CustomerPaymentMethodsCreate:                        true,
	CustomerPaymentMethodsRevoke:                        true,
	CustomerPaymentMethodsUpdate:                        true,
	CustomersCreate:                                     true,
	CustomersDataRequest:                                true,
	CustomersDelete:                                     true,
	CustomersDisable:                                    true,
	CustomersEnable:                                     true,
	CustomersMerge:                                      true,
	CustomersPurchasingSummary:                          true,
	CustomersRedact:                                     true,
	CustomersUpdate:                                     true,
	CustomersEmailMarketingConsentUpdate:                true,
	CustomersMarketingConsentUpdate:                     true,
	DeliveryPromiseSettingsUpdate:                       true,
	DiscountsCreate:                                     true,
	DiscountsDelete:                                     true,
	DiscountsRedeemcodeAdded:                            true,
	DiscountsRedeemcodeRemoved:                          true,
	DiscountsUpdate:                                     true,
	DisputesCreate:                                      true,
	DisputesUpdate:                                      true,
	DomainsCreate:                                       true,
	DomainsDestroy:                                      true,
	DomainsUpdate:                                       true,
	DraftOrdersCreate:                                   true,
	DraftOrdersDelete:                                   true,
	DraftOrdersUpdate:                                   true,
	FulfillmentEventsCreate:                             true,
	FulfillmentEventsDelete:                             true,
	FulfillmentHoldsAdded:                               true,
	FulfillmentHoldsReleased:                            true,
	FulfillmentOrdersCancellationRequestAccepted:        true,
	FulfillmentOrdersCancellationRequestRejected:        true,
	FulfillmentOrdersCancellationRequestSubmitted:       true,
	FulfillmentOrdersCancelled:                          true,
	FulfillmentOrdersFulfillmentRequestAccepted:         true,
	FulfillmentOrdersFulfillmentRequestRejected:         true,
	FulfillmentOrdersFulfillmentRequestSubmitted:        true,
	FulfillmentOrdersFulfillmentServiceFailedToComplete: true,
	FulfillmentOrdersHoldReleased:                       true,
	FulfillmentOrdersLineItemsPreparedForLocalDelivery:  true,
	FulfillmentOrdersLineItemsPreparedForPickup:         true,
	FulfillmentOrdersMerged:                             true,
	FulfillmentOrdersMoved:                              true,
	FulfillmentOrdersOrderRoutingComplete:               true,
	FulfillmentOrdersPlacedOnHold:                       true,
	FulfillmentOrdersRescheduled:                        true,
	FulfillmentOrdersScheduledFulfillmentOrderReady:     true,
	FulfillmentOrdersSplit:                              true,
	FulfillmentsCreate:                                  true,
	FulfillmentsUpdate:                                  true,
	InventoryItemsCreate:                                true,
	InventoryItemsDelete:                                true,
	InventoryItemsUpdate:                                true,
	InventoryLevelsConnect:                              true,
	InventoryLevelsDisconnect:                           true,
	InventoryLevelsUpdate:                               true,
	LocalesCreate:                                       true,
	LocalesUpdate:                                       true,
	LocationsActivate:                                   true,
	LocationsCreate:                                     true,
	LocationsDeactivate:                                 true,
	LocationsDelete:                                     true,
	LocationsUpdate:                                     true,
	MarketsCreate:                                       true,
	MarketsDelete:                                       true,
	MarketsUpdate:                                       true,
	MetafieldDefinitionsCreate:                          true,
	MetafieldDefinitionsDelete:                          true,
	MetafieldDefinitionsUpdate:                          true,
	MetaobjectsCreate:                                   true,
	MetaobjectsDelete:                                   true,
	MetaobjectsUpdate:                                   true,
	OrderTransactionsCreate:                             true,
	OrdersCancelled:                                     true,
	OrdersCreate:                                        true,
	OrdersDelete:                                        true,
	OrdersEdited:                                        true,
	OrdersFulfilled:                                     true,
	OrdersPaid:                                          true,
	OrdersPartiallyFulfilled:                            true,
	OrdersRiskAssessmentChanged:                         true,
	OrdersShopifyProtectEligibilityChanged:              true,
	OrdersUpdated:                                       true,
	PaymentSchedulesDue:                                 true,
	PaymentTermsCreate:                                  true,
	PaymentTermsDelete:                                  true,
	PaymentTermsUpdate:                                  true,
	ProductFeedsCreate:                                  true,
	ProductFeedsFullSync:                                true,
	ProductFeedsFullSyncFinish:                          true,
	ProductFeedsIncrementalSync:                         true,
	ProductFeedsUpdate:                                  true,
	ProductListingsAdd:                                  true,
	ProductListingsRemove:                               true,
	ProductListingsUpdate:                               true,
	ProductPublicationsCreate:                           true,
	ProductPublicationsDelete:                           true,
	ProductPublicationsUpdate:                           true,
	ProductsCreate:                                      true,
	ProductsDelete:                                      true,
	ProductsUpdate:                                      true,
	ProfilesCreate:                                      true,
	ProfilesDelete:                                      true,
	ProfilesUpdate:                                      true,
	RefundsCreate:                                       true,
	ReturnsApprove:                                      true,
	ReturnsCancel:                                       true,
	ReturnsClose:                                        true,
	ReturnsDecline:                                      true,
	ReturnsReopen:                                       true,
	ReturnsRequest:                                      true,
	ReturnsUpdate:                                       true,
	ReverseDeliveriesAttachDeliverable:                  true,
	ReverseFulfillmentOrdersDispose:                     true,
	ScheduledProductListingsAdd:                         true,
	ScheduledProductListingsRemove:                      true,
	ScheduledProductListingsUpdate:                      true,
	SegmentsCreate:                                      true,
	SegmentsDelete:                                      true,
	SegmentsUpdate:                                      true,
	SellingPlanGroupsCreate:                             true,
	SellingPlanGroupsDelete:                             true,
	SellingPlanGroupsUpdate:                             true,
	ShippingAddressesCreate:                             true,
	ShippingAddressesUpdate:                             true,
	ShopRedact:                                          true,
	ShopUpdate:                                          true,
	SubscriptionBillingAttemptsChallenged:               true,
	SubscriptionBillingAttemptsFailure:                  true,
	SubscriptionBillingAttemptsSuccess:                  true,
	SubscriptionBillingCycleEditsCreate:                 true,
	SubscriptionBillingCycleEditsDelete:                 true,
	SubscriptionBillingCycleEditsUpdate:                 true,
	SubscriptionBillingCyclesSkip:                       true,
	SubscriptionBillingCyclesUnskip:                     true,
	SubscriptionContractsActivate:                       true,
	SubscriptionContractsCancel:                         true,
	SubscriptionContractsCreate:                         true,
	SubscriptionContractsExpire:                         true,
	SubscriptionContractsFail:                           true,
	SubscriptionContractsPause:                          true,
	SubscriptionContractsUpdate:                         true,
	TaxServicesCreate:                                   true,
	TaxServicesUpdate:                                   true,
	TenderTransactionsCreate:                            true,
	ThemesCreate:                                        true,
	ThemesDelete:                                        true,
	ThemesPublish:                                       true,
	ThemesUpdate:                                        true,
	VariantsInStock:                                     true,
	VariantsOutOfStock:                                  true,
}

func ValidateTopic(topic string) (Topic, error) {
	webhookTopic := Topic(topic)
	if !valids[webhookTopic] {
		return "", fmt.Errorf("invalid webhook topic: %s", topic)
	}
	return webhookTopic, nil
}
