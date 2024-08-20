package matrixaddition

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	arr1 = [TESTCOUNT][][]int{
		{{1, 2}, {2, 3}},
		{{1}, {2}},
	}
	arr2 = [TESTCOUNT][][]int{
		{{3, 4}, {5, 6}},
		{{3}, {4}},
	}
	result = [TESTCOUNT][][]int{
		{{4, 6}, {7, 9}},
		{{4}, {6}},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("arr1 : %v\n", arr1[i])
		fmt.Printf("arr2 : %v\n", arr2[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(arr1[i], arr2[i]))
	}
}
