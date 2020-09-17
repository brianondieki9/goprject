package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//struct for employee
type Employee struct {
	Id   int
	Name string
	City string
}

//function to connect to the local mysql driver
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbpass := "Green2013[]"
	dbName := "goblog"

	db, err := sql.Open(dbDriver, dbUser+":"+dbpass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db

}

func main() {
	dbConn()

}
