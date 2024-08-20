package main

import (
	"log"

	"github.com/mauroandocilla/cli-squeezy-pix/cmd"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
