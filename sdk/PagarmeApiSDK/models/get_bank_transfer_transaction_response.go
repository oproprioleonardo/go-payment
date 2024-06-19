package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// GetBankTransferTransactionResponse represents a GetBankTransferTransactionResponse struct.
// Response object for getting a bank transfer transaction
type GetBankTransferTransactionResponse struct {
	GetTransactionResponse
	// Payment url
	Url *string `json:"url,omitempty"`
	// Transaction identifier for the bank
	BankTid *string `json:"bank_tid,omitempty"`
	// Bank
	Bank *string `json:"bank,omitempty"`
	// Payment date
	PaidAt *time.Time `json:"paid_at,omitempty"`
	// Paid amount
	PaidAmount *int `json:"paid_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for GetBankTransferTransactionResponse.
// It customizes the JSON marshaling process for GetBankTransferTransactionResponse objects.
func (g *GetBankTransferTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetBankTransferTransactionResponse object to a map representation for JSON marshaling.
func (g *GetBankTransferTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "bank_transfer"
	}
	if g.Url != nil {
		structMap["url"] = g.Url
	}
	if g.BankTid != nil {
		structMap["bank_tid"] = g.BankTid
	}
	if g.Bank != nil {
		structMap["bank"] = g.Bank
	}
	if g.PaidAt != nil {
		structMap["paid_at"] = g.PaidAt.Format(time.RFC3339)
	}
	if g.PaidAmount != nil {
		structMap["paid_amount"] = g.PaidAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBankTransferTransactionResponse.
// It customizes the JSON unmarshaling process for GetBankTransferTransactionResponse objects.
func (g *GetBankTransferTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Url        *string `json:"url,omitempty"`
		BankTid    *string `json:"bank_tid,omitempty"`
		Bank       *string `json:"bank,omitempty"`
		PaidAt     *string `json:"paid_at,omitempty"`
		PaidAmount *int    `json:"paid_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.BankTid = temp.BankTid
	g.Bank = temp.Bank
	if temp.PaidAt != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, *temp.PaidAt)
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt = &PaidAtVal
	}
	g.PaidAmount = temp.PaidAmount
	return nil
}

// url returns the url from the GetBankTransferTransactionResponse.
func (g *GetBankTransferTransactionResponse) GetUrl() *string {
	return g.Url
}

// bank_tid returns the bank_tid from the GetBankTransferTransactionResponse.
func (g *GetBankTransferTransactionResponse) GetBankTid() *string {
	return g.BankTid
}

// bank returns the bank from the GetBankTransferTransactionResponse.
func (g *GetBankTransferTransactionResponse) GetBank() *string {
	return g.Bank
}

// paid_at returns the paid_at from the GetBankTransferTransactionResponse.
func (g *GetBankTransferTransactionResponse) GetPaidAt() *time.Time {
	return g.PaidAt
}

// paid_amount returns the paid_amount from the GetBankTransferTransactionResponse.
func (g *GetBankTransferTransactionResponse) GetPaidAmount() *int {
	return g.PaidAmount
}

// UnmarshalGetBankTransferTransactionResponse is a utility function used to unmarshal JSON into a GetBankTransferTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetBankTransferTransactionResponseInterface and an error, if any.
func UnmarshalGetBankTransferTransactionResponse(
	inputType string,
	input []byte) (
	GetBankTransferTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getBankTransferTransactionResponse, ok := getTransactionResponse.(GetBankTransferTransactionResponseInterface); ok {
		return getBankTransferTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getBankTransferTransactionResponse %v", getTransactionResponse)
}

// ToGetBankTransferTransactionResponseArray converts an array of GetBankTransferTransactionResponseField to an array of GetBankTransferTransactionResponseInterface.
func ToGetBankTransferTransactionResponseArray(getBankTransferTransactionResponse []GetBankTransferTransactionResponseField) []GetBankTransferTransactionResponseInterface {
	var items []GetBankTransferTransactionResponseInterface
	for _, temp := range getBankTransferTransactionResponse {
		items = append(items, temp.GetBankTransferTransactionResponseInterface)
	}
	return items
}

// GetBankTransferTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetBankTransferTransactionResponseField struct {
	GetBankTransferTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBankTransferTransactionResponseField.
// It customizes the JSON unmarshaling process for GetBankTransferTransactionResponseField objects.
func (g *GetBankTransferTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetBankTransferTransactionResponseInterface, err = UnmarshalGetBankTransferTransactionResponse("bank_transfer", input)
	return err
}

// GetBankTransferTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetBankTransferTransactionResponseSlice struct {
	Slice []GetBankTransferTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBankTransferTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetBankTransferTransactionResponseSlice objects.
func (ga *GetBankTransferTransactionResponseSlice) UnmarshalJSON(input []byte) error {
	tempStruct := struct {
		Slice []json.RawMessage
	}{}
	temp := []json.RawMessage{}
	err := json.Unmarshal(input, &tempStruct)
	if err != nil {
		err := json.Unmarshal(input, &temp)
		if err != nil {
			return err
		}
	} else {
		temp = tempStruct.Slice
	}

	ga.Slice = nil
	if temp != nil {
		ga.Slice = []GetBankTransferTransactionResponseInterface{}
		for _, getBankTransferTransactionResponse := range temp {
			if getBankTransferTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetBankTransferTransactionResponseInterface
			obj, err := UnmarshalGetBankTransferTransactionResponse("bank_transfer", getBankTransferTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
