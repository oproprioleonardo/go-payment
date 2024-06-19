package models

import (
	"encoding/json"
)

// CreateOrderRequest represents a CreateOrderRequest struct.
// Request for creating an order
type CreateOrderRequest struct {
	// Items
	Items []CreateOrderItemRequest `json:"items"`
	// Customer
	Customer CreateCustomerRequest `json:"customer"`
	// Payment data
	Payments []CreatePaymentRequest `json:"payments"`
	// The order code
	Code string `json:"code"`
	// The customer id
	CustomerId Optional[string] `json:"customer_id"`
	// Shipping data
	Shipping *CreateShippingRequest `json:"shipping,omitempty"`
	// Metadata
	Metadata Optional[map[string]string] `json:"metadata"`
	// Defines whether the order will go through anti-fraud
	AntifraudEnabled *bool `json:"antifraud_enabled,omitempty"`
	// Ip address
	Ip *string `json:"ip,omitempty"`
	// Session id
	SessionId *string `json:"session_id,omitempty"`
	// Request's location
	Location *CreateLocationRequest `json:"location,omitempty"`
	// Device's informations
	Device *CreateDeviceRequest `json:"device,omitempty"`
	Closed bool                 `json:"closed"`
	// Currency
	Currency  *string                 `json:"currency,omitempty"`
	Antifraud *CreateAntifraudRequest `json:"antifraud,omitempty"`
	// SubMerchant
	Submerchant *CreateSubMerchantRequest `json:"submerchant,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrderRequest.
// It customizes the JSON marshaling process for CreateOrderRequest objects.
func (c *CreateOrderRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrderRequest object to a map representation for JSON marshaling.
func (c *CreateOrderRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["items"] = c.Items
	structMap["customer"] = c.Customer
	structMap["payments"] = c.Payments
	structMap["code"] = c.Code
	if c.CustomerId.IsValueSet() {
		structMap["customer_id"] = c.CustomerId.Value()
	}
	if c.Shipping != nil {
		structMap["shipping"] = c.Shipping
	}
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	if c.AntifraudEnabled != nil {
		structMap["antifraud_enabled"] = c.AntifraudEnabled
	}
	if c.Ip != nil {
		structMap["ip"] = c.Ip
	}
	if c.SessionId != nil {
		structMap["session_id"] = c.SessionId
	}
	if c.Location != nil {
		structMap["location"] = c.Location
	}
	if c.Device != nil {
		structMap["device"] = c.Device
	}
	structMap["closed"] = c.Closed
	if c.Currency != nil {
		structMap["currency"] = c.Currency
	}
	if c.Antifraud != nil {
		structMap["antifraud"] = c.Antifraud
	}
	if c.Submerchant != nil {
		structMap["submerchant"] = c.Submerchant
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrderRequest.
// It customizes the JSON unmarshaling process for CreateOrderRequest objects.
func (c *CreateOrderRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Items            []CreateOrderItemRequest    `json:"items"`
		Customer         CreateCustomerRequest       `json:"customer"`
		Payments         []CreatePaymentRequest      `json:"payments"`
		Code             string                      `json:"code"`
		CustomerId       Optional[string]            `json:"customer_id"`
		Shipping         *CreateShippingRequest      `json:"shipping,omitempty"`
		Metadata         Optional[map[string]string] `json:"metadata"`
		AntifraudEnabled *bool                       `json:"antifraud_enabled,omitempty"`
		Ip               *string                     `json:"ip,omitempty"`
		SessionId        *string                     `json:"session_id,omitempty"`
		Location         *CreateLocationRequest      `json:"location,omitempty"`
		Device           *CreateDeviceRequest        `json:"device,omitempty"`
		Closed           bool                        `json:"closed"`
		Currency         *string                     `json:"currency,omitempty"`
		Antifraud        *CreateAntifraudRequest     `json:"antifraud,omitempty"`
		Submerchant      *CreateSubMerchantRequest   `json:"submerchant,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Items = temp.Items
	c.Customer = temp.Customer
	c.Payments = temp.Payments
	c.Code = temp.Code
	c.CustomerId = temp.CustomerId
	c.Shipping = temp.Shipping
	c.Metadata = temp.Metadata
	c.AntifraudEnabled = temp.AntifraudEnabled
	c.Ip = temp.Ip
	c.SessionId = temp.SessionId
	c.Location = temp.Location
	c.Device = temp.Device
	c.Closed = temp.Closed
	c.Currency = temp.Currency
	c.Antifraud = temp.Antifraud
	c.Submerchant = temp.Submerchant
	return nil
}
