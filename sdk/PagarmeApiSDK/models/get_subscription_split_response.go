package models

import (
	"encoding/json"
)

// GetSubscriptionSplitResponse represents a GetSubscriptionSplitResponse struct.
type GetSubscriptionSplitResponse struct {
	// Defines if the split is enabled
	Enabled Optional[bool] `json:"enabled"`
	// Split
	Rules Optional[[]GetSplitResponse] `json:"rules"`
}

// MarshalJSON implements the json.Marshaler interface for GetSubscriptionSplitResponse.
// It customizes the JSON marshaling process for GetSubscriptionSplitResponse objects.
func (g *GetSubscriptionSplitResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSubscriptionSplitResponse object to a map representation for JSON marshaling.
func (g *GetSubscriptionSplitResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Enabled.IsValueSet() {
		structMap["enabled"] = g.Enabled.Value()
	}
	if g.Rules.IsValueSet() {
		structMap["rules"] = g.Rules.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSubscriptionSplitResponse.
// It customizes the JSON unmarshaling process for GetSubscriptionSplitResponse objects.
func (g *GetSubscriptionSplitResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled Optional[bool]               `json:"enabled"`
		Rules   Optional[[]GetSplitResponse] `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Enabled = temp.Enabled
	g.Rules = temp.Rules
	return nil
}
