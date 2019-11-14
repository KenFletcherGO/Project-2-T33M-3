package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"

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

	//Runs the http service in a goroutine as to not freeze the code.
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", loginpage)
	http.HandleFunc("/main", mainpage)
	http.ListenAndServe(":8080", nil)

	//The following code is used to initialize the users database containing usernames and passwords.
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, dbport, dbusername, dbpassword, dbname)
	db, err := sql.Open("postgres", datasource)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	/*Use the following code to test if the database is working for you or not. We will remove the test later.
	rows, _ := db.Query("SELECT * FROM USERS")

	for rows.Next() {
		var id string
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}*/
}

func loginpage(response http.ResponseWriter, request *http.Request) {
	login, _ := template.ParseFiles("webpage/loginpage.html")
	login.Execute(response, nil)
}

func mainpage(response http.ResponseWriter, request *http.Request) {
	mainpage, _ := template.ParseFiles("webpage/mainpage.html")
	mainpage.Execute(response, nil)
}

func logincss(response http.ResponseWriter, request *http.Request) {
	login, _ := template.ParseFiles("webpage/loginpage.css")
	login.Execute(response, nil)
}

//CreateFolder creates a new folder inside the server.
func CreateFolder(n string) {
	exec.Command("mkdir " + n)
}

//DeleteFolder will delete a folder inside the server.
func DeleteFolder(n string) {
	exec.Command("rm -rf " + n)
}

//NewFile will create a new file inside the current folder.
func NewFile(n string) {
	exec.Command("touch " + n)
}

//DeleteFile will delete a file inside the current folder.
func DeleteFile(n string) {
	exec.Command("rm " + n)
}
