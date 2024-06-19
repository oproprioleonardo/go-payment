package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateChargeRequest represents a CreateChargeRequest struct.
// Request for creating a new charge
type CreateChargeRequest struct {
	// Code
	Code Optional[string] `json:"code"`
	// The amount of the charge, in cents
	Amount int `json:"amount"`
	// The customer's id
	CustomerId Optional[string] `json:"customer_id"`
	// Customer data
	Customer Optional[CreateCustomerRequest] `json:"customer"`
	// Payment data
	Payment CreatePaymentRequest `json:"payment"`
	// Metadata
	Metadata Optional[map[string]string] `json:"metadata"`
	// The charge due date
	DueAt     Optional[time.Time]              `json:"due_at"`
	Antifraud Optional[CreateAntifraudRequest] `json:"antifraud"`
	// Order Id
	OrderId string `json:"order_id"`
}

// MarshalJSON implements the json.Marshaler interface for CreateChargeRequest.
// It customizes the JSON marshaling process for CreateChargeRequest objects.
func (c *CreateChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateChargeRequest object to a map representation for JSON marshaling.
func (c *CreateChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Code.IsValueSet() {
		structMap["code"] = c.Code.Value()
	}
	structMap["amount"] = c.Amount
	if c.CustomerId.IsValueSet() {
		structMap["customer_id"] = c.CustomerId.Value()
	}
	if c.Customer.IsValueSet() {
		structMap["customer"] = c.Customer.Value()
	}
	structMap["payment"] = c.Payment
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	if c.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if c.DueAt.Value() != nil {
			val := c.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	if c.Antifraud.IsValueSet() {
		structMap["antifraud"] = c.Antifraud.Value()
	}
	structMap["order_id"] = c.OrderId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateChargeRequest.
// It customizes the JSON unmarshaling process for CreateChargeRequest objects.
func (c *CreateChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code       Optional[string]                 `json:"code"`
		Amount     int                              `json:"amount"`
		CustomerId Optional[string]                 `json:"customer_id"`
		Customer   Optional[CreateCustomerRequest]  `json:"customer"`
		Payment    CreatePaymentRequest             `json:"payment"`
		Metadata   Optional[map[string]string]      `json:"metadata"`
		DueAt      Optional[string]                 `json:"due_at"`
		Antifraud  Optional[CreateAntifraudRequest] `json:"antifraud"`
		OrderId    string                           `json:"order_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Code = temp.Code
	c.Amount = temp.Amount
	c.CustomerId = temp.CustomerId
	c.Customer = temp.Customer
	c.Payment = temp.Payment
	c.Metadata = temp.Metadata
	c.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		c.DueAt.SetValue(&DueAtVal)
	}
	c.Antifraud = temp.Antifraud
	c.OrderId = temp.OrderId
	return nil
}
