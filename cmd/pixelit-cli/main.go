package main

import (
	"fmt"
	"os"
)

func main() {
	// Execute()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
