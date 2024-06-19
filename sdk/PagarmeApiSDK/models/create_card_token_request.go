package models

import (
	"encoding/json"
)

// CreateCardTokenRequest represents a CreateCardTokenRequest struct.
// Card token data
type CreateCardTokenRequest struct {
	// Credit card number
	Number string `json:"number"`
	// Holder name, as written on the card
	HolderName string `json:"holder_name"`
	// The expiration month
	ExpMonth int `json:"exp_month"`
	// The expiration year, that can be informed with 2 or 4 digits
	ExpYear int `json:"exp_year"`
	// The card's security code
	Cvv string `json:"cvv"`
	// Card brand
	Brand string `json:"brand"`
	Label string `json:"label"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCardTokenRequest.
// It customizes the JSON marshaling process for CreateCardTokenRequest objects.
func (c *CreateCardTokenRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCardTokenRequest object to a map representation for JSON marshaling.
func (c *CreateCardTokenRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = c.Number
	structMap["holder_name"] = c.HolderName
	structMap["exp_month"] = c.ExpMonth
	structMap["exp_year"] = c.ExpYear
	structMap["cvv"] = c.Cvv
	structMap["brand"] = c.Brand
	structMap["label"] = c.Label
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCardTokenRequest.
// It customizes the JSON unmarshaling process for CreateCardTokenRequest objects.
func (c *CreateCardTokenRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number     string `json:"number"`
		HolderName string `json:"holder_name"`
		ExpMonth   int    `json:"exp_month"`
		ExpYear    int    `json:"exp_year"`
		Cvv        string `json:"cvv"`
		Brand      string `json:"brand"`
		Label      string `json:"label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Number = temp.Number
	c.HolderName = temp.HolderName
	c.ExpMonth = temp.ExpMonth
	c.ExpYear = temp.ExpYear
	c.Cvv = temp.Cvv
	c.Brand = temp.Brand
	c.Label = temp.Label
	return nil
}
