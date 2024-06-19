package models

import (
	"encoding/json"
)

// CreateCancelChargeRequest represents a CreateCancelChargeRequest struct.
// Request for canceling a charge.
type CreateCancelChargeRequest struct {
	// The amount that will be canceled.
	Amount *int `json:"amount,omitempty"`
	// The split rules request
	SplitRules []CreateCancelChargeSplitRulesRequest `json:"split_rules,omitempty"`
	// Splits
	Split              []CreateSplitRequest           `json:"split,omitempty"`
	OperationReference string                         `json:"operation_reference"`
	BankAccount        *CreateBankAccountRefundingDTO `json:"bank_account,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCancelChargeRequest.
// It customizes the JSON marshaling process for CreateCancelChargeRequest objects.
func (c *CreateCancelChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCancelChargeRequest object to a map representation for JSON marshaling.
func (c *CreateCancelChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.SplitRules != nil {
		structMap["split_rules"] = c.SplitRules
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	structMap["operation_reference"] = c.OperationReference
	if c.BankAccount != nil {
		structMap["bank_account"] = c.BankAccount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCancelChargeRequest.
// It customizes the JSON unmarshaling process for CreateCancelChargeRequest objects.
func (c *CreateCancelChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount             *int                                  `json:"amount,omitempty"`
		SplitRules         []CreateCancelChargeSplitRulesRequest `json:"split_rules,omitempty"`
		Split              []CreateSplitRequest                  `json:"split,omitempty"`
		OperationReference string                                `json:"operation_reference"`
		BankAccount        *CreateBankAccountRefundingDTO        `json:"bank_account,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.SplitRules = temp.SplitRules
	c.Split = temp.Split
	c.OperationReference = temp.OperationReference
	c.BankAccount = temp.BankAccount
	return nil
}
