package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joberly/imports"
)

func main() {
	// Check for the target project directory.
	dirName := os.Args[1]
	if dirName == "" {
		fmt.Printf("No project directory specified.\n")
		os.Exit(-1)
	}

	// Get imports list
	imps, err := imports.Get(dirName)
	if err != nil {
		fmt.Printf("Failed to get imports for %s: %s\n", dirName, err)
		os.Exit(-1)
	}

	// Generate JSON object text.
	outBytes, err := json.MarshalIndent(&imps, "", "  ")
	if err != nil {
		fmt.Printf("Error creating JSON output: %s\n", err)
		os.Exit(-1)
	}

	// Print JSON to stdout.
	fmt.Println(string(outBytes))
}
