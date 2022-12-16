package dbutil

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-training/bankingapp/bankaccount"
	"github.com/go-training/bankingapp/customer"
)

const (
	username = "root"
	password = "India@123"
	hostname = "127.0.0.1:3306"
	dbname   = "bankdb"
)

var DB *sql.DB

func ConnectDb() (*sql.DB, error) {
	DB, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	createDb(DB, "bankdb")

	DB.Close()
	DB, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}

	//Setting connection pool options
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(time.Minute * 5)

	//ping created database to verify the connection
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancelfunc()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return DB, nil
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func createDb(db *sql.DB, dbName string) {
	// Create database if not exists
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	response, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
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

func CreateTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS customer(customer_id int primary key, customer_name text, 
        customer_type text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
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
	res, err = db.ExecContext(ctx, query)
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

func InsertCustomer(db *sql.DB, c customer.Customer) (customerId int64, err error) {
	query := "INSERT INTO customer(customer_id, customer_name, customer_type) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
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

func InsertBankAccount(db *sql.DB, a bankaccount.BankAccount) error {
	query := "INSERT INTO bank_account(account_id, account_type, balance, customer_id) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, a.AccountId, a.AccountType, a.OpeningBalance, a.AccountHolder.CustomerId)
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
	accountId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted bank account", err)
		return err
	}
	log.Printf("Bank account with ID %d created", accountId)
	return nil
}

func MultipleInsertCutomer(db *sql.DB, customers []customer.Customer) error {
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
	stmt, err := db.PrepareContext(ctx, query)
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

func MultipleInsertBankAccount(db *sql.DB, bacnkAccounts []bankaccount.BankAccount) error {
	query := "INSERT INTO bank_account(account_id, account_type, balance, customer_id) VALUES"
	var inserts []string
	var params []interface{}
	for _, v := range bacnkAccounts {
		inserts = append(inserts, "(?, ?, ?, ?)")
		params = append(params, v.AccountId, v.AccountType, v.OpeningBalance, v.AccountHolder.CustomerId)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	log.Println("query is", query)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
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
