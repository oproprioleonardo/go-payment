package models

import (
	"encoding/json"
	"fmt"
)

// GetCashTransactionResponse represents a GetCashTransactionResponse struct.
// Response object for getting a cash transaction
type GetCashTransactionResponse struct {
	GetTransactionResponse
	// Description
	Description Optional[string] `json:"description"`
}

// MarshalJSON implements the json.Marshaler interface for GetCashTransactionResponse.
// It customizes the JSON marshaling process for GetCashTransactionResponse objects.
func (g *GetCashTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCashTransactionResponse object to a map representation for JSON marshaling.
func (g *GetCashTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "cash"
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCashTransactionResponse.
// It customizes the JSON unmarshaling process for GetCashTransactionResponse objects.
func (g *GetCashTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		Description Optional[string] `json:"description"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Description = temp.Description
	return nil
}

// description returns the description from the GetCashTransactionResponse.
func (g *GetCashTransactionResponse) GetDescription() Optional[string] {
	return g.Description
}

// UnmarshalGetCashTransactionResponse is a utility function used to unmarshal JSON into a GetCashTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetCashTransactionResponseInterface and an error, if any.
func UnmarshalGetCashTransactionResponse(
	inputType string,
	input []byte) (
	GetCashTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getCashTransactionResponse, ok := getTransactionResponse.(GetCashTransactionResponseInterface); ok {
		return getCashTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getCashTransactionResponse %v", getTransactionResponse)
}

// ToGetCashTransactionResponseArray converts an array of GetCashTransactionResponseField to an array of GetCashTransactionResponseInterface.
func ToGetCashTransactionResponseArray(getCashTransactionResponse []GetCashTransactionResponseField) []GetCashTransactionResponseInterface {
	var items []GetCashTransactionResponseInterface
	for _, temp := range getCashTransactionResponse {
		items = append(items, temp.GetCashTransactionResponseInterface)
	}
	return items
}

// GetCashTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetCashTransactionResponseField struct {
	GetCashTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCashTransactionResponseField.
// It customizes the JSON unmarshaling process for GetCashTransactionResponseField objects.
func (g *GetCashTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetCashTransactionResponseInterface, err = UnmarshalGetCashTransactionResponse("cash", input)
	return err
}

// GetCashTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetCashTransactionResponseSlice struct {
	Slice []GetCashTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCashTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetCashTransactionResponseSlice objects.
func (ga *GetCashTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetCashTransactionResponseInterface{}
		for _, getCashTransactionResponse := range temp {
			if getCashTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetCashTransactionResponseInterface
			obj, err := UnmarshalGetCashTransactionResponse("cash", getCashTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
