package util

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	bankaccount "github.com/chauhansantosh/go-training/bankingapp/model/account"
	customer "github.com/chauhansantosh/go-training/bankingapp/model/customer"
	mysql "github.com/chauhansantosh/go-training/bankingapp/mysqldb"
)

const (
	GETQUERY = `SELECT account_id, customer_id, account_type, balance, 
	created_at, updated_at, IFNULL(account_pan, ''), is_active, is_locked, IFNULL(lock_period_fd, 0), 
	locked_until, IFNULL(penalty_fd, 0)  
	FROM bankdb.bank_account `
)

func CreateDb(dbName string) {
	// Create database if not exists
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	response, err := mysql.DB.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	rowsAffected, err := response.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", rowsAffected)
}

func CreateTables() error {
	query := `CREATE TABLE IF NOT EXISTS customer(customer_id int primary key, customer_name text, 
        customer_type text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := mysql.DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)

	query = `CREATE TABLE IF NOT EXISTS bank_account(account_id int primary key, customer_id int, balance float, 
            account_type text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP, 
            CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customer(customer_id))`
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err = mysql.DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating table", err)
		return err
	}
	rows, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

func InsertCustomer(c customer.Customer) (customerId int64, err error) {
	query := "INSERT INTO customer(customer_id, customer_name, customer_type) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, c.CustomerId, c.CustomerName, c.CustomerType)
	if err != nil {
		log.Printf("Error %s when inserting row into customer table", err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	log.Printf("%d customers created ", rows)
	custId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted customer", err)
		return 0, err
	}
	log.Printf("Customer with ID %d created", custId)
	return custId, nil
}

func InsertBankAccount(a bankaccount.BankAccount) error {
	query := `INSERT INTO bank_account(
	account_id, account_type, balance, customer_id, account_pan, is_locked, locked_until, lock_period_fd, penalty_fd, odallowed, odamount) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	var penaltyFd float32
	if a.AccountType == "FIXED" {
		lockedUntil := time.Now().AddDate(a.LockPeriodFd, 0, 0)
		a.IsLocked = true
		a.LockedUntil = &lockedUntil
		if a.PenaltyFd > 0 {
			penaltyFd = a.PenaltyFd
		} else {
			penaltyFd = 0.1
		}
	}

	if a.AccountType == "SAVINGS" && (a.OpeningBalance) > 10000000 {
		log.Printf("Error: Savings account cannot have more than 10 millions")
		return errors.New("Balance limit reached. Can not have more than 10 millions in SAVINGS account.")
	}

	if a.AccountType == "CURRENT" && a.OverdraftAllowed && a.OdAmount <= 0 {
		log.Printf("Error: Overdraft amount not specified")
		return errors.New("overdraft amount is mandatory when overdraft allowed is set to true")
	}
	if a.AccountType == "CURRENT" && !a.OverdraftAllowed && a.OdAmount > 0 {
		a.OdAmount = 0
		log.Printf("Overdraft amount is set to 0 as overdraft flag is set to false")
	}

	res, err := stmt.ExecContext(ctx, a.AccountId, a.AccountType, a.OpeningBalance, a.CustomerId,
		a.AccountPan, a.IsLocked, a.LockedUntil, a.LockPeriodFd, penaltyFd, a.OverdraftAllowed, a.OdAmount)
	if err != nil {
		log.Printf("Error %s when inserting row into bank_account table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d bank accounts created ", rows)
	return nil
}

func MultipleInsertCutomer(customers []customer.Customer) error {
	query := "INSERT INTO customer(customer_id, customer_name, customer_type) VALUES "
	var inserts []string
	var params []interface{}
	for _, v := range customers {
		inserts = append(inserts, "(?, ?, ?)")
		params = append(params, v.CustomerId, v.CustomerName, v.CustomerType)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	log.Println("query is", query)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into customer table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d customers created simulatneously", rows)
	return nil
}

func MultipleInsertBankAccount(bacnkAccounts []bankaccount.BankAccount) error {
	query := "INSERT INTO bank_account(account_id, account_type, balance, customer_id) VALUES"
	var inserts []string
	var params []interface{}
	for _, v := range bacnkAccounts {
		inserts = append(inserts, "(?, ?, ?, ?)")
		params = append(params, v.AccountId, v.AccountType, v.OpeningBalance, v.CustomerId)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	log.Println("query is", query)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := mysql.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		log.Printf("Error %s when inserting row into bank_account table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d bank accounts created simulatneously", rows)
	return nil
}
