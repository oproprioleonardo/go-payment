package models

import (
	"encoding/json"
)

// CreateVoucherPaymentRequest represents a CreateVoucherPaymentRequest struct.
// The settings for creating a voucher payment
type CreateVoucherPaymentRequest struct {
	// The text that will be shown on the voucher's statement
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Card id
	CardId *string `json:"card_id,omitempty"`
	// Card token
	CardToken *string `json:"card_token,omitempty"`
	// Card info
	Card *CreateCardRequest `json:"Card,omitempty"`
	// Defines whether the card has been used one or more times.
	RecurrencyCycle *string `json:"recurrency_cycle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateVoucherPaymentRequest.
// It customizes the JSON marshaling process for CreateVoucherPaymentRequest objects.
func (c *CreateVoucherPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateVoucherPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateVoucherPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.CardId != nil {
		structMap["card_id"] = c.CardId
	}
	if c.CardToken != nil {
		structMap["card_token"] = c.CardToken
	}
	if c.Card != nil {
		structMap["Card"] = c.Card
	}
	if c.RecurrencyCycle != nil {
		structMap["recurrency_cycle"] = c.RecurrencyCycle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateVoucherPaymentRequest.
// It customizes the JSON unmarshaling process for CreateVoucherPaymentRequest objects.
func (c *CreateVoucherPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string            `json:"statement_descriptor,omitempty"`
		CardId              *string            `json:"card_id,omitempty"`
		CardToken           *string            `json:"card_token,omitempty"`
		Card                *CreateCardRequest `json:"Card,omitempty"`
		RecurrencyCycle     *string            `json:"recurrency_cycle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.CardId = temp.CardId
	c.CardToken = temp.CardToken
	c.Card = temp.Card
	c.RecurrencyCycle = temp.RecurrencyCycle
	return nil
}
