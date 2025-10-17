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
	// It’s only possible to call w.WriteHeader() once per response, and after the status code has been written it can’t be changed.
	// If we don’t call w.WriteHeader() explicitly, then the first call to w.Write() will automatically send a 200 OK status code to the user.
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
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
HTTP
====
- It’s the language (or set of rules) that web browsers and web servers use to communicate with each other.
- Every HTTP message has two sides
	- HTTP Request (from client → server) - What the client asks for
	- HTTP Response (from server → client) - What the server sends back

Each of HTTP request and HTTP response messages has two main parts:
	- Headers → metadata (information about the request or response)
		- Each header is a key-value pair.
		- Headers tell the server important details about the request.
	- Body → the actual content (like HTML, JSON, image, etc.)

ResponseWriter
==============
- The Client (browser) Sends an HTTP Request
- The Go HTTP server parses that raw request and automatically builds two objects
	- w → a ResponseWriter
	- r → a Request
- then it calls the handler - snippetView(w, r)
- handler function — it’s called by Go’s HTTP server whenever a request matches a route
- w http.ResponseWriter — The “Output Pipe”
	- w is an interface provided by Go’s HTTP package.
	- It represents the connection back to the client.
	- we don’t “return” anything from the handler function
	- instead, we write data into the ResponseWriter (w), which represents a network connection to the client.

“When does Go actually send that data back to the client?”
	- w.Write([]byte("Viewing..."))
	- Go buffers your data (it doesn’t necessarily flush immediately).
		=> A buffer is a small piece of memory used to store data temporarily before sending it somewhere — like a network, disk, or printer.
		=> Flush means: “Send whatever is currently in the buffer right now.”
	- writes the body bytes into an internal buffer connected to the client.
	- The response is fully sent right after the handler function ends — unless Go already flushed data earlier (e.g., if the buffer filled up).
*/
