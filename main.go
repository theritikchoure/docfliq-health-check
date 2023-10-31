package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const version = "v1.2.0"

var (
	websiteURL string
	maxRetries int
)

var (
	retryInterval  time.Duration
	requestTimeout time.Duration
)

func main() {
	helpFlag := flag.Bool("help", false, "Display usage information")
	versionFlag := flag.Bool("version", false, "Print the program version")
	websites := flag.String("site", "", "Comma-separated list of website URLs")
	maxRetriesValue := flag.Int("maxretries", 3, "Maximum number of retries")
	retryIntervalValue := flag.Duration("retryinterval", 5*time.Second, "Time to wait between retries")
	requestTimeoutValue := flag.Duration("requesttimeout", 10*time.Second, "Timeout for HTTP requests")

	flag.Parse()

	if *helpFlag {
		displayUsage()
	}

	if *versionFlag {
		displayVersion()
	}

	if *websites == "" {
		fmt.Println("No websites specified. Example commands:")
		fmt.Println("$ websentry -site https://example.com -maxretries 5 -retryinterval 10s -requesttimeout 15s")
		fmt.Println("$ websentry -site https://example.com,https://example2.com,https://example3.com -maxretries 5 -retryinterval 10s -requesttimeout 15s")
		return
	}

	maxRetries = *maxRetriesValue
	retryInterval = *retryIntervalValue
	requestTimeout = *requestTimeoutValue

	websiteURLs := strings.Split(*websites, ",")

	for _, url := range websiteURLs {
		if !strings.Contains(strings.ToLower(url), "http") {
			websiteURL = "https://" + url
		} else {
			websiteURL = url
		}
		checkHealth()
	}
}

func checkHealth() {
	fmt.Println("----------------------------------------")
	fmt.Printf("Checking the health of %s...\n", websiteURL)
	fmt.Println("----------------------------------------")

	client := &http.Client{
		Timeout: requestTimeout,
	}

	for i := 0; i < maxRetries; i++ {
		startTime := time.Now()
		response, err := client.Get(websiteURL)
		if err != nil || response.StatusCode != http.StatusOK {
			fmt.Printf("Attempt %d: Website %s is down or not responding.\n", i+1, websiteURL)
			if i < maxRetries-1 {
				fmt.Printf("Retrying in %v...\n", retryInterval)
				time.Sleep(retryInterval)
			} else {
				fmt.Println("Max retries reached. Website is still down.")
				os.Exit(1)
			}
		} else {
			sslEnabled := response.Header.Get("Strict-Transport-Security") != ""
			responseTime := time.Since(startTime)
			contentLength := response.ContentLength
			redirectURL := response.Request.URL
			serverHeader := response.Header.Get("Server")
			contentType := response.Header.Get("Content-Type")

			fmt.Println("✅ Website", websiteURL, "is up and running")
			fmt.Printf("⭐ Response Time: %v\n", responseTime)
			fmt.Printf("⭐ Content Length: %d bytes\n", contentLength)
			fmt.Printf("⭐ Redirect URL: %s\n", redirectURL)
			fmt.Printf("⭐ Server: %s\n", serverHeader)
			fmt.Printf("⭐ Content Type: %s\n", contentType)
			fmt.Printf("⭐ SSL enabled = %t\n", sslEnabled)
			break
		}
	}
}

func displayUsage() {
	fmt.Println("Websentry - Check the health of a website and report status")
	fmt.Println("Usage:")
	fmt.Println("websentry [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	os.Exit(0)
}

func displayVersion() {
	fmt.Printf("Websentry %s\n", version)
	os.Exit(0)
}
