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
	
	url := flag.String("url", "", "URL to fetch")
	requestType := flag.String("request", "GET", "HTTP method: GET, POST, etc.") // new flag
	data := flag.String("data", "", "Data to send with POST")
	verbose := flag.Bool("verbose", false, "Show verbose output")

	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: go run main.go -url https://example.com [-request GET/POST] [-data 'body']")
		return
	}

	// Prepare the request
	var req *http.Request
	var err error

	if strings.ToUpper(*requestType) == "POST" {
		req, err = http.NewRequest("POST", *url, strings.NewReader(*data)) // handle POST body
	} else {
		req, err = http.NewRequest(strings.ToUpper(*requestType), *url, nil)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Creating request failed: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Println("Request Method:", req.Method)
		fmt.Println("Request URL:", req.URL.String())
		if *data != "" {
			fmt.Println("Request Body:", *data)
		}
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print response
	fmt.Println("Status:", resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading body failed: %v\n", err)
	}
}
