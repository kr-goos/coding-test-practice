package harshadnumber

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 4

var (
	x      = [TESTCOUNT]int{10, 12, 11, 13}
	result = [TESTCOUNT]bool{true, true, false, false}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("x : %d\n", x[i])
		fmt.Printf("result : %t\n", result[i])
		fmt.Printf("solution result : %t\n", solution(x[i]))
	}
}

func TestSolution2(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("x : %d\n", x[i])
		fmt.Printf("result : %t\n", result[i])
		fmt.Printf("solution result : %t\n", solution2(x[i]))
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(x[0])
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2(x[0])
	}
}
