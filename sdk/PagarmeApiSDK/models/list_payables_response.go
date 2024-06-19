package models

import (
	"encoding/json"
)

// ListPayablesResponse represents a ListPayablesResponse struct.
// Response object for listing payable objects
type ListPayablesResponse struct {
	// The payable object
	Data Optional[[]GetPayableResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListPayablesResponse.
// It customizes the JSON marshaling process for ListPayablesResponse objects.
func (l *ListPayablesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListPayablesResponse object to a map representation for JSON marshaling.
func (l *ListPayablesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPayablesResponse.
// It customizes the JSON unmarshaling process for ListPayablesResponse objects.
func (l *ListPayablesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetPayableResponse] `json:"data"`
		Paging Optional[PagingResponse]       `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
