package main

import (
	"net/http"

	"./database"
	_ "github.com/go-sql-driver/mysql"
)

type Logic struct {
	db database.MyDb
}

func (l *Logic) signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	msg, err := l.db.SetUser(username, password)

	if err != false {
		http.Error(res, msg, 500)
	}
	if msg != "" {
		res.Write([]byte(msg))
		return
	}
	http.Redirect(res, req, "/", 301)

}

func (l *Logic) loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	err := l.db.LoginUser(username, password)

	if err != false {
		http.Redirect(res, req, "/loginerror", 301)
	}

	res.Write([]byte("Hello " + username))

}

func (l *Logic) loginErrorPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/loginerror.html")
		return
	}
}

func (l *Logic) homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html")
}

func handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(res, req)
	})
}

func main() {

	l := Logic{db: database.New()}

	//http.HandleFunc("/", handle(homePage))
	http.HandleFunc("/", l.homePage)
	http.HandleFunc("/login", l.loginPage)
	http.HandleFunc("/loginerror", l.loginErrorPage)
	http.HandleFunc("/signup", l.signupPage)

	http.ListenAndServe(":8080", nil)
}
