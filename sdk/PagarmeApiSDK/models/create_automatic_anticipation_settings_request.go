package models

import (
	"encoding/json"
)

// CreateAutomaticAnticipationSettingsRequest represents a CreateAutomaticAnticipationSettingsRequest struct.
type CreateAutomaticAnticipationSettingsRequest struct {
	Enabled          bool   `json:"enabled"`
	Type             string `json:"type"`
	VolumePercentage int    `json:"volume_percentage"`
	Delay            int    `json:"delay"`
	Days             []int  `json:"days"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAutomaticAnticipationSettingsRequest.
// It customizes the JSON marshaling process for CreateAutomaticAnticipationSettingsRequest objects.
func (c *CreateAutomaticAnticipationSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAutomaticAnticipationSettingsRequest object to a map representation for JSON marshaling.
func (c *CreateAutomaticAnticipationSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["enabled"] = c.Enabled
	structMap["type"] = c.Type
	structMap["volume_percentage"] = c.VolumePercentage
	structMap["delay"] = c.Delay
	structMap["days"] = c.Days
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAutomaticAnticipationSettingsRequest.
// It customizes the JSON unmarshaling process for CreateAutomaticAnticipationSettingsRequest objects.
func (c *CreateAutomaticAnticipationSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          bool   `json:"enabled"`
		Type             string `json:"type"`
		VolumePercentage int    `json:"volume_percentage"`
		Delay            int    `json:"delay"`
		Days             []int  `json:"days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Enabled = temp.Enabled
	c.Type = temp.Type
	c.VolumePercentage = temp.VolumePercentage
	c.Delay = temp.Delay
	c.Days = temp.Days
	return nil
}
