
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/operations"
)

// Recuperer et convertir les arguments
func Arguments(args []string) ([]int, error) {
	var numbers []int

	for _, arg := range args {
		values := strings.Fields(arg)

		for _, value := range values {
			n, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, n)
		}
	}

	return numbers, nil
}

// Verifier les doublons
func CheckDuplicates(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				return false
			}
		}
	}

	return true
}

// Verifier si la pile est deja triee
func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}

	return true
}

// Trier les nombres
func Sort() {
	// Ne rien faire si la pile est deja triee
	if IsSorted(operations.A) {
		return
	}

	// Trier deux nombres avec une seule instruction
	if len(operations.A) == 2 {
		operations.Sa()
		return
	}

	// Trier trois nombres
	if len(operations.A) == 3 {
		a := operations.A

		if a[0] > a[1] && a[1] < a[2] && a[0] < a[2] {
			operations.Sa()
		} else if a[0] > a[1] && a[1] > a[2] {
			operations.Sa()
			operations.Rra()
		} else if a[0] > a[1] && a[1] < a[2] {
			operations.Ra()
		} else if a[0] < a[1] && a[1] > a[2] && a[0] < a[2] {
			operations.Sa()
			operations.Ra()
		} else if a[0] < a[1] && a[1] > a[2] {
			operations.Rra()
		}

		return
	}

	// Trier les piles de plus de trois nombres
	for len(operations.A) > 0 {
		minIndex := 0

		// Chercher le plus petit nombre
		for i := 1; i < len(operations.A); i++ {
			if operations.A[i] < operations.A[minIndex] {
				minIndex = i
			}
		}

		// Choisir la rotation la plus courte
		if minIndex <= len(operations.A)/2 {
			for minIndex > 0 {
				operations.Ra()
				minIndex--
			}
		} else {
			rotations := len(operations.A) - minIndex

			for rotations > 0 {
				operations.Rra()
				rotations--
			}
		}

		// Envoyer le plus petit nombre dans B
		operations.Pb()
	}

	// Remettre les nombres dans A
	for len(operations.B) > 0 {
		operations.Pa()
	}
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	numbers, err := Arguments(os.Args[1:])
	if err != nil || len(numbers) == 0 || !CheckDuplicates(numbers) {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Initialiser les deux piles
	operations.A = numbers
	operations.B = []int{}
	operations.Result = []string{}

	// Lancer le tri
	Sort()

	// Afficher les instructions
	for _, instruction := range operations.Result {
		fmt.Println(instruction)
	}
}
