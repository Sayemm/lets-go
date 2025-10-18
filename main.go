package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Restricting the root url pattern

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		fmt.Println("hELLO")
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "ID = %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // net/http constants
		w.Header().Set("Allow", "POST")

		// calls the w.WriteHeader() and w.Write() methods behind the scenes
		// http.Error(w, "Method Not Allowed", 405)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=3111111")

	// Retrieve the first value for the "Cache-Control" header.
	ans := w.Header().Get("Cache-Control")
	fmt.Println("====>", ans)

	// Retrieve a slice of all values for the "Cache-Control" header.
	ansArr := w.Header().Values("Cache-Control")
	fmt.Println("====>", ansArr)

	// Delete all values for the "Cache-Control" header.
	w.Header().Del("Cache-Control")

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
- fmt.Fprintf() function you’ll notice that it takes an io.Writer as the first parameter…
- but we passed it our http.ResponseWriter object instead — and it worked fine.
*** We’re able to do this because the io.Writer type is an interface, and the
    http.ResponseWriter object satisfies the interface because it has a w.Write() method.
*/
