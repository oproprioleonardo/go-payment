package models

import (
	"encoding/json"
)

// GetInterestResponse represents a GetInterestResponse struct.
// Interest Response
type GetInterestResponse struct {
	// Days
	Days Optional[int] `json:"days"`
	// Type
	Type Optional[string] `json:"type"`
	// Amount
	Amount Optional[int] `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetInterestResponse.
// It customizes the JSON marshaling process for GetInterestResponse objects.
func (g *GetInterestResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetInterestResponse object to a map representation for JSON marshaling.
func (g *GetInterestResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Days.IsValueSet() {
		structMap["days"] = g.Days.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetInterestResponse.
// It customizes the JSON unmarshaling process for GetInterestResponse objects.
func (g *GetInterestResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Days   Optional[int]    `json:"days"`
		Type   Optional[string] `json:"type"`
		Amount Optional[int]    `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Days = temp.Days
	g.Type = temp.Type
	g.Amount = temp.Amount
	return nil
}
