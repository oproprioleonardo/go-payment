package models

import (
	"encoding/json"
)

// CreateApplePayRequest represents a CreateApplePayRequest struct.
// The ApplePay Token Payment Request
type CreateApplePayRequest struct {
	// The token version
	Version string `json:"version"`
	// The cryptography data
	Data string `json:"data"`
	// The ApplePay header request
	Header CreateApplePayHeaderRequest `json:"header"`
	// Detached PKCS #7 signature, Base64 encoded as string
	Signature string `json:"signature"`
	// ApplePay Merchant identifier
	MerchantIdentifier string `json:"merchant_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for CreateApplePayRequest.
// It customizes the JSON marshaling process for CreateApplePayRequest objects.
func (c *CreateApplePayRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateApplePayRequest object to a map representation for JSON marshaling.
func (c *CreateApplePayRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["version"] = c.Version
	structMap["data"] = c.Data
	structMap["header"] = c.Header
	structMap["signature"] = c.Signature
	structMap["merchant_identifier"] = c.MerchantIdentifier
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateApplePayRequest.
// It customizes the JSON unmarshaling process for CreateApplePayRequest objects.
func (c *CreateApplePayRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Version            string                      `json:"version"`
		Data               string                      `json:"data"`
		Header             CreateApplePayHeaderRequest `json:"header"`
		Signature          string                      `json:"signature"`
		MerchantIdentifier string                      `json:"merchant_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Version = temp.Version
	c.Data = temp.Data
	c.Header = temp.Header
	c.Signature = temp.Signature
	c.MerchantIdentifier = temp.MerchantIdentifier
	return nil
}
