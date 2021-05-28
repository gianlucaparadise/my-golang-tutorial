package exercises

import (
	"fmt"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, z = %v", i, z)
	}
	return z
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	result := make(map[string]int)

	for _, w := range words {
		result[w] = result[w] + 1
	}

	return result
}

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	prev := 0
	current := 0
	return func() int {

		if current == 0 && prev == 0 {
			current = 1
			return 1
		}

		result := prev + current
		prev, current = current, result

		return result
	}
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}
