package models

import (
	"encoding/json"
)

// CreateTransfer represents a CreateTransfer struct.
type CreateTransfer struct {
	Amount   int      `json:"amount"`
	SourceId string   `json:"source_id"`
	TargetId string   `json:"target_id"`
	Metadata []string `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateTransfer.
// It customizes the JSON marshaling process for CreateTransfer objects.
func (c *CreateTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateTransfer object to a map representation for JSON marshaling.
func (c *CreateTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["source_id"] = c.SourceId
	structMap["target_id"] = c.TargetId
	if c.Metadata != nil {
		structMap["metadata"] = c.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateTransfer.
// It customizes the JSON unmarshaling process for CreateTransfer objects.
func (c *CreateTransfer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount   int      `json:"amount"`
		SourceId string   `json:"source_id"`
		TargetId string   `json:"target_id"`
		Metadata []string `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.SourceId = temp.SourceId
	c.TargetId = temp.TargetId
	c.Metadata = temp.Metadata
	return nil
}
