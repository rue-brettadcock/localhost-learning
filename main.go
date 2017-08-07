package main

import (
	"net/http"

	"./database"

	_ "github.com/go-sql-driver/mysql"
)

func signupPageHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	msg, err := database.SetUser(username, password)

	if err == true {
		http.Error(res, msg, 500)
	}
	if msg != "" {
		res.Write([]byte(msg))
	} else {
		http.Redirect(res, req, "/", 301)
		return
	}

}

func loginPageHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	err := database.LoginUser(username, password)

	if err != false {
		http.Redirect(res, req, "/login", 301)
		return
	}
	res.Write([]byte("Hello " + username))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html")
}

func main() {

	database.OpenDatabaseConnection()
	defer database.CloseDatabaseConnection()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPageHandler)
	http.HandleFunc("/signup", signupPageHandler)

	http.ListenAndServe(":8080", nil)
}
