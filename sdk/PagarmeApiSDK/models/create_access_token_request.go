package models

import (
	"encoding/json"
)

// CreateAccessTokenRequest represents a CreateAccessTokenRequest struct.
// Request for creating a new Access Token
type CreateAccessTokenRequest struct {
	// Minutes to expire the token
	ExpiresIn *int `json:"expires_in,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAccessTokenRequest.
// It customizes the JSON marshaling process for CreateAccessTokenRequest objects.
func (c *CreateAccessTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAccessTokenRequest object to a map representation for JSON marshaling.
func (c *CreateAccessTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAccessTokenRequest.
// It customizes the JSON unmarshaling process for CreateAccessTokenRequest objects.
func (c *CreateAccessTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresIn *int `json:"expires_in,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ExpiresIn = temp.ExpiresIn
	return nil
}
