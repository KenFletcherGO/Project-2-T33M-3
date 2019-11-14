package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os/exec"

	dbconnection "Project-2-T33M-3-nadine/dbConnection"
)

//localstruct
type localstruct struct {
	FILES []string
}

//Users a
type Users struct {
	Username      string
	Password      string
	Nametooshort  bool
	Namenotunique bool
	Pwtooshort    bool
	//
}

//LoginInfo a
type LoginInfo struct {
	CurrentUser string
	Loggedin    bool
	Invalidname bool
	Invalidpw   bool
	//
}

//ViewInfo a
type ViewInfo struct {
	Usr        []Users
	Singleuser Users
	Login      LoginInfo
}

var userName, password, enter string

//fmt.Scanln(&localuser)
//
//fmt.Scanln(&hostname)
const hostname = "192.168.1.33"
const localuser = "garner"

//Signin a
var Signin = LoginInfo{}

func main() {

	db := dbconnection.DbConnection()
	ping(db)
	getAll(db)

	/*
		//These are variables for HOME/local computer:
		fmt.Printf("Enter your user name(for local computer):  ")
		fmt.Scanln(&localuser)
		//These are login variables for the REMOTE/target computer.
		fmt.Printf("Enter a hostname(IP) of target remote computer:  ")
		fmt.Scanln(&hostname)
		fmt.Printf("Enter a username of this target remote computer:  ")
		fmt.Scanln(&userName)
		fmt.Printf("Enter a password for this target remote computer:  ")
		fmt.Scanln(&password)
		enter = userName + "@" + hostname */

	http.HandleFunc("/", index)
	http.HandleFunc("/remotefiles.html", remotefiles)
	http.HandleFunc("/localfiles.html", localfiles)
	http.HandleFunc("/downloader", downloader)
	http.HandleFunc("/uploader", uploader)
	http.HandleFunc("/registrationform", registrationForm)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	Signin.Loggedin = false

	fmt.Println(" Open browser to localhost:7004")
	http.ListenAndServe(":7004", nil)
}

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	user.Username = request.FormValue("uname")
	user.Password = request.FormValue("pwd")
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
		user.Username = request.FormValue("uname")
		user.Password = request.FormValue("pwd")
		if uniqueName(user.Username) == true {
			login.Invalidname = true
		} else if passwordMatches(user.Username, user.Password) == false {
			login.Invalidpw = true
		} else {
			Signin.CurrentUser = user.Username
			Signin.Loggedin = true

			userName = user.Username
			enter = userName + "@" + hostname

		}
	} else {

		user.Username = Signin.CurrentUser

	}

	view.Singleuser = user
	view.Login = login

	//HARDCODE
	//connect to this socket
	conn, _ := net.Dial("tcp", "192.168.1.33:8081")

	// send to socket
	fmt.Fprintf(conn, user.Username+"\n")

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

//remotefiles displays remote computer files to html page.
func remotefiles(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/remotefiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	enter = userName + "@" + hostname

	remote3 := exec.Command("ssh", enter, "mkdir", "/home/"+userName+"/servercatchbox")
	remote3.Run()

	remote, err := exec.Command("ssh", enter, "ls", "/home/"+userName+"/servercatchbox", ">", "file1", ";", "cat", "file1").Output()
	g := localstruct{FILES: make([]string, 1)}
	length := 0
	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}
	temp.Execute(response, g)
}

//localfiles displays host computer files to html page.
func localfiles(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/localfiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	remote3 := exec.Command("mkdir", "/home/"+localuser+"/servercatchbox")
	remote3.Run()

	remote, err := exec.Command("ls", "/home/"+localuser+"/servercatchbox", ">", "file1", ";", "cat", "file1").Output()
	g := localstruct{FILES: make([]string, 1)}
	length := 0

	if err != nil {

		fmt.Println(err)
		fmt.Print("some error")
	}

	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}
	temp.Execute(response, g)
}

//prototype UPLOAD
func uploader(response http.ResponseWriter, request *http.Request) {
	var upload1 = request.FormValue("upload1")
	fmt.Println("The file to upload is: " + upload1)
	enter = userName + "@" + hostname
	remote2 := exec.Command("scp", "/home/"+localuser+"/servercatchbox/"+upload1, enter+":/home/"+userName+"/servercatchbox")
	remote2.Run()

	temp, _ := template.ParseFiles("html/remotefiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	remote, err := exec.Command("ssh", enter, "ls", "/home/"+userName+"/servercatchbox", ">", "file1", ";", "cat", "file1").Output()

	g := localstruct{FILES: make([]string, 1)}
	length := 0
	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}
	remote2.Run()
	temp.Execute(response, g)
}

//downloader
func downloader(response http.ResponseWriter, request *http.Request) {
	var download1 = request.FormValue("download1")
	fmt.Println("The file to download is: " + download1)
	enter = userName + "@" + hostname
	remote2 := exec.Command("scp", enter+":/home/"+userName+"/servercatchbox/"+download1, "/home/"+localuser+"/servercatchbox")
	remote2.Run()
	temp, _ := template.ParseFiles("html/localfiles.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	remote, err := exec.Command("ls", "/home/"+localuser+"/servercatchbox", ">", "file1", ";", "cat", "file1").Output()

	g := localstruct{FILES: make([]string, 1)}
	length := 0
	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(remote); l = l + 1 {
		if remote[l] != 10 {
			g.FILES[length] = g.FILES[length] + string(remote[l])
		} else {
			g.FILES = append(g.FILES, "\n")
			length = length + 1
		}
	}

	temp.Execute(response, g)
}
