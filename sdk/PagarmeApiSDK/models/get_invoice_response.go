package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetInvoiceResponse represents a GetInvoiceResponse struct.
// Response object for getting an invoice
type GetInvoiceResponse struct {
	Id             Optional[string]                    `json:"id"`
	Code           Optional[string]                    `json:"code"`
	Url            Optional[string]                    `json:"url"`
	Amount         Optional[int]                       `json:"amount"`
	Status         Optional[string]                    `json:"status"`
	PaymentMethod  Optional[string]                    `json:"payment_method"`
	CreatedAt      Optional[time.Time]                 `json:"created_at"`
	Items          Optional[[]GetInvoiceItemResponse]  `json:"items"`
	Customer       Optional[GetCustomerResponse]       `json:"customer"`
	Charge         Optional[GetChargeResponse]         `json:"charge"`
	Installments   Optional[int]                       `json:"installments"`
	BillingAddress Optional[GetBillingAddressResponse] `json:"billing_address"`
	Subscription   Optional[GetSubscriptionResponse]   `json:"subscription"`
	Cycle          Optional[GetPeriodResponse]         `json:"cycle"`
	Shipping       Optional[GetShippingResponse]       `json:"shipping"`
	Metadata       Optional[map[string]string]         `json:"metadata"`
	DueAt          Optional[time.Time]                 `json:"due_at"`
	CanceledAt     Optional[time.Time]                 `json:"canceled_at"`
	BillingAt      Optional[time.Time]                 `json:"billing_at"`
	SeenAt         Optional[time.Time]                 `json:"seen_at"`
	// Total discounted value
	TotalDiscount Optional[int] `json:"total_discount"`
	// Total discounted value
	TotalIncrement Optional[int] `json:"total_increment"`
	// Subscription Id
	SubscriptionId Optional[string] `json:"subscription_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetInvoiceResponse.
// It customizes the JSON marshaling process for GetInvoiceResponse objects.
func (g *GetInvoiceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetInvoiceResponse object to a map representation for JSON marshaling.
func (g *GetInvoiceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.PaymentMethod.IsValueSet() {
		structMap["payment_method"] = g.PaymentMethod.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.Items.IsValueSet() {
		structMap["items"] = g.Items.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Charge.IsValueSet() {
		structMap["charge"] = g.Charge.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	if g.BillingAddress.IsValueSet() {
		structMap["billing_address"] = g.BillingAddress.Value()
	}
	if g.Subscription.IsValueSet() {
		structMap["subscription"] = g.Subscription.Value()
	}
	if g.Cycle.IsValueSet() {
		structMap["cycle"] = g.Cycle.Value()
	}
	if g.Shipping.IsValueSet() {
		structMap["shipping"] = g.Shipping.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if g.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if g.CanceledAt.Value() != nil {
			val := g.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if g.BillingAt.IsValueSet() {
		var BillingAtVal *string = nil
		if g.BillingAt.Value() != nil {
			val := g.BillingAt.Value().Format(time.RFC3339)
			BillingAtVal = &val
		}
		structMap["billing_at"] = BillingAtVal
	}
	if g.SeenAt.IsValueSet() {
		var SeenAtVal *string = nil
		if g.SeenAt.Value() != nil {
			val := g.SeenAt.Value().Format(time.RFC3339)
			SeenAtVal = &val
		}
		structMap["seen_at"] = SeenAtVal
	}
	if g.TotalDiscount.IsValueSet() {
		structMap["total_discount"] = g.TotalDiscount.Value()
	}
	if g.TotalIncrement.IsValueSet() {
		structMap["total_increment"] = g.TotalIncrement.Value()
	}
	if g.SubscriptionId.IsValueSet() {
		structMap["subscription_id"] = g.SubscriptionId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetInvoiceResponse.
// It customizes the JSON unmarshaling process for GetInvoiceResponse objects.
func (g *GetInvoiceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id             Optional[string]                    `json:"id"`
		Code           Optional[string]                    `json:"code"`
		Url            Optional[string]                    `json:"url"`
		Amount         Optional[int]                       `json:"amount"`
		Status         Optional[string]                    `json:"status"`
		PaymentMethod  Optional[string]                    `json:"payment_method"`
		CreatedAt      Optional[string]                    `json:"created_at"`
		Items          Optional[[]GetInvoiceItemResponse]  `json:"items"`
		Customer       Optional[GetCustomerResponse]       `json:"customer"`
		Charge         Optional[GetChargeResponse]         `json:"charge"`
		Installments   Optional[int]                       `json:"installments"`
		BillingAddress Optional[GetBillingAddressResponse] `json:"billing_address"`
		Subscription   Optional[GetSubscriptionResponse]   `json:"subscription"`
		Cycle          Optional[GetPeriodResponse]         `json:"cycle"`
		Shipping       Optional[GetShippingResponse]       `json:"shipping"`
		Metadata       Optional[map[string]string]         `json:"metadata"`
		DueAt          Optional[string]                    `json:"due_at"`
		CanceledAt     Optional[string]                    `json:"canceled_at"`
		BillingAt      Optional[string]                    `json:"billing_at"`
		SeenAt         Optional[string]                    `json:"seen_at"`
		TotalDiscount  Optional[int]                       `json:"total_discount"`
		TotalIncrement Optional[int]                       `json:"total_increment"`
		SubscriptionId Optional[string]                    `json:"subscription_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.Url = temp.Url
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.PaymentMethod = temp.PaymentMethod
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.Items = temp.Items
	g.Customer = temp.Customer
	g.Charge = temp.Charge
	g.Installments = temp.Installments
	g.BillingAddress = temp.BillingAddress
	g.Subscription = temp.Subscription
	g.Cycle = temp.Cycle
	g.Shipping = temp.Shipping
	g.Metadata = temp.Metadata
	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
	g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		g.CanceledAt.SetValue(&CanceledAtVal)
	}
	g.BillingAt.ShouldSetValue(temp.BillingAt.IsValueSet())
	if temp.BillingAt.Value() != nil {
		BillingAtVal, err := time.Parse(time.RFC3339, (*temp.BillingAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse billing_at as % s format.", time.RFC3339)
		}
		g.BillingAt.SetValue(&BillingAtVal)
	}
	g.SeenAt.ShouldSetValue(temp.SeenAt.IsValueSet())
	if temp.SeenAt.Value() != nil {
		SeenAtVal, err := time.Parse(time.RFC3339, (*temp.SeenAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse seen_at as % s format.", time.RFC3339)
		}
		g.SeenAt.SetValue(&SeenAtVal)
	}
	g.TotalDiscount = temp.TotalDiscount
	g.TotalIncrement = temp.TotalIncrement
	g.SubscriptionId = temp.SubscriptionId
	return nil
}
