package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateBoletoPaymentRequest represents a CreateBoletoPaymentRequest struct.
// Contains the settings for creating a boleto payment
type CreateBoletoPaymentRequest struct {
	// Number of retries
	Retries int `json:"retries"`
	// The bank code, containing three characters. The available codes are on the API specification
	Bank Optional[string] `json:"bank"`
	// The instructions field that will be printed on the boleto.
	Instructions string `json:"instructions"`
	// Boleto due date
	DueAt Optional[time.Time] `json:"due_at"`
	// Card's billing address
	BillingAddress CreateAddressRequest `json:"billing_address"`
	// The address id for the billing address
	BillingAddressId Optional[string] `json:"billing_address_id"`
	// Customer identification number with the bank
	NossoNumero Optional[string] `json:"nosso_numero"`
	// Boleto identification
	DocumentNumber string `json:"document_number"`
	// Soft Descriptor
	StatementDescriptor string                          `json:"statement_descriptor"`
	Interest            Optional[CreateInterestRequest] `json:"interest"`
	Fine                Optional[CreateFineRequest]     `json:"fine"`
	MaxDaysToPayPastDue Optional[int]                   `json:"max_days_to_pay_past_due"`
}

// MarshalJSON implements the json.Marshaler interface for CreateBoletoPaymentRequest.
// It customizes the JSON marshaling process for CreateBoletoPaymentRequest objects.
func (c *CreateBoletoPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateBoletoPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateBoletoPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["retries"] = c.Retries
	if c.Bank.IsValueSet() {
		structMap["bank"] = c.Bank.Value()
	}
	structMap["instructions"] = c.Instructions
	if c.DueAt.IsValueSet() {
		var DueAtVal *string = nil
		if c.DueAt.Value() != nil {
			val := c.DueAt.Value().Format(time.RFC3339)
			DueAtVal = &val
		}
		structMap["due_at"] = DueAtVal
	}
	structMap["billing_address"] = c.BillingAddress
	if c.BillingAddressId.IsValueSet() {
		structMap["billing_address_id"] = c.BillingAddressId.Value()
	}
	if c.NossoNumero.IsValueSet() {
		structMap["nosso_numero"] = c.NossoNumero.Value()
	}
	structMap["document_number"] = c.DocumentNumber
	structMap["statement_descriptor"] = c.StatementDescriptor
	if c.Interest.IsValueSet() {
		structMap["interest"] = c.Interest.Value()
	}
	if c.Fine.IsValueSet() {
		structMap["fine"] = c.Fine.Value()
	}
	if c.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = c.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateBoletoPaymentRequest.
// It customizes the JSON unmarshaling process for CreateBoletoPaymentRequest objects.
func (c *CreateBoletoPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Retries             int                             `json:"retries"`
		Bank                Optional[string]                `json:"bank"`
		Instructions        string                          `json:"instructions"`
		DueAt               Optional[string]                `json:"due_at"`
		BillingAddress      CreateAddressRequest            `json:"billing_address"`
		BillingAddressId    Optional[string]                `json:"billing_address_id"`
		NossoNumero         Optional[string]                `json:"nosso_numero"`
		DocumentNumber      string                          `json:"document_number"`
		StatementDescriptor string                          `json:"statement_descriptor"`
		Interest            Optional[CreateInterestRequest] `json:"interest"`
		Fine                Optional[CreateFineRequest]     `json:"fine"`
		MaxDaysToPayPastDue Optional[int]                   `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Retries = temp.Retries
	c.Bank = temp.Bank
	c.Instructions = temp.Instructions
	c.DueAt.ShouldSetValue(temp.DueAt.IsValueSet())
	if temp.DueAt.Value() != nil {
		DueAtVal, err := time.Parse(time.RFC3339, (*temp.DueAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		c.DueAt.SetValue(&DueAtVal)
	}
	c.BillingAddress = temp.BillingAddress
	c.BillingAddressId = temp.BillingAddressId
	c.NossoNumero = temp.NossoNumero
	c.DocumentNumber = temp.DocumentNumber
	c.StatementDescriptor = temp.StatementDescriptor
	c.Interest = temp.Interest
	c.Fine = temp.Fine
	c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
