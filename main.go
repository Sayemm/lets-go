package main

import (
	"fmt"
	"log"
	"net/http"
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
	w.Write([]byte("Viewing..."))
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
When sending a response Go will automatically set three system-generated headers
	- Date
	- Content-Length
	- Content-Type.
		- Go will attempt to set the correct one for you by content sniffing the response body with the http.DetectContentType() function.
		- If this function can’t guess the content type, Go will fall back to setting the header Content-Type: application/octet-stream instead.
		- it can’t distinguish JSON from plain text.
		- So, by default, JSON responses will be sent with a Content-Type: text/plain; charset=utf-8 header.


=> The Header() map (which you access through w.Header()) exists on the server side before the response is sent.
=> what does w.Header().Get("Cache-Control") do?
---It reads a header value from the response header map — on the server side, before the response has been written to the client.


=> The word canonicalization means - Converting data into a standard, consistent form.
   - When you’re using the Set(), Add(), Del(), Get() and Values() methods on the header map, the header name will always be canonicalized using the textproto.CanonicalMIMEHeaderKey() function.
   - This converts the first letter and any letter following a hyphen to upper case, and the rest of the letters to lowercase.
   - This has the practical implication that when calling these methods the header name is case-insensitive.


=> The Del() method doesn’t remove system-generated headers. To suppress these, you need
to access the underlying header map directly and set the value to nil. If you want to
suppress the Date header, for example, you need to write:
	-> w.Header()["Date"] = nil
*/
