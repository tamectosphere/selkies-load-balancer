package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received connection from %s", r.RemoteAddr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("No PORT environment variable detected. Defaulting to port 8080")
		port = "8080"
	} else {
		if _, err := strconv.Atoi(port); err != nil {
			log.Fatalf("Invalid PORT environment variable: %v", err)
		}
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
