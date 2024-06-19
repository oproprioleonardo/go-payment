package models

import (
	"encoding/json"
)

// GetChargesSummaryResponse represents a GetChargesSummaryResponse struct.
type GetChargesSummaryResponse struct {
	Total Optional[int] `json:"total"`
}

// MarshalJSON implements the json.Marshaler interface for GetChargesSummaryResponse.
// It customizes the JSON marshaling process for GetChargesSummaryResponse objects.
func (g *GetChargesSummaryResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetChargesSummaryResponse object to a map representation for JSON marshaling.
func (g *GetChargesSummaryResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Total.IsValueSet() {
		structMap["total"] = g.Total.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetChargesSummaryResponse.
// It customizes the JSON unmarshaling process for GetChargesSummaryResponse objects.
func (g *GetChargesSummaryResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Total Optional[int] `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Total = temp.Total
	return nil
}
