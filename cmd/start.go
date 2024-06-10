package main

import (
	"github.com/koleaby4/muzz_rest_api/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/user/create", handlers.CreateUserHandler)
	http.HandleFunc("/login", handlers.LoginPost)
	http.ListenAndServe(":8080", nil)
}
