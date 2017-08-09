package main

import "./service"

func main() {
	// toggle these two and the import statement to switch
	// between html and cURL views
	// html.Start()
	service.ListenAndServe()
}
