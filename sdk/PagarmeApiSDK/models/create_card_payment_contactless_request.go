package models

import (
	"encoding/json"
)

// CreateCardPaymentContactlessRequest represents a CreateCardPaymentContactlessRequest struct.
// The card payment contactless request
type CreateCardPaymentContactlessRequest struct {
	// The authentication type
	Type string `json:"type"`
	// The ApplePay encrypted request
	ApplePay *CreateApplePayRequest `json:"apple_pay,omitempty"`
	// The GooglePay encrypted request
	GooglePay *CreateGooglePayRequest `json:"google_pay,omitempty"`
	// The Emv encrypted request
	Emv *CreateEmvDecryptRequest `json:"emv,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCardPaymentContactlessRequest.
// It customizes the JSON marshaling process for CreateCardPaymentContactlessRequest objects.
func (c *CreateCardPaymentContactlessRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCardPaymentContactlessRequest object to a map representation for JSON marshaling.
func (c *CreateCardPaymentContactlessRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	if c.ApplePay != nil {
		structMap["apple_pay"] = c.ApplePay
	}
	if c.GooglePay != nil {
		structMap["google_pay"] = c.GooglePay
	}
	if c.Emv != nil {
		structMap["emv"] = c.Emv
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCardPaymentContactlessRequest.
// It customizes the JSON unmarshaling process for CreateCardPaymentContactlessRequest objects.
func (c *CreateCardPaymentContactlessRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      string                   `json:"type"`
		ApplePay  *CreateApplePayRequest   `json:"apple_pay,omitempty"`
		GooglePay *CreateGooglePayRequest  `json:"google_pay,omitempty"`
		Emv       *CreateEmvDecryptRequest `json:"emv,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.ApplePay = temp.ApplePay
	c.GooglePay = temp.GooglePay
	c.Emv = temp.Emv
	return nil
}
