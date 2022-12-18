package customerhandler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-training/bankingapp/customer"
	"github.com/go-training/bankingapp/dbutil"
)

var start time.Time

func CreateCustomer(c *gin.Context) {
	errorRespList := []customer.ErrorResponse{}
	var customer customer.Customer

	if err := c.BindJSON(&customer); err != nil {
		log.Printf("Error - Invalid Data in request.")
		errorResponse := constructErrorResponse(err.Error(), "1001", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, c, nil)
		return
	}

	v := validator.New()
	err := v.Struct(customer)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorRespList = constructErrorResponse(fmt.Sprint(e), "1002", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, c, nil)
		return
	}

	_, err = dbutil.InsertCustomer(customer)
	switch err {
	case nil:
		constructResponse(http.StatusOK, false, nil, c, &customer)
		return

	default:
		log.Printf("Error %s when inserting customer in db", err)
		errorResponse := constructErrorResponse(err.Error(), "1003", errorRespList)
		constructResponse(http.StatusInternalServerError, true,
			&errorResponse, c, nil)
		return
	}
}

func GetCustomers(ctx *gin.Context) {

	start = time.Now()
	errorRespList := []customer.ErrorResponse{}
	customerList := []customer.Customer{}

	var customerResponse customer.Customer

	rows, err := dbutil.DB.Query(`SELECT customer_id, customer_name, customer_type, created_at, updated_at FROM bankdb.customer`)

	if err != nil {
		log.Println("Error while fetching customers", err)
		errorRespList = constructErrorResponse(err.Error(), "1004", errorRespList)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&customerResponse.CustomerId,
			&customerResponse.CustomerName,
			&customerResponse.CustomerType,
			&customerResponse.CreatedAt,
			&customerResponse.UpdatedAt); err != nil {
			log.Println("Error while scanning for customers columns", err)
			errorRespList = constructErrorResponse(err.Error(), "1005", errorRespList)
		}
		customerList = append(customerList, customerResponse)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while fetching customers", err)
		errorRespList = constructErrorResponse(err.Error(), "1006", errorRespList)
	}

	//Return response based on the data retrieved from DB else throw error
	switch {
	case len(errorRespList) == 0:
		constructGetResponse(http.StatusOK, false,
			nil, ctx, &customerList)
		return
	default:
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	}

}

func GetCustomerById(ctx *gin.Context) {

	start = time.Now()
	errorRespList := []customer.ErrorResponse{}
	customerList := []customer.Customer{}
	var customerResponse customer.Customer
	customerId, _ := strconv.ParseInt(ctx.Param("customerId"), 10, 64)

	rows := dbutil.DB.QueryRow(`SELECT 
	customer_id, customer_name, customer_type, created_at, updated_at 
	FROM bankdb.customer
	WHERE customer_id = ?`, customerId)

	scanError := rows.Scan(&customerResponse.CustomerId,
		&customerResponse.CustomerName,
		&customerResponse.CustomerType,
		&customerResponse.CreatedAt,
		&customerResponse.UpdatedAt)

	switch scanError {
	case nil:
		customerList = append(customerList, customerResponse)
		constructGetResponse(http.StatusOK, false,
			nil, ctx, &customerList)
		return
	case sql.ErrNoRows:
		log.Printf("No customer with customerId %d\n", customerId)
		errorRespList = constructErrorResponse(scanError.Error(), "1007", errorRespList)
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	default:
		log.Println("Error while fetching customer", scanError)
		errorRespList = constructErrorResponse(scanError.Error(), "1008", errorRespList)
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	}
}

/*
construct response
*/
func constructResponse(statusCode int, haserror bool, errorResponse *[]customer.ErrorResponse,
	c *gin.Context, customerObj *customer.Customer) {
	response := customer.CustomerResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start).Milliseconds(),
		HasErrorResponse: haserror,
		ErrorResponse:    errorResponse,
		Customer:         customerObj,
	}
	c.JSON(statusCode, response)
}

/*
construct response
*/
func constructGetResponse(statusCode int, haserror bool, errorResponse *[]customer.ErrorResponse,
	c *gin.Context, customerObj *[]customer.Customer) {
	response := customer.CustomerGetResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start).Milliseconds(),
		HasErrorResponse: haserror,
		ErrorResponse:    errorResponse,
		Customers:        customerObj,
	}
	c.JSON(statusCode, response)
}

/*
construct error response
*/
func constructErrorResponse(message string, code string, errorList []customer.ErrorResponse) []customer.ErrorResponse {
	errorResponse := customer.ErrorResponse{
		ErrorMessage: message,
		ErrorCode:    code,
	}
	errorList = append(errorList, errorResponse)
	return errorList
}
