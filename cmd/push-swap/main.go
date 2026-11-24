package main

import (
	"fmt"
	"os"

	"pushswap/internal/pushswap"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	numbers, err := pushswap.ParseNumbers(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	for _, instruction := range pushswap.Generate(numbers) {
		fmt.Println(instruction)
	}
}