package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/faridlan/nostra-api/helper"
)

func NewConnection() *sql.DB {

	host := os.Getenv("host")
	port := os.Getenv("port")

	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s:%s)/nostra", host, port))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
