package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

/* func RunServer(port string, routes map[string]func(http.ResponseWriter, *http.Request)) {
	http.Handle("/", http.FileServer(http.Dir("web/static")))

	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}

	fmt.Printf("Server is running on port: %s\n", port)

	err := http.ListenAndServe(port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err == nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

} */

func RunServer(port string, handler http.Handler) {
	http.Handle("/", handler)

	fmt.Printf("Server is running on port: %s\n", port)
	err := http.ListenAndServe(port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err == nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
