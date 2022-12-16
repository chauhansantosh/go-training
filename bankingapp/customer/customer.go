package customer

import "fmt"

type CustomerInterface interface {
	Create(customerId int64, customerName string, customerType string)
	DisplayDetails()
}

type Customer struct {
	CustomerId   int64  `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	CustomerType string `json:"customer_type"`
}

func (customer *Customer) Create(customerId int64, customerName string, customerType string) {
	customer.CustomerId = customerId
	customer.CustomerName = customerName
	customer.CustomerType = customerType
}

func (customer *Customer) DisplayDetails() {
	fmt.Println("Customer ", customer.CustomerName, customer.CustomerId, customer.CustomerType)
}

var Customers = []Customer{
	{CustomerId: 12345, CustomerName: "Santosh", CustomerType: "INDIVIDUAL"},
	{CustomerId: 23456, CustomerName: "Deepika", CustomerType: "INDIVIDUAL"},
	{CustomerId: 34567, CustomerName: "HCL", CustomerType: "COMPANY"},
}
