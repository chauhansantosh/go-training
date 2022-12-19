package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
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

	defer DB.Close()
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
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}
