package customerhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-training/bankingapp/customer"
	"github.com/go-training/bankingapp/dbutil"
)

func GetCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customer.Customers)
}

func CreateCustomer(c *gin.Context) {
	db, err := dbutil.ConnectDb()
	var customer customer.Customer
	if err := c.BindJSON(&customer); err != nil {
		log.Printf("Error - Invalid Data in request: ")
		return
	}
	customerId, err := dbutil.InsertCustomer(db, customer)
	if err != nil {
		log.Printf("Error %s when inserting customer in db", err)
		return
	}
	c.IndentedJSON(http.StatusOK, customerId)
}
