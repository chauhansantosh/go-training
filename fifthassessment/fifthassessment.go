package fifthassessment

import (
	"errors"
	"fmt"
)

func Hello() {
	fmt.Println("Hello from fifth assessment!")
}

type BankAccountInterface interface {
	Init(accHolderName string, accountId int64, accountType string, openingBalance float64)
	Withdraw(withdrawAmount float64) error
	Deposit(depositAmount float64)
	BalanceCheck()
}

type BankAccount struct {
	AccountHolderName string
	AccountId         int64
	AccountType       string
	OpeningBalance    float64
}

func (account *BankAccount) Init(accountHolderName string, accountId int64, accountType string, openingBalance float64) {
	account.AccountHolderName = accountHolderName
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
	fmt.Printf("Balance for the %s account = %f\n", account.AccountHolderName, account.OpeningBalance)
}
