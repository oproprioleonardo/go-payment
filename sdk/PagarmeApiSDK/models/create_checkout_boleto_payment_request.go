package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateCheckoutBoletoPaymentRequest represents a CreateCheckoutBoletoPaymentRequest struct.
type CreateCheckoutBoletoPaymentRequest struct {
	// Bank identifier
	Bank string `json:"bank"`
	// Instructions
	Instructions string `json:"instructions"`
	// Due date
	DueAt time.Time `json:"due_at"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutBoletoPaymentRequest.
// It customizes the JSON marshaling process for CreateCheckoutBoletoPaymentRequest objects.
func (c *CreateCheckoutBoletoPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutBoletoPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutBoletoPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank"] = c.Bank
	structMap["instructions"] = c.Instructions
	structMap["due_at"] = c.DueAt.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutBoletoPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutBoletoPaymentRequest objects.
func (c *CreateCheckoutBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank         string `json:"bank"`
		Instructions string `json:"instructions"`
		DueAt        string `json:"due_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Bank = temp.Bank
	c.Instructions = temp.Instructions
	DueAtVal, err := time.Parse(time.RFC3339, temp.DueAt)
	if err != nil {
		log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
	}
	c.DueAt = DueAtVal
	return nil
}
