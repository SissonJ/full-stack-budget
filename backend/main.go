package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/volatiletech/authboss/v3"
)

func main() {
	authboss_init()
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func authboss_init() {
	ab = authboss.New()
}