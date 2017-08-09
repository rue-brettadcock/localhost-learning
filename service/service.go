package service

import (
	"flag"
	"net/http"

	"../logic"
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
	mux.HandleFunc(pat.Get("/homePage/:name"), handler.homePage)
	mux.HandleFunc(pat.Get("/loginPage/:username/:password"), handler.loginPage)
	http.ListenAndServe(bindTo, mux)
}
