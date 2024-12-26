package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	componentType := flag.String("type", "", "Type of component to run (origin, edge, loadbalancer, client)")
	port := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

	switch *componentType {
	case "origin":
		runOriginServer(*port)
	case "edge":
		runEdgeServer(*port)
	case "loadbalancer":
		runLoadBalancer(*port)
	case "client":
		runClient()
	default:
		log.Fatalf("Invalid component type. Use -type=origin|edge|loadbalancer|client")
	}
}

func runOriginServer(port int) {
	server := NewOriginServer()
	log.Printf("Starting Origin Server on port %d", port)
	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%d", port)))
}

func runEdgeServer(port int) {
	server := NewEdgeServer("http://localhost:8080") // Assuming origin server is on 8080
	log.Printf("Starting Edge Server on port %d", port)
	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%d", port)))
}

func runLoadBalancer(port int) {
	edgeServers := []string{"http://localhost:8081", "http://localhost:8082"}
	server := NewLoadBalancer(edgeServers)
	log.Printf("Starting Load Balancer on port %d", port)
	log.Fatal(server.ListenAndServe(fmt.Sprintf(":%d", port)))
}

func runClient() {
	client := NewClient("http://localhost:8090") // Assuming load balancer is on 8090
	client.Run()
}

