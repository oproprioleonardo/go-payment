package models

import (
	"encoding/json"
)

// CreateCancelSubscriptionRequest represents a CreateCancelSubscriptionRequest struct.
// Request for canceling a subscription
type CreateCancelSubscriptionRequest struct {
	// Indicates if the pending invoices must also be canceled.
	CancelPendingInvoices bool `json:"cancel_pending_invoices"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCancelSubscriptionRequest.
// It customizes the JSON marshaling process for CreateCancelSubscriptionRequest objects.
func (c *CreateCancelSubscriptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCancelSubscriptionRequest object to a map representation for JSON marshaling.
func (c *CreateCancelSubscriptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["cancel_pending_invoices"] = c.CancelPendingInvoices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCancelSubscriptionRequest.
// It customizes the JSON unmarshaling process for CreateCancelSubscriptionRequest objects.
func (c *CreateCancelSubscriptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CancelPendingInvoices bool `json:"cancel_pending_invoices"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CancelPendingInvoices = temp.CancelPendingInvoices
	return nil
}
