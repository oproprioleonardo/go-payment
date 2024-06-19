package entity

import "errors"

type Customer struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Document Document `json:"document"`
	Phone    Phone    `json:"phone"`
	Address  Address  `json:"address"`
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.New("customer name is required")
	}
	if c.Email == "" {
		return errors.New("customer email is required")
	}
	if c.Document == (Document{}) {
		return errors.New("customer document is required")
	} else if err := c.Document.Validate(); err != nil {
		return err
	}
	if c.Phone == (Phone{}) {
		return errors.New("customer phone is required")
	} else if err := c.Phone.Validate(); err != nil {
		return err
	}
	if c.Address == (Address{}) {
		return errors.New("customer address is required")
	} else if err := c.Address.Validate(); err != nil {
		return err
	}
	return nil
}

type Document struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (d *Document) Validate() error {
	if d.Type == "" {
		return errors.New("customer document type is required")
	}
	if d.Value == "" {
		return errors.New("customer document value is required")
	}
	return nil
}

type Address struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	Line1   string `json:"line_1"`
	Line2   string `json:"line_2"`
}

func (a *Address) Validate() error {
	if a.Country == "" {
		return errors.New("country is required")
	}
	if a.State == "" {
		return errors.New("state is required")
	}
	if a.City == "" {
		return errors.New("city is required")
	}
	if a.ZipCode == "" {
		return errors.New("zip_code is required")
	}
	if a.Line1 == "" {
		return errors.New("line_1 is required")
	}
	return nil

}

type Phone struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
	Type        string `json:"type"`
}

func (p *Phone) Validate() error {
	if p.CountryCode == "" {
		return errors.New("country_code is required")
	}
	if p.AreaCode == "" {
		return errors.New("area_code is required")
	}
	if p.Number == "" {
		return errors.New("number is required")
	}
	if p.Type == "" {
		return errors.New("type is required")
	}
	if p.Type != "mobile" && p.Type != "landline" {
		return errors.New("type must be mobile or landline")
	}
	return nil
}
