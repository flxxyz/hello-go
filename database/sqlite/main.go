package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//sql.Register("sqlite3", &sqlite3.SQLiteDriver{})
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}

const path = "/Users/flxxyz/.dbeaver4/.metadata/sample-database-sqlite-1/Chinook.db"

func main()  {
	db, err := sql.Open("sqlite3", path)
	checkErr(err)

	rows, err := db.Query("select CustomerId, FirstName, LastName from customer")
	checkErr(err)

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		err = rows.Scan(&id, &firstName, &lastName)
		checkErr(err)

		fmt.Printf("id=%d, firstName=%s, lastName=%s", id, firstName, lastName)
		fmt.Println()
	}

	db.Close()
}
