package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// create a file server to serve the assests from the given folder
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// attach handler functions
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return secureHeaders(mux)
}
