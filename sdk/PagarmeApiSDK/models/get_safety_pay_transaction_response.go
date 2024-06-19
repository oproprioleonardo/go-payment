package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// GetSafetyPayTransactionResponse represents a GetSafetyPayTransactionResponse struct.
// Response object for getting a safety pay transaction
type GetSafetyPayTransactionResponse struct {
	GetTransactionResponse
	// Payment url
	Url Optional[string] `json:"url"`
	// Transaction identifier on bank
	BankTid Optional[string] `json:"bank_tid"`
	// Payment date
	PaidAt Optional[time.Time] `json:"paid_at"`
	// Paid amount
	PaidAmount Optional[int] `json:"paid_amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetSafetyPayTransactionResponse.
// It customizes the JSON marshaling process for GetSafetyPayTransactionResponse objects.
func (g *GetSafetyPayTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSafetyPayTransactionResponse object to a map representation for JSON marshaling.
func (g *GetSafetyPayTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "safetypay"
	}
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.BankTid.IsValueSet() {
		structMap["bank_tid"] = g.BankTid.Value()
	}
	if g.PaidAt.IsValueSet() {
		var PaidAtVal *string = nil
		if g.PaidAt.Value() != nil {
			val := g.PaidAt.Value().Format(time.RFC3339)
			PaidAtVal = &val
		}
		structMap["paid_at"] = PaidAtVal
	}
	if g.PaidAmount.IsValueSet() {
		structMap["paid_amount"] = g.PaidAmount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSafetyPayTransactionResponse.
// It customizes the JSON unmarshaling process for GetSafetyPayTransactionResponse objects.
func (g *GetSafetyPayTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Url        Optional[string] `json:"url"`
		BankTid    Optional[string] `json:"bank_tid"`
		PaidAt     Optional[string] `json:"paid_at"`
		PaidAmount Optional[int]    `json:"paid_amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.BankTid = temp.BankTid
	g.PaidAt.ShouldSetValue(temp.PaidAt.IsValueSet())
	if temp.PaidAt.Value() != nil {
		PaidAtVal, err := time.Parse(time.RFC3339, (*temp.PaidAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse paid_at as % s format.", time.RFC3339)
		}
		g.PaidAt.SetValue(&PaidAtVal)
	}
	g.PaidAmount = temp.PaidAmount
	return nil
}

// url returns the url from the GetSafetyPayTransactionResponse.
func (g *GetSafetyPayTransactionResponse) GetUrl() Optional[string] {
	return g.Url
}

// bank_tid returns the bank_tid from the GetSafetyPayTransactionResponse.
func (g *GetSafetyPayTransactionResponse) GetBankTid() Optional[string] {
	return g.BankTid
}

// paid_at returns the paid_at from the GetSafetyPayTransactionResponse.
func (g *GetSafetyPayTransactionResponse) GetPaidAt() Optional[time.Time] {
	return g.PaidAt
}

// paid_amount returns the paid_amount from the GetSafetyPayTransactionResponse.
func (g *GetSafetyPayTransactionResponse) GetPaidAmount() Optional[int] {
	return g.PaidAmount
}

// UnmarshalGetSafetyPayTransactionResponse is a utility function used to unmarshal JSON into a GetSafetyPayTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetSafetyPayTransactionResponseInterface and an error, if any.
func UnmarshalGetSafetyPayTransactionResponse(
	inputType string,
	input []byte) (
	GetSafetyPayTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getSafetyPayTransactionResponse, ok := getTransactionResponse.(GetSafetyPayTransactionResponseInterface); ok {
		return getSafetyPayTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getSafetyPayTransactionResponse %v", getTransactionResponse)
}

// ToGetSafetyPayTransactionResponseArray converts an array of GetSafetyPayTransactionResponseField to an array of GetSafetyPayTransactionResponseInterface.
func ToGetSafetyPayTransactionResponseArray(getSafetyPayTransactionResponse []GetSafetyPayTransactionResponseField) []GetSafetyPayTransactionResponseInterface {
	var items []GetSafetyPayTransactionResponseInterface
	for _, temp := range getSafetyPayTransactionResponse {
		items = append(items, temp.GetSafetyPayTransactionResponseInterface)
	}
	return items
}

// GetSafetyPayTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetSafetyPayTransactionResponseField struct {
	GetSafetyPayTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSafetyPayTransactionResponseField.
// It customizes the JSON unmarshaling process for GetSafetyPayTransactionResponseField objects.
func (g *GetSafetyPayTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetSafetyPayTransactionResponseInterface, err = UnmarshalGetSafetyPayTransactionResponse("safetypay", input)
	return err
}

// GetSafetyPayTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetSafetyPayTransactionResponseSlice struct {
	Slice []GetSafetyPayTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSafetyPayTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetSafetyPayTransactionResponseSlice objects.
func (ga *GetSafetyPayTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetSafetyPayTransactionResponseInterface{}
		for _, getSafetyPayTransactionResponse := range temp {
			if getSafetyPayTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetSafetyPayTransactionResponseInterface
			obj, err := UnmarshalGetSafetyPayTransactionResponse("safetypay", getSafetyPayTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
