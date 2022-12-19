package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/nostra-api/helper"
)

func NewConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/nostra")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
