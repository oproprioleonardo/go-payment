package models

import (
	"encoding/json"
	"log"
	"time"
)

// CreateAnticipationRequest represents a CreateAnticipationRequest struct.
// Request for creating an anticipation
type CreateAnticipationRequest struct {
	// Amount requested for the anticipation
	Amount int `json:"amount"`
	// Timeframe
	Timeframe string `json:"timeframe"`
	// Payment date
	PaymentDate time.Time `json:"payment_date"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAnticipationRequest.
// It customizes the JSON marshaling process for CreateAnticipationRequest objects.
func (c *CreateAnticipationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAnticipationRequest object to a map representation for JSON marshaling.
func (c *CreateAnticipationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["timeframe"] = c.Timeframe
	structMap["payment_date"] = c.PaymentDate.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAnticipationRequest.
// It customizes the JSON unmarshaling process for CreateAnticipationRequest objects.
func (c *CreateAnticipationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int    `json:"amount"`
		Timeframe   string `json:"timeframe"`
		PaymentDate string `json:"payment_date"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Timeframe = temp.Timeframe
	PaymentDateVal, err := time.Parse(time.RFC3339, temp.PaymentDate)
	if err != nil {
		log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
	}
	c.PaymentDate = PaymentDateVal
	return nil
}
