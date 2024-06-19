package models

import (
	"encoding/json"
)

// GetWithdrawTargetResponse represents a GetWithdrawTargetResponse struct.
type GetWithdrawTargetResponse struct {
	TargetId Optional[string] `json:"target_id"`
	Type     Optional[string] `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for GetWithdrawTargetResponse.
// It customizes the JSON marshaling process for GetWithdrawTargetResponse objects.
func (g *GetWithdrawTargetResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetWithdrawTargetResponse object to a map representation for JSON marshaling.
func (g *GetWithdrawTargetResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.TargetId.IsValueSet() {
		structMap["target_id"] = g.TargetId.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetWithdrawTargetResponse.
// It customizes the JSON unmarshaling process for GetWithdrawTargetResponse objects.
func (g *GetWithdrawTargetResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TargetId Optional[string] `json:"target_id"`
		Type     Optional[string] `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.TargetId = temp.TargetId
	g.Type = temp.Type
	return nil
}
