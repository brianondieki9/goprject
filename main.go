package main

import (
	"database/sql"
	"fmt"
<<<<<<< HEAD
	"net/http"
	"text/template"
=======
>>>>>>> 5f36de8b784b8aa428221a561764e0f490ae64b7

	_ "github.com/go-sql-driver/mysql"
)

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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var city, name string

		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err := selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.City = city
		emp.Name = name
	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func main() {
	db := dbConn()
<<<<<<< HEAD

	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}

=======

	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}

>>>>>>> 5f36de8b784b8aa428221a561764e0f490ae64b7
	fmt.Println(res)
}
