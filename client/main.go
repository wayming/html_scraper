package main

import (
	"context"
	"fmt"
	"log"

	// Assuming generated Go files from .proto
	HtmlScraper "github.com/wayming/HtmlScraper/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a connection to the gRPC server using grpc.NewClient
	conn, err := grpc.NewClient(
		"localhost:50051", // Server address
		grpc.WithTransportCredentials(insecure.NewCredentials()), // For insecure connection (no TLS)
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new HtmlScraper client
	client := HtmlScraper.NewHtmlScraperClient(conn)

	// Create a sample Request to send
	request := &HtmlScraper.Request{
		HtmlText: "Sample HTML Content",
		DataType: "json",
	}

	// Call ProcessData method from the HtmlScraper service
	response, err := client.ProcessData(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling ProcessData: %v", err)
	}

	// Print the response
	fmt.Println("Response Status:", response.GetStatus())
	fmt.Println("Response JSON Data:", response.GetJsonData())
}
