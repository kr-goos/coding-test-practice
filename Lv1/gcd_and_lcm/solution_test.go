package gcdandlcm

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	n      = [TESTCOUNT]int{3, 2}
	m      = [TESTCOUNT]int{12, 5}
	result = [TESTCOUNT][]int{
		{3, 12},
		{1, 10},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("n : %d\n", n[i])
		fmt.Printf("m : %d\n", m[i])

		r := solution(n[i], m[i])
		for j, answer := range result[i] {
			if answer != r[j] {
				t.Errorf("test %d %d index failure", i, j)
			}
		}
	}
}
