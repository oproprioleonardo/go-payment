package models

import (
	"encoding/json"
)

// CreatePlanRequest represents a CreatePlanRequest struct.
// Request for creating a plan
type CreatePlanRequest struct {
	// Plan's name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Text that will be printed on the credit card's statement
	StatementDescriptor string `json:"statement_descriptor"`
	// Plan items
	Items []CreatePlanItemRequest `json:"items"`
	// Indicates if the plan is shippable
	Shippable bool `json:"shippable"`
	// Allowed payment methods for the plan
	PaymentMethods []string `json:"payment_methods"`
	// Number of installments
	Installments []int `json:"installments"`
	// Currency
	Currency string `json:"currency"`
	// Interval
	Interval string `json:"interval"`
	// Interval counts between two charges. For instance, if the interval is 'month' and count is 2, the customer will be charged once every two months.
	IntervalCount int `json:"interval_count"`
	// Allowed billings days for the subscription, in case the plan type is 'exact_day'
	BillingDays []int `json:"billing_days"`
	// Billing type
	BillingType string `json:"billing_type"`
	// Plan's pricing scheme
	PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	// Minimum price that will be charged
	MinimumPrice *int `json:"minimum_price,omitempty"`
	// Number of cycles
	Cycles *int `json:"cycles,omitempty"`
	// Quantity
	Quantity *int `json:"quantity,omitempty"`
	// Trial period, where the customer will not be charged.
	TrialPeriodDays *int `json:"trial_period_days,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePlanRequest.
// It customizes the JSON marshaling process for CreatePlanRequest objects.
func (c *CreatePlanRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePlanRequest object to a map representation for JSON marshaling.
func (c *CreatePlanRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["description"] = c.Description
	structMap["statement_descriptor"] = c.StatementDescriptor
	structMap["items"] = c.Items
	structMap["shippable"] = c.Shippable
	structMap["payment_methods"] = c.PaymentMethods
	structMap["installments"] = c.Installments
	structMap["currency"] = c.Currency
	structMap["interval"] = c.Interval
	structMap["interval_count"] = c.IntervalCount
	structMap["billing_days"] = c.BillingDays
	structMap["billing_type"] = c.BillingType
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["metadata"] = c.Metadata
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.TrialPeriodDays != nil {
		structMap["trial_period_days"] = c.TrialPeriodDays
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePlanRequest.
// It customizes the JSON unmarshaling process for CreatePlanRequest objects.
func (c *CreatePlanRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                string                     `json:"name"`
		Description         string                     `json:"description"`
		StatementDescriptor string                     `json:"statement_descriptor"`
		Items               []CreatePlanItemRequest    `json:"items"`
		Shippable           bool                       `json:"shippable"`
		PaymentMethods      []string                   `json:"payment_methods"`
		Installments        []int                      `json:"installments"`
		Currency            string                     `json:"currency"`
		Interval            string                     `json:"interval"`
		IntervalCount       int                        `json:"interval_count"`
		BillingDays         []int                      `json:"billing_days"`
		BillingType         string                     `json:"billing_type"`
		PricingScheme       CreatePricingSchemeRequest `json:"pricing_scheme"`
		Metadata            map[string]string          `json:"metadata"`
		MinimumPrice        *int                       `json:"minimum_price,omitempty"`
		Cycles              *int                       `json:"cycles,omitempty"`
		Quantity            *int                       `json:"quantity,omitempty"`
		TrialPeriodDays     *int                       `json:"trial_period_days,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Description = temp.Description
	c.StatementDescriptor = temp.StatementDescriptor
	c.Items = temp.Items
	c.Shippable = temp.Shippable
	c.PaymentMethods = temp.PaymentMethods
	c.Installments = temp.Installments
	c.Currency = temp.Currency
	c.Interval = temp.Interval
	c.IntervalCount = temp.IntervalCount
	c.BillingDays = temp.BillingDays
	c.BillingType = temp.BillingType
	c.PricingScheme = temp.PricingScheme
	c.Metadata = temp.Metadata
	c.MinimumPrice = temp.MinimumPrice
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	c.TrialPeriodDays = temp.TrialPeriodDays
	return nil
}
