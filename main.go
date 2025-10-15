package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Println("Listening to port 4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/*
handlers
========
- handlers as being a bit like controllers.
- They’re responsible for executing the application logic and for writing HTTP response headers and bodies.

router (or servemux in Go terminology).
=======================================
- This stores a mapping between the URL patterns for your application and the corresponding handlers.
- Usually you have one servemux for your application containing all your routes.

server
======
- One of the great things about Go is that you can
establish a web server and listen for incoming requests as part of your application itself.
- You don’t need an external third-party server like Nginx, Apache or Caddy.
-  Each time the server receives a new HTTP request it will pass the request on to
the servemux and — in turn — the servemux will check the URL path and dispatch the
request to the matching handler.


=> The TCP network address that you pass to http.ListenAndServe() should be in the format "host:port"
- If you omit the host (like we did with ":4000") then the server will listen on all
your computer’s available network interfaces.
- Generally, you only need to specify a host in the address if your computer has multiple network interfaces
and you want to listen on just one of them.


=> A network interface is simply a connection point between your computer and a network.
   A network interface is a hardware or virtual device that lets your computer send and receive data on a network.
   It’s how your computer “talks” to other machines.
- house (computer) - different door (ip / interface) - different rooms (port)

http.ListenAndServe(":4000", mux)
---------------------------------
- Create a room (port 4000) behind every door that exists on this computer.
- So now, that same service (your Go web server) can be reached:
	through the Wi-Fi door (e.g. 192.168.1.12:4000)
	through the Ethernet door (e.g. 10.0.0.5:4000)
	through the loopback door (e.g. 127.0.0.1:4000)
- There is no single main door for port 4000 — each interface has its own version of port 4000.

Technically what happens Under the hood:
----------------------------------------
- The operating system (OS) listens for packets on each interface’s IP and port 4000.
- When data arrives at 192.168.1.12:4000, the OS delivers it to your Go program.
- The same happens for 127.0.0.1:4000, etc.
- So each interface is a separate “entry route” to the same listening service.
*/
