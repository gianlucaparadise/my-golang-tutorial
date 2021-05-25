package exercises

import (
	"reflect"
	"testing"
)

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
