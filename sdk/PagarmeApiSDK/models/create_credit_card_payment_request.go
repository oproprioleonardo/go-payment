package models

import (
	"encoding/json"
)

// CreateCreditCardPaymentRequest represents a CreateCreditCardPaymentRequest struct.
// The settings for creating a credit card payment
type CreateCreditCardPaymentRequest struct {
	// Number of installments
	Installments *int `json:"installments,omitempty"`
	// The text that will be shown on the credit card's statement
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Credit card data
	Card *CreateCardRequest `json:"card,omitempty"`
	// The credit card id
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
	// Customer business segment code
	MerchantCategoryCode *int64 `json:"merchant_category_code,omitempty"`
	// The payment authentication request
	Authentication *CreatePaymentAuthenticationRequest `json:"authentication,omitempty"`
	// The Credit card payment contactless request
	Contactless *CreateCardPaymentContactlessRequest `json:"contactless,omitempty"`
	// Indicates whether a particular payment will enter the offline retry flow
	AutoRecovery *bool `json:"auto_recovery,omitempty"`
	// AuthOnly, AuthAndCapture, PreAuth
	OperationType *string `json:"operation_type,omitempty"`
	// Defines whether the card has been used one or more times.
	RecurrencyCycle *string                   `json:"recurrency_cycle,omitempty"`
	Payload         *CreateCardPayloadRequest `json:"payload,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCreditCardPaymentRequest.
// It customizes the JSON marshaling process for CreateCreditCardPaymentRequest objects.
func (c *CreateCreditCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCreditCardPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCreditCardPaymentRequest) toMap() map[string]any {
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
	if c.MerchantCategoryCode != nil {
		structMap["merchant_category_code"] = c.MerchantCategoryCode
	}
	if c.Authentication != nil {
		structMap["authentication"] = c.Authentication
	}
	if c.Contactless != nil {
		structMap["contactless"] = c.Contactless
	}
	if c.AutoRecovery != nil {
		structMap["auto_recovery"] = c.AutoRecovery
	}
	if c.OperationType != nil {
		structMap["operation_type"] = c.OperationType
	}
	if c.RecurrencyCycle != nil {
		structMap["recurrency_cycle"] = c.RecurrencyCycle
	}
	if c.Payload != nil {
		structMap["payload"] = c.Payload
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCreditCardPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCreditCardPaymentRequest objects.
func (c *CreateCreditCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Installments         *int                                 `json:"installments,omitempty"`
		StatementDescriptor  *string                              `json:"statement_descriptor,omitempty"`
		Card                 *CreateCardRequest                   `json:"card,omitempty"`
		CardId               *string                              `json:"card_id,omitempty"`
		CardToken            *string                              `json:"card_token,omitempty"`
		Recurrence           *bool                                `json:"recurrence,omitempty"`
		Capture              *bool                                `json:"capture,omitempty"`
		ExtendedLimitEnabled *bool                                `json:"extended_limit_enabled,omitempty"`
		ExtendedLimitCode    *string                              `json:"extended_limit_code,omitempty"`
		MerchantCategoryCode *int64                               `json:"merchant_category_code,omitempty"`
		Authentication       *CreatePaymentAuthenticationRequest  `json:"authentication,omitempty"`
		Contactless          *CreateCardPaymentContactlessRequest `json:"contactless,omitempty"`
		AutoRecovery         *bool                                `json:"auto_recovery,omitempty"`
		OperationType        *string                              `json:"operation_type,omitempty"`
		RecurrencyCycle      *string                              `json:"recurrency_cycle,omitempty"`
		Payload              *CreateCardPayloadRequest            `json:"payload,omitempty"`
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
	c.MerchantCategoryCode = temp.MerchantCategoryCode
	c.Authentication = temp.Authentication
	c.Contactless = temp.Contactless
	c.AutoRecovery = temp.AutoRecovery
	c.OperationType = temp.OperationType
	c.RecurrencyCycle = temp.RecurrencyCycle
	c.Payload = temp.Payload
	return nil
}
