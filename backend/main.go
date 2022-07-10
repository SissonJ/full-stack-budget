package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
//	"github.com/volatiletech/authboss/v3"
)

type passwordRequest struct {
	User string `json:"user"`
	Password string `json:"password"`
}

func main() {
//	authboss_init()
	http.HandleFunc("/api/password", passwordHandler)

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	var decoded passwordRequest
	var response = "false"

	// Try to decode the request into the thumbnailRequest struct.
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if decoded.User == "sissonj" && decoded.Password == "passworD" {
		response = "true"
	}

	fmt.Printf("Got the following info: %s %s\n", decoded.User, decoded.Password)
	_, err = fmt.Fprintf(w, `{ "status": "%s" }`, response)
}

//func authboss_init() {
//	ab = authboss.New()
//}