package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlPassword string
var mysqlPort string
var mysqlDbName string
var mysqlDbHost string

var connectionString string

var mainConnection *sql.DB

func init() {
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlPort = os.Getenv("MYSQL_DB_PORT")
	mysqlDbName = os.Getenv("MYSQL_DB_NAME")
	mysqlDbHost = os.Getenv("MYSQL_DB_HOST")

	connectionString = fmt.Sprintf("root:%s@tcp(%s:%s)/%s", mysqlPassword, mysqlDbHost, mysqlPort, mysqlDbName)
	mainConnection = makeMainConnection()
}

func makeMainConnection() *sql.DB {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Println("Successfully connected to MySql database")

	return db
}

func GetMySqlConnection(ctx context.Context) *sql.Conn {
	newConnection, err := mainConnection.Conn(ctx)
	if err != nil {
		panic(err.Error())
	}

	return newConnection
}
