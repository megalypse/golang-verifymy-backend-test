package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlPassword string
var mysqlPort string
var mysqlDbName string
var mysqlDbHost string

var mainConnection *sql.DB

func init() {
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlPort = os.Getenv("MYSQL_DB_PORT")
	mysqlDbName = os.Getenv("MYSQL_DB_NAME")
	mysqlDbHost = os.Getenv("MYSQL_DB_HOST")

	connectionString := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?parseTime=true", mysqlPassword, mysqlDbHost, mysqlPort, mysqlDbName)
	mainConnection = makeMainConnection(connectionString)
}

func makeMainConnection(connectionStr string) *sql.DB {
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		panic(err.Error())
	}

	const healthcheckAmt int = 5
	for i := 0; i < healthcheckAmt; i++ {
		err = db.Ping()
		if err == nil {
			break
		} else if err != nil && i < healthcheckAmt {
			log.Printf("Failed connecting to database (%d/%d)", i+1, healthcheckAmt)
		} else {
			panic(err.Error())
		}

		time.Sleep(time.Second * 5)
	}

	log.Println("Successfully connected to MySql database")

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetConnMaxIdleTime(time.Second * 5)

	return db
}

func GetMySqlConnection(ctx context.Context) *sql.Conn {
	newConnection, err := mainConnection.Conn(ctx)
	if err != nil {
		panic(err.Error())
	}

	return newConnection
}
