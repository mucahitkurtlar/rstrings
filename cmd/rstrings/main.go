package main

import (
	"fmt"
	"os"

	"github.com/mucahitkurtlar/rstrings/internal/rstrings"
)

func main() {
	args := os.Args[1:]

	output, err := rstrings.Rstrings(args...)
	if err != nil {
		fmt.Printf("%s: %s\n", os.Args[0], err)
		fmt.Println(rstrings.HelpText())
		os.Exit(1)
	}
	fmt.Println(output)
}
