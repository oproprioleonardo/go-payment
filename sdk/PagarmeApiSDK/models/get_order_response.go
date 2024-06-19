package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetOrderResponse represents a GetOrderResponse struct.
// Response object for getting an Order
type GetOrderResponse struct {
	Id       Optional[string] `json:"id"`
	Code     Optional[string] `json:"code"`
	Amount   Optional[int]    `json:"amount"`
	Currency Optional[string] `json:"currency"`
	// Indicates whether the order is closed
	Closed     Optional[bool]                   `json:"closed"`
	Items      Optional[[]GetOrderItemResponse] `json:"items"`
	Customer   Optional[GetCustomerResponse]    `json:"customer"`
	Status     Optional[string]                 `json:"status"`
	CreatedAt  Optional[time.Time]              `json:"created_at"`
	UpdatedAt  Optional[time.Time]              `json:"updated_at"`
	ClosedAt   Optional[time.Time]              `json:"closed_at"`
	Charges    Optional[[]GetChargeResponse]    `json:"charges"`
	InvoiceUrl Optional[string]                 `json:"invoice_url"`
	Shipping   Optional[GetShippingResponse]    `json:"shipping"`
	Metadata   Optional[map[string]string]      `json:"metadata"`
	// Checkout Payment Settings Response
	Checkouts Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
	// Ip address
	Ip Optional[string] `json:"ip"`
	// Session id
	SessionId Optional[string] `json:"session_id"`
	// Location
	Location Optional[GetLocationResponse] `json:"location"`
	// Device's informations
	Device Optional[GetDeviceResponse] `json:"device"`
}

// MarshalJSON implements the json.Marshaler interface for GetOrderResponse.
// It customizes the JSON marshaling process for GetOrderResponse objects.
func (g *GetOrderResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetOrderResponse object to a map representation for JSON marshaling.
func (g *GetOrderResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.Closed.IsValueSet() {
		structMap["closed"] = g.Closed.Value()
	}
	if g.Items.IsValueSet() {
		structMap["items"] = g.Items.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
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
	if g.ClosedAt.IsValueSet() {
		var ClosedAtVal *string = nil
		if g.ClosedAt.Value() != nil {
			val := g.ClosedAt.Value().Format(time.RFC3339)
			ClosedAtVal = &val
		}
		structMap["closed_at"] = ClosedAtVal
	}
	if g.Charges.IsValueSet() {
		structMap["charges"] = g.Charges.Value()
	}
	if g.InvoiceUrl.IsValueSet() {
		structMap["invoice_url"] = g.InvoiceUrl.Value()
	}
	if g.Shipping.IsValueSet() {
		structMap["shipping"] = g.Shipping.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Checkouts.IsValueSet() {
		structMap["checkouts"] = g.Checkouts.Value()
	}
	if g.Ip.IsValueSet() {
		structMap["ip"] = g.Ip.Value()
	}
	if g.SessionId.IsValueSet() {
		structMap["session_id"] = g.SessionId.Value()
	}
	if g.Location.IsValueSet() {
		structMap["location"] = g.Location.Value()
	}
	if g.Device.IsValueSet() {
		structMap["device"] = g.Device.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrderResponse.
// It customizes the JSON unmarshaling process for GetOrderResponse objects.
func (g *GetOrderResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id         Optional[string]                       `json:"id"`
		Code       Optional[string]                       `json:"code"`
		Amount     Optional[int]                          `json:"amount"`
		Currency   Optional[string]                       `json:"currency"`
		Closed     Optional[bool]                         `json:"closed"`
		Items      Optional[[]GetOrderItemResponse]       `json:"items"`
		Customer   Optional[GetCustomerResponse]          `json:"customer"`
		Status     Optional[string]                       `json:"status"`
		CreatedAt  Optional[string]                       `json:"created_at"`
		UpdatedAt  Optional[string]                       `json:"updated_at"`
		ClosedAt   Optional[string]                       `json:"closed_at"`
		Charges    Optional[[]GetChargeResponse]          `json:"charges"`
		InvoiceUrl Optional[string]                       `json:"invoice_url"`
		Shipping   Optional[GetShippingResponse]          `json:"shipping"`
		Metadata   Optional[map[string]string]            `json:"metadata"`
		Checkouts  Optional[[]GetCheckoutPaymentResponse] `json:"checkouts"`
		Ip         Optional[string]                       `json:"ip"`
		SessionId  Optional[string]                       `json:"session_id"`
		Location   Optional[GetLocationResponse]          `json:"location"`
		Device     Optional[GetDeviceResponse]            `json:"device"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Code = temp.Code
	g.Amount = temp.Amount
	g.Currency = temp.Currency
	g.Closed = temp.Closed
	g.Items = temp.Items
	g.Customer = temp.Customer
	g.Status = temp.Status
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
	g.ClosedAt.ShouldSetValue(temp.ClosedAt.IsValueSet())
	if temp.ClosedAt.Value() != nil {
		ClosedAtVal, err := time.Parse(time.RFC3339, (*temp.ClosedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse closed_at as % s format.", time.RFC3339)
		}
		g.ClosedAt.SetValue(&ClosedAtVal)
	}
	g.Charges = temp.Charges
	g.InvoiceUrl = temp.InvoiceUrl
	g.Shipping = temp.Shipping
	g.Metadata = temp.Metadata
	g.Checkouts = temp.Checkouts
	g.Ip = temp.Ip
	g.SessionId = temp.SessionId
	g.Location = temp.Location
	g.Device = temp.Device
	return nil
}
