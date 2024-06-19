package models

import (
	"encoding/json"
)

// GetMovementObjectBaseResponse represents a GetMovementObjectBaseResponse struct.
// Generic response object for getting a MovementObjectBase.
type GetMovementObjectBaseResponse struct {
	Object    *string          `json:"object,omitempty"`
	Id        Optional[string] `json:"id"`
	Status    Optional[string] `json:"status"`
	Amount    Optional[string] `json:"amount"`
	CreatedAt Optional[string] `json:"created_at"`
	Type      Optional[string] `json:"type"`
	ChargeId  Optional[string] `json:"charge_id"`
	GatewayId Optional[string] `json:"gateway_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetMovementObjectBaseResponse.
// It customizes the JSON marshaling process for GetMovementObjectBaseResponse objects.
func (g *GetMovementObjectBaseResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetMovementObjectBaseResponse object to a map representation for JSON marshaling.
func (g *GetMovementObjectBaseResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Object != nil {
		structMap["object"] = *g.Object
	} else {
		structMap["object"] = "MovementObject"
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.ChargeId.IsValueSet() {
		structMap["charge_id"] = g.ChargeId.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectBaseResponse.
// It customizes the JSON unmarshaling process for GetMovementObjectBaseResponse objects.
func (g *GetMovementObjectBaseResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Object    *string          `json:"object,omitempty"`
		Id        Optional[string] `json:"id"`
		Status    Optional[string] `json:"status"`
		Amount    Optional[string] `json:"amount"`
		CreatedAt Optional[string] `json:"created_at"`
		Type      Optional[string] `json:"type"`
		ChargeId  Optional[string] `json:"charge_id"`
		GatewayId Optional[string] `json:"gateway_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Object = temp.Object
	g.Id = temp.Id
	g.Status = temp.Status
	g.Amount = temp.Amount
	g.CreatedAt = temp.CreatedAt
	g.Type = temp.Type
	g.ChargeId = temp.ChargeId
	g.GatewayId = temp.GatewayId
	return nil
}

// object returns the object from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetObject() *string {
	return g.Object
}

// id returns the id from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetId() Optional[string] {
	return g.Id
}

// status returns the status from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetStatus() Optional[string] {
	return g.Status
}

// amount returns the amount from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetAmount() Optional[string] {
	return g.Amount
}

// created_at returns the created_at from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetCreatedAt() Optional[string] {
	return g.CreatedAt
}

// type returns the type from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetType() Optional[string] {
	return g.Type
}

// charge_id returns the charge_id from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetChargeId() Optional[string] {
	return g.ChargeId
}

// gateway_id returns the gateway_id from the GetMovementObjectBaseResponse.
func (g *GetMovementObjectBaseResponse) GetGatewayId() Optional[string] {
	return g.GatewayId
}

// UnmarshalGetMovementObjectBaseResponse is a utility function used to unmarshal JSON into a GetMovementObjectBaseResponseInterface.
// It takes a JSON input []byte and returns the corresponding GetMovementObjectBaseResponseInterface and an error, if any.
func UnmarshalGetMovementObjectBaseResponse(
	inputType string,
	input []byte) (
	GetMovementObjectBaseResponseInterface,
	error) {
	discrim := &struct {
		Object *string
	}{}

	err := json.Unmarshal(input, &discrim)
	if err != nil {
		return nil, err
	}
	if discrim.Object == nil {
		discrim.Object = &inputType
	}

	var res GetMovementObjectBaseResponseInterface
	switch *discrim.Object {
	case "refund":
		res = &GetMovementObjectRefundResponse{}
	case "feeCollection":
		res = &GetMovementObjectFeeCollectionResponse{}
	case "payable":
		res = &GetMovementObjectPayableResponse{}
	case "transfer":
		res = &GetMovementObjectTransferResponse{}
	default:
		res = &GetMovementObjectBaseResponse{}
	}
	json.Unmarshal(input, res)
	return res, nil
}

// ToGetMovementObjectBaseResponseArray converts an array of GetMovementObjectBaseResponseField to an array of GetMovementObjectBaseResponseInterface.
func ToGetMovementObjectBaseResponseArray(getMovementObjectBaseResponse []GetMovementObjectBaseResponseField) []GetMovementObjectBaseResponseInterface {
	var items []GetMovementObjectBaseResponseInterface
	for _, temp := range getMovementObjectBaseResponse {
		items = append(items, temp.GetMovementObjectBaseResponseInterface)
	}
	return items
}

// GetMovementObjectBaseResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetMovementObjectBaseResponseField struct {
	GetMovementObjectBaseResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectBaseResponseField.
// It customizes the JSON unmarshaling process for GetMovementObjectBaseResponseField objects.
func (g *GetMovementObjectBaseResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectBaseResponseInterface, err = UnmarshalGetMovementObjectBaseResponse("MovementObject", input)
	return err
}

// GetMovementObjectBaseResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetMovementObjectBaseResponseSlice struct {
	Slice []GetMovementObjectBaseResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectBaseResponseSlice.
// It customizes the JSON unmarshaling process for GetMovementObjectBaseResponseSlice objects.
func (ga *GetMovementObjectBaseResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectBaseResponseInterface{}
		for _, getMovementObjectBaseResponse := range temp {
			if getMovementObjectBaseResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectBaseResponseInterface
			obj, err := UnmarshalGetMovementObjectBaseResponse("MovementObject", getMovementObjectBaseResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
