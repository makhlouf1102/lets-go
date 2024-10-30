package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/login", LoginPage)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/index.html")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/login.html")
}
