/**
 * Copyright 2024 (c) Paul DESPLANQUES
 *
 * This file is part of shifumi_backend.
 * This main function is the entry point of the application.
*/
 
package main

import (
	"net/http"
	"shifumi_backend/src/handlers"
	"shifumi_backend/src/handlers/auth"
	"shifumi_backend/src/handlers/game"
	"shifumi_backend/src/utils"
)

func main() {
	mux := http.NewServeMux()

	/* HTTP EndPoints */
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/register", auth.RegisterHandler)

	/* Use internally in python game */
	mux.HandleFunc("/internal/game/play", game.PlayHandler)
	mux.HandleFunc("/internal/game/finish", game.FinishHandler)

	mux.HandleFunc("/pub", game.PubHandler)
	mux.HandleFunc("/user", handlers.UserHandler)

	/* Default EndPoints */
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ErrorResponseHandler(w, "Route not found")
	}))

	/* Connect to DB in GoRoutine */
	go utils.ConnectDB()

	/* Start HTTP Server */
	err := http.ListenAndServe(":12234", mux)
	if err != nil {
		panic(err)
	}
}
