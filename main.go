package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	websiteURL := "https://docfliq.com" // Replace with your website URL

	response, err := http.Get(websiteURL)
	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Printf("Website %s is down or not responding.\n", websiteURL)
		os.Exit(1)
	}
	fmt.Printf("Website %s is up and running!\n", websiteURL)
}
