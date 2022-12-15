package fourthassessment

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello from fourth assessment!")
}

type BankAccount struct {
	AccountHolderName string
	AccountId         int64
	AccountType       string
	OpeningBalance    float64
}

func (account *BankAccount) Initialization(accountHolderName string, accountId int64, accountType string, openingBalance float64) {
	account.AccountHolderName = accountHolderName
	account.AccountId = accountId
	account.AccountType = accountType
	account.OpeningBalance = openingBalance
}

func (account *BankAccount) Withdraw() {
	var withdrawAmount float64
	fmt.Print("\nEnter the amount to withdraw: ")
	fmt.Scanf("%f", &withdrawAmount)

	if withdrawAmount > account.OpeningBalance {
		fmt.Println("Insufficient fund")
		return
	}
	account.OpeningBalance -= withdrawAmount
	account.DisplayBalance()
}

func (account *BankAccount) Deposit() {
	var depositAmount float64
	fmt.Print("\nEnter the amount to deposit: ")
	fmt.Scanf("%f", &depositAmount)
	account.OpeningBalance += depositAmount
	account.DisplayBalance()
}

func (account *BankAccount) DisplayBalance() {
	fmt.Printf("Balance for the %s account = %f\n", account.AccountHolderName, account.OpeningBalance)
}
