
package checker 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/operations"
)

// Arguments recupere les nombres saisis dans le terminal.
func Arguments(args []string) ([]int, error) {
	var numbers []int

	for _, arg := range args {
		values := strings.Fields(arg)

		for _, value := range values {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}

// CheckDuplicates verifie les doublons.
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

// Check lit et execute les instructions.
func Check() error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())

		if instruction == "" {
			continue
		}

		if !operations.Execute(instruction) {
			return fmt.Errorf("instruction invalide")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if !IsSorted(operations.A) || len(operations.B) != 0 {
		fmt.Println("KO")
		return nil
	}

	fmt.Println("OK")
	return nil
}

// IsSorted verifie si la pile est triee.
func IsSorted(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] > stack[i+1] {
			return false
		}
	}

	return true
}

// Run lance le checker.
func Run() {
	if err := Check(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
}

// Initialiser les piles.
func Init(numbers []int) {
	operations.A = numbers
	operations.B = []int{}
	operations.Result = []string{}
}
