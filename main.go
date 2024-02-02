package main

import (
	"net/http"
	"os"
)

// handlers in Go defining controllers
// so it carry out application logic and writing response headers and bodies
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world!</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// mux equal to the router in general terms
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
