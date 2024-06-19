package models

import (
	"encoding/json"
)

// GetThreeDSecureResponse represents a GetThreeDSecureResponse struct.
// 3D-S payment authentication response
type GetThreeDSecureResponse struct {
	// MPI Vendor
	Mpi Optional[string] `json:"mpi"`
	// Electronic Commerce Indicator (ECI) (Opcional)
	Eci Optional[string] `json:"eci"`
	// Online payment cryptogram, definido pelo 3-D Secure.
	Cavv Optional[string] `json:"cavv"`
	// Identificador da transação (XID)
	TransactionId Optional[string] `json:"transaction_Id"`
	// Url de redirecionamento de sucessso
	SuccessUrl Optional[string] `json:"success_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetThreeDSecureResponse.
// It customizes the JSON marshaling process for GetThreeDSecureResponse objects.
func (g *GetThreeDSecureResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetThreeDSecureResponse object to a map representation for JSON marshaling.
func (g *GetThreeDSecureResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Mpi.IsValueSet() {
		structMap["mpi"] = g.Mpi.Value()
	}
	if g.Eci.IsValueSet() {
		structMap["eci"] = g.Eci.Value()
	}
	if g.Cavv.IsValueSet() {
		structMap["cavv"] = g.Cavv.Value()
	}
	if g.TransactionId.IsValueSet() {
		structMap["transaction_Id"] = g.TransactionId.Value()
	}
	if g.SuccessUrl.IsValueSet() {
		structMap["success_url"] = g.SuccessUrl.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetThreeDSecureResponse.
// It customizes the JSON unmarshaling process for GetThreeDSecureResponse objects.
func (g *GetThreeDSecureResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Mpi           Optional[string] `json:"mpi"`
		Eci           Optional[string] `json:"eci"`
		Cavv          Optional[string] `json:"cavv"`
		TransactionId Optional[string] `json:"transaction_Id"`
		SuccessUrl    Optional[string] `json:"success_url"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Mpi = temp.Mpi
	g.Eci = temp.Eci
	g.Cavv = temp.Cavv
	g.TransactionId = temp.TransactionId
	g.SuccessUrl = temp.SuccessUrl
	return nil
}
