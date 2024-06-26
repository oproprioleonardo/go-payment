/*
Package pagarmeapisdk

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
)

// ListInvoicesResponse represents a ListInvoicesResponse struct.
// Response object for listing invoices
type ListInvoicesResponse struct {
	// The Invoice objects
	Data Optional[[]GetInvoiceResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListInvoicesResponse.
// It customizes the JSON marshaling process for ListInvoicesResponse objects.
func (l *ListInvoicesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListInvoicesResponse object to a map representation for JSON marshaling.
func (l *ListInvoicesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListInvoicesResponse.
// It customizes the JSON unmarshaling process for ListInvoicesResponse objects.
func (l *ListInvoicesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetInvoiceResponse] `json:"data"`
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
