package exercises

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	const count = 4
	inputs := [count]float64{4.0, 16.0, 25.0, 121.0}
	expectedList := [count]float64{2.0, 4.0, 5.0, 11.0}

	for i := 0; i < count; i++ {
		input := inputs[i]
		actual := Sqrt(input)
		expected := expectedList[i]

		if actual != expected {
			t.Fatalf(`Expected and Actual differ: Expected = %v   Actual = %v`, expected, actual)
		}
	}
}

func TestSqrtFaster(t *testing.T) {
	const count = 4
	inputs := [count]float64{4.0, 16.0, 25.0, 121.0}
	expectedList := [count]float64{2.0, 4.0, 5.0, 11.0}

	for i := 0; i < count; i++ {
		input := inputs[i]
		actual := SqrtFaster(input)
		expected := expectedList[i]

		if math.Abs(actual-expected) > 0.000001 {
			t.Fatalf(`Expected and Actual differ: Expected = %v   Actual = %v`, expected, actual)
		}
	}
}

func TestWordCount(t *testing.T) {
	testString1 := "I ate a donut. Then I ate another donut."
	expectedResult1 := map[string]int{"I": 2, "Then": 1, "a": 1, "another": 1, "ate": 2, "donut.": 2}
	actualResult1 := WordCount(testString1)

	if !reflect.DeepEqual(expectedResult1, actualResult1) {
		t.Fatalf(`f("%v") failed`, testString1)
	}

	testString2 := "A man a plan a canal panama."
	expectedResult2 := map[string]int{"A": 1, "a": 2, "canal": 1, "man": 1, "panama.": 1, "plan": 1}
	actualResult2 := WordCount(testString2)

	if !reflect.DeepEqual(expectedResult2, actualResult2) {
		t.Fatalf(`f("%v") failed`, testString2)
	}
}

func TestFibonacciCount(t *testing.T) {
	f := Fibonacci()
	fibonacciExpected := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	for i := 0; i < 10; i++ {
		actual := f()
		expected := fibonacciExpected[i]

		t.Logf(`Index: %v Actual = %v - Expected = %v`, i, actual, expected)

		if actual != expected {
			t.Fatalf(`Values at index "%v" differ. Expected: "%v" - Actual: "%v"`, i, expected, actual)
		}
	}
}

func TestStringer(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	hostsExpected := map[string]string{
		"loopback":  "127.0.0.1",
		"googleDNS": "8.8.8.8",
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		hostActual := fmt.Sprint(ip)
		hostExpected := hostsExpected[name]

		if hostActual != hostExpected {
			t.Fatalf(`Values with name "%v" differ. Expected: %v - Actual: %v`, name, hostExpected, hostActual)
		}
	}
}

func TestSqrtWithError(t *testing.T) {
	_, err := SqrtWithError(-2)

	if err == nil {
		t.Fatalf(`Expected error not returned`)
	}
}

// This function is taken from "golang.org/x/tour/reader" v0.1.0
func ValidateReader(r io.Reader) error {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				return fmt.Errorf("got byte %x at offset %v, want 'A'\n", v, o+i)
			}
		}
		o += n
		if err != nil {
			return fmt.Errorf("read error: %v\n", err)
		}
	}
	if o == 0 {
		return fmt.Errorf("read zero bytes after %d Read calls\n", i)
	}
	return nil
}

func TestMyReader(t *testing.T) {
	error := ValidateReader(MyReader{})

	if error != nil {
		t.Fatal(error)
	}
}
