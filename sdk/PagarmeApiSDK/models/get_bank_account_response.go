package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetBankAccountResponse represents a GetBankAccountResponse struct.
type GetBankAccountResponse struct {
	// Id
	Id Optional[string] `json:"id"`
	// Holder name
	HolderName Optional[string] `json:"holder_name"`
	// Holder type
	HolderType Optional[string] `json:"holder_type"`
	// Bank
	Bank Optional[string] `json:"bank"`
	// Branch number
	BranchNumber Optional[string] `json:"branch_number"`
	// Branch check digit
	BranchCheckDigit Optional[string] `json:"branch_check_digit"`
	// Account number
	AccountNumber Optional[string] `json:"account_number"`
	// Account check digit
	AccountCheckDigit Optional[string] `json:"account_check_digit"`
	// Bank account type
	Type Optional[string] `json:"type"`
	// Bank account status
	Status Optional[string] `json:"status"`
	// Creation date
	CreatedAt Optional[time.Time] `json:"created_at"`
	// Last update date
	UpdatedAt Optional[time.Time] `json:"updated_at"`
	// Deletion date
	DeletedAt Optional[time.Time] `json:"deleted_at"`
	// Recipient
	Recipient Optional[GetRecipientResponse] `json:"recipient"`
	// Metadata
	Metadata Optional[map[string]string] `json:"metadata"`
	// Pix Key
	PixKey Optional[string] `json:"pix_key"`
}

// MarshalJSON implements the json.Marshaler interface for GetBankAccountResponse.
// It customizes the JSON marshaling process for GetBankAccountResponse objects.
func (g *GetBankAccountResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetBankAccountResponse object to a map representation for JSON marshaling.
func (g *GetBankAccountResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.HolderName.IsValueSet() {
		structMap["holder_name"] = g.HolderName.Value()
	}
	if g.HolderType.IsValueSet() {
		structMap["holder_type"] = g.HolderType.Value()
	}
	if g.Bank.IsValueSet() {
		structMap["bank"] = g.Bank.Value()
	}
	if g.BranchNumber.IsValueSet() {
		structMap["branch_number"] = g.BranchNumber.Value()
	}
	if g.BranchCheckDigit.IsValueSet() {
		structMap["branch_check_digit"] = g.BranchCheckDigit.Value()
	}
	if g.AccountNumber.IsValueSet() {
		structMap["account_number"] = g.AccountNumber.Value()
	}
	if g.AccountCheckDigit.IsValueSet() {
		structMap["account_check_digit"] = g.AccountCheckDigit.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
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
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.PixKey.IsValueSet() {
		structMap["pix_key"] = g.PixKey.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBankAccountResponse.
// It customizes the JSON unmarshaling process for GetBankAccountResponse objects.
func (g *GetBankAccountResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                Optional[string]               `json:"id"`
		HolderName        Optional[string]               `json:"holder_name"`
		HolderType        Optional[string]               `json:"holder_type"`
		Bank              Optional[string]               `json:"bank"`
		BranchNumber      Optional[string]               `json:"branch_number"`
		BranchCheckDigit  Optional[string]               `json:"branch_check_digit"`
		AccountNumber     Optional[string]               `json:"account_number"`
		AccountCheckDigit Optional[string]               `json:"account_check_digit"`
		Type              Optional[string]               `json:"type"`
		Status            Optional[string]               `json:"status"`
		CreatedAt         Optional[string]               `json:"created_at"`
		UpdatedAt         Optional[string]               `json:"updated_at"`
		DeletedAt         Optional[string]               `json:"deleted_at"`
		Recipient         Optional[GetRecipientResponse] `json:"recipient"`
		Metadata          Optional[map[string]string]    `json:"metadata"`
		PixKey            Optional[string]               `json:"pix_key"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.HolderName = temp.HolderName
	g.HolderType = temp.HolderType
	g.Bank = temp.Bank
	g.BranchNumber = temp.BranchNumber
	g.BranchCheckDigit = temp.BranchCheckDigit
	g.AccountNumber = temp.AccountNumber
	g.AccountCheckDigit = temp.AccountCheckDigit
	g.Type = temp.Type
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
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.Recipient = temp.Recipient
	g.Metadata = temp.Metadata
	g.PixKey = temp.PixKey
	return nil
}
