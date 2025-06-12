package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arngrimur/randomfile/internal/pkg/command"
)

func main() {
	logFile, err := os.OpenFile("/tmp/randomfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open log file: %s\n", err)
		os.Exit(1)
	}
	log.SetOutput(logFile)
	image, err := command.GetImage(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		printHelp()
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, image)
}

func printHelp() {
	fmt.Println("Usage: randomfile <directory>")
	fmt.Println("Prints the path to a random file in the directory")
}
