package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateShippingRequest represents a CreateShippingRequest struct.
// Shipping data
type CreateShippingRequest struct {
	// Shipping amount
	Amount int `json:"amount"`
	// Description
	Description string `json:"description"`
	// Recipient name
	RecipientName string `json:"recipient_name"`
	// Recipient phone number
	RecipientPhone string `json:"recipient_phone"`
	// The id of the address that will be used for shipping
	AddressId string `json:"address_id"`
	// Address data
	Address CreateAddressRequest `json:"address"`
	// Data máxima de entrega
	MaxDeliveryDate *time.Time `json:"max_delivery_date,omitempty"`
	// Prazo estimado de entrega
	EstimatedDeliveryDate *time.Time `json:"estimated_delivery_date,omitempty"`
	// Shipping type
	Type string `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for CreateShippingRequest.
// It customizes the JSON marshaling process for CreateShippingRequest objects.
func (c *CreateShippingRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateShippingRequest object to a map representation for JSON marshaling.
func (c *CreateShippingRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["description"] = c.Description
	structMap["recipient_name"] = c.RecipientName
	structMap["recipient_phone"] = c.RecipientPhone
	structMap["address_id"] = c.AddressId
	structMap["address"] = c.Address
	if c.MaxDeliveryDate != nil {
		structMap["max_delivery_date"] = c.MaxDeliveryDate.Format(time.RFC3339)
	}
	if c.EstimatedDeliveryDate != nil {
		structMap["estimated_delivery_date"] = c.EstimatedDeliveryDate.Format(time.RFC3339)
	}
	structMap["type"] = c.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateShippingRequest.
// It customizes the JSON unmarshaling process for CreateShippingRequest objects.
func (c *CreateShippingRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount                int                  `json:"amount"`
		Description           string               `json:"description"`
		RecipientName         string               `json:"recipient_name"`
		RecipientPhone        string               `json:"recipient_phone"`
		AddressId             string               `json:"address_id"`
		Address               CreateAddressRequest `json:"address"`
		MaxDeliveryDate       *string              `json:"max_delivery_date,omitempty"`
		EstimatedDeliveryDate *string              `json:"estimated_delivery_date,omitempty"`
		Type                  string               `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Description = temp.Description
	c.RecipientName = temp.RecipientName
	c.RecipientPhone = temp.RecipientPhone
	c.AddressId = temp.AddressId
	c.Address = temp.Address
	if temp.MaxDeliveryDate != nil {
		MaxDeliveryDateVal, err := time.Parse(time.RFC3339, *temp.MaxDeliveryDate)
		if err != nil {
			log.Fatalf("Cannot Parse max_delivery_date as % s format.", time.RFC3339)
		}
		c.MaxDeliveryDate = &MaxDeliveryDateVal
	}
	if temp.EstimatedDeliveryDate != nil {
		EstimatedDeliveryDateVal, err := time.Parse(time.RFC3339, *temp.EstimatedDeliveryDate)
		if err != nil {
			log.Fatalf("Cannot Parse estimated_delivery_date as % s format.", time.RFC3339)
		}
		c.EstimatedDeliveryDate = &EstimatedDeliveryDateVal
	}
	c.Type = temp.Type
	return nil
}
