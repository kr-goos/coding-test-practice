package sumbetween

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	a = [TESTCOUNT]int{
		3,
		3,
		5,
	}

	b = [TESTCOUNT]int{
		5,
		3,
		3,
	}

	result = [TESTCOUNT]int{
		12,
		3,
		12,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("a ", a[i])
		fmt.Println("b ", b[i])
		fmt.Println("result ", result[i])
		fmt.Println("solution result : ", solution(a[i], b[i]))
	}
}
