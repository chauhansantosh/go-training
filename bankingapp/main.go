package main

import (
	"fmt"
	"log"

	"github.com/chauhansantosh/go-training/bankingapp/util"

	"github.com/gin-gonic/gin"

	bankaccount "github.com/chauhansantosh/go-training/bankingapp/model/account"
	accounthandler "github.com/chauhansantosh/go-training/bankingapp/handler/account"
	customer "github.com/chauhansantosh/go-training/bankingapp/model/customer"
	customerhandler "github.com/chauhansantosh/go-training/bankingapp/handler/customer"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := util.ConnectDb()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	util.DB = db
	defer db.Close()
	log.Printf("Successfully connected to database")

	err = util.CreateTables()
	if err != nil {
		log.Printf("Create table failed with error %s", err)
		return
	}

	var customer1, customer2, customer3 customer.Customer
	customer1 = customer.Customer{}
	customer1.Create(12345, "Santosh", "INDIVIDUAL")
	customer1.DisplayDetails()

	customer2 = customer.Customer{}
	customer2.Create(23456, "Deepika", "INDIVIDUAL")
	customer2.DisplayDetails()

	customer3 = customer.Customer{}
	customer3.Create(34567, "HCL", "COMPANY")
	customer3.DisplayDetails()

	var account1, account2, account3 bankaccount.BankAccount
	account1 = bankaccount.BankAccount{}
	account1.Init(customer1.CustomerId, 51691852, "SAVINGS", 10000.5)
	account1.BalanceCheck()
	//err = account1.Withdraw(1000)
	if err != nil {
		fmt.Println(err)
	}

	account2 = bankaccount.BankAccount{}
	account2.Init(customer2.CustomerId, 51691853, "FIXED", 20000.5)
	account2.BalanceCheck()
	//account2.Deposit(5000)
	//err = account2.Withdraw(10000)
	if err != nil {
		fmt.Println(err)
	}

	account3 = bankaccount.BankAccount{}
	account3.Init(customer3.CustomerId, 51691854, "CURRENT", 30000.0)
	account3.BalanceCheck()
	//err = account3.Withdraw(5000)
	if err != nil {
		fmt.Println(err)
	}
	//err = account3.Withdraw(30000)
	if err != nil {
		fmt.Println(err)
	}

	/* err = multipleInsertCutomer(db, []customer.Customer{customer1, customer2, customer3})
	if err != nil {
		log.Printf("Multiple insert failed with error %s", err)
		return
	}

	err = multipleInsertBankAccount(db, []bankaccount.BankAccount{account1, account2, account3})
	if err != nil {
		log.Printf("Multiple insert failed with error %s", err)
		return
	} */

	router := gin.Default()
	router.GET("/customers", customerhandler.GetCustomers)
	router.GET("/customers/:customerId", customerhandler.GetCustomerById)
	router.PUT("/customer", customerhandler.CreateCustomer)

	router.GET("/accounts", accounthandler.GetAccounts)
	router.GET("/accounts/:accountId", accounthandler.GetAccountById)
	router.GET("/accounts/customer/:customerId/getallaccounts", accounthandler.GetAccountsByCustomerId)
	router.PUT("/account", accounthandler.CreateAccount)
	router.PUT("/accounts/:accountId/withdraw", accounthandler.Withdraw)
	router.PUT("/accounts/:accountId/deposit", accounthandler.Deposit)

	router.Run("localhost:8080")
}
