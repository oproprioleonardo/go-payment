package models

import (
	"encoding/json"
)

// GetCheckoutCardInstallmentOptionsResponse represents a GetCheckoutCardInstallmentOptionsResponse struct.
type GetCheckoutCardInstallmentOptionsResponse struct {
	// Número de parcelas
	Number *string `json:"number"`
	// Valor total da compra
	Total *int `json:"total"`
}

// MarshalJSON implements the json.Marshaler interface for GetCheckoutCardInstallmentOptionsResponse.
// It customizes the JSON marshaling process for GetCheckoutCardInstallmentOptionsResponse objects.
func (g *GetCheckoutCardInstallmentOptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCheckoutCardInstallmentOptionsResponse object to a map representation for JSON marshaling.
func (g *GetCheckoutCardInstallmentOptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = g.Number
	structMap["total"] = g.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCheckoutCardInstallmentOptionsResponse.
// It customizes the JSON unmarshaling process for GetCheckoutCardInstallmentOptionsResponse objects.
func (g *GetCheckoutCardInstallmentOptionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number *string `json:"number"`
		Total  *int    `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Number = temp.Number
	g.Total = temp.Total
	return nil
}
