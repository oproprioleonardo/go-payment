package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetChargeResponse represents a GetChargeResponse struct.
// Response object for getting a charge
type GetChargeResponse struct {
	Id              Optional[string]                          `json:"id"`
	Code            Optional[string]                          `json:"code"`
	GatewayId       Optional[string]                          `json:"gateway_id"`
	Amount          Optional[int]                             `json:"amount"`
	Status          Optional[string]                          `json:"status"`
	Currency        Optional[string]                          `json:"currency"`
	PaymentMethod   Optional[string]                          `json:"payment_method"`
	DueAt           Optional[time.Time]                       `json:"due_at"`
	CreatedAt       Optional[time.Time]                       `json:"created_at"`
	UpdatedAt       Optional[time.Time]                       `json:"updated_at"`
	LastTransaction Optional[GetTransactionResponseInterface] `json:"last_transaction"`
	Invoice         Optional[GetInvoiceResponse]              `json:"invoice"`
	Order           Optional[GetOrderResponse]                `json:"order"`
	Customer        Optional[GetCustomerResponse]             `json:"customer"`
	Metadata        Optional[map[string]string]               `json:"metadata"`
	PaidAt          Optional[time.Time]                       `json:"paid_at"`
	CanceledAt      Optional[time.Time]                       `json:"canceled_at"`
	// Canceled Amount
	CanceledAmount Optional[int] `json:"canceled_amount"`
	// Paid amount
	PaidAmount Optional[int] `json:"paid_amount"`
	// interest and fine paid
	InterestAndFinePaid Optional[int] `json:"interest_and_fine_paid"`
	// Defines whether the card has been used one or more times.
	RecurrencyCycle Optional[string] `json:"recurrency_cycle"`
}

// MarshalJSON implements the json.Marshaler interface for GetChargeResponse.
// It customizes the JSON marshaling process for GetChargeResponse objects.
func (g *GetChargeResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetChargeResponse object to a map representation for JSON marshaling.
func (g *GetChargeResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.PaymentMethod.IsValueSet() {
		structMap["payment_method"] = g.PaymentMethod.Value()
	}
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.LastTransaction.IsValueSet() {
		structMap["last_transaction"] = g.LastTransaction.Value()
	}
	if g.Invoice.IsValueSet() {
		structMap["invoice"] = g.Invoice.Value()
	}
	if g.Order.IsValueSet() {
		structMap["order"] = g.Order.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.PaidAt.IsValueSet() {
		var PaidAtVal *string = nil
		if g.PaidAt.Value() != nil {
			val := g.PaidAt.Value().Format(time.RFC3339)
			PaidAtVal = &val
		}
		structMap["paid_at"] = PaidAtVal
	}
	if g.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if g.CanceledAt.Value() != nil {
			val := g.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if g.CanceledAmount.IsValueSet() {
		structMap["canceled_amount"] = g.CanceledAmount.Value()
	}
	if g.PaidAmount.IsValueSet() {
		structMap["paid_amount"] = g.PaidAmount.Value()
	}
	if g.InterestAndFinePaid.IsValueSet() {
		structMap["interest_and_fine_paid"] = g.InterestAndFinePaid.Value()
	}
	if g.RecurrencyCycle.IsValueSet() {
		structMap["recurrency_cycle"] = g.RecurrencyCycle.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetChargeResponse.
// It customizes the JSON unmarshaling process for GetChargeResponse objects.
func (g *GetChargeResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                  Optional[string]                          `json:"id"`
		Code                Optional[string]                          `json:"code"`
		GatewayId           Optional[string]                          `json:"gateway_id"`
		Amount              Optional[int]                             `json:"amount"`
		Status              Optional[string]                          `json:"status"`
		Currency            Optional[string]                          `json:"currency"`
		PaymentMethod       Optional[string]                          `json:"payment_method"`
		DueAt               Optional[string]                          `json:"due_at"`
		CreatedAt           Optional[string]                          `json:"created_at"`
		UpdatedAt           Optional[string]                          `json:"updated_at"`
		LastTransaction     Optional[GetTransactionResponseInterface] `json:"last_transaction"`
		Invoice             Optional[GetInvoiceResponse]              `json:"invoice"`
		Order               Optional[GetOrderResponse]                `json:"order"`
		Customer            Optional[GetCustomerResponse]             `json:"customer"`
		Metadata            Optional[map[string]string]               `json:"metadata"`
		PaidAt              Optional[string]                          `json:"paid_at"`
		CanceledAt          Optional[string]                          `json:"canceled_at"`
		CanceledAmount      Optional[int]                             `json:"canceled_amount"`
		PaidAmount          Optional[int]                             `json:"paid_amount"`
		InterestAndFinePaid Optional[int]                             `json:"interest_and_fine_paid"`
		RecurrencyCycle     Optional[string]                          `json:"recurrency_cycle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.GatewayId = temp.GatewayId
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.Currency = temp.Currency
	g.PaymentMethod = temp.PaymentMethod
	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.LastTransaction = temp.LastTransaction
	g.Invoice = temp.Invoice
	g.Order = temp.Order
	g.Customer = temp.Customer
	g.Metadata = temp.Metadata
	g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
	if temp.PaidAt.Value() != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt.SetValue(&PaidAtVal)
	}
	g.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		g.CanceledAt.SetValue(&CanceledAtVal)
	}
	g.CanceledAmount = temp.CanceledAmount
	g.PaidAmount = temp.PaidAmount
	g.InterestAndFinePaid = temp.InterestAndFinePaid
	g.RecurrencyCycle = temp.RecurrencyCycle
	return nil
}
