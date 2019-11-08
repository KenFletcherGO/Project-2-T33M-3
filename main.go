package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host       = "localhost"
	dbport     = 5432
	dbusername = "postgres"
	dbpassword = "postgres"
	dbname     = "postgres"
)

func main() {
	//The following code is used to initialize the users database containing usernames and passwords.
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, dbport, dbusername, dbpassword, dbname)
	db, err := sql.Open("postgres", datasource)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	//Use the following code to test if the database is working for you or not. We will remove the test later.
	rows, _ := db.Query("SELECT * FROM USERS")

	for rows.Next() {
		var id string
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
