package sumofdivisors

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	n      = [TESTCOUNT]int{12, 5}
	result = [TESTCOUNT]int{28, 6}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("n : ", n[i])
		fmt.Println("result : ", result[i])
		r := solution(n[i])
		fmt.Println("solution result : ", r)
		if result[i] != r {
			t.Errorf("result %d != solution result %d", result[i], r)
		}
	}
}
