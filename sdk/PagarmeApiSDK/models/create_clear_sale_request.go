package models

import (
	"encoding/json"
)

// CreateClearSaleRequest represents a CreateClearSaleRequest struct.
type CreateClearSaleRequest struct {
	CustomSla int `json:"custom_sla"`
}

// MarshalJSON implements the json.Marshaler interface for CreateClearSaleRequest.
// It customizes the JSON marshaling process for CreateClearSaleRequest objects.
func (c *CreateClearSaleRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateClearSaleRequest object to a map representation for JSON marshaling.
func (c *CreateClearSaleRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["custom_sla"] = c.CustomSla
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateClearSaleRequest.
// It customizes the JSON unmarshaling process for CreateClearSaleRequest objects.
func (c *CreateClearSaleRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CustomSla int `json:"custom_sla"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CustomSla = temp.CustomSla
	return nil
}
