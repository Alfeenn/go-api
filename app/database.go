package app

import (
	"database/sql"
	"fmt"

	"github.com/Alfeenn/api-go/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golangdb")
	helper.PanicIfErr(err)
	fmt.Println("Connected")
	return db
}
