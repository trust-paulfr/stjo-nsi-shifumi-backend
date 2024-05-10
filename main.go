package main

import (
	"log"
	"net/http"
	"shifumi_backend/src/handlers"
	"shifumi_backend/src/handlers/auth"
	"shifumi_backend/src/utils"
)

func main() {
	log.Println("Starting server on :8080")

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/register", auth.RegisterHandler)

	mux.HandleFunc("/user", handlers.UserHandler)

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ErrorResponseHandler(w, "Route not found")
	}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
