package main

import (
	grpc_client "orchestrator/internal/grpc_clinet"
)


func main() {
	// send a request and print the response
	grpc_client.Client()
}