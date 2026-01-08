package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	port := ":8000"

	log.Println("Starting server on port", port)
	http.HandleFunc("/", helloWorldHandler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error: %s", err)
	}
	log.Println("Finished execution")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string
	}{
		Message: "Hello world",
	}

	json.NewEncoder(w).Encode(response)
}
