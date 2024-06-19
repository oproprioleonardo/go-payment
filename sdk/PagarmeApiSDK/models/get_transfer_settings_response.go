package models

import (
	"encoding/json"
)

// GetTransferSettingsResponse represents a GetTransferSettingsResponse struct.
type GetTransferSettingsResponse struct {
	TransferEnabled  Optional[bool]   `json:"transfer_enabled"`
	TransferInterval Optional[string] `json:"transfer_interval"`
	TransferDay      Optional[int]    `json:"transfer_day"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransferSettingsResponse.
// It customizes the JSON marshaling process for GetTransferSettingsResponse objects.
func (g *GetTransferSettingsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetTransferSettingsResponse object to a map representation for JSON marshaling.
func (g *GetTransferSettingsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.TransferEnabled.IsValueSet() {
		structMap["transfer_enabled"] = g.TransferEnabled.Value()
	}
	if g.TransferInterval.IsValueSet() {
		structMap["transfer_interval"] = g.TransferInterval.Value()
	}
	if g.TransferDay.IsValueSet() {
		structMap["transfer_day"] = g.TransferDay.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransferSettingsResponse.
// It customizes the JSON unmarshaling process for GetTransferSettingsResponse objects.
func (g *GetTransferSettingsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransferEnabled  Optional[bool]   `json:"transfer_enabled"`
		TransferInterval Optional[string] `json:"transfer_interval"`
		TransferDay      Optional[int]    `json:"transfer_day"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.TransferEnabled = temp.TransferEnabled
	g.TransferInterval = temp.TransferInterval
	g.TransferDay = temp.TransferDay
	return nil
}
