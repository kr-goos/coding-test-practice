package removesmallestnumber

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	arr = [TESTCOUNT][]int{
		{4, 3, 2, 1},
		{10},
	}

	result = [TESTCOUNT][]int{
		{4, 3, 2},
		{-1},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("arr : %v\n", arr[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(arr[i]))
	}

}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(arr[0])
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2(arr[0])
	}
}
