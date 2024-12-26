# Go CDN Implementation

This project is a simple Content Delivery Network (CDN) implementation in Go. It includes an origin server, edge servers with LRU caching, a load balancer using consistent hashing, and a client for testing.

## Project Structure

- `main.go`: Entry point of the application, handles command-line arguments and starts the appropriate component
- `origin_server.go`: Implementation of the origin server that stores and serves the original content
- `edge_server.go`: Implementation of the edge server with LRU cache for faster content delivery
- `load_balancer.go`: Implementation of the load balancer using consistent hashing to distribute requests
- `lru_cache.go`: LRU (Least Recently Used) cache implementation for efficient content caching
- `consistent_hash.go`: Consistent hashing implementation for load balancing
- `client.go`: Simple client for testing the CDN functionality

## Prerequisites

- Go 1.16 or higher

To check your Go version, run:
\`\`\`
go version
\`\`\`

## Building the Project

To build the project, follow these steps:

1. Ensure you have Go installed on your system. You can check your Go version by running:
   \`\`\`
   go version
   \`\`\`
   If Go is not installed, download and install it from the official Go website: https://golang.org/dl/

2. Clone the repository or download the project files to your local machine.

3. Open a terminal or command prompt and navigate to the project root directory:
   \`\`\`
   cd path/to/GoCDN
   \`\`\`

4. Initialize the Go module (if not already done):
   \`\`\`
   go mod init github.com/axshatj/GoCDN
   \`\`\`

5. Download any required dependencies (although this project doesn't have external dependencies, it's a good practice):
   \`\`\`
   go mod tidy
   \`\`\`

6. Build the project by running:
   \`\`\`
   go build -o GoCDN
   \`\`\`
   This command compiles the Go code and creates an executable named `GoCDN` (or `GoCDN.exe` on Windows) in your project directory.

7. Verify that the build was successful by checking the existence of the executable:
   \`\`\`
   ls GoCDN    # On Unix-like systems (Linux, macOS)
   dir GoCDN*  # On Windows
   \`\`\`

You should now have a `GoCDN` executable in your project directory, ready to run.

