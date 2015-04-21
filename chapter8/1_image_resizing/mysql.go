package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var globalMySQLDB *sql.DB

func NewMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
