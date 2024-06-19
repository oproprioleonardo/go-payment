package models

import (
	"encoding/json"
)

// GetCheckoutPaymentSettingsResponse represents a GetCheckoutPaymentSettingsResponse struct.
// Checkout Payment Settings Response
type GetCheckoutPaymentSettingsResponse struct {
	// Success Url
	SuccessUrl Optional[string] `json:"success_url"`
	// Payment Url
	PaymentUrl Optional[string] `json:"payment_url"`
	// Accepted Payment Methods
	AcceptedPaymentMethods Optional[[]string] `json:"accepted_payment_methods"`
	// Status
	Status Optional[string] `json:"status"`
	// Customer
	Customer Optional[GetCustomerResponse] `json:"customer"`
	// Payment amount
	Amount Optional[int] `json:"amount"`
	// Default Payment Method
	DefaultPaymentMethod Optional[string] `json:"default_payment_method"`
	// Gateway Affiliation Id
	GatewayAffiliationId Optional[string] `json:"gateway_affiliation_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutPaymentSettingsResponse.
// It customizes the JSON marshaling process for GetCheckoutPaymentSettingsResponse objects.
func (g *GetCheckoutPaymentSettingsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutPaymentSettingsResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutPaymentSettingsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.SuccessUrl.IsValueSet() {
		structMap["success_url"] = g.SuccessUrl.Value()
	}
	if g.PaymentUrl.IsValueSet() {
		structMap["payment_url"] = g.PaymentUrl.Value()
	}
	if g.AcceptedPaymentMethods.IsValueSet() {
		structMap["accepted_payment_methods"] = g.AcceptedPaymentMethods.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.DefaultPaymentMethod.IsValueSet() {
		structMap["default_payment_method"] = g.DefaultPaymentMethod.Value()
	}
	if g.GatewayAffiliationId.IsValueSet() {
		structMap["gateway_affiliation_id"] = g.GatewayAffiliationId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutPaymentSettingsResponse.
// It customizes the JSON unmarshaling process for GetCheckoutPaymentSettingsResponse objects.
func (g *GetCheckoutPaymentSettingsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SuccessUrl             Optional[string]              `json:"success_url"`
		PaymentUrl             Optional[string]              `json:"payment_url"`
		AcceptedPaymentMethods Optional[[]string]            `json:"accepted_payment_methods"`
		Status                 Optional[string]              `json:"status"`
		Customer               Optional[GetCustomerResponse] `json:"customer"`
		Amount                 Optional[int]                 `json:"amount"`
		DefaultPaymentMethod   Optional[string]              `json:"default_payment_method"`
		GatewayAffiliationId   Optional[string]              `json:"gateway_affiliation_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SuccessUrl = temp.SuccessUrl
	g.PaymentUrl = temp.PaymentUrl
	g.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
	g.Status = temp.Status
	g.Customer = temp.Customer
	g.Amount = temp.Amount
	g.DefaultPaymentMethod = temp.DefaultPaymentMethod
	g.GatewayAffiliationId = temp.GatewayAffiliationId
	return nil
}
