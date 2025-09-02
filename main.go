package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// CLI flag
	url := flag.String("url", "", "URL to fetch")

	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: go run main.go -url https://example.com")
		return
	}

	// send GET request
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close() // close when done

	// print status
	fmt.Println("Status:", resp.Status)

	// print body
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading body failed: %v\n", err)
	}
}
