package models

import (
	"encoding/json"
)

// ListTransactionsFilesResponse represents a ListTransactionsFilesResponse struct.
// Response object for listing of transactions files
type ListTransactionsFilesResponse struct {
	Data Optional[[]GetTransactionReportFileResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListTransactionsFilesResponse.
// It customizes the JSON marshaling process for ListTransactionsFilesResponse objects.
func (l *ListTransactionsFilesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListTransactionsFilesResponse object to a map representation for JSON marshaling.
func (l *ListTransactionsFilesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListTransactionsFilesResponse.
// It customizes the JSON unmarshaling process for ListTransactionsFilesResponse objects.
func (l *ListTransactionsFilesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetTransactionReportFileResponse] `json:"data"`
		Paging Optional[PagingResponse]                     `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
