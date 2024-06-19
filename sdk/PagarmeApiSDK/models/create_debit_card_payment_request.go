package models

import (
	"encoding/json"
)

// CreateDebitCardPaymentRequest represents a CreateDebitCardPaymentRequest struct.
// The settings for creating a debit card payment
type CreateDebitCardPaymentRequest struct {
	// The text that will be shown on the debit card's statement
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Debit card data
	Card *CreateCardRequest `json:"card,omitempty"`
	// The debit card id
	CardId *string `json:"card_id,omitempty"`
	// The debit card token
	CardToken *string `json:"card_token,omitempty"`
	// Indicates a recurrence
	Recurrence *bool `json:"recurrence,omitempty"`
	// The payment authentication request
	Authentication *CreatePaymentAuthenticationRequest `json:"authentication,omitempty"`
	// The Debit card payment token request
	Token *CreateCardPaymentContactlessRequest `json:"token,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateDebitCardPaymentRequest.
// It customizes the JSON marshaling process for CreateDebitCardPaymentRequest objects.
func (c *CreateDebitCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateDebitCardPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateDebitCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.Card != nil {
		structMap["card"] = c.Card
	}
	if c.CardId != nil {
		structMap["card_id"] = c.CardId
	}
	if c.CardToken != nil {
		structMap["card_token"] = c.CardToken
	}
	if c.Recurrence != nil {
		structMap["recurrence"] = c.Recurrence
	}
	if c.Authentication != nil {
		structMap["authentication"] = c.Authentication
	}
	if c.Token != nil {
		structMap["token"] = c.Token
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateDebitCardPaymentRequest.
// It customizes the JSON unmarshaling process for CreateDebitCardPaymentRequest objects.
func (c *CreateDebitCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                              `json:"statement_descriptor,omitempty"`
		Card                *CreateCardRequest                   `json:"card,omitempty"`
		CardId              *string                              `json:"card_id,omitempty"`
		CardToken           *string                              `json:"card_token,omitempty"`
		Recurrence          *bool                                `json:"recurrence,omitempty"`
		Authentication      *CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
		Token               *CreateCardPaymentContactlessRequest `json:"token,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Card = temp.Card
	c.CardId = temp.CardId
	c.CardToken = temp.CardToken
	c.Recurrence = temp.Recurrence
	c.Authentication = temp.Authentication
	c.Token = temp.Token
	return nil
}
