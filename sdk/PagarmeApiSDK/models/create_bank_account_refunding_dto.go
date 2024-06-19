package models

import (
	"encoding/json"
)

// CreateBankAccountRefundingDTO represents a CreateBankAccountRefundingDTO struct.
// Bank Account
type CreateBankAccountRefundingDTO struct {
	// Nome/razão social do favorecido
	HolderName string `json:"holder_name"`
	// Tipo de titular (pessoa física ou jurídica)
	HolderType string `json:"holder_type"`
	// CPF ou CNPJ do favorecido
	HolderDocument string `json:"holder_document"`
	// Dígitos que identificam cada banco.
	Bank string `json:"bank"`
	// Número da agência bancária
	BranchNumber string `json:"branch_number"`
	// Dígito da agência bancária
	BranchCheckDigit string `json:"branch_check_digit"`
	// Número da conta
	AccountNumber string `json:"account_number"`
	// Dígito verificador da conta
	AccountCheckDigit string `json:"account_check_digit"`
	// Tipo de conta
	Type string `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for CreateBankAccountRefundingDTO.
// It customizes the JSON marshaling process for CreateBankAccountRefundingDTO objects.
func (c *CreateBankAccountRefundingDTO) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateBankAccountRefundingDTO object to a map representation for JSON marshaling.
func (c *CreateBankAccountRefundingDTO) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["holder_name"] = c.HolderName
	structMap["holder_type"] = c.HolderType
	structMap["holder_document"] = c.HolderDocument
	structMap["bank"] = c.Bank
	structMap["branch_number"] = c.BranchNumber
	structMap["branch_check_digit"] = c.BranchCheckDigit
	structMap["account_number"] = c.AccountNumber
	structMap["account_check_digit"] = c.AccountCheckDigit
	structMap["type"] = c.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateBankAccountRefundingDTO.
// It customizes the JSON unmarshaling process for CreateBankAccountRefundingDTO objects.
func (c *CreateBankAccountRefundingDTO) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HolderName        string `json:"holder_name"`
		HolderType        string `json:"holder_type"`
		HolderDocument    string `json:"holder_document"`
		Bank              string `json:"bank"`
		BranchNumber      string `json:"branch_number"`
		BranchCheckDigit  string `json:"branch_check_digit"`
		AccountNumber     string `json:"account_number"`
		AccountCheckDigit string `json:"account_check_digit"`
		Type              string `json:"type"`
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
	return nil
}
