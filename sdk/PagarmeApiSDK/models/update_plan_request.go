package models

import (
	"encoding/json"
)

// UpdatePlanRequest represents a UpdatePlanRequest struct.
// Request for updating a plan
type UpdatePlanRequest struct {
	// Plan's name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Number os installments
	Installments []int `json:"installments"`
	// Text that will be shown on the credit card's statement
	StatementDescriptor string `json:"statement_descriptor"`
	// Currency
	Currency string `json:"currency"`
	// Interval
	Interval string `json:"interval"`
	// Interval count
	IntervalCount int `json:"interval_count"`
	// Payment methods accepted by the plan
	PaymentMethods []string `json:"payment_methods"`
	// Billing type
	BillingType string `json:"billing_type"`
	// Plan status
	Status string `json:"status"`
	// Indicates if the plan is shippable
	Shippable bool `json:"shippable"`
	// Billing days accepted by the plan
	BillingDays []int `json:"billing_days"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	// Minimum price
	MinimumPrice *int `json:"minimum_price,omitempty"`
	// Number of trial period in days, where the customer will not be charged
	TrialPeriodDays *int `json:"trial_period_days,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePlanRequest.
// It customizes the JSON marshaling process for UpdatePlanRequest objects.
func (u *UpdatePlanRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdatePlanRequest object to a map representation for JSON marshaling.
func (u *UpdatePlanRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = u.Name
	structMap["description"] = u.Description
	structMap["installments"] = u.Installments
	structMap["statement_descriptor"] = u.StatementDescriptor
	structMap["currency"] = u.Currency
	structMap["interval"] = u.Interval
	structMap["interval_count"] = u.IntervalCount
	structMap["payment_methods"] = u.PaymentMethods
	structMap["billing_type"] = u.BillingType
	structMap["status"] = u.Status
	structMap["shippable"] = u.Shippable
	structMap["billing_days"] = u.BillingDays
	structMap["metadata"] = u.Metadata
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	if u.TrialPeriodDays != nil {
		structMap["trial_period_days"] = u.TrialPeriodDays
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePlanRequest.
// It customizes the JSON unmarshaling process for UpdatePlanRequest objects.
func (u *UpdatePlanRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                string            `json:"name"`
		Description         string            `json:"description"`
		Installments        []int             `json:"installments"`
		StatementDescriptor string            `json:"statement_descriptor"`
		Currency            string            `json:"currency"`
		Interval            string            `json:"interval"`
		IntervalCount       int               `json:"interval_count"`
		PaymentMethods      []string          `json:"payment_methods"`
		BillingType         string            `json:"billing_type"`
		Status              string            `json:"status"`
		Shippable           bool              `json:"shippable"`
		BillingDays         []int             `json:"billing_days"`
		Metadata            map[string]string `json:"metadata"`
		MinimumPrice        *int              `json:"minimum_price,omitempty"`
		TrialPeriodDays     *int              `json:"trial_period_days,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Description = temp.Description
	u.Installments = temp.Installments
	u.StatementDescriptor = temp.StatementDescriptor
	u.Currency = temp.Currency
	u.Interval = temp.Interval
	u.IntervalCount = temp.IntervalCount
	u.PaymentMethods = temp.PaymentMethods
	u.BillingType = temp.BillingType
	u.Status = temp.Status
	u.Shippable = temp.Shippable
	u.BillingDays = temp.BillingDays
	u.Metadata = temp.Metadata
	u.MinimumPrice = temp.MinimumPrice
	u.TrialPeriodDays = temp.TrialPeriodDays
	return nil
}
