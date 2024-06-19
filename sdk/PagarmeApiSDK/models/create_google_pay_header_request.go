package models

import (
	"encoding/json"
)

// CreateGooglePayHeaderRequest represents a CreateGooglePayHeaderRequest struct.
// The GooglePay header request
type CreateGooglePayHeaderRequest struct {
	// X.509 encoded key bytes, Base64 encoded as a string
	EphemeralPublicKey string `json:"ephemeral_public_key"`
}

// MarshalJSON implements the json.Marshaler interface for CreateGooglePayHeaderRequest.
// It customizes the JSON marshaling process for CreateGooglePayHeaderRequest objects.
func (c *CreateGooglePayHeaderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateGooglePayHeaderRequest object to a map representation for JSON marshaling.
func (c *CreateGooglePayHeaderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["ephemeral_public_key"] = c.EphemeralPublicKey
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateGooglePayHeaderRequest.
// It customizes the JSON unmarshaling process for CreateGooglePayHeaderRequest objects.
func (c *CreateGooglePayHeaderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EphemeralPublicKey string `json:"ephemeral_public_key"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.EphemeralPublicKey = temp.EphemeralPublicKey
	return nil
}
