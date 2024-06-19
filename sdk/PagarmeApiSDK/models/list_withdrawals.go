package models

import (
	"encoding/json"
)

// ListWithdrawals represents a ListWithdrawals struct.
type ListWithdrawals struct {
	// The Increments response
	Data []GetWithdrawResponse `json:"data"`
	// Paging object
	Paging PagingResponse `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListWithdrawals.
// It customizes the JSON marshaling process for ListWithdrawals objects.
func (l *ListWithdrawals) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListWithdrawals object to a map representation for JSON marshaling.
func (l *ListWithdrawals) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["data"] = l.Data
	structMap["paging"] = l.Paging
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListWithdrawals.
// It customizes the JSON unmarshaling process for ListWithdrawals objects.
func (l *ListWithdrawals) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   []GetWithdrawResponse `json:"data"`
		Paging PagingResponse        `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
