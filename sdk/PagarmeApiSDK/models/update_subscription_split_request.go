package models

import (
	"encoding/json"
)

// UpdateSubscriptionSplitRequest represents a UpdateSubscriptionSplitRequest struct.
type UpdateSubscriptionSplitRequest struct {
	// Defines if the split is enabled
	Enabled bool `json:"enabled"`
	// Split
	Rules []CreateSplitRequest `json:"rules"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionSplitRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionSplitRequest objects.
func (u *UpdateSubscriptionSplitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionSplitRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionSplitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = u.Enabled
	structMap["rules"] = u.Rules
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionSplitRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionSplitRequest objects.
func (u *UpdateSubscriptionSplitRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled bool                 `json:"enabled"`
		Rules   []CreateSplitRequest `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Enabled = temp.Enabled
	u.Rules = temp.Rules
	return nil
}
