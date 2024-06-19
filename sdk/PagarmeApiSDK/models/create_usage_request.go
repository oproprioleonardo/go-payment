package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateUsageRequest represents a CreateUsageRequest struct.
// Request for creating a usage
type CreateUsageRequest struct {
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	UsedAt      time.Time `json:"used_at"`
	// Identification code in the client system
	Code *string `json:"code,omitempty"`
	// identification group in the client system
	Group *string `json:"group,omitempty"`
	// Field used in item scheme type 'Percent'
	Amount *int `json:"amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateUsageRequest.
// It customizes the JSON marshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateUsageRequest object to a map representation for JSON marshaling.
func (c *CreateUsageRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["quantity"] = c.Quantity
	structMap["description"] = c.Description
	structMap["used_at"] = c.UsedAt.Format(time.RFC3339)
	if c.Code != nil {
		structMap["code"] = c.Code
	}
	if c.Group != nil {
		structMap["group"] = c.Group
	}
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsageRequest.
// It customizes the JSON unmarshaling process for CreateUsageRequest objects.
func (c *CreateUsageRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Quantity    int     `json:"quantity"`
		Description string  `json:"description"`
		UsedAt      string  `json:"used_at"`
		Code        *string `json:"code,omitempty"`
		Group       *string `json:"group,omitempty"`
		Amount      *int    `json:"amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Quantity = temp.Quantity
	c.Description = temp.Description
	UsedAtVal, err := time.Parse(time.RFC3339, temp.UsedAt)
	if err != nil {
		log.Fatalf("Cannot Parse used_at as % s format.", time.RFC3339)
	}
	c.UsedAt = UsedAtVal
	c.Code = temp.Code
	c.Group = temp.Group
	c.Amount = temp.Amount
	return nil
}
