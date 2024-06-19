package models

import (
	"encoding/json"
)

// CreateCaptureChargeRequest represents a CreateCaptureChargeRequest struct.
// Request for capturing a charge
type CreateCaptureChargeRequest struct {
	// Code for the charge. Sending this field will update the code send on the charge and order creation.
	Code string `json:"code"`
	// The amount that will be captured
	Amount *int `json:"amount,omitempty"`
	// Splits
	Split              []CreateSplitRequest `json:"split,omitempty"`
	OperationReference string               `json:"operation_reference"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCaptureChargeRequest.
// It customizes the JSON marshaling process for CreateCaptureChargeRequest objects.
func (c *CreateCaptureChargeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCaptureChargeRequest object to a map representation for JSON marshaling.
func (c *CreateCaptureChargeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["code"] = c.Code
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.Split != nil {
		structMap["split"] = c.Split
	}
	structMap["operation_reference"] = c.OperationReference
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCaptureChargeRequest.
// It customizes the JSON unmarshaling process for CreateCaptureChargeRequest objects.
func (c *CreateCaptureChargeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code               string               `json:"code"`
		Amount             *int                 `json:"amount,omitempty"`
		Split              []CreateSplitRequest `json:"split,omitempty"`
		OperationReference string               `json:"operation_reference"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Code = temp.Code
	c.Amount = temp.Amount
	c.Split = temp.Split
	c.OperationReference = temp.OperationReference
	return nil
}
