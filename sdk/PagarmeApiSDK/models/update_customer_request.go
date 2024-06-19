package models

import (
	"encoding/json"
)

// UpdateCustomerRequest represents a UpdateCustomerRequest struct.
// Request for updating a customer
type UpdateCustomerRequest struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// Document number
	Document *string `json:"document,omitempty"`
	// Person type
	Type *string `json:"type,omitempty"`
	// Address
	Address *CreateAddressRequest `json:"address,omitempty"`
	// Metadata
	Metadata map[string]string    `json:"metadata,omitempty"`
	Phones   *CreatePhonesRequest `json:"phones,omitempty"`
	// Código de referência do cliente no sistema da loja. Max: 52 caracteres
	Code *string `json:"code,omitempty"`
	// Gênero do cliente
	Gender       *string `json:"gender,omitempty"`
	DocumentType *string `json:"document_type,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCustomerRequest.
// It customizes the JSON marshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCustomerRequest object to a map representation for JSON marshaling.
func (u *UpdateCustomerRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Name != nil {
		structMap["name"] = u.Name
	}
	if u.Email != nil {
		structMap["email"] = u.Email
	}
	if u.Document != nil {
		structMap["document"] = u.Document
	}
	if u.Type != nil {
		structMap["type"] = u.Type
	}
	if u.Address != nil {
		structMap["address"] = u.Address
	}
	if u.Metadata != nil {
		structMap["metadata"] = u.Metadata
	}
	if u.Phones != nil {
		structMap["phones"] = u.Phones
	}
	if u.Code != nil {
		structMap["code"] = u.Code
	}
	if u.Gender != nil {
		structMap["gender"] = u.Gender
	}
	if u.DocumentType != nil {
		structMap["document_type"] = u.DocumentType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCustomerRequest.
// It customizes the JSON unmarshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name         *string               `json:"name,omitempty"`
		Email        *string               `json:"email,omitempty"`
		Document     *string               `json:"document,omitempty"`
		Type         *string               `json:"type,omitempty"`
		Address      *CreateAddressRequest `json:"address,omitempty"`
		Metadata     map[string]string     `json:"metadata,omitempty"`
		Phones       *CreatePhonesRequest  `json:"phones,omitempty"`
		Code         *string               `json:"code,omitempty"`
		Gender       *string               `json:"gender,omitempty"`
		DocumentType *string               `json:"document_type,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Email = temp.Email
	u.Document = temp.Document
	u.Type = temp.Type
	u.Address = temp.Address
	u.Metadata = temp.Metadata
	u.Phones = temp.Phones
	u.Code = temp.Code
	u.Gender = temp.Gender
	u.DocumentType = temp.DocumentType
	return nil
}
