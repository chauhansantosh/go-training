package customer

import (
	"fmt"
	"time"
)

type CustomerInterface interface {
	Create(customerId int64, customerName string, customerType string)
	DisplayDetails()
}

type Customer struct {
	CustomerId   int64      `json:"customer_id,omitempty" validate:"required"`
	CustomerName string     `json:"customer_name,omitempty" validate:"required"`
	CustomerType string     `json:"customer_type,omitempty" validate:"required,oneof=INDIVIDUAL COMPANY"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

func (customer *Customer) Create(customerId int64, customerName string, customerType string) {
	customer.CustomerId = customerId
	customer.CustomerName = customerName
	customer.CustomerType = customerType
}

func (customer *Customer) DisplayDetails() {
	fmt.Println("Customer ", customer.CustomerName, customer.CustomerId, customer.CustomerType)
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
}

type CustomerResponse struct {
	StatusCode       int              `json:"statusCode,omitempty"`
	TimeElapsed      int64            `json:"timeElapsed,omitempty"`
	HasErrorResponse bool             `json:"hasErrorResponse"`
	ErrorResponse    *[]ErrorResponse `json:"errorResponse,omitempty"`
	Customer         *Customer        `json:"customer,omitempty"`
}

type CustomerGetResponse struct {
	StatusCode       int              `json:"statusCode,omitempty"`
	TimeElapsed      int64            `json:"timeElapsed,omitempty"`
	HasErrorResponse bool             `json:"hasErrorResponse"`
	ErrorResponse    *[]ErrorResponse `json:"errorResponse,omitempty"`
	Customers        *[]Customer      `json:"customers,omitempty"`
}
