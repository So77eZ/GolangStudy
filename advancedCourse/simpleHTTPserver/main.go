package main

import (
	"fmt"
	"net/http"
)

func main() {
	// url :=
	router := http.NewServeMux()
	NewHellowHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listennig on port: 8081")
	server.ListenAndServe()
}
