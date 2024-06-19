package models

import (
	"encoding/json"
)

// CreateInvoiceRequest represents a CreateInvoiceRequest struct.
// Request for creating a new Invoice
type CreateInvoiceRequest struct {
	// Metadata
	Metadata map[string]string `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceRequest.
// It customizes the JSON marshaling process for CreateInvoiceRequest objects.
func (c *CreateInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceRequest object to a map representation for JSON marshaling.
func (c *CreateInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["metadata"] = c.Metadata
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceRequest.
// It customizes the JSON unmarshaling process for CreateInvoiceRequest objects.
func (c *CreateInvoiceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Metadata map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Metadata = temp.Metadata
	return nil
}
