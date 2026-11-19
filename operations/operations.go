
package operations

// Les deux piles A et B
var A []int
var B []int

// Le resultat des instructions
var Result []string

// Echanger les deux premiers elements
func swap(stack []int) {
	if len(stack) > 1 {
		stack[0], stack[1] = stack[1], stack[0]
	}
}

// Echanger les deux premiers elements de A
func Sa() {
	if len(A) > 1 {
		swap(A)
		Result = append(Result, "sa")
	}
}

// Echanger les deux premiers elements de B
func Sb() {
	if len(B) > 1 {
		swap(B)
		Result = append(Result, "sb")
	}
}

// Echanger les deux piles
func Ss() {
	if len(A) > 1 || len(B) > 1 {
		swap(A)
		swap(B)
		Result = append(Result, "ss")
	}
}

// Faire passer le premier element en dernier
func rotate(stack []int) {
	if len(stack) > 1 {
		first := stack[0]

		for i := 0; i < len(stack)-1; i++ {
			stack[i] = stack[i+1]
		}

		stack[len(stack)-1] = first
	}
}

// Rotation de A
func Ra() {
	if len(A) > 1 {
		rotate(A)
		Result = append(Result, "ra")
	}
}

// Rotation de B
func Rb() {
	if len(B) > 1 {
		rotate(B)
		Result = append(Result, "rb")
	}
}

// Rotation des deux piles
func Rr() {
	if len(A) > 1 || len(B) > 1 {
		rotate(A)
		rotate(B)
		Result = append(Result, "rr")
	}
}

// Faire passer le dernier element en premier
func reverseRotate(stack []int) {
	if len(stack) > 1 {
		last := stack[len(stack)-1]

		for i := len(stack) - 1; i > 0; i-- {
			stack[i] = stack[i-1]
		}

		stack[0] = last
	}
}

// Reverse rotation de A
func Rra() {
	if len(A) > 1 {
		reverseRotate(A)
		Result = append(Result, "rra")
	}
}

// Reverse rotation de B
func Rrb() {
	if len(B) > 1 {
		reverseRotate(B)
		Result = append(Result, "rrb")
	}
}

// Reverse rotation des deux piles
func Rrr() {
	if len(A) > 1 || len(B) > 1 {
		reverseRotate(A)
		reverseRotate(B)
		Result = append(Result, "rrr")
	}
}

// Envoyer le premier element de B vers A
func Pa() {
	if len(B) > 0 {
		A = append([]int{B[0]}, A...)
		B = B[1:]
		Result = append(Result, "pa")
	}
}

// Envoyer le premier element de A vers B
func Pb() {
	if len(A) > 0 {
		B = append([]int{A[0]}, B...)
		A = A[1:]
		Result = append(Result, "pb")
	}
}

// Executer une instruction
func Execute(name string) bool {
	switch name {
	case "pa":
		Pa()
	case "pb":
		Pb()
	case "sa":
		Sa()
	case "sb":
		Sb()
	case "ss":
		Ss()
	case "ra":
		Ra()
	case "rb":
		Rb()
	case "rr":
		Rr()
	case "rra":
		Rra()
	case "rrb":
		Rrb()
	case "rrr":
		Rrr()
	default:
		return false
	}

	return true
}
