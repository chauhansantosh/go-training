package bankaccounthandler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-training/bankingapp/bankaccount"
	"github.com/go-training/bankingapp/dbutil"
)

var start time.Time

var timeType = reflect.TypeOf(time.Time{})

func GetAccounts(ctx *gin.Context) {

	start = time.Now()
	errorRespList := []bankaccount.ErrorResponse{}
	bankaccountList := []bankaccount.BankAccount{}

	var bankAccountResponse bankaccount.BankAccount

	rows, err := dbutil.DB.Query(`SELECT account_id, customer_id, account_type, balance, created_at, updated_at, account_pan FROM bankdb.bank_account`)

	if err != nil {
		log.Println("Error while fetching bankaccounts", err)
		errorRespList = constructErrorResponse(err.Error(), "1001", errorRespList)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&bankAccountResponse.AccountId,
			&bankAccountResponse.CustomerId,
			&bankAccountResponse.AccountType,
			&bankAccountResponse.OpeningBalance,
			&bankAccountResponse.CreatedAt,
			&bankAccountResponse.UpdatedAt,
			&bankAccountResponse.AccountPan,
		); err != nil {
			log.Println("Error while scanning for bankaccounts columns", err)
			errorRespList = constructErrorResponse(err.Error(), "1002", errorRespList)
		}
		bankaccountList = append(bankaccountList, bankAccountResponse)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error while fetching bankaccounts", err)
		errorRespList = constructErrorResponse(err.Error(), "1003", errorRespList)
	}

	//Return response based on the data retrieved from DB else throw error
	switch {
	case len(errorRespList) == 0:
		constructGetResponse(http.StatusOK, false,
			nil, ctx, &bankaccountList)
		return
	default:
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	}

}

func CreateAccount(c *gin.Context) {
	errorRespList := []bankaccount.ErrorResponse{}
	var bankaccount bankaccount.BankAccount

	if err := c.BindJSON(&bankaccount); err != nil {
		log.Printf("Error - Invalid Data in request.")
		errorResponse := constructErrorResponse(err.Error(), "1004", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, c, nil)
		return
	}

	v := validator.New()
	err := v.Struct(bankaccount)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorRespList = constructErrorResponse(fmt.Sprint(e), "1005", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, c, nil)
		return
	}

	err = dbutil.InsertBankAccount(bankaccount)
	switch err {
	case nil:
		constructResponse(http.StatusOK, false, nil, c, &bankaccount)
		return

	default:
		log.Printf("Error %s when inserting bank account in db", err)
		errorResponse := constructErrorResponse(err.Error(), "1006", errorRespList)
		constructResponse(http.StatusInternalServerError, true,
			&errorResponse, c, nil)
		return
	}
}

func GetAccountById(ctx *gin.Context) {

	start = time.Now()
	errorRespList := []bankaccount.ErrorResponse{}
	bankAccountList := []bankaccount.BankAccount{}
	var bankAccountResponse bankaccount.BankAccount
	bankAccountId, _ := strconv.ParseInt(ctx.Param("accountId"), 10, 64)

	rows := dbutil.DB.QueryRow(`SELECT 
	account_id, customer_id, account_type, balance, created_at, updated_at, account_pan 
	FROM bankdb.bank_account
	WHERE account_id = ?`, bankAccountId)

	scanError := rows.Scan(&bankAccountResponse.AccountId,
		&bankAccountResponse.CustomerId,
		&bankAccountResponse.AccountType,
		&bankAccountResponse.OpeningBalance,
		&bankAccountResponse.CreatedAt,
		&bankAccountResponse.UpdatedAt,
		&bankAccountResponse.AccountPan,
	)

	bankAccountResponse.BalanceCheck()

	switch scanError {
	case nil:
		bankAccountList = append(bankAccountList, bankAccountResponse)
		constructGetResponse(http.StatusOK, false,
			nil, ctx, &bankAccountList)
		return
	case sql.ErrNoRows:
		log.Printf("No bank account with accountId %d\n", bankAccountId)
		errorRespList = constructErrorResponse(scanError.Error(), "1007", errorRespList)
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	default:
		log.Println("Error while fetching bank account", scanError)
		errorRespList = constructErrorResponse(scanError.Error(), "1008", errorRespList)
		constructGetResponse(http.StatusInternalServerError, true, &errorRespList, ctx, nil)
		return
	}
}

