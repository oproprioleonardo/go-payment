package models

import (
	"encoding/json"
	"fmt"
)

// GetMovementObjectFeeCollectionResponse represents a GetMovementObjectFeeCollectionResponse struct.
// Generic response object for getting a MovementObjectFeeCollection.
type GetMovementObjectFeeCollectionResponse struct {
	GetMovementObjectBaseResponse
	Description Optional[string] `json:"description"`
	PaymentDate Optional[string] `json:"payment_date"`
	RecipientId Optional[string] `json:"recipient_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetMovementObjectFeeCollectionResponse.
// It customizes the JSON marshaling process for GetMovementObjectFeeCollectionResponse objects.
func (g *GetMovementObjectFeeCollectionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetMovementObjectFeeCollectionResponse object to a map representation for JSON marshaling.
func (g *GetMovementObjectFeeCollectionResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.GetMovementObjectBaseResponse.Object != nil {
		structMap["object"] = *g.GetMovementObjectBaseResponse.Object
	} else {
		structMap["object"] = "feeCollection"
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.PaymentDate.IsValueSet() {
		structMap["payment_date"] = g.PaymentDate.Value()
	}
	if g.RecipientId.IsValueSet() {
		structMap["recipient_id"] = g.RecipientId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectFeeCollectionResponse.
// It customizes the JSON unmarshaling process for GetMovementObjectFeeCollectionResponse objects.
func (g *GetMovementObjectFeeCollectionResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		Description Optional[string] `json:"description"`
		PaymentDate Optional[string] `json:"payment_date"`
		RecipientId Optional[string] `json:"recipient_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Description = temp.Description
	g.PaymentDate = temp.PaymentDate
	g.RecipientId = temp.RecipientId
	return nil
}

// description returns the description from the GetMovementObjectFeeCollectionResponse.
func (g *GetMovementObjectFeeCollectionResponse) GetDescription() Optional[string] {
	return g.Description
}

// payment_date returns the payment_date from the GetMovementObjectFeeCollectionResponse.
func (g *GetMovementObjectFeeCollectionResponse) GetPaymentDate() Optional[string] {
	return g.PaymentDate
}

// recipient_id returns the recipient_id from the GetMovementObjectFeeCollectionResponse.
func (g *GetMovementObjectFeeCollectionResponse) GetRecipientId() Optional[string] {
	return g.RecipientId
}

// UnmarshalGetMovementObjectFeeCollectionResponse is a utility function used to unmarshal JSON into a GetMovementObjectFeeCollectionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetMovementObjectFeeCollectionResponseInterface and an error, if any.
func UnmarshalGetMovementObjectFeeCollectionResponse(
	inputType string,
	input []byte) (
	GetMovementObjectFeeCollectionResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectFeeCollectionResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectFeeCollectionResponseInterface); ok {
		return getMovementObjectFeeCollectionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectFeeCollectionResponse %v", getMovementObjectBaseResponse)
}

// ToGetMovementObjectFeeCollectionResponseArray converts an array of GetMovementObjectFeeCollectionResponseField to an array of GetMovementObjectFeeCollectionResponseInterface.
func ToGetMovementObjectFeeCollectionResponseArray(getMovementObjectFeeCollectionResponse []GetMovementObjectFeeCollectionResponseField) []GetMovementObjectFeeCollectionResponseInterface {
	var items []GetMovementObjectFeeCollectionResponseInterface
	for _, temp := range getMovementObjectFeeCollectionResponse {
		items = append(items, temp.GetMovementObjectFeeCollectionResponseInterface)
	}
	return items
}

// GetMovementObjectFeeCollectionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetMovementObjectFeeCollectionResponseField struct {
	GetMovementObjectFeeCollectionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectFeeCollectionResponseField.
// It customizes the JSON unmarshaling process for GetMovementObjectFeeCollectionResponseField objects.
func (g *GetMovementObjectFeeCollectionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectFeeCollectionResponseInterface, err = UnmarshalGetMovementObjectFeeCollectionResponse("feeCollection", input)
	return err
}

// GetMovementObjectFeeCollectionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetMovementObjectFeeCollectionResponseSlice struct {
	Slice []GetMovementObjectFeeCollectionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectFeeCollectionResponseSlice.
// It customizes the JSON unmarshaling process for GetMovementObjectFeeCollectionResponseSlice objects.
func (ga *GetMovementObjectFeeCollectionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectFeeCollectionResponseInterface{}
		for _, getMovementObjectFeeCollectionResponse := range temp {
			if getMovementObjectFeeCollectionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectFeeCollectionResponseInterface
			obj, err := UnmarshalGetMovementObjectFeeCollectionResponse("feeCollection", getMovementObjectFeeCollectionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
