package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// CLI flags
	url := flag.String("url", "", "URL to fetch")
	requestType := flag.String("request", "GET", "HTTP method: GET, POST, etc.") 
	data := flag.String("data", "", "Data to send with POST (raw JSON supported)")
	verbose := flag.Bool("verbose", false, "Show verbose output")

	flag.Parse()

	// Validate URL
	if *url == "" {
		fmt.Println("Usage: go run main.go -url https://example.com [-request GET/POST] [-data '{\"key\":\"value\"}'] [-verbose]")
		return
	}

	// Prepare request body
	var bodyReader io.Reader
	if *data != "" {
		bodyReader = strings.NewReader(*data)
	}

	// Create HTTP request
	req, err := http.NewRequest(strings.ToUpper(*requestType), *url, bodyReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating request failed: %v\n", err)
		os.Exit(1)
	}

	// Automatically set JSON header for POST/PUT if data is provided
	if bodyReader != nil && (*requestType == "POST" || *requestType == "PUT") {
		req.Header.Set("Content-Type", "application/json")
	}

	// Verbose output
	if *verbose {
		fmt.Println("Request Method:", req.Method)
		fmt.Println("Request URL:", req.URL.String())
		if bodyReader != nil {
			fmt.Println("Request Body:", *data)
		}
		fmt.Println("Request Headers:", req.Header)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print response status
	fmt.Println("Status:", resp.Status, "✅")

	// Print response body
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading body failed: %v\n", err)
	}
}
