package models

import (
	"encoding/json"
)

// GetTransferSourceResponse represents a GetTransferSourceResponse struct.
type GetTransferSourceResponse struct {
	SourceId Optional[string] `json:"source_id"`
	Type     Optional[string] `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransferSourceResponse.
// It customizes the JSON marshaling process for GetTransferSourceResponse objects.
func (g *GetTransferSourceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetTransferSourceResponse object to a map representation for JSON marshaling.
func (g *GetTransferSourceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.SourceId.IsValueSet() {
		structMap["source_id"] = g.SourceId.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransferSourceResponse.
// It customizes the JSON unmarshaling process for GetTransferSourceResponse objects.
func (g *GetTransferSourceResponse) UnmarshalJSON(input []byte) error {
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
