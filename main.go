package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	GetMySqlConnection()
}

func GetMySqlConnection() {
	db, err := sql.Open("mysql", "root:mock-extremely-secure-password@tcp(127.0.0.1:3301)/verifymy_backend_test")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("New MySql connection created")
}
