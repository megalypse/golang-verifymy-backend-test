package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlPassword string
var mysqlPort string
var mysqlDbName string

var connectionString string

func init() {
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlPort = os.Getenv("MYSQL_DB_PORT")
	mysqlDbName = os.Getenv("MYSQL_DB_NAME")

	connectionString = fmt.Sprintf("root:%s@tcp(127.0.0.1:%s)/%s", mysqlPassword, mysqlPort, mysqlDbName)
}

func GetMySqlConnection() *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!")

	return db
}
