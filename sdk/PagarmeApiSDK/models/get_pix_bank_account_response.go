package models

import (
	"encoding/json"
)

// GetPixBankAccountResponse represents a GetPixBankAccountResponse struct.
// Payer's bank details.
type GetPixBankAccountResponse struct {
	BankName      Optional[string] `json:"bank_name"`
	Ispb          Optional[string] `json:"ispb"`
	BranchCode    Optional[string] `json:"branch_code"`
	AccountNumber Optional[string] `json:"account_number"`
}

// MarshalJSON implements the json.Marshaler interface for GetPixBankAccountResponse.
// It customizes the JSON marshaling process for GetPixBankAccountResponse objects.
func (g *GetPixBankAccountResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPixBankAccountResponse object to a map representation for JSON marshaling.
func (g *GetPixBankAccountResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.BankName.IsValueSet() {
		structMap["bank_name"] = g.BankName.Value()
	}
	if g.Ispb.IsValueSet() {
		structMap["ispb"] = g.Ispb.Value()
	}
	if g.BranchCode.IsValueSet() {
		structMap["branch_code"] = g.BranchCode.Value()
	}
	if g.AccountNumber.IsValueSet() {
		structMap["account_number"] = g.AccountNumber.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixBankAccountResponse.
// It customizes the JSON unmarshaling process for GetPixBankAccountResponse objects.
func (g *GetPixBankAccountResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankName      Optional[string] `json:"bank_name"`
		Ispb          Optional[string] `json:"ispb"`
		BranchCode    Optional[string] `json:"branch_code"`
		AccountNumber Optional[string] `json:"account_number"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.BankName = temp.BankName
	g.Ispb = temp.Ispb
	g.BranchCode = temp.BranchCode
	g.AccountNumber = temp.AccountNumber
	return nil
}
