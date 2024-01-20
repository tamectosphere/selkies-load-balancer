package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"

	selkiesloadbalancer "load-balancer/internal/selkies-load-balancer"
	selkieswebserver "load-balancer/internal/selkies-webserver"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide a command (selkies-lb or selkies-server)")
	}

	command := os.Args[1]
	os.Args = os.Args[1:] // Shift the arguments

	port, healthCheckInterval := getFlagOptions()

	switch command {
	case "selkies-lb":
		validateHealthCheckInterval(healthCheckInterval)
		selkiesloadbalancer.Start(healthCheckInterval)
	case "selkies-server":
		validatePort(port)
		selkieswebserver.Start(port)
	default:
		log.Fatal("Unknown command. Use 'selkies-lb' or 'selkies-server'")
	}
}

func getFlagOptions() (int, int) {
	var port int
	var healthCheckInterval int
	flag.IntVar(&port, "port", 0, "Choose a port: 8282, 8383, 8484")
	flag.IntVar(&healthCheckInterval, "health-check-interval", 3, "Input health check interval")
	flag.Parse()

	return port, healthCheckInterval
}

func validatePort(port int) {
	if port == 0 {
		log.Fatal("No port specified. Please provide a port: 8282, 8383, 8484")
	}

	allowedPorts := map[int]bool{
		8282: true,
		8383: true,
		8484: true,
	}

	if _, ok := allowedPorts[port]; !ok {
		log.Fatalf("Invalid port: %d. Allowed ports are 8282, 8383, 8484", port)
	}
}

func validateHealthCheckInterval(healthCheckInterval int) {
	if healthCheckInterval <= 0 {
		log.Fatalf("Invalid health check interval: must be greater than 0")
	}
}
