package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/theritikchoure/logx"
)

var (
	websiteURL     = ""               // Replace with your website URL
	maxRetries     = 3                // Maximum number of retries
	retryInterval  = 5 * time.Second  // Time to wait between retries
	requestTimeout = 10 * time.Second // Timeout for HTTP requests
)

var options = make(map[string]string)

func main() {

	if len(os.Args) <= 2 {
		log.Fatalf("Please specify a valid arguments")
	}

	for i := 1; i < len(os.Args); i += 2 {
		key := os.Args[i]

		if i+1 >= len(os.Args) {
			log.Fatal("Invalid arguments, please specify value for ", key)
		}

		value := os.Args[i+1]
		if strings.Contains(strings.ToLower(value), "--") {
			log.Fatal("Please specify a valid argument for ", key)
		}

		options[key] = value
	}

	if options["--site"] != "" {
		websiteURL = options["--site"]

		if !strings.Contains(strings.ToLower(websiteURL), "http") {
			websiteURL = "https://" + websiteURL
		}
	} else {
		log.Fatal("Site is not specified")
	}

	if options["--maxretry"] != "" {
		retryStr := options["--maxretry"]
		maxRetries, _ = strconv.Atoi(retryStr)
	}

	if options["--retryin"] != "" {
		retryInStr, found := options["--retryin"]
		if found {
			retryIn, err := strconv.Atoi(retryInStr)

			if err == nil {
				retryInterval = time.Duration(retryIn) * time.Second
			} else {
				log.Fatal("Error converting '--retryin' to an integer: \n", err)
			}
		}
	}

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
