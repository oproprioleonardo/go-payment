package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetTransfer represents a GetTransfer struct.
type GetTransfer struct {
	Id                   string                    `json:"id"`
	GatewayId            string                    `json:"gateway_id"`
	Amount               int                       `json:"amount"`
	Status               string                    `json:"status"`
	CreatedAt            time.Time                 `json:"created_at"`
	UpdatedAt            time.Time                 `json:"updated_at"`
	Metadata             map[string]string         `json:"metadata,omitempty"`
	Fee                  *int                      `json:"fee,omitempty"`
	FundingDate          *time.Time                `json:"funding_date,omitempty"`
	FundingEstimatedDate *time.Time                `json:"funding_estimated_date,omitempty"`
	Type                 string                    `json:"type"`
	Source               GetTransferSourceResponse `json:"source"`
	Target               GetTransferTargetResponse `json:"target"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransfer.
// It customizes the JSON marshaling process for GetTransfer objects.
func (g *GetTransfer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetTransfer object to a map representation for JSON marshaling.
func (g *GetTransfer) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = g.Id
	structMap["gateway_id"] = g.GatewayId
	structMap["amount"] = g.Amount
	structMap["status"] = g.Status
	structMap["created_at"] = g.CreatedAt.Format(time.RFC3339)
	structMap["updated_at"] = g.UpdatedAt.Format(time.RFC3339)
	if g.Metadata != nil {
		structMap["metadata"] = g.Metadata
	}
	if g.Fee != nil {
		structMap["fee"] = g.Fee
	}
	if g.FundingDate != nil {
		structMap["funding_date"] = g.FundingDate.Format(time.RFC3339)
	}
	if g.FundingEstimatedDate != nil {
		structMap["funding_estimated_date"] = g.FundingEstimatedDate.Format(time.RFC3339)
	}
	structMap["type"] = g.Type
	structMap["source"] = g.Source
	structMap["target"] = g.Target
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransfer.
// It customizes the JSON unmarshaling process for GetTransfer objects.
func (g *GetTransfer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                   string                    `json:"id"`
		GatewayId            string                    `json:"gateway_id"`
		Amount               int                       `json:"amount"`
		Status               string                    `json:"status"`
		CreatedAt            string                    `json:"created_at"`
		UpdatedAt            string                    `json:"updated_at"`
		Metadata             map[string]string         `json:"metadata,omitempty"`
		Fee                  *int                      `json:"fee,omitempty"`
		FundingDate          *string                   `json:"funding_date,omitempty"`
		FundingEstimatedDate *string                   `json:"funding_estimated_date,omitempty"`
		Type                 string                    `json:"type"`
		Source               GetTransferSourceResponse `json:"source"`
		Target               GetTransferTargetResponse `json:"target"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.GatewayId = temp.GatewayId
	g.Amount = temp.Amount
	g.Status = temp.Status
	CreatedAtVal, err := time.Parse(time.RFC3339, temp.CreatedAt)
	if err != nil {
		log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
	}
	g.CreatedAt = CreatedAtVal
	UpdatedAtVal, err := time.Parse(time.RFC3339, temp.UpdatedAt)
	if err != nil {
		log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
	}
	g.UpdatedAt = UpdatedAtVal
	g.Metadata = temp.Metadata
	g.Fee = temp.Fee
	if temp.FundingDate != nil {
		FundingDateVal, err := time.Parse(time.RFC3339, *temp.FundingDate)
		if err != nil {
			log.Fatalf("Cannot Parse funding_date as % s format.", time.RFC3339)
		}
		g.FundingDate = &FundingDateVal
	}
	if temp.FundingEstimatedDate != nil {
		FundingEstimatedDateVal, err := time.Parse(time.RFC3339, *temp.FundingEstimatedDate)
		if err != nil {
			log.Fatalf("Cannot Parse funding_estimated_date as % s format.", time.RFC3339)
		}
		g.FundingEstimatedDate = &FundingEstimatedDateVal
	}
	g.Type = temp.Type
	g.Source = temp.Source
	g.Target = temp.Target
	return nil
}
