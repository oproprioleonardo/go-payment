package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateCheckoutPixPaymentRequest represents a CreateCheckoutPixPaymentRequest struct.
// Checkout pix payment request
type CreateCheckoutPixPaymentRequest struct {
	// Expires at
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// Expires in
	ExpiresIn *int `json:"expires_in,omitempty"`
	// Additional information
	AdditionalInformation []PixAdditionalInformation `json:"additional_information,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutPixPaymentRequest.
// It customizes the JSON marshaling process for CreateCheckoutPixPaymentRequest objects.
func (c *CreateCheckoutPixPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutPixPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutPixPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ExpiresAt != nil {
		structMap["expires_at"] = c.ExpiresAt.Format(time.RFC3339)
	}
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	if c.AdditionalInformation != nil {
		structMap["additional_information"] = c.AdditionalInformation
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutPixPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutPixPaymentRequest objects.
func (c *CreateCheckoutPixPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresAt             *string                    `json:"expires_at,omitempty"`
		ExpiresIn             *int                       `json:"expires_in,omitempty"`
		AdditionalInformation []PixAdditionalInformation `json:"additional_information,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.ExpiresAt != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		c.ExpiresAt = &ExpiresAtVal
	}
	c.ExpiresIn = temp.ExpiresIn
	c.AdditionalInformation = temp.AdditionalInformation
	return nil
}
