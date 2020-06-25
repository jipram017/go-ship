package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewConnection() (*sqlx.DB, error) {
	//host := os.Getenv("DB_HOST")
	host := "127.0.0.1"
	//user := os.Getenv("DB_USER")
	user := "postgres"
	//dbName := os.Getenv("DB_NAME")
	dbName := "shippy"
	//password := os.Getenv("DB_PASSWORD")
	password := "postgres"
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
