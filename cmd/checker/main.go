
package main

import (
	"fmt"
	"os"

	"push-swap/checker"
)

func main() {
	// Vérifier les arguments
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Récupérer les nombres
	numbers, err := checker.Arguments(os.Args[1:])
	if err != nil || len(numbers) == 0 {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Vérifier les doublons
	if !checker.CheckDuplicates(numbers) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Initialiser les piles
	checker.Init(numbers)

	// Lancer le checker
	checker.Run()
}
