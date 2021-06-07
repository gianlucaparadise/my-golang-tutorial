package exercises

import (
	"fmt"
	"io"
	"math"
	"strings"

	"golang.org/x/tour/tree"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, z = %v", i, z)
	}
	return z
}

func SqrtFaster(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		newZ := z - (z*z-x)/(2*z)
		fmt.Printf("Iteration %v, z = %v", i, newZ)
		if math.Abs(z-newZ) < 0.000001 {
			break
		}

		z = newZ
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

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func SqrtWithError(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	result := Sqrt(x)
	return result, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {

	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLM"
	lowercase := "abcdefghijklmnopqrstuvwxyzabcdefghijklm"

	n, err := rot13.r.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			var letter = b[i]
			if 'A' <= letter && letter <= 'Z' {
				b[i] = uppercase[letter-'A'+13]
			} else if 'a' <= letter && letter <= 'z' {
				b[i] = lowercase[letter-'a'+13]
			}
		}
	}
	return n, err
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.Value

	if t.Left != nil {
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

func StartWalking(t *tree.Tree, ch chan int) {
	go Walk(t, ch)

	close(ch)
}
