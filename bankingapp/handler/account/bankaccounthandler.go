package accounthandler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	bankaccount "github.com/chauhansantosh/go-training/bankingapp/model/account"
	mysql "github.com/chauhansantosh/go-training/bankingapp/mysqldb"
	"github.com/chauhansantosh/go-training/bankingapp/util"
	"github.com/gin-gonic/gin"
)

var start time.Time

func GetAccounts(ctx *gin.Context) {

	start = time.Now()
	errorRespList := []bankaccount.ErrorResponse{}
	bankaccountList := []bankaccount.BankAccount{}

	var bankAccountResponse bankaccount.BankAccount
	customerId, _ := strconv.ParseInt(ctx.Param("customerId"), 10, 64)

	query := util.GETQUERY + `WHERE (? > 0 AND customer_id = ?) OR (? = 0)`
	fmt.Println("query ==>", query)

	rows, err := mysql.DB.Query(query, customerId, customerId, customerId)
	/*
		rows, err := mysql.DB.Query(`SELECT account_id, customer_id, account_type, balance,
		created_at, updated_at, IFNULL(account_pan, '')  FROM bankdb.bank_account`) */

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
			&bankAccountResponse.IsActive,
			&bankAccountResponse.IsLocked,
			&bankAccountResponse.LockPeriodFd,
			&bankAccountResponse.LockedUntil,
			&bankAccountResponse.PenaltyFd,
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
	//Validate request payload
	if errors, err := util.ValidateRequest(c, bankaccount); err != nil {
		for _, e := range errors {
			errorRespList = constructErrorResponse(e, "1005", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, c, nil)
		return
	}

	err := util.InsertBankAccount(bankaccount)
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

	rows := mysql.DB.QueryRow(util.GETQUERY+`WHERE account_id = ?`, bankAccountId)

	scanError := rows.Scan(&bankAccountResponse.AccountId,
		&bankAccountResponse.CustomerId,
		&bankAccountResponse.AccountType,
		&bankAccountResponse.OpeningBalance,
		&bankAccountResponse.CreatedAt,
		&bankAccountResponse.UpdatedAt,
		&bankAccountResponse.AccountPan,
		&bankAccountResponse.IsActive,
		&bankAccountResponse.IsLocked,
		&bankAccountResponse.LockPeriodFd,
		&bankAccountResponse.LockedUntil,
		&bankAccountResponse.PenaltyFd,
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

func GetAccountsByCustomerId(ctx *gin.Context) {
	GetAccounts(ctx)
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
	preMatureWithdrawal := req.PreMatureWithdrawal

	//Validate requst payload
	if errors, err := util.ValidateRequest(ctx, req); err != nil {
		for _, e := range errors {
			errorRespList = constructErrorResponse(e, "1002", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, ctx, nil)
		return
	}

	// Create a function for preparing failure results.
	fail := func(err error) {
		fmt.Printf("Withdraw error: %v", err)
		errorResponse := constructErrorResponse(err.Error(), "1011", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
	}

	// Get a Tx for making transaction requests.
	tx, err := mysql.DB.BeginTx(ctx, nil)
	if err != nil {
		fail(err)
		return
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Confirm that account have enough funds.
	var enough, overdraftallowed bool
	var openingBalance, overdraftamount, newBalance, newOdAmount float64
	var customerId int64
	//changes for fixed a/c withdrawl
	var accType string
	var isAccActive int
	var lockPeriodFD string
	var fdPenalty float64
	var lockedUntil string
	if err = tx.QueryRowContext(ctx, `SELECT (balance >= ?), balance, customer_id, account_type, is_active, 
	IFNULL(lock_period_fd, ''), penalty_fd, IFNULL(locked_until, ''), odallowed, odamount 
	from bankdb.bank_account where account_id = ?`,
		withdrawAmount, accountId).Scan(&enough, &openingBalance, &customerId,
		&accType, &isAccActive, &lockPeriodFD, &fdPenalty, &lockedUntil, &overdraftallowed, &overdraftamount); err != nil {

		if err == sql.ErrNoRows {
			fail(errors.New("No account found"))
			return
		}
		fail(err)
		return
	}

	if !enough && !overdraftallowed {
		fail(errors.New("Insufficient fund. Please enter again"))
		return
	} else if !enough && overdraftallowed && (overdraftamount+openingBalance) < withdrawAmount {
		fail(errors.New("Insufficient fund. Withdraw amount cannot be covered by overdraft amount, please enter again"))
		return
	}

	if accType == "FIXED" {
		//calculate FD maturity date
		currentDate := time.Now()
		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, lockedUntil)
		if err != nil {
			fail(errors.New("Error in converting FD Maturity date"))
			return
		}
		//fdMaturityDate := t.Format(layout)

		if currentDate.Before(t) && !preMatureWithdrawal {
			fail(fmt.Errorf("Locking period of your FD is still not complete. You need to pass preMatureWithdrawal as true to withdraw. Penalty will be applied in case of premature withdrawal."))
			return
		} else {
			if preMatureWithdrawal && withdrawAmount == openingBalance {
				penalty := openingBalance * (fdPenalty / 100)
				withdrawAmount = openingBalance - penalty
				newBalance = 0
			} else if withdrawAmount != openingBalance {
				fail(fmt.Errorf("All amount has to be withdrawn from FD account. Your account balance is %f. Penalty will be applied in case of premature withdrawal.", openingBalance))
				return
			}
			// close the FD a/c by withdrawing entire amount
			_, err = tx.ExecContext(ctx, "UPDATE bankdb.bank_account SET balance = 0, is_active = 0 WHERE account_id = ?",
				accountId)
			if err != nil {
				fail(err)
				return
			}
		}
	} else if enough {
		// Update the account with new balance
		_, err = tx.ExecContext(ctx, "UPDATE bankdb.bank_account SET balance = balance - ? WHERE account_id = ?",
			withdrawAmount, accountId)
		if err != nil {
			fail(err)
			return
		}
		newBalance = openingBalance - withdrawAmount
	} else if !enough && overdraftallowed && (overdraftamount+openingBalance) >= withdrawAmount {
		// Update the account with new balance
		newOdAmount = overdraftamount - (withdrawAmount - openingBalance)
		_, err = tx.ExecContext(ctx, "UPDATE bankdb.bank_account SET balance = 0, odamount = ? WHERE account_id = ?",
			newOdAmount, accountId)
		if err != nil {
			fail(err)
			return
		}
		newBalance = 0
	}

	// Create a new row in the transaction table.
	result, err := tx.ExecContext(ctx, `INSERT INTO 
	bankdb.transaction(account_id, customer_id,opening_balance, amount, new_balance, transaction_type, new_odamount) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		accountId, customerId, openingBalance, withdrawAmount, newBalance, "DEBIT", newOdAmount)
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
	if overdraftallowed {
		accountRes.OdAmount = newOdAmount
		accountRes.OverdraftAllowed = overdraftallowed
	}
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

	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Error - Invalid Data in request.")
		errorResponse := constructErrorResponse(err.Error(), "1012", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
		return
	}
	depositAmount := req.Amount
	accountPan := req.AccountPan

	//Validate request payload
	if errors, err := util.ValidateRequest(ctx, req); err != nil {
		for _, e := range errors {
			errorRespList = constructErrorResponse(e, "1002", errorRespList)
		}
		constructResponse(http.StatusBadRequest, true, &errorRespList, ctx, nil)
		return
	}

	// Create a function for preparing failure results.
	fail := func(err error) {
		fmt.Printf("Deposit error: %v", err)
		errorResponse := constructErrorResponse(err.Error(), "1014", errorRespList)
		constructResponse(http.StatusBadRequest, true, &errorResponse, ctx, nil)
	}

	// Get a Tx for making transaction requests.
	tx, err := mysql.DB.BeginTx(ctx, nil)
	if err != nil {
		fail(err)
		return
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Check if account exists
	var openingBalance float64
	var customerId int64
	var accountType string
	var isLocked bool
	if err = tx.QueryRowContext(ctx, "SELECT balance, customer_id, account_type, is_locked from bankdb.bank_account where account_id = ?",
		accountId).Scan(&openingBalance, &customerId, &accountType, &isLocked); err != nil {
		if err == sql.ErrNoRows {
			fail(errors.New("No account found"))
			return
		}
		fail(err)
		return
	}

	if accountType == "SAVINGS" && depositAmount > 50000 && accountPan == "" {
		fail(errors.New("PAN is mandatory to deposit the amount more than 50000"))
		return
	}

	if accountType == "SAVINGS" && (openingBalance+depositAmount) > 10000000 {
		fail(errors.New("Balance limit reached. Can not have more than 10 millions in SAVINGS account."))
		return
	}

	if accountType == "CURRENT" && depositAmount > 250000 && accountPan == "" {
		fail(errors.New("PAN is mandatory to deposit the amount more than 250000"))
		return
	}

	if accountType == "FIXED" {
		fail(errors.New("Deposit in fixed account is not allowed more than once."))
		return
	}
	if isLocked {
		fail(errors.New("Account is Locked. Transaction is not allowed."))
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
