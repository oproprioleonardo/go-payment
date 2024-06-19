package models

import (
	"encoding/json"
)

// CreateSubscriptionBoletoRequest represents a CreateSubscriptionBoletoRequest struct.
// Information about fines and interest on the "boleto" used from payment
type CreateSubscriptionBoletoRequest struct {
	Interest            *CreateInterestRequest `json:"interest,omitempty"`
	Fine                *CreateFineRequest     `json:"fine,omitempty"`
	MaxDaysToPayPastDue Optional[int]          `json:"max_days_to_pay_past_due"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionBoletoRequest.
// It customizes the JSON marshaling process for CreateSubscriptionBoletoRequest objects.
func (c *CreateSubscriptionBoletoRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionBoletoRequest object to a map representation for JSON marshaling.
func (c *CreateSubscriptionBoletoRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Interest != nil {
		structMap["interest"] = c.Interest
	}
	if c.Fine != nil {
		structMap["fine"] = c.Fine
	}
	if c.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = c.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionBoletoRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionBoletoRequest objects.
func (c *CreateSubscriptionBoletoRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Interest            *CreateInterestRequest `json:"interest,omitempty"`
		Fine                *CreateFineRequest     `json:"fine,omitempty"`
		MaxDaysToPayPastDue Optional[int]          `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Interest = temp.Interest
	c.Fine = temp.Fine
	c.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
