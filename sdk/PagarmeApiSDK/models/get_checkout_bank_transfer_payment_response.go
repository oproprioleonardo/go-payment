package models

import (
	"encoding/json"
)

// GetCheckoutBankTransferPaymentResponse represents a GetCheckoutBankTransferPaymentResponse struct.
// Bank transfer checkout response
type GetCheckoutBankTransferPaymentResponse struct {
	// bank list response
	Bank Optional[[]string] `json:"bank"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutBankTransferPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutBankTransferPaymentResponse objects.
func (g *GetCheckoutBankTransferPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutBankTransferPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutBankTransferPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Bank.IsValueSet() {
		structMap["bank"] = g.Bank.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutBankTransferPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutBankTransferPaymentResponse objects.
func (g *GetCheckoutBankTransferPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank Optional[[]string] `json:"bank"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Bank = temp.Bank
	return nil
}
