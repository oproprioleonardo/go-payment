package models

import (
	"encoding/json"
)

// CreateAntifraudRequest represents a CreateAntifraudRequest struct.
type CreateAntifraudRequest struct {
	Type      string                 `json:"type"`
	Clearsale CreateClearSaleRequest `json:"clearsale"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAntifraudRequest.
// It customizes the JSON marshaling process for CreateAntifraudRequest objects.
func (c *CreateAntifraudRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAntifraudRequest object to a map representation for JSON marshaling.
func (c *CreateAntifraudRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = c.Type
	structMap["clearsale"] = c.Clearsale
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAntifraudRequest.
// It customizes the JSON unmarshaling process for CreateAntifraudRequest objects.
func (c *CreateAntifraudRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      string                 `json:"type"`
		Clearsale CreateClearSaleRequest `json:"clearsale"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.Clearsale = temp.Clearsale
	return nil
}
