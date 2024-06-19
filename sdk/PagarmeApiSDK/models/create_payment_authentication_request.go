package models

import (
	"encoding/json"
)

// CreatePaymentAuthenticationRequest represents a CreatePaymentAuthenticationRequest struct.
// The payment authentication request
type CreatePaymentAuthenticationRequest struct {
	// The Authentication type
	Type string `json:"type"`
	// The 3D-S authentication object
	ThreedSecure CreateThreeDSecureRequest `json:"threed_secure"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentAuthenticationRequest.
// It customizes the JSON marshaling process for CreatePaymentAuthenticationRequest objects.
func (c *CreatePaymentAuthenticationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentAuthenticationRequest object to a map representation for JSON marshaling.
func (c *CreatePaymentAuthenticationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["threed_secure"] = c.ThreedSecure
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentAuthenticationRequest.
// It customizes the JSON unmarshaling process for CreatePaymentAuthenticationRequest objects.
func (c *CreatePaymentAuthenticationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type         string                    `json:"type"`
		ThreedSecure CreateThreeDSecureRequest `json:"threed_secure"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.ThreedSecure = temp.ThreedSecure
	return nil
}
