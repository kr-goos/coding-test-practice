package collatzconjecture

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	n = [TESTCOUNT]int{
		6,
		16,
		626331,
	}

	result = [TESTCOUNT]int{
		8,
		4,
		-1,
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("n : %d\n", n[i])
		fmt.Printf("result : %d\n", result[i])
		fmt.Printf("solution result : %d\n", solution(n[i]))
	}
}
