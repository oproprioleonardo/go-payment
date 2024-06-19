package models

import (
	"encoding/json"
)

// CreateBankAccountRequest represents a CreateBankAccountRequest struct.
// Request for creating a bank account
type CreateBankAccountRequest struct {
	// Bank account holder name
	HolderName string `json:"holder_name"`
	// Bank account holder type
	HolderType string `json:"holder_type"`
	// Bank account holder document
	HolderDocument string `json:"holder_document"`
	// Bank
	Bank string `json:"bank"`
	// Branch number
	BranchNumber string `json:"branch_number"`
	// Branch check digit
	BranchCheckDigit Optional[string] `json:"branch_check_digit"`
	// Account number
	AccountNumber string `json:"account_number"`
	// Account check digit
	AccountCheckDigit string `json:"account_check_digit"`
	// Bank account type
	Type string `json:"type"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	// Pix key
	PixKey Optional[string] `json:"pix_key"`
}

// MarshalJSON implements the json.Marshaler interface for CreateBankAccountRequest.
// It customizes the JSON marshaling process for CreateBankAccountRequest objects.
func (c *CreateBankAccountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateBankAccountRequest object to a map representation for JSON marshaling.
func (c *CreateBankAccountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["holder_name"] = c.HolderName
	structMap["holder_type"] = c.HolderType
	structMap["holder_document"] = c.HolderDocument
	structMap["bank"] = c.Bank
	structMap["branch_number"] = c.BranchNumber
	if c.BranchCheckDigit.IsValueSet() {
		structMap["branch_check_digit"] = c.BranchCheckDigit.Value()
	}
	structMap["account_number"] = c.AccountNumber
	structMap["account_check_digit"] = c.AccountCheckDigit
	structMap["type"] = c.Type
	structMap["metadata"] = c.Metadata
	if c.PixKey.IsValueSet() {
		structMap["pix_key"] = c.PixKey.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateBankAccountRequest.
// It customizes the JSON unmarshaling process for CreateBankAccountRequest objects.
func (c *CreateBankAccountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HolderName        string            `json:"holder_name"`
		HolderType        string            `json:"holder_type"`
		HolderDocument    string            `json:"holder_document"`
		Bank              string            `json:"bank"`
		BranchNumber      string            `json:"branch_number"`
		BranchCheckDigit  Optional[string]  `json:"branch_check_digit"`
		AccountNumber     string            `json:"account_number"`
		AccountCheckDigit string            `json:"account_check_digit"`
		Type              string            `json:"type"`
		Metadata          map[string]string `json:"metadata"`
		PixKey            Optional[string]  `json:"pix_key"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.HolderName = temp.HolderName
	c.HolderType = temp.HolderType
	c.HolderDocument = temp.HolderDocument
	c.Bank = temp.Bank
	c.BranchNumber = temp.BranchNumber
	c.BranchCheckDigit = temp.BranchCheckDigit
	c.AccountNumber = temp.AccountNumber
	c.AccountCheckDigit = temp.AccountCheckDigit
	c.Type = temp.Type
	c.Metadata = temp.Metadata
	c.PixKey = temp.PixKey
	return nil
}
