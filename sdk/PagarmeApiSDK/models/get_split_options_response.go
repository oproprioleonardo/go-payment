package models

import (
	"encoding/json"
)

// GetSplitOptionsResponse represents a GetSplitOptionsResponse struct.
type GetSplitOptionsResponse struct {
	Liable              Optional[bool]   `json:"liable"`
	ChargeProcessingFee Optional[bool]   `json:"charge_processing_fee"`
	ChargeRemainderFee  Optional[string] `json:"charge_remainder_fee"`
}

// MarshalJSON implements the json.Marshaler interface for GetSplitOptionsResponse.
// It customizes the JSON marshaling process for GetSplitOptionsResponse objects.
func (g *GetSplitOptionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSplitOptionsResponse object to a map representation for JSON marshaling.
func (g *GetSplitOptionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Liable.IsValueSet() {
		structMap["liable"] = g.Liable.Value()
	}
	if g.ChargeProcessingFee.IsValueSet() {
		structMap["charge_processing_fee"] = g.ChargeProcessingFee.Value()
	}
	if g.ChargeRemainderFee.IsValueSet() {
		structMap["charge_remainder_fee"] = g.ChargeRemainderFee.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSplitOptionsResponse.
// It customizes the JSON unmarshaling process for GetSplitOptionsResponse objects.
func (g *GetSplitOptionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Liable              Optional[bool]   `json:"liable"`
		ChargeProcessingFee Optional[bool]   `json:"charge_processing_fee"`
		ChargeRemainderFee  Optional[string] `json:"charge_remainder_fee"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Liable = temp.Liable
	g.ChargeProcessingFee = temp.ChargeProcessingFee
	g.ChargeRemainderFee = temp.ChargeRemainderFee
	return nil
}
