package models

import (
	"encoding/json"
)

// UpdateSubscriptionPaymentMethodRequest represents a UpdateSubscriptionPaymentMethodRequest struct.
// Request for updating a subscription's payment method
type UpdateSubscriptionPaymentMethodRequest struct {
	// The new payment method
	PaymentMethod string `json:"payment_method"`
	// Card id
	CardId string `json:"card_id"`
	// Card data
	Card CreateCardRequest `json:"card"`
	// The Card Token
	CardToken *string `json:"card_token,omitempty"`
	// Information about fines and interest on the "boleto" used from payment
	Boleto *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionPaymentMethodRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionPaymentMethodRequest objects.
func (u *UpdateSubscriptionPaymentMethodRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionPaymentMethodRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionPaymentMethodRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_method"] = u.PaymentMethod
	structMap["card_id"] = u.CardId
	structMap["card"] = u.Card
	if u.CardToken != nil {
		structMap["card_token"] = u.CardToken
	}
	if u.Boleto != nil {
		structMap["boleto"] = u.Boleto
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionPaymentMethodRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionPaymentMethodRequest objects.
func (u *UpdateSubscriptionPaymentMethodRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentMethod string                           `json:"payment_method"`
		CardId        string                           `json:"card_id"`
		Card          CreateCardRequest                `json:"card"`
		CardToken     *string                          `json:"card_token,omitempty"`
		Boleto        *CreateSubscriptionBoletoRequest `json:"boleto,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.PaymentMethod = temp.PaymentMethod
	u.CardId = temp.CardId
	u.Card = temp.Card
	u.CardToken = temp.CardToken
	u.Boleto = temp.Boleto
	return nil
}
