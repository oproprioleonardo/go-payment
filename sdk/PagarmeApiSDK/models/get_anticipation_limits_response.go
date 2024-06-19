package models

import (
	"encoding/json"
)

// GetAnticipationLimitsResponse represents a GetAnticipationLimitsResponse struct.
// Anticipation limits
type GetAnticipationLimitsResponse struct {
	// Max limit
	Max Optional[GetAnticipationLimitResponse] `json:"max"`
	// Min limit
	Min Optional[GetAnticipationLimitResponse] `json:"min"`
}

// MarshalJSON implements the json.Marshaler interface for GetAnticipationLimitsResponse.
// It customizes the JSON marshaling process for GetAnticipationLimitsResponse objects.
func (g *GetAnticipationLimitsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetAnticipationLimitsResponse object to a map representation for JSON marshaling.
func (g *GetAnticipationLimitsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Max.IsValueSet() {
		structMap["max"] = g.Max.Value()
	}
	if g.Min.IsValueSet() {
		structMap["min"] = g.Min.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAnticipationLimitsResponse.
// It customizes the JSON unmarshaling process for GetAnticipationLimitsResponse objects.
func (g *GetAnticipationLimitsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Max Optional[GetAnticipationLimitResponse] `json:"max"`
		Min Optional[GetAnticipationLimitResponse] `json:"min"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Max = temp.Max
	g.Min = temp.Min
	return nil
}
