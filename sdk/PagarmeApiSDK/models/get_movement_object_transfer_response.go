package models

import (
	"encoding/json"
	"fmt"
)

// GetMovementObjectTransferResponse represents a GetMovementObjectTransferResponse struct.
type GetMovementObjectTransferResponse struct {
	GetMovementObjectBaseResponse
	SourceType           Optional[string] `json:"source_type"`
	SourceId             Optional[string] `json:"source_id"`
	TargetType           Optional[string] `json:"target_type"`
	TargetId             Optional[string] `json:"target_id"`
	Fee                  Optional[string] `json:"fee"`
	FundingDate          Optional[string] `json:"funding_date"`
	FundingEstimatedDate Optional[string] `json:"funding_estimated_date"`
	BankAccount          Optional[string] `json:"bank_account"`
}

// MarshalJSON implements the json.Marshaler interface for GetMovementObjectTransferResponse.
// It customizes the JSON marshaling process for GetMovementObjectTransferResponse objects.
func (g *GetMovementObjectTransferResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetMovementObjectTransferResponse object to a map representation for JSON marshaling.
func (g *GetMovementObjectTransferResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.GetMovementObjectBaseResponse.Object != nil {
		structMap["object"] = *g.GetMovementObjectBaseResponse.Object
	} else {
		structMap["object"] = "transfer"
	}
	if g.SourceType.IsValueSet() {
		structMap["source_type"] = g.SourceType.Value()
	}
	if g.SourceId.IsValueSet() {
		structMap["source_id"] = g.SourceId.Value()
	}
	if g.TargetType.IsValueSet() {
		structMap["target_type"] = g.TargetType.Value()
	}
	if g.TargetId.IsValueSet() {
		structMap["target_id"] = g.TargetId.Value()
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	if g.FundingDate.IsValueSet() {
		structMap["funding_date"] = g.FundingDate.Value()
	}
	if g.FundingEstimatedDate.IsValueSet() {
		structMap["funding_estimated_date"] = g.FundingEstimatedDate.Value()
	}
	if g.BankAccount.IsValueSet() {
		structMap["bank_account"] = g.BankAccount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectTransferResponse.
// It customizes the JSON unmarshaling process for GetMovementObjectTransferResponse objects.
func (g *GetMovementObjectTransferResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		SourceType           Optional[string] `json:"source_type"`
		SourceId             Optional[string] `json:"source_id"`
		TargetType           Optional[string] `json:"target_type"`
		TargetId             Optional[string] `json:"target_id"`
		Fee                  Optional[string] `json:"fee"`
		FundingDate          Optional[string] `json:"funding_date"`
		FundingEstimatedDate Optional[string] `json:"funding_estimated_date"`
		BankAccount          Optional[string] `json:"bank_account"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.SourceType = temp.SourceType
	g.SourceId = temp.SourceId
	g.TargetType = temp.TargetType
	g.TargetId = temp.TargetId
	g.Fee = temp.Fee
	g.FundingDate = temp.FundingDate
	g.FundingEstimatedDate = temp.FundingEstimatedDate
	g.BankAccount = temp.BankAccount
	return nil
}

// source_type returns the source_type from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetSourceType() Optional[string] {
	return g.SourceType
}

// source_id returns the source_id from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetSourceId() Optional[string] {
	return g.SourceId
}

// target_type returns the target_type from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetTargetType() Optional[string] {
	return g.TargetType
}

// target_id returns the target_id from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetTargetId() Optional[string] {
	return g.TargetId
}

// fee returns the fee from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetFee() Optional[string] {
	return g.Fee
}

// funding_date returns the funding_date from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetFundingDate() Optional[string] {
	return g.FundingDate
}

// funding_estimated_date returns the funding_estimated_date from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetFundingEstimatedDate() Optional[string] {
	return g.FundingEstimatedDate
}

// bank_account returns the bank_account from the GetMovementObjectTransferResponse.
func (g *GetMovementObjectTransferResponse) GetBankAccount() Optional[string] {
	return g.BankAccount
}

// UnmarshalGetMovementObjectTransferResponse is a utility function used to unmarshal JSON into a GetMovementObjectTransferResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetMovementObjectTransferResponseInterface and an error, if any.
func UnmarshalGetMovementObjectTransferResponse(
	inputType string,
	input []byte) (
	GetMovementObjectTransferResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectTransferResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectTransferResponseInterface); ok {
		return getMovementObjectTransferResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectTransferResponse %v", getMovementObjectBaseResponse)
}

// ToGetMovementObjectTransferResponseArray converts an array of GetMovementObjectTransferResponseField to an array of GetMovementObjectTransferResponseInterface.
func ToGetMovementObjectTransferResponseArray(getMovementObjectTransferResponse []GetMovementObjectTransferResponseField) []GetMovementObjectTransferResponseInterface {
	var items []GetMovementObjectTransferResponseInterface
	for _, temp := range getMovementObjectTransferResponse {
		items = append(items, temp.GetMovementObjectTransferResponseInterface)
	}
	return items
}

// GetMovementObjectTransferResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetMovementObjectTransferResponseField struct {
	GetMovementObjectTransferResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectTransferResponseField.
// It customizes the JSON unmarshaling process for GetMovementObjectTransferResponseField objects.
func (g *GetMovementObjectTransferResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectTransferResponseInterface, err = UnmarshalGetMovementObjectTransferResponse("transfer", input)
	return err
}

// GetMovementObjectTransferResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetMovementObjectTransferResponseSlice struct {
	Slice []GetMovementObjectTransferResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectTransferResponseSlice.
// It customizes the JSON unmarshaling process for GetMovementObjectTransferResponseSlice objects.
func (ga *GetMovementObjectTransferResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectTransferResponseInterface{}
		for _, getMovementObjectTransferResponse := range temp {
			if getMovementObjectTransferResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectTransferResponseInterface
			obj, err := UnmarshalGetMovementObjectTransferResponse("transfer", getMovementObjectTransferResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
