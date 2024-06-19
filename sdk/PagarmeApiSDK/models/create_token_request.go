package models

import (
	"encoding/json"
)

// CreateTokenRequest represents a CreateTokenRequest struct.
// Token data
type CreateTokenRequest struct {
	// Token type
	Type string `json:"type"`
	// Card data
	Card CreateCardTokenRequest `json:"card"`
}

// MarshalJSON implements the json.Marshaler interface for CreateTokenRequest.
// It customizes the JSON marshaling process for CreateTokenRequest objects.
func (c *CreateTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateTokenRequest object to a map representation for JSON marshaling.
func (c *CreateTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["card"] = c.Card
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateTokenRequest.
// It customizes the JSON unmarshaling process for CreateTokenRequest objects.
func (c *CreateTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type string                 `json:"type"`
		Card CreateCardTokenRequest `json:"card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.Card = temp.Card
	return nil
}
