package models

import (
	"encoding/json"
)

// ListChargeTransactionsResponse represents a ListChargeTransactionsResponse struct.
// Response object for listing charge transactions
type ListChargeTransactionsResponse struct {
	// The charge transactions objects
	Data Optional[[]GetTransactionResponseInterface] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListChargeTransactionsResponse.
// It customizes the JSON marshaling process for ListChargeTransactionsResponse objects.
func (l *ListChargeTransactionsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListChargeTransactionsResponse object to a map representation for JSON marshaling.
func (l *ListChargeTransactionsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListChargeTransactionsResponse.
// It customizes the JSON unmarshaling process for ListChargeTransactionsResponse objects.
func (l *ListChargeTransactionsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetTransactionResponseInterface] `json:"data"`
		Paging Optional[PagingResponse]                    `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
