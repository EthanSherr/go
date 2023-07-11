package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Error: Expected 1 argument, but got %v", len(args))
		os.Exit(1)
	}
	filePath := os.Args[1]
	file, fErr := os.Open(filePath)
	if fErr != nil {
		fmt.Printf("Error: %v", fErr)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
