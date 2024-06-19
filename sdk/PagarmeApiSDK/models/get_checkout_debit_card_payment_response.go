package models

import (
	"encoding/json"
)

// GetCheckoutDebitCardPaymentResponse represents a GetCheckoutDebitCardPaymentResponse struct.
type GetCheckoutDebitCardPaymentResponse struct {
	// Descrição na fatura
	StatementDescriptor Optional[string] `json:"statement_descriptor"`
	// Payment Authentication response object data
	Authentication Optional[GetPaymentAuthenticationResponse] `json:"authentication"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutDebitCardPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutDebitCardPaymentResponse objects.
func (g *GetCheckoutDebitCardPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutDebitCardPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutDebitCardPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.Authentication.IsValueSet() {
		structMap["authentication"] = g.Authentication.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutDebitCardPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutDebitCardPaymentResponse objects.
func (g *GetCheckoutDebitCardPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor Optional[string]                           `json:"statement_descriptor"`
		Authentication      Optional[GetPaymentAuthenticationResponse] `json:"authentication"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StatementDescriptor = temp.StatementDescriptor
	g.Authentication = temp.Authentication
	return nil
}
