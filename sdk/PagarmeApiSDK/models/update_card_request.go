package models

import (
	"encoding/json"
)

// UpdateCardRequest represents a UpdateCardRequest struct.
// Request for updating a card
type UpdateCardRequest struct {
	// Holder name
	HolderName string `json:"holder_name"`
	// Expiration month
	ExpMonth int `json:"exp_month"`
	// Expiration year
	ExpYear int `json:"exp_year"`
	// Id of the address to be used as billing address
	BillingAddressId Optional[string] `json:"billing_address_id"`
	// Billing address
	BillingAddress CreateAddressRequest `json:"billing_address"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	Label    string            `json:"label"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCardRequest.
// It customizes the JSON marshaling process for UpdateCardRequest objects.
func (u *UpdateCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCardRequest object to a map representation for JSON marshaling.
func (u *UpdateCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["holder_name"] = u.HolderName
	structMap["exp_month"] = u.ExpMonth
	structMap["exp_year"] = u.ExpYear
	if u.BillingAddressId.IsValueSet() {
		structMap["billing_address_id"] = u.BillingAddressId.Value()
	}
	structMap["billing_address"] = u.BillingAddress
	structMap["metadata"] = u.Metadata
	structMap["label"] = u.Label
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCardRequest.
// It customizes the JSON unmarshaling process for UpdateCardRequest objects.
func (u *UpdateCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HolderName       string               `json:"holder_name"`
		ExpMonth         int                  `json:"exp_month"`
		ExpYear          int                  `json:"exp_year"`
		BillingAddressId Optional[string]     `json:"billing_address_id"`
		BillingAddress   CreateAddressRequest `json:"billing_address"`
		Metadata         map[string]string    `json:"metadata"`
		Label            string               `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.HolderName = temp.HolderName
	u.ExpMonth = temp.ExpMonth
	u.ExpYear = temp.ExpYear
	u.BillingAddressId = temp.BillingAddressId
	u.BillingAddress = temp.BillingAddress
	u.Metadata = temp.Metadata
	u.Label = temp.Label
	return nil
}