// Withdraw money from bank account
func Withdraw(ctx *gin.Context) {
	errorRespList := []bankaccount.ErrorResponse{}
	accountId, _ := strconv.ParseInt(ctx.Param("accountId"), 10, 64)

	var req bankaccount.TransactionRequest
	var accountRes bankaccount.BankAccount

	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Error - Invalid Data in request.")
		errorResponse := constructErrorResponse(err.Error(), "1009", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
		return
	}
	withdrawAmount := req.Amount

	v := validator.New()
	err := v.Struct(req)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorRespList = constructErrorResponse(fmt.Sprint(e), "1010", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, ctx, nil)
		return
	}

	// Create a function for preparing failure results.
	fail := func(err error) {
		fmt.Printf("Deposit error: %v", err)
		errorResponse := constructErrorResponse(err.Error(), "1011", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
		return
	}

	// Get a Tx for making transaction requests.
	tx, err := dbutil.DB.BeginTx(ctx, nil)
	if err != nil {
		fail(err)
		return
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Confirm that account have enough funds.
	var enough bool
	var openingBalance float64
	var customerId int64
	if err = tx.QueryRowContext(ctx, "SELECT (balance >= ?), balance, customer_id from bankdb.bank_account where account_id = ?",
		withdrawAmount, accountId).Scan(&enough, &openingBalance, &customerId); err != nil {
		if err == sql.ErrNoRows {
			fail(fmt.Errorf("No account found"))
			return
		}
		fail(err)
		return
	}

	if !enough {
		fail(fmt.Errorf("Insufficient fund. Please enter again"))
		return
	}

	// Update the account with new balance
	_, err = tx.ExecContext(ctx, "UPDATE bankdb.bank_account SET balance = balance - ? WHERE account_id = ?",
		withdrawAmount, accountId)
	if err != nil {
		fail(err)
		return
	}

	newBalance := openingBalance - withdrawAmount

	// Create a new row in the transaction table.
	result, err := tx.ExecContext(ctx, `INSERT INTO 
	bankdb.transaction(account_id, customer_id,opening_balance, amount, new_balance,transaction_type) 
	VALUES (?, ?, ?, ?, ?, ?)`,
		accountId, customerId, openingBalance, withdrawAmount, newBalance, "DEBIT")
	if err != nil {
		fail(err)
		return
	}
	// Get the transaction id of just created transaction.
	transId, err := result.LastInsertId()
	fmt.Println("Transaction Id of the last row inserted", transId)
	if err != nil {
		fail(err)
		return
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		fail(err)
		return
	}
	accountRes.OpeningBalance = newBalance
	accountRes.CustomerId = customerId
	accountRes.AccountId = accountId
	constructResponse(http.StatusOK, false, nil, ctx, &accountRes)
}

// Deposit money to a bank account
func Deposit(ctx *gin.Context) {
	errorRespList := []bankaccount.ErrorResponse{}
	accountId, _ := strconv.ParseInt(ctx.Param("accountId"), 10, 64)

	var req bankaccount.TransactionRequest
	var accountRes bankaccount.BankAccount

	//v, ok := binding.Validator.Engine().(*validator.Validate)

	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Error - Invalid Data in request.")
		errorResponse := constructErrorResponse(err.Error(), "1012", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
		return
	}
	depositAmount := req.Amount
	accountPan := req.AccountPan

	v := validator.New()
	err := v.Struct(req)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorRespList = constructErrorResponse(fmt.Sprint(e), "1013", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, ctx, nil)
		return
	}

	// Create a function for preparing failure results.
	fail := func(err error) {
		fmt.Printf("Deposit error: %v", err)
		errorResponse := constructErrorResponse(err.Error(), "1014", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
		return
	}

	// Get a Tx for making transaction requests.
	tx, err := dbutil.DB.BeginTx(ctx, nil)
	if err != nil {
		fail(err)
		return
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Check if account exists
	var openingBalance float64
	var customerId int64
	if err = tx.QueryRowContext(ctx, "SELECT balance, customer_id from bankdb.bank_account where account_id = ?",
		accountId).Scan(&openingBalance, &customerId); err != nil {
		if err == sql.ErrNoRows {
			fail(fmt.Errorf("No account found"))
			return
		}
		fail(err)
		return
	}

	// Update the account with new balance
	_, err = tx.ExecContext(ctx, "UPDATE bankdb.bank_account SET account_pan = ?, balance = balance + ? WHERE account_id = ?",
		accountPan, depositAmount, accountId)
	if err != nil {
		fail(err)
		return
	}

	newBalance := openingBalance + depositAmount

	// Create a new row in the transaction table.
	result, err := tx.ExecContext(ctx, `INSERT INTO 
	bankdb.transaction(account_id, customer_id,opening_balance, amount, new_balance,transaction_type) 
	VALUES (?, ?, ?, ?, ?, ?)`,
		accountId, customerId, openingBalance, depositAmount, newBalance, "CREDIT")
	if err != nil {
		fail(err)
		return
	}
	// Get the transaction id of just created transaction.
	transId, err := result.LastInsertId()
	fmt.Println("Transaction Id of the last row inserted", transId)
	if err != nil {
		fail(err)
		return
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		fail(err)
		return
	}
	accountRes.OpeningBalance = newBalance
	accountRes.CustomerId = customerId
	accountRes.AccountId = accountId
	accountRes.AccountPan = accountPan
	constructResponse(http.StatusOK, false, nil, ctx, &accountRes)
}

/*
construct response
*/
func constructGetResponse(statusCode int, haserror bool, errorResponse *[]bankaccount.ErrorResponse,
	c *gin.Context, bankaccountObj *[]bankaccount.BankAccount) {
	response := bankaccount.AccountGetResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start).Milliseconds(),
		HasErrorResponse: haserror,
		ErrorResponse:    errorResponse,
		BankAccounts:     bankaccountObj,
	}
	c.JSON(statusCode, response)
}

/*
construct response
*/
func constructResponse(statusCode int, haserror bool, errorResponse *[]bankaccount.ErrorResponse,
	c *gin.Context, bankAccountObj *bankaccount.BankAccount) {
	response := bankaccount.AccountResponse{
		StatusCode:       statusCode,
		TimeElapsed:      time.Since(start).Milliseconds(),
		HasErrorResponse: haserror,
		ErrorResponse:    errorResponse,
		BankAccount:      bankAccountObj,
	}
	c.JSON(statusCode, response)
}

/*
construct error response
*/
func constructErrorResponse(message string, code string, errorList []bankaccount.ErrorResponse) []bankaccount.ErrorResponse {
	errorResponse := bankaccount.ErrorResponse{
		ErrorMessage: message,
		ErrorCode:    code,
	}
	errorList = append(errorList, errorResponse)
	return errorList
}