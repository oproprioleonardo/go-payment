package models

import (
	"encoding/json"
)

// ListOrderResponse represents a ListOrderResponse struct.
// Response object for listing order objects
type ListOrderResponse struct {
	// The order object
	Data Optional[[]GetOrderResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListOrderResponse.
// It customizes the JSON marshaling process for ListOrderResponse objects.
func (l *ListOrderResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListOrderResponse object to a map representation for JSON marshaling.
func (l *ListOrderResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListOrderResponse.
// It customizes the JSON unmarshaling process for ListOrderResponse objects.
func (l *ListOrderResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetOrderResponse] `json:"data"`
		Paging Optional[PagingResponse]     `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
