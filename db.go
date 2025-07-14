package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() error {
	connStr := "user=postgres password=secret dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping db: %w", err)
	}
	return createUsersTable()
}

func createUsersTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`
	_, err := db.Exec(query)
	return err
}
