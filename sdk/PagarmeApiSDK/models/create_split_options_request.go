package models

import (
	"encoding/json"
)

// CreateSplitOptionsRequest represents a CreateSplitOptionsRequest struct.
// The Split Options Request
type CreateSplitOptionsRequest struct {
	// Liable options
	Liable *bool `json:"liable,omitempty"`
	// Charge processing fee
	ChargeProcessingFee *bool `json:"charge_processing_fee,omitempty"`
	ChargeRemainderFee  *bool `json:"charge_remainder_fee,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSplitOptionsRequest.
// It customizes the JSON marshaling process for CreateSplitOptionsRequest objects.
func (c *CreateSplitOptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSplitOptionsRequest object to a map representation for JSON marshaling.
func (c *CreateSplitOptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Liable != nil {
		structMap["liable"] = c.Liable
	}
	if c.ChargeProcessingFee != nil {
		structMap["charge_processing_fee"] = c.ChargeProcessingFee
	}
	if c.ChargeRemainderFee != nil {
		structMap["charge_remainder_fee"] = c.ChargeRemainderFee
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSplitOptionsRequest.
// It customizes the JSON unmarshaling process for CreateSplitOptionsRequest objects.
func (c *CreateSplitOptionsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Liable              *bool `json:"liable,omitempty"`
		ChargeProcessingFee *bool `json:"charge_processing_fee,omitempty"`
		ChargeRemainderFee  *bool `json:"charge_remainder_fee,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Liable = temp.Liable
	c.ChargeProcessingFee = temp.ChargeProcessingFee
	c.ChargeRemainderFee = temp.ChargeRemainderFee
	return nil
}
