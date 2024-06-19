package models

import (
	"encoding/json"
)

// GetAutomaticAnticipationResponse represents a GetAutomaticAnticipationResponse struct.
type GetAutomaticAnticipationResponse struct {
	Enabled          Optional[bool]   `json:"enabled"`
	Type             Optional[string] `json:"type"`
	VolumePercentage Optional[int]    `json:"volume_percentage"`
	Delay            Optional[int]    `json:"delay"`
	Days             Optional[[]int]  `json:"days"`
}

// MarshalJSON implements the json.Marshaler interface for GetAutomaticAnticipationResponse.
// It customizes the JSON marshaling process for GetAutomaticAnticipationResponse objects.
func (g *GetAutomaticAnticipationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetAutomaticAnticipationResponse object to a map representation for JSON marshaling.
func (g *GetAutomaticAnticipationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Enabled.IsValueSet() {
		structMap["enabled"] = g.Enabled.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.VolumePercentage.IsValueSet() {
		structMap["volume_percentage"] = g.VolumePercentage.Value()
	}
	if g.Delay.IsValueSet() {
		structMap["delay"] = g.Delay.Value()
	}
	if g.Days.IsValueSet() {
		structMap["days"] = g.Days.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAutomaticAnticipationResponse.
// It customizes the JSON unmarshaling process for GetAutomaticAnticipationResponse objects.
func (g *GetAutomaticAnticipationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Enabled          Optional[bool]   `json:"enabled"`
		Type             Optional[string] `json:"type"`
		VolumePercentage Optional[int]    `json:"volume_percentage"`
		Delay            Optional[int]    `json:"delay"`
		Days             Optional[[]int]  `json:"days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Enabled = temp.Enabled
	g.Type = temp.Type
	g.VolumePercentage = temp.VolumePercentage
	g.Delay = temp.Delay
	g.Days = temp.Days
	return nil
}
