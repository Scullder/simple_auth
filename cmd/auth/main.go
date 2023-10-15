package main

import (
	"auth/internal/handlers"
	"auth/internal/middlewares"
	"auth/pkg/server"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT string = "localhost:8181"

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.Handle("/hello", middlewares.AuthMiddleware(http.HandlerFunc(handlers.Hello)))

	server.RunServer(PORT, mux)
}
