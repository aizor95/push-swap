package main

import (
	"bufio"
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

	piles := pushswap.NewCheckerPiles(numbers)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if err := piles.ApplyInstruction(scanner.Text()); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if piles.IsSorted() && piles.BEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}