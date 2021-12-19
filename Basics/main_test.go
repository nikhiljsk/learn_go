package main

import (
	"fmt"
	"testing"
)

func ExampleDiffAll() {
	fmt.Println(DiffAll(6, 7, 9))
	// Output:
	// -22
}

func ExampleDiffAll_two() {
	fmt.Println(DiffAll(5, 7, 9))
	// Output:
	// -21
}

func TestSumAll(t *testing.T) {
	result := sumAll(1, 2, 3)
	if result != 6 {
		t.Error("Nah! That's bad. Expected 6, got:", result)
	}
	fmt.Println("Aha! Nice!", result)
}

func TestTableTest(t *testing.T) {
	// Table tests
	type tests struct {
		numbers []int
		answer  int
	}

	testcases := []tests{
		{
			numbers: []int{1, 2, 3},
			answer:  6,
		},
		{
			numbers: []int{1, 27, 3},
			answer:  31,
		},
	}

	for _, testcase := range testcases {
		result := sumAll(testcase.numbers...)
		if result != testcase.answer {
			t.Errorf("Nananana! What a mistake. Expected %v got %v", testcase.answer, result)
		}
	}
	fmt.Println("Great work mate!")
}

func BenchmarkDiffAll(b *testing.B) { // Benchmark is for testing the time for code to execute
	for i := 0; i < b.N; i++ { // Runs multiple times, and b.N is determined dynamically, and runs until the ToE is statistically accurate
		DiffAll(1, 3000, 43399, 3512391358)
	}
}
