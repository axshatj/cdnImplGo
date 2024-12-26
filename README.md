# Go CDN Implementation

This project is a simple Content Delivery Network (CDN) implementation in Go. It includes an origin server, edge servers with LRU caching, a load balancer using consistent hashing, and a client for testing.

## Project Structure

- `main.go`: Entry point of the application
- `origin_server.go`: Implementation of the origin server
- `edge_server.go`: Implementation of the edge server with LRU cache
- `load_balancer.go`: Implementation of the load balancer with consistent hashing
- `lru_cache.go`: LRU cache implementation
- `consistent_hash.go`: Consistent hashing implementation
- `client.go`: Simple client for testing the CDN

## Prerequisites

- Go 1.16 or higher

## Building the Project

To build the project, run the following command in the project root directory:

