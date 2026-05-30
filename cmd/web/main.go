package main

import (
	"fmt"
	"net/http"
<<<<<<< HEAD

	"github.com/Pyspas/Projet_site/internal/handles"
=======
>>>>>>> 62a6b09e83a75b01378ebeee091e7b6c499e1245
)

const port = ":4200"

func main() {

<<<<<<< HEAD
	http.HandleFunc("/", handles.Home)
	http.HandleFunc("/about", handles.About)
=======
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
>>>>>>> 62a6b09e83a75b01378ebeee091e7b6c499e1245

	fmt.Println("Server is running on http://localhost:4200")
	http.ListenAndServe(port, nil)
}
