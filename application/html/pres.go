package html

import (
	"net/http"

	"../logic"
)

type Presentation struct {
	logic *logic.Logic
}

func Start() {
	p := Presentation{logic: logic.New()}

	http.HandleFunc("/login", p.loginPage)
	http.HandleFunc("/loginerror", p.loginErrorPage)
	http.HandleFunc("/signup", p.signupPage)
	http.HandleFunc("/", p.homePage)
	http.ListenAndServe(":8080", nil)
}

func (p *Presentation) signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	msg, err := p.logic.Register(username, password)

	if err != false {
		http.Error(res, msg, 500)
	} else if msg != "" {
		res.Write([]byte(msg))
		return
	}
	http.Redirect(res, req, "/", 301)

}

func (p *Presentation) loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	err := p.logic.SignIn(username, password)

	if err != false {
		http.Redirect(res, req, "/loginerror", 301)
	}

	res.Write([]byte("Hello " + username))

}

func (p *Presentation) loginErrorPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "html/loginerror.html")
		return
	}
}

func (p *Presentation) homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "html/index.html")
}
