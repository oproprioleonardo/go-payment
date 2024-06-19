package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetCheckoutBoletoPaymentResponse represents a GetCheckoutBoletoPaymentResponse struct.
type GetCheckoutBoletoPaymentResponse struct {
	// Data de vencimento do boleto
	DueAt Optional[time.Time] `json:"due_at"`
	// Instruções do boleto
	Instructions Optional[string] `json:"instructions"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutBoletoPaymentResponse.
// It customizes the JSON marshaling process for GetCheckoutBoletoPaymentResponse objects.
func (g *GetCheckoutBoletoPaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutBoletoPaymentResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutBoletoPaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if g.DueAt.Value() != nil {
			val := g.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if g.Instructions.IsValueSet() {
		structMap["instructions"] = g.Instructions.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutBoletoPaymentResponse.
// It customizes the JSON unmarshaling process for GetCheckoutBoletoPaymentResponse objects.
func (g *GetCheckoutBoletoPaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		DueAt        Optional[string] `json:"due_at"`
		Instructions Optional[string] `json:"instructions"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		g.DueAt.SetValue(&DueAtVal)
	}
	g.Instructions = temp.Instructions
	return nil
}
