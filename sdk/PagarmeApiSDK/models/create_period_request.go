package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreatePeriodRequest represents a CreatePeriodRequest struct.
type CreatePeriodRequest struct {
	EndAt *time.Time `json:"end_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePeriodRequest.
// It customizes the JSON marshaling process for CreatePeriodRequest objects.
func (c *CreatePeriodRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePeriodRequest object to a map representation for JSON marshaling.
func (c *CreatePeriodRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.EndAt != nil {
		structMap["end_at"] = c.EndAt.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePeriodRequest.
// It customizes the JSON unmarshaling process for CreatePeriodRequest objects.
func (c *CreatePeriodRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EndAt *string `json:"end_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.EndAt != nil {
		EndAtVal, err := time.Parse(time.RFC3339, *temp.EndAt)
		if err != nil {
			log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
		}
		c.EndAt = &EndAtVal
	}
	return nil
}
