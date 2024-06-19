package models

import (
	"encoding/json"
)

// ListChargesResponse represents a ListChargesResponse struct.
// Response object for listing charges
type ListChargesResponse struct {
	// The charge objects
	Data Optional[[]GetChargeResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListChargesResponse.
// It customizes the JSON marshaling process for ListChargesResponse objects.
func (l *ListChargesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListChargesResponse object to a map representation for JSON marshaling.
func (l *ListChargesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListChargesResponse.
// It customizes the JSON unmarshaling process for ListChargesResponse objects.
func (l *ListChargesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetChargeResponse] `json:"data"`
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
