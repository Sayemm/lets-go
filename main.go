package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Listening to port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/*
Go’s servemux supports two different types of URL patterns:
	- fixed paths and
	- subtree paths.

fixed path
----------
- Fixed paths don’t end with a trailing slash
- fixed path patterns like these are only matched (and the corresponding handler called)
  when the request URL path exactly matches the fixed path.
- "/snippet/view" and "/snippet/create"

subtree paths
-------------
- Subtree path patterns are matched (and the corresponding handler called)
whenever the start of a request URL path matches the subtree path.
- "/" and "/static/"
- subtree paths as acting a bit like they have a wildcard at the end, like "/**" or "/static/**".
*/
