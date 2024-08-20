package repeatedstring

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	n      = [TESTCOUNT]int{3, 4}
	result = [TESTCOUNT]string{
		"수박수",
		"수박수박",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("n : ", n[i])
		fmt.Println("result : ", result[i])
		fmt.Println("solution result : ", solution(n[i]))
	}
}
