package models

import (
	"encoding/json"
)

// CreateApplePayHeaderRequest represents a CreateApplePayHeaderRequest struct.
// The ApplePay header request
type CreateApplePayHeaderRequest struct {
	// SHA–256 hash, Base64 string codified
	PublicKeyHash *string `json:"public_key_hash,omitempty"`
	// X.509 encoded key bytes, Base64 encoded as a string
	EphemeralPublicKey string `json:"ephemeral_public_key"`
	// Transaction identifier, generated on Device
	TransactionId *string `json:"transaction_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateApplePayHeaderRequest.
// It customizes the JSON marshaling process for CreateApplePayHeaderRequest objects.
func (c *CreateApplePayHeaderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateApplePayHeaderRequest object to a map representation for JSON marshaling.
func (c *CreateApplePayHeaderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.PublicKeyHash != nil {
		structMap["public_key_hash"] = c.PublicKeyHash
	}
	structMap["ephemeral_public_key"] = c.EphemeralPublicKey
	if c.TransactionId != nil {
		structMap["transaction_id"] = c.TransactionId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateApplePayHeaderRequest.
// It customizes the JSON unmarshaling process for CreateApplePayHeaderRequest objects.
func (c *CreateApplePayHeaderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PublicKeyHash      *string `json:"public_key_hash,omitempty"`
		EphemeralPublicKey string  `json:"ephemeral_public_key"`
		TransactionId      *string `json:"transaction_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PublicKeyHash = temp.PublicKeyHash
	c.EphemeralPublicKey = temp.EphemeralPublicKey
	c.TransactionId = temp.TransactionId
	return nil
}
