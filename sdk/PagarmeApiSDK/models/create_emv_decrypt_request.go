package models

import (
	"encoding/json"
)

// CreateEmvDecryptRequest represents a CreateEmvDecryptRequest struct.
type CreateEmvDecryptRequest struct {
	IccData            string                                  `json:"icc_data"`
	CardSequenceNumber string                                  `json:"card_sequence_number"`
	Data               CreateEmvDataDecryptRequest             `json:"data"`
	Poi                *CreateCardPaymentContactlessPOIRequest `json:"poi,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEmvDecryptRequest.
// It customizes the JSON marshaling process for CreateEmvDecryptRequest objects.
func (c *CreateEmvDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEmvDecryptRequest object to a map representation for JSON marshaling.
func (c *CreateEmvDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["icc_data"] = c.IccData
	structMap["card_sequence_number"] = c.CardSequenceNumber
	structMap["data"] = c.Data
	if c.Poi != nil {
		structMap["poi"] = c.Poi
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEmvDecryptRequest.
// It customizes the JSON unmarshaling process for CreateEmvDecryptRequest objects.
func (c *CreateEmvDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		IccData            string                                  `json:"icc_data"`
		CardSequenceNumber string                                  `json:"card_sequence_number"`
		Data               CreateEmvDataDecryptRequest             `json:"data"`
		Poi                *CreateCardPaymentContactlessPOIRequest `json:"poi,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.IccData = temp.IccData
	c.CardSequenceNumber = temp.CardSequenceNumber
	c.Data = temp.Data
	c.Poi = temp.Poi
	return nil
}
