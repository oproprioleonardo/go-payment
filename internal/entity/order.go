package entity

import "errors"

type OrderRequest struct {
	OrderID       string   `json:"order_id"`
	Customer      Customer `json:"customer"`
	CardHash      string   `json:"card_hash"`
	PaymentMethod string   `json:"payment_method"`
	Total         int      `json:"total"`
}

func (o *OrderRequest) Validate() error {
	if o.OrderID == "" {
		return errors.New("order_id is required")
	}
	if o.Customer == (Customer{}) {
		return errors.New("customer is required")
	}
	if err := o.Customer.Validate(); err != nil {
		return err
	}
	if o.PaymentMethod == "" {
		return errors.New("payment_method is required")
	}
	if o.PaymentMethod == "credit_card" || o.PaymentMethod == "debit_card" {
		if o.CardHash == "" {
			return errors.New("card_hash is required")
		}
	}
	if o.Total <= 0 {
		return errors.New("total must be greater than 0")
	}
	return nil
}

func (o *OrderRequest) Process() (*OrderResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	orderResponse := NewOrderResponse(o.OrderID, "failed")
	if o.Total > 100 {
		orderResponse.Status = "paid"
	} else {
		orderResponse.Status = "refused"
	}
	return orderResponse, nil
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"` // paid, failed
}

func NewOrderResponse(orderID string, status string) *OrderResponse {
	return &OrderResponse{
		OrderID: orderID,
		Status:  status,
	}
}
