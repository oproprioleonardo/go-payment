package models

import (
	"encoding/json"
)

// CreateCustomerRequest represents a CreateCustomerRequest struct.
// Request for creating a new customer
type CreateCustomerRequest struct {
	// Name
	Name string `json:"name"`
	// Email
	Email string `json:"email"`
	// Document number. Only numbers, no special characters.
	Document string `json:"document"`
	// Person type. Can be either 'individual' or 'company'
	Type string `json:"type"`
	// The customer's address
	Address CreateAddressRequest `json:"address"`
	// Metadata
	Metadata map[string]string   `json:"metadata"`
	Phones   CreatePhonesRequest `json:"phones"`
	// Customer code
	Code string `json:"code"`
	// Customer Gender
	Gender       *string `json:"gender,omitempty"`
	DocumentType *string `json:"document_type,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCustomerRequest.
// It customizes the JSON marshaling process for CreateCustomerRequest objects.
func (c *CreateCustomerRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCustomerRequest object to a map representation for JSON marshaling.
func (c *CreateCustomerRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["email"] = c.Email
	structMap["document"] = c.Document
	structMap["type"] = c.Type
	structMap["address"] = c.Address
	structMap["metadata"] = c.Metadata
	structMap["phones"] = c.Phones
	structMap["code"] = c.Code
	if c.Gender != nil {
		structMap["gender"] = c.Gender
	}
	if c.DocumentType != nil {
		structMap["document_type"] = c.DocumentType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCustomerRequest.
// It customizes the JSON unmarshaling process for CreateCustomerRequest objects.
func (c *CreateCustomerRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name         string               `json:"name"`
		Email        string               `json:"email"`
		Document     string               `json:"document"`
		Type         string               `json:"type"`
		Address      CreateAddressRequest `json:"address"`
		Metadata     map[string]string    `json:"metadata"`
		Phones       CreatePhonesRequest  `json:"phones"`
		Code         string               `json:"code"`
		Gender       *string              `json:"gender,omitempty"`
		DocumentType *string              `json:"document_type,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.Email = temp.Email
	c.Document = temp.Document
	c.Type = temp.Type
	c.Address = temp.Address
	c.Metadata = temp.Metadata
	c.Phones = temp.Phones
	c.Code = temp.Code
	c.Gender = temp.Gender
	c.DocumentType = temp.DocumentType
	return nil
}
