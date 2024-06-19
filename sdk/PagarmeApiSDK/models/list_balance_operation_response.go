package models

import (
	"encoding/json"
)

// ListBalanceOperationResponse represents a ListBalanceOperationResponse struct.
// Response object for listing BalanceOperation objects
type ListBalanceOperationResponse struct {
	// The BalanceOperation object
	Data   Optional[[]GetBalanceOperationResponse] `json:"data"`
	Paging Optional[PagingResponse]                `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListBalanceOperationResponse.
// It customizes the JSON marshaling process for ListBalanceOperationResponse objects.
func (l *ListBalanceOperationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListBalanceOperationResponse object to a map representation for JSON marshaling.
func (l *ListBalanceOperationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListBalanceOperationResponse.
// It customizes the JSON unmarshaling process for ListBalanceOperationResponse objects.
func (l *ListBalanceOperationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetBalanceOperationResponse] `json:"data"`
		Paging Optional[PagingResponse]                `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
