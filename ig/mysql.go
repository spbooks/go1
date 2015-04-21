package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var globalMySQLDB *sql.DB

func init() {
	db, err := NewMySQLDB("root:root@tcp(127.0.0.1:3306)/gophr")
	if err != nil {
		panic(err)
	}
	globalMySQLDB = db
}

func NewMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
