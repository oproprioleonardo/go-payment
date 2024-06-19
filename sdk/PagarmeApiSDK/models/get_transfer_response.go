package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetTransferResponse represents a GetTransferResponse struct.
// Transfer response
type GetTransferResponse struct {
	// Id
	Id Optional[string] `json:"id"`
	// Transfer amount
	Amount Optional[int] `json:"amount"`
	// Transfer status
	Status Optional[string] `json:"status"`
	// Transfer creation date
	CreatedAt Optional[time.Time] `json:"created_at"`
	// Transfer last update date
	UpdatedAt Optional[time.Time] `json:"updated_at"`
	// Bank account
	BankAccount Optional[GetBankAccountResponse] `json:"bank_account"`
	// Metadata
	Metadata Optional[map[string]string] `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransferResponse.
// It customizes the JSON marshaling process for GetTransferResponse objects.
func (g *GetTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetTransferResponse object to a map representation for JSON marshaling.
func (g *GetTransferResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransferResponse.
// It customizes the JSON unmarshaling process for GetTransferResponse objects.
func (g *GetTransferResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          Optional[string]                 `json:"id"`
		Amount      Optional[int]                    `json:"amount"`
		Status      Optional[string]                 `json:"status"`
		CreatedAt   Optional[string]                 `json:"created_at"`
		UpdatedAt   Optional[string]                 `json:"updated_at"`
		BankAccount Optional[GetBankAccountResponse] `json:"bank_account"`
		Metadata    Optional[map[string]string]      `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Amount = temp.Amount
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.BankAccount = temp.BankAccount
	g.Metadata = temp.Metadata
	return nil
}
