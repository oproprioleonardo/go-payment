package models

import (
	"encoding/json"
)

// UpdateAutomaticAnticipationSettingsRequest represents a UpdateAutomaticAnticipationSettingsRequest struct.
type UpdateAutomaticAnticipationSettingsRequest struct {
	Enabled          *bool   `json:"enabled,omitempty"`
	Type             *string `json:"type,omitempty"`
	VolumePercentage *int    `json:"volume_percentage,omitempty"`
	Delay            *int    `json:"delay,omitempty"`
	Days             *int    `json:"days,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateAutomaticAnticipationSettingsRequest.
// It customizes the JSON marshaling process for UpdateAutomaticAnticipationSettingsRequest objects.
func (u *UpdateAutomaticAnticipationSettingsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateAutomaticAnticipationSettingsRequest object to a map representation for JSON marshaling.
func (u *UpdateAutomaticAnticipationSettingsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Enabled != nil {
		structMap["enabled"] = u.Enabled
	}
	if u.Type != nil {
		structMap["type"] = u.Type
	}
	if u.VolumePercentage != nil {
		structMap["volume_percentage"] = u.VolumePercentage
	}
	if u.Delay != nil {
		structMap["delay"] = u.Delay
	}
	if u.Days != nil {
		structMap["days"] = u.Days
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateAutomaticAnticipationSettingsRequest.
// It customizes the JSON unmarshaling process for UpdateAutomaticAnticipationSettingsRequest objects.
func (u *UpdateAutomaticAnticipationSettingsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          *bool   `json:"enabled,omitempty"`
		Type             *string `json:"type,omitempty"`
		VolumePercentage *int    `json:"volume_percentage,omitempty"`
		Delay            *int    `json:"delay,omitempty"`
		Days             *int    `json:"days,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Enabled = temp.Enabled
	u.Type = temp.Type
	u.VolumePercentage = temp.VolumePercentage
	u.Delay = temp.Delay
	u.Days = temp.Days
	return nil
}
