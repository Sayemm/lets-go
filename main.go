package main

import "fmt"

func main() {
	fmt.Println("Starting to learn go from the book")
}

/*
- module path as basically being a canonical name or identifier for the project.
- we want to pick a module path that is globally unique and unlikely to be used by anything else.
- In the Go community, a common convention is to base the module paths on a URL that we own.
- when there is a valid go.mod file in the root of the project directory, the project is a module.
- Setting up your project as a module has a number of advantages — including
making it much easier to manage third-party dependencies, avoid supply-chain attacks,
and ensure reproducible builds of your application in the future.

- go mod init snippetbox
*/
