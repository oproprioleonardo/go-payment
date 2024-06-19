package models

import (
	"encoding/json"
)

// ListTransferResponse represents a ListTransferResponse struct.
// List of paginated transfer objects
type ListTransferResponse struct {
	// Transfers
	Data Optional[[]GetTransferResponse] `json:"data"`
	// Paging
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListTransferResponse.
// It customizes the JSON marshaling process for ListTransferResponse objects.
func (l *ListTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListTransferResponse object to a map representation for JSON marshaling.
func (l *ListTransferResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListTransferResponse.
// It customizes the JSON unmarshaling process for ListTransferResponse objects.
func (l *ListTransferResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetTransferResponse] `json:"data"`
		Paging Optional[PagingResponse]        `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
