package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateTransactionReportFileRequest represents a CreateTransactionReportFileRequest struct.
type CreateTransactionReportFileRequest struct {
	Name    string     `json:"name"`
	StartAt *time.Time `json:"start_at,omitempty"`
	EndAt   *string    `json:"end_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateTransactionReportFileRequest.
// It customizes the JSON marshaling process for CreateTransactionReportFileRequest objects.
func (c *CreateTransactionReportFileRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateTransactionReportFileRequest object to a map representation for JSON marshaling.
func (c *CreateTransactionReportFileRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	if c.StartAt != nil {
		structMap["start_at"] = c.StartAt.Format(time.RFC3339)
	}
	if c.EndAt != nil {
		structMap["end_at"] = c.EndAt
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateTransactionReportFileRequest.
// It customizes the JSON unmarshaling process for CreateTransactionReportFileRequest objects.
func (c *CreateTransactionReportFileRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name    string  `json:"name"`
		StartAt *string `json:"start_at,omitempty"`
		EndAt   *string `json:"end_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	if temp.StartAt != nil {
		StartAtVal, err := time.Parse(time.RFC3339, *temp.StartAt)
		if err != nil {
			log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
		}
		c.StartAt = &StartAtVal
	}
	c.EndAt = temp.EndAt
	return nil
}
