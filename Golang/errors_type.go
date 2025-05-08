package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// fetchURL makes an HTTP GET request and returns the response body or an error
func fetchURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	data, err := fetchURL(url)
	if err != nil {
		fmt.Printf("*******************\n")
		log.Fatalf("server error : (%s)", err, "\n \n")
		//fmt.Errorf("Error fetching URL: %v\n", url)
		os.Exit(1)
	}

	fmt.Printf("Fetched %d bytes from %s:\n\n%s\n", len(data), url, data[:100]) // Print first 100 chars
}
