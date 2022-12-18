package bankaccount

import (
	"errors"
	"fmt"
	"time"
)

type BankAccountInterface interface {
	Init(customerId int64, accountId int64, accountType string, openingBalance float64)
	Withdraw(withdrawAmount float64) error
	Deposit(depositAmount float64)
	BalanceCheck()
}

type BankAccount struct {
	CustomerId     int64      `json:"customer_id,omitempty" validate:"required"`
	AccountId      int64      `json:"account_id,omitempty" validate:"required"`
	AccountType    string     `json:"account_type,omitempty" validate:"required,oneof=SAVINGS CURRENT FIXED"`
	OpeningBalance float64    `json:"balance,omitempty" validate:"required,min=1"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	AccountPan     string     `json:"account_pan,omitempty"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
}

type AccountGetResponse struct {
	StatusCode       int              `json:"statusCode,omitempty"`
	TimeElapsed      int64            `json:"timeElapsed,omitempty"`
	HasErrorResponse bool             `json:"hasErrorResponse"`
	ErrorResponse    *[]ErrorResponse `json:"errorResponse,omitempty"`
	BankAccounts     *[]BankAccount   `json:"bank_accounts,omitempty"`
}

type AccountResponse struct {
	StatusCode       int              `json:"statusCode,omitempty"`
	TimeElapsed      int64            `json:"timeElapsed,omitempty"`
	HasErrorResponse bool             `json:"hasErrorResponse"`
	ErrorResponse    *[]ErrorResponse `json:"errorResponse,omitempty"`
	BankAccount      *BankAccount     `json:"bank_account,omitempty"`
}

type TransactionRequest struct {
	Amount     float64 `json:"amount" validate:"required,min=1"`
	AccountPan string  `json:"account_pan,omitempty"`
}

func (account *BankAccount) Init(customerId int64, accountId int64, accountType string, openingBalance float64) {
	account.CustomerId = customerId
	account.AccountId = accountId
	account.AccountType = accountType
	account.OpeningBalance = openingBalance
}

func (account *BankAccount) Withdraw(withdrawAmount float64) error {
	fmt.Println("Amount withdrawn", withdrawAmount)
	if withdrawAmount > account.OpeningBalance {
		return errors.New("insufficient fund")
	}
	account.OpeningBalance -= withdrawAmount
	account.BalanceCheck()
	return nil
}

func (account *BankAccount) Deposit(depositAmount float64) {
	account.OpeningBalance += depositAmount
	fmt.Println("Amount deposited", depositAmount)
	account.BalanceCheck()
}

func (account *BankAccount) BalanceCheck() {
	fmt.Printf("Balance for the %d account = %f\n", account.CustomerId, account.OpeningBalance)
}
