package main

import (
	"flag"
	"fmt"
	"os"
)

func DisplayUsage() {
	fmt.Println("Websentry - Check the health of a website and report status")
	fmt.Println("Usage:")
	fmt.Println("websentry [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	os.Exit(0)
}
