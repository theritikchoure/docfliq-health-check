package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/theritikchoure/logx"
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
		logx.Log("No websites specified. Example commands:", logx.FGRED, "")
		logx.Log("$ websentry -site https://example.com -maxretries 5 -retryinterval 10s -requesttimeout 15s", "", "")
		logx.Log("$ websentry -site https://example.com,https://example2.com,https://example3.com -maxretries 5 -retryinterval 10s -requesttimeout 15s", "", "")
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
	logx.Logf("----------------------------------------", logx.FGMAGENTA, "")
	logx.Logf("Checking the health of %s...", "", "", websiteURL)
	logx.Logf("----------------------------------------", logx.FGMAGENTA, "")

	client := &http.Client{
		Timeout: requestTimeout,
	}

	for i := 0; i < maxRetries; i++ {
		startTime := time.Now()
		response, err := client.Get(websiteURL)
		if err != nil || response.StatusCode != http.StatusOK {
			logx.Logf("Attempt %d: Website %s is down or not responding.", logx.FGRED, "", i+1, websiteURL)
			if i < maxRetries-1 {
				fmt.Printf("Retrying in %v...\n", retryInterval)
				time.Sleep(retryInterval)
			} else {
				fmt.Printf("Max retries reached. Website is still down.\n")
				os.Exit(1)
			}
		} else {
			sslEnabled := response.Header.Get("Strict-Transport-Security") != ""
			responseTime := time.Since(startTime)
			contentLength := response.ContentLength
			redirectURL := response.Request.URL
			serverHeader := response.Header.Get("Server")
			contentType := response.Header.Get("Content-Type")

			logx.Logf("✅ Website %s is up and running\n", "", "", websiteURL)
			logx.Logf("⭐ Response Time: %v", logx.FGGREEN, "", responseTime)
			logx.Logf("⭐ Content Length: %d bytes", logx.FGGREEN, "", contentLength)
			logx.Logf("⭐ Redirect URL: %s", logx.FGGREEN, "", redirectURL)
			logx.Logf("⭐ Server: %s", logx.FGGREEN, "", serverHeader)
			logx.Logf("⭐ Content Type: %s", logx.FGGREEN, "", contentType)
			logx.Logf("⭐ SSL enabled = %t", logx.FGGREEN, "", sslEnabled)
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
