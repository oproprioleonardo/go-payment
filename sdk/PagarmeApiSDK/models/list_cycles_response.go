package models

import (
	"encoding/json"
)

// ListCyclesResponse represents a ListCyclesResponse struct.
// Response object for listing subscription cycles
type ListCyclesResponse struct {
	// The subscription cycles objects
	Data Optional[[]GetPeriodResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListCyclesResponse.
// It customizes the JSON marshaling process for ListCyclesResponse objects.
func (l *ListCyclesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListCyclesResponse object to a map representation for JSON marshaling.
func (l *ListCyclesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCyclesResponse.
// It customizes the JSON unmarshaling process for ListCyclesResponse objects.
func (l *ListCyclesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetPeriodResponse] `json:"data"`
		Paging Optional[PagingResponse]      `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
