package models

import (
	"encoding/json"
)

// UpdateRecipientBankAccountRequest represents a UpdateRecipientBankAccountRequest struct.
// Updates the default bank account for a recipient
type UpdateRecipientBankAccountRequest struct {
	// Bank account
	BankAccount CreateBankAccountRequest `json:"bank_account"`
	// Payment mode
	PaymentMode string `json:"payment_mode"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateRecipientBankAccountRequest.
// It customizes the JSON marshaling process for UpdateRecipientBankAccountRequest objects.
func (u *UpdateRecipientBankAccountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateRecipientBankAccountRequest object to a map representation for JSON marshaling.
func (u *UpdateRecipientBankAccountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank_account"] = u.BankAccount
	structMap["payment_mode"] = u.PaymentMode
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateRecipientBankAccountRequest.
// It customizes the JSON unmarshaling process for UpdateRecipientBankAccountRequest objects.
func (u *UpdateRecipientBankAccountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankAccount CreateBankAccountRequest `json:"bank_account"`
		PaymentMode string                   `json:"payment_mode"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.BankAccount = temp.BankAccount
	u.PaymentMode = temp.PaymentMode
	return nil
}
