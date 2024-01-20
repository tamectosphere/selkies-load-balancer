# Selkies Load Balancer

## Why?

https://medium.com/@TamEctosphere/why-developing-load-balancers-is-key-to-future-proofing-your-tech-skills-443436cdc0f5

## Overview

This project implements a simple load balancer (`selkies-lb`) and web server (`selkies-server`) using Go. The load balancer distributes incoming HTTP requests across multiple backend servers using a round-robin method. Each backend server can handle requests and report its health status.

This project was inspired by the [Build You Own Load Balancer](https://codingchallenges.fyi/challenges/challenge-load-balancer)

## Features

- **Load Balancer (`selkies-lb`)**: Distributes incoming requests across multiple backend servers.
- **Backend Server (`selkies-server`)**: Handles requests and responds to health checks.
- **Health Check**: Periodically checks the health of each backend server and removes unhealthy servers from the load balancing rotation.
- **Round-Robin Algorithm**: Routes incoming requests to backend servers in a round-robin fashion.

## Getting Started

These instructions will help you get the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.x or higher)

### Installation

1. Clone the repository:
```bash
git clone [URL to the repository]
```

2. Navigate to the cloned directory:
```bash
cd [repository name]
```

3. Build the project (this will create executables for the load balancer and backend server):
```bash
go build -o selkies cmd/main.go
```
## Configuration

Before running the application, you need to set up the environment variables:

- `LOAD_BALANCER_PORT`: The port on which the load balancer will listen. This should not be the same as the ports used by backend servers (8282, 8383, 8484). For example:
```bash
LOAD_BALANCER_PORT=8181
```
Ensure that this port is different from the ones used by the backend servers to avoid port conflicts.

### Running the Application

1. Start the Load Balancer:
To start the load balancer without specifying the health check interval (defaults to 3 seconds):
```bash
./selkies selkies-lb
```

To start the load balancer with a custom health check interval (for example, 5 seconds):
```bash
./selkies selkies-lb -health-check-interval 5
```


2. Start Backend Servers on different ports:
```bash
./selkies selkies-server -port 8282
./selkies selkies-server -port 8383
./selkies selkies-server -port 8484
```

## Usage

- The load balancer will start on port 8181 and route incoming requests to the backend servers.
- Access the load balancer through `http://localhost:8181`.
- The backend servers will respond to requests routed from the load balancer.
- Health checks are performed periodically to ensure backend servers are operational.
