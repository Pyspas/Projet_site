package main

import (
	"fmt"
	"net/http"

	"github.com/Pyspas/Projet_site/internal/handles"
)

const port = ":4200"

func main() {

	http.HandleFunc("/", handles.Home)
	http.HandleFunc("/about", handles.About)

	fmt.Println("Server is running on http://localhost:4200")
	http.ListenAndServe(port, nil)
}
