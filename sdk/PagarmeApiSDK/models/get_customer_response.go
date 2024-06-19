package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetCustomerResponse represents a GetCustomerResponse struct.
// Response object for getting a customer
type GetCustomerResponse struct {
	Id            Optional[string]             `json:"id"`
	Name          Optional[string]             `json:"name"`
	Email         Optional[string]             `json:"email"`
	Delinquent    Optional[bool]               `json:"delinquent"`
	CreatedAt     Optional[time.Time]          `json:"created_at"`
	UpdatedAt     Optional[time.Time]          `json:"updated_at"`
	Document      Optional[string]             `json:"document"`
	Type          Optional[string]             `json:"type"`
	FbAccessToken Optional[string]             `json:"fb_access_token"`
	Address       Optional[GetAddressResponse] `json:"address"`
	Metadata      Optional[map[string]string]  `json:"metadata"`
	Phones        Optional[GetPhonesResponse]  `json:"phones"`
	FbId          Optional[int64]              `json:"fb_id"`
	// Código de referência do cliente no sistema da loja. Max: 52 caracteres
	Code         Optional[string] `json:"code"`
	DocumentType Optional[string] `json:"document_type"`
}

// MarshalJSON implements the json.Marshaler interface for GetCustomerResponse.
// It customizes the JSON marshaling process for GetCustomerResponse objects.
func (g *GetCustomerResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCustomerResponse object to a map representation for JSON marshaling.
func (g *GetCustomerResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Email.IsValueSet() {
		structMap["email"] = g.Email.Value()
	}
	if g.Delinquent.IsValueSet() {
		structMap["delinquent"] = g.Delinquent.Value()
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
	if g.Document.IsValueSet() {
		structMap["document"] = g.Document.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.FbAccessToken.IsValueSet() {
		structMap["fb_access_token"] = g.FbAccessToken.Value()
	}
	if g.Address.IsValueSet() {
		structMap["address"] = g.Address.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Phones.IsValueSet() {
		structMap["phones"] = g.Phones.Value()
	}
	if g.FbId.IsValueSet() {
		structMap["fb_id"] = g.FbId.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.DocumentType.IsValueSet() {
		structMap["document_type"] = g.DocumentType.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCustomerResponse.
// It customizes the JSON unmarshaling process for GetCustomerResponse objects.
func (g *GetCustomerResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            Optional[string]             `json:"id"`
		Name          Optional[string]             `json:"name"`
		Email         Optional[string]             `json:"email"`
		Delinquent    Optional[bool]               `json:"delinquent"`
		CreatedAt     Optional[string]             `json:"created_at"`
		UpdatedAt     Optional[string]             `json:"updated_at"`
		Document      Optional[string]             `json:"document"`
		Type          Optional[string]             `json:"type"`
		FbAccessToken Optional[string]             `json:"fb_access_token"`
		Address       Optional[GetAddressResponse] `json:"address"`
		Metadata      Optional[map[string]string]  `json:"metadata"`
		Phones        Optional[GetPhonesResponse]  `json:"phones"`
		FbId          Optional[int64]              `json:"fb_id"`
		Code          Optional[string]             `json:"code"`
		DocumentType  Optional[string]             `json:"document_type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Name = temp.Name
	g.Email = temp.Email
	g.Delinquent = temp.Delinquent
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
	g.Document = temp.Document
	g.Type = temp.Type
	g.FbAccessToken = temp.FbAccessToken
	g.Address = temp.Address
	g.Metadata = temp.Metadata
	g.Phones = temp.Phones
	g.FbId = temp.FbId
	g.Code = temp.Code
	g.DocumentType = temp.DocumentType
	return nil
}
