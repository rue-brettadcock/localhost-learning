package service

import (
	"flag"
	"net/http"

	"github.com/rue-brettadcock/localhost-learning/logic"
	"goji.io"
	"goji.io/pat"
)

var (
	bindTo string
)

func init() {
	flag.StringVar(&bindTo, "listen", ":8080", "host:port to bind to")
}

// ListenAndServe initializes and starts the service
func ListenAndServe() {
	handler := Presentation{logic: logic.New()}
	mux := goji.NewMux()

	mux.HandleFunc(pat.Get("/signupPage/:username/:password"), handler.signupPage)
	mux.HandleFunc(pat.Get("/loginPage/:username/:password"), handler.loginPage)
	mux.HandleFunc(pat.Get("/:name"), handler.homePage)
	http.ListenAndServe(bindTo, mux)
}
