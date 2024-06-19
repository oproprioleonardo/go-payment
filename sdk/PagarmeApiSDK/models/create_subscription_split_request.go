package models

import (
	"encoding/json"
)

// CreateSubscriptionSplitRequest represents a CreateSubscriptionSplitRequest struct.
type CreateSubscriptionSplitRequest struct {
	// Defines if the split is enabled
	Enabled bool `json:"enabled"`
	// Split
	Rules []CreateSplitRequest `json:"rules"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionSplitRequest.
// It customizes the JSON marshaling process for CreateSubscriptionSplitRequest objects.
func (c *CreateSubscriptionSplitRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionSplitRequest object to a map representation for JSON marshaling.
func (c *CreateSubscriptionSplitRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = c.Enabled
	structMap["rules"] = c.Rules
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionSplitRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionSplitRequest objects.
func (c *CreateSubscriptionSplitRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled bool                 `json:"enabled"`
		Rules   []CreateSplitRequest `json:"rules"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Enabled = temp.Enabled
	c.Rules = temp.Rules
	return nil
}
