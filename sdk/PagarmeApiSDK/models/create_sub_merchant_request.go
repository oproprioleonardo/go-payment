package models

import (
	"encoding/json"
)

// CreateSubMerchantRequest represents a CreateSubMerchantRequest struct.
// SubMerchant
type CreateSubMerchantRequest struct {
	// Payment Facilitator Code
	PaymentFacilitatorCode string `json:"payment_facilitator_code"`
	// Code
	Code string `json:"code"`
	// Name
	Name string `json:"name"`
	// Merchant Category Code
	MerchantCategoryCode string `json:"merchant_category_code"`
	// Document number. Only numbers, no special characters.
	Document string `json:"document"`
	// Document type. Can be either 'individual' or 'company'
	Type string `json:"type"`
	// Phone
	Phone CreatePhoneRequest `json:"phone"`
	// Address
	Address CreateAddressRequest `json:"address"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubMerchantRequest.
// It customizes the JSON marshaling process for CreateSubMerchantRequest objects.
func (c *CreateSubMerchantRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubMerchantRequest object to a map representation for JSON marshaling.
func (c *CreateSubMerchantRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_facilitator_code"] = c.PaymentFacilitatorCode
	structMap["code"] = c.Code
	structMap["name"] = c.Name
	structMap["merchant_category_code"] = c.MerchantCategoryCode
	structMap["document"] = c.Document
	structMap["type"] = c.Type
	structMap["phone"] = c.Phone
	structMap["address"] = c.Address
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubMerchantRequest.
// It customizes the JSON unmarshaling process for CreateSubMerchantRequest objects.
func (c *CreateSubMerchantRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentFacilitatorCode string               `json:"payment_facilitator_code"`
		Code                   string               `json:"code"`
		Name                   string               `json:"name"`
		MerchantCategoryCode   string               `json:"merchant_category_code"`
		Document               string               `json:"document"`
		Type                   string               `json:"type"`
		Phone                  CreatePhoneRequest   `json:"phone"`
		Address                CreateAddressRequest `json:"address"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PaymentFacilitatorCode = temp.PaymentFacilitatorCode
	c.Code = temp.Code
	c.Name = temp.Name
	c.MerchantCategoryCode = temp.MerchantCategoryCode
	c.Document = temp.Document
	c.Type = temp.Type
	c.Phone = temp.Phone
	c.Address = temp.Address
	return nil
}
