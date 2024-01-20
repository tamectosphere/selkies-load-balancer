package selkieswebserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Start(port int) {

	defineRoute(port)

	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}

}

func defineRoute(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s", r.RemoteAddr)
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
		log.Printf("Host: %s", r.Host)

		responseString := fmt.Sprintf("Hello from a server: %d \n", port)
		io.WriteString(w, responseString)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK") // Simple response to indicate the server is up
	})
}
