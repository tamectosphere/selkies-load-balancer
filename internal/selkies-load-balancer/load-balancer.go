package selkiesloadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"
)

var currentServerIndex int = 0

type BackendServer struct {
	URL     string
	IsAlive bool
}

func (s *BackendServer) CheckHealth() {
	resp, err := http.Get(s.URL + "/health")
	if err != nil || resp.StatusCode != 200 {
		s.IsAlive = false
		log.Printf("%s is dead", s.URL)
		return
	}
	log.Printf("%s is alive", s.URL)
	s.IsAlive = true
}

func Start(healthCheckInterval int) {

	backends := getBackendServers()

	periodicallyHealthCheck(backends, healthCheckInterval)

	defineRoute(backends)

	port := getPort()
	log.Printf("Loadbalancer: Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func getBackendServers() []*BackendServer {
	return []*BackendServer{
		{URL: "http://localhost:8282", IsAlive: true},
		{URL: "http://localhost:8383", IsAlive: true},
		{URL: "http://localhost:8484", IsAlive: true},
	}
}

func periodicallyHealthCheck(backends []*BackendServer, numSecond int) {
	go func() {
		for {
			for _, backend := range backends {
				backend.CheckHealth()
				time.Sleep(time.Duration(numSecond) * time.Second)
			}
		}
	}()
}

func defineRoute(backends []*BackendServer) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		backendServer := getNextServer(backends)
		if backendServer == nil {
			http.Error(w, "No available servers", http.StatusServiceUnavailable)
			return
		}

		// Parse the backend server URL
		url, err := url.Parse(backendServer.URL)
		if err != nil {
			log.Printf("Failed to parse backend URL: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)

		log.Printf("Forwarding request to %s", backendServer.URL)
	})
}

func getPort() string {
	port := os.Getenv("LOAD_BALANCER_PORT")
	if port == "" {
		log.Fatalf("LOAD_BALANCER_PORT is missing")
	} else {
		if _, err := strconv.Atoi(port); err != nil {
			log.Fatalf("Invalid PORT environment variable: %v", err)
		}
	}

	return port
}

func getNextServer(backends []*BackendServer) *BackendServer {
	// Locking mechanisms like mutex should be used here if you're handling concurrent requests

	startingIndex := currentServerIndex
	for {
		currentServer := backends[currentServerIndex]
		currentServerIndex = (currentServerIndex + 1) % len(backends)

		currentServer.CheckHealth()

		if currentServer.IsAlive {
			return currentServer
		}

		// If we've looped through all servers and none are alive, return nil or handle the case appropriately
		if currentServerIndex == startingIndex {
			return nil
		}
	}
}
