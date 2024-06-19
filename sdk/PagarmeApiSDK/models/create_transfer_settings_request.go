package models

import (
	"encoding/json"
)

// CreateTransferSettingsRequest represents a CreateTransferSettingsRequest struct.
// Informações de transferência do recebedor
type CreateTransferSettingsRequest struct {
	TransferEnabled  bool   `json:"transfer_enabled"`
	TransferInterval string `json:"transfer_interval"`
	TransferDay      int    `json:"transfer_day"`
}

// MarshalJSON implements the json.Marshaler interface for CreateTransferSettingsRequest.
// It customizes the JSON marshaling process for CreateTransferSettingsRequest objects.
func (c *CreateTransferSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateTransferSettingsRequest object to a map representation for JSON marshaling.
func (c *CreateTransferSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["transfer_enabled"] = c.TransferEnabled
	structMap["transfer_interval"] = c.TransferInterval
	structMap["transfer_day"] = c.TransferDay
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateTransferSettingsRequest.
// It customizes the JSON unmarshaling process for CreateTransferSettingsRequest objects.
func (c *CreateTransferSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransferEnabled  bool   `json:"transfer_enabled"`
		TransferInterval string `json:"transfer_interval"`
		TransferDay      int    `json:"transfer_day"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.TransferEnabled = temp.TransferEnabled
	c.TransferInterval = temp.TransferInterval
	c.TransferDay = temp.TransferDay
	return nil
}
