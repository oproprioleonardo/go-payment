package models

import (
	"encoding/json"
)

// ListTransfers represents a ListTransfers struct.
type ListTransfers struct {
	// The Increments response
	Data []GetTransfer `json:"data"`
	// Paging object
	Paging PagingResponse `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListTransfers.
// It customizes the JSON marshaling process for ListTransfers objects.
func (l *ListTransfers) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListTransfers object to a map representation for JSON marshaling.
func (l *ListTransfers) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = l.Data
	structMap["paging"] = l.Paging
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListTransfers.
// It customizes the JSON unmarshaling process for ListTransfers objects.
func (l *ListTransfers) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   []GetTransfer  `json:"data"`
		Paging PagingResponse `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
