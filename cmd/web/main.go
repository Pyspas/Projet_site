package main

import (
	"fmt"
	"net/http"
)

const port = ":4200"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server is running on http://localhost:4200")
	http.ListenAndServe(port, nil)
}
