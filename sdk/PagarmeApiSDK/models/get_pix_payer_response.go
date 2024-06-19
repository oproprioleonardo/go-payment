package models

import (
	"encoding/json"
)

// GetPixPayerResponse represents a GetPixPayerResponse struct.
// Pix payer data.
type GetPixPayerResponse struct {
	Name         Optional[string]                    `json:"name"`
	Document     Optional[string]                    `json:"document"`
	DocumentType Optional[string]                    `json:"document_type"`
	BankAccount  Optional[GetPixBankAccountResponse] `json:"bank_account"`
}

// MarshalJSON implements the json.Marshaler interface for GetPixPayerResponse.
// It customizes the JSON marshaling process for GetPixPayerResponse objects.
func (g *GetPixPayerResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPixPayerResponse object to a map representation for JSON marshaling.
func (g *GetPixPayerResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Document.IsValueSet() {
		structMap["document"] = g.Document.Value()
	}
	if g.DocumentType.IsValueSet() {
		structMap["document_type"] = g.DocumentType.Value()
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPixPayerResponse.
// It customizes the JSON unmarshaling process for GetPixPayerResponse objects.
func (g *GetPixPayerResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name         Optional[string]                    `json:"name"`
		Document     Optional[string]                    `json:"document"`
		DocumentType Optional[string]                    `json:"document_type"`
		BankAccount  Optional[GetPixBankAccountResponse] `json:"bank_account"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Name = temp.Name
	g.Document = temp.Document
	g.DocumentType = temp.DocumentType
	g.BankAccount = temp.BankAccount
	return nil
}
