package service

import (
	"fmt"
	"net/http"

	"goji.io/pat"

	"../logic"
)

type Presentation struct {
	logic *logic.Logic
}

type httpHandler interface {
	echo(req *http.Request) (int, string)
}

func echo(in string) string {
	if in == "" {
		in = "Empty URL path"
	}
	return fmt.Sprintf("Specified path: %s!\n", in)
}

func (p *Presentation) signupPage(res http.ResponseWriter, req *http.Request) {

	username := pat.Param(req, "username")
	password := pat.Param(req, "password")

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	output, _ := p.logic.Register(username, password)
	fmt.Fprintf(res, "%s\n", output)

}

func (p *Presentation) loginPage(res http.ResponseWriter, req *http.Request) {

	username := pat.Param(req, "username")
	password := pat.Param(req, "password")

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	err := p.logic.SignIn(username, password)
	if err != false {
		fmt.Fprint(res, "Invalid Credentials\n")
	} else {
		fmt.Fprintf(res, "Login Successful: Hello %s\n", username)
	}

}

func (p *Presentation) homePage(res http.ResponseWriter, req *http.Request) {

	name := pat.Param(req, "name")

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	output := echo(name)
	fmt.Fprintf(res, "%s\n", output)
}
