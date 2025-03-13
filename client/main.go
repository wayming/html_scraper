package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// Assuming generated Go files from .proto
	HtmlScraper "github.com/wayming/HtmlScraper/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a connection to the gRPC server using grpc.NewClient
	var host = os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}
	conn, err := grpc.NewClient(
		host+":50051", // Server address
		grpc.WithTransportCredentials(insecure.NewCredentials()), // For insecure connection (no TLS)
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a new HtmlScraper client
	client := HtmlScraper.NewHtmlScraperClient(conn)

	// Read a html file
	content, err := os.ReadFile("input/income_statement.html")
	if err != nil {
		log.Fatal(err)
	}

	// Create a sample Request to send
	request := &HtmlScraper.Request{
		HtmlText: string(content),
		PageType: "income_statement",
	}

	// Call ProcessPage method from the HtmlScraper service
	response, err := client.ProcessPage(context.Background(), request)
	if err != nil {
		log.Fatalf("Error calling ProcessPage: %v", err)
	}

	// Unmarshal the JSON string
	var obj []map[string]interface{}
	err = json.Unmarshal([]byte(string(response.GetJsonData())), &obj)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal the object back into a pretty-printed JSON string
	prettyJSON, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Print the response
	fmt.Println("Response Status:", response.GetStatus())
	fmt.Println("Response JSON Data:", string(prettyJSON))
}
