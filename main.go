package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	dbconnection "github.com/NGKlaure/Project-2-T33M-3/dbConnection"
)

type Users struct {
	Username      string
	Password      string
	Nametooshort  bool
	Namenotunique bool
	Pwtooshort    bool
	//
}

type LoginInfo struct {
	CurrentUser string
	Loggedin    bool
	Invalidname bool
	Invalidpw   bool
	//
}

type ViewInfo struct {
	Usr        []Users
	Singleuser Users
	Login      LoginInfo
}

var Signin = LoginInfo{}

func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	temp.Execute(response, Signin)
}

func registrationForm(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/registrationform.html")
	temp.Execute(response, nil)
}

func register(response http.ResponseWriter, request *http.Request) {
	db := dbconnection.DbConnection()
	defer db.Close()

	temp, _ := template.ParseFiles("templates/register.html")
	user := Users{}
	user.Username = request.FormValue("name")
	user.Password = request.FormValue("pw")
	// check if username is too short or is already taken or if password is too short
	if len(user.Username) < 3 {
		user.Nametooshort = true

	} else if uniqueName(user.Username) == false {
		user.Namenotunique = true
	} else if len(user.Password) < 3 {
		user.Pwtooshort = true
		// insert username and password into database if acceptable
	} else {
		statement := "INSERT INTO users (username, password)"
		statement += " VALUES ($1, $2);"
		_, err := db.Exec(statement, user.Username, user.Password)
		if err != nil {
			panic(err)
		}
	}

	temp.Execute(response, user)
}

//handler for loging
func login(response http.ResponseWriter, request *http.Request) {
	db := dbconnection.DbConnection()
	defer db.Close()

	temp, _ := template.ParseFiles("templates/login.html")

	user := Users{}
	view := ViewInfo{}
	login := LoginInfo{}
	// if not logged in then check if username and password are in database
	if !Signin.Loggedin {
		user.Username = request.FormValue("name")
		user.Password = request.FormValue("pw")
		if uniqueName(user.Username) == true {
			login.Invalidname = true
		} else if passwordMatches(user.Username, user.Password) == false {
			login.Invalidpw = true
		} else {
			Signin.CurrentUser = user.Username
			Signin.Loggedin = true

		}
	} else {

		user.Username = Signin.CurrentUser
	}

	view.Singleuser = user
	view.Login = login
	temp.Execute(response, view)
}

//handle logout listener
func logout(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("templates/index.html")
	Signin.Loggedin = false

	temp.Execute(response, Signin)
}

// check if the username does not already exist in database
func uniqueName(name string) bool {
	db := dbconnection.DbConnection()
	defer db.Close()

	rows, _ := db.Query("SELECT username FROM users")
	for rows.Next() {
		var username string
		rows.Scan(&username)
		if name == username {
			return false
		}
	}
	return true // name is not already in the db
}

//check if password match username
func passwordMatches(name string, password string) bool {
	db := dbconnection.DbConnection()
	defer db.Close()
	var pw string
	row := db.QueryRow("SELECT password FROM users WHERE username = $1", name)
	row.Scan(&pw)
	if password == pw {
		return true
	}
	return false
}

// methode to test if connected with database
func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected3!")
}

func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {

		var userName string
		var password string
		rows.Scan(&userName, &password)
		fmt.Println(userName, password)
	}
}

func main() {
	fmt.Println("hello")
	db := dbconnection.DbConnection()
	ping(db)
	getAll(db)

	http.HandleFunc("/", index)
	http.HandleFunc("/registrationform", registrationForm)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	Signin.Loggedin = false

	fmt.Println("Open Localhost:12345")
	http.ListenAndServe(":12345", nil)
}
