package bankaccount

import (
	"errors"
	"fmt"

	"github.com/go-training/bankingapp/customer"
)

type BankAccountInterface interface {
	Init(customer customer.Customer, accountId int64, accountType string, openingBalance float64)
	Withdraw(withdrawAmount float64) error
	Deposit(depositAmount float64)
	BalanceCheck()
}

type BankAccount struct {
	AccountHolder  customer.Customer `json:"customer"`
	AccountId      int64             `json:"account_id"`
	AccountType    string            `json:"account_type"`
	OpeningBalance float64           `json:"balance"`
}

func (account *BankAccount) Init(accountHolder customer.Customer, accountId int64, accountType string, openingBalance float64) {
	account.AccountHolder = accountHolder
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
	fmt.Printf("Balance for the %s account = %f\n", account.AccountHolder.CustomerName, account.OpeningBalance)
}
