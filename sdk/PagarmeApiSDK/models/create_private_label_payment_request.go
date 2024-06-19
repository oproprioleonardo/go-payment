package models

import (
	"encoding/json"
)

// CreatePrivateLabelPaymentRequest represents a CreatePrivateLabelPaymentRequest struct.
// The settings for creating a private label payment
type CreatePrivateLabelPaymentRequest struct {
	// Number of installments
	Installments *int `json:"installments,omitempty"`
	// The text that will be shown on the private label's statement
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Card data
	Card *CreateCardRequest `json:"card,omitempty"`
	// The Card id
	CardId    *string `json:"card_id,omitempty"`
	CardToken *string `json:"card_token,omitempty"`
	// Indicates a recurrence
	Recurrence *bool `json:"recurrence,omitempty"`
	// Indicates if the operation should be only authorization or auth and capture.
	Capture *bool `json:"capture,omitempty"`
	// Indicates whether the extended label (private label) is enabled
	ExtendedLimitEnabled *bool `json:"extended_limit_enabled,omitempty"`
	// Extended Limit Code
	ExtendedLimitCode *string `json:"extended_limit_code,omitempty"`
	// Defines whether the card has been used one or more times.
	RecurrencyCycle *string `json:"recurrency_cycle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrivateLabelPaymentRequest.
// It customizes the JSON marshaling process for CreatePrivateLabelPaymentRequest objects.
func (c *CreatePrivateLabelPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePrivateLabelPaymentRequest object to a map representation for JSON marshaling.
func (c *CreatePrivateLabelPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Installments != nil {
		structMap["installments"] = c.Installments
	}
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
	if c.Capture != nil {
		structMap["capture"] = c.Capture
	}
	if c.ExtendedLimitEnabled != nil {
		structMap["extended_limit_enabled"] = c.ExtendedLimitEnabled
	}
	if c.ExtendedLimitCode != nil {
		structMap["extended_limit_code"] = c.ExtendedLimitCode
	}
	if c.RecurrencyCycle != nil {
		structMap["recurrency_cycle"] = c.RecurrencyCycle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrivateLabelPaymentRequest.
// It customizes the JSON unmarshaling process for CreatePrivateLabelPaymentRequest objects.
func (c *CreatePrivateLabelPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Installments         *int               `json:"installments,omitempty"`
		StatementDescriptor  *string            `json:"statement_descriptor,omitempty"`
		Card                 *CreateCardRequest `json:"card,omitempty"`
		CardId               *string            `json:"card_id,omitempty"`
		CardToken            *string            `json:"card_token,omitempty"`
		Recurrence           *bool              `json:"recurrence,omitempty"`
		Capture              *bool              `json:"capture,omitempty"`
		ExtendedLimitEnabled *bool              `json:"extended_limit_enabled,omitempty"`
		ExtendedLimitCode    *string            `json:"extended_limit_code,omitempty"`
		RecurrencyCycle      *string            `json:"recurrency_cycle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Installments = temp.Installments
	c.StatementDescriptor = temp.StatementDescriptor
	c.Card = temp.Card
	c.CardId = temp.CardId
	c.CardToken = temp.CardToken
	c.Recurrence = temp.Recurrence
	c.Capture = temp.Capture
	c.ExtendedLimitEnabled = temp.ExtendedLimitEnabled
	c.ExtendedLimitCode = temp.ExtendedLimitCode
	c.RecurrencyCycle = temp.RecurrencyCycle
	return nil
}
