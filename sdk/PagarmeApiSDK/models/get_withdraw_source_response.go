package models

import (
	"encoding/json"
)

// GetWithdrawSourceResponse represents a GetWithdrawSourceResponse struct.
type GetWithdrawSourceResponse struct {
	SourceId Optional[string] `json:"source_id"`
	Type     Optional[string] `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for GetWithdrawSourceResponse.
// It customizes the JSON marshaling process for GetWithdrawSourceResponse objects.
func (g *GetWithdrawSourceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetWithdrawSourceResponse object to a map representation for JSON marshaling.
func (g *GetWithdrawSourceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.SourceId.IsValueSet() {
		structMap["source_id"] = g.SourceId.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetWithdrawSourceResponse.
// It customizes the JSON unmarshaling process for GetWithdrawSourceResponse objects.
func (g *GetWithdrawSourceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SourceId Optional[string] `json:"source_id"`
		Type     Optional[string] `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SourceId = temp.SourceId
	g.Type = temp.Type
	return nil
}
