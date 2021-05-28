package exercises

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	expected := 2.0
	actual := Sqrt(4)

	if actual != expected {
		t.Fatalf(`Expected and Actual differ: Expected = %v   Actual = %v`, expected, actual)
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
