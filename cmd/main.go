package main

import (
	"fmt"
	"os"

	"github.com/arngrimur/randomfile/internal/pkg/command"
)

func main() {
	image, err := command.GetImage(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, image)
}
