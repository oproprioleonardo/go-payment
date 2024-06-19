package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetAddressResponse represents a GetAddressResponse struct.
// Response object for getting an Address
type GetAddressResponse struct {
	Id           Optional[string]              `json:"id"`
	Street       Optional[string]              `json:"street"`
	Number       Optional[string]              `json:"number"`
	Complement   Optional[string]              `json:"complement"`
	ZipCode      Optional[string]              `json:"zip_code"`
	Neighborhood Optional[string]              `json:"neighborhood"`
	City         Optional[string]              `json:"city"`
	State        Optional[string]              `json:"state"`
	Country      Optional[string]              `json:"country"`
	Status       Optional[string]              `json:"status"`
	CreatedAt    Optional[time.Time]           `json:"created_at"`
	UpdatedAt    Optional[time.Time]           `json:"updated_at"`
	Customer     Optional[GetCustomerResponse] `json:"customer"`
	Metadata     Optional[map[string]string]   `json:"metadata"`
	// Line 1 for address
	Line1 Optional[string] `json:"line_1"`
	// Line 2 for address
	Line2     Optional[string]    `json:"line_2"`
	DeletedAt Optional[time.Time] `json:"deleted_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetAddressResponse.
// It customizes the JSON marshaling process for GetAddressResponse objects.
func (g *GetAddressResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetAddressResponse object to a map representation for JSON marshaling.
func (g *GetAddressResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Street.IsValueSet() {
		structMap["street"] = g.Street.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.Complement.IsValueSet() {
		structMap["complement"] = g.Complement.Value()
	}
	if g.ZipCode.IsValueSet() {
		structMap["zip_code"] = g.ZipCode.Value()
	}
	if g.Neighborhood.IsValueSet() {
		structMap["neighborhood"] = g.Neighborhood.Value()
	}
	if g.City.IsValueSet() {
		structMap["city"] = g.City.Value()
	}
	if g.State.IsValueSet() {
		structMap["state"] = g.State.Value()
	}
	if g.Country.IsValueSet() {
		structMap["country"] = g.Country.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
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
	if g.Customer.IsValueSet() {
		structMap["customer"] = g.Customer.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.Line1.IsValueSet() {
		structMap["line_1"] = g.Line1.Value()
	}
	if g.Line2.IsValueSet() {
		structMap["line_2"] = g.Line2.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAddressResponse.
// It customizes the JSON unmarshaling process for GetAddressResponse objects.
func (g *GetAddressResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id           Optional[string]              `json:"id"`
		Street       Optional[string]              `json:"street"`
		Number       Optional[string]              `json:"number"`
		Complement   Optional[string]              `json:"complement"`
		ZipCode      Optional[string]              `json:"zip_code"`
		Neighborhood Optional[string]              `json:"neighborhood"`
		City         Optional[string]              `json:"city"`
		State        Optional[string]              `json:"state"`
		Country      Optional[string]              `json:"country"`
		Status       Optional[string]              `json:"status"`
		CreatedAt    Optional[string]              `json:"created_at"`
		UpdatedAt    Optional[string]              `json:"updated_at"`
		Customer     Optional[GetCustomerResponse] `json:"customer"`
		Metadata     Optional[map[string]string]   `json:"metadata"`
		Line1        Optional[string]              `json:"line_1"`
		Line2        Optional[string]              `json:"line_2"`
		DeletedAt    Optional[string]              `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Street = temp.Street
	g.Number = temp.Number
	g.Complement = temp.Complement
	g.ZipCode = temp.ZipCode
	g.Neighborhood = temp.Neighborhood
	g.City = temp.City
	g.State = temp.State
	g.Country = temp.Country
	g.Status = temp.Status
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
	g.Customer = temp.Customer
	g.Metadata = temp.Metadata
	g.Line1 = temp.Line1
	g.Line2 = temp.Line2
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	return nil
}
