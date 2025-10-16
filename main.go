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

Servemux features and quirks
============================
- In Go’s servemux, longer URL patterns always take precedence over shorter ones.
- So, if a servemux contains multiple patterns which match a request, it will always dispatch the
  request to the handler corresponding to the longest pattern.

- Request URL paths are automatically sanitized. If the request path contains any . or ..
elements or repeated slashes, the user will automatically be redirected to an equivalent clean URL.

- If a subtree path has been registered and a request is received for that subtree path
without a trailing slash, then the user will automatically be sent a
301 Permanent Redirect to the subtree path with the slash added. For example, if you
have registered the subtree path /foo/, then any request to /foo will be redirected to
/foo/.


-> It’s possible to include host names in your URL patterns. This can be useful when you want
to redirect all HTTP requests to a canonical URL, or if your application is acting as the back
end for multiple sites or services.
*/
