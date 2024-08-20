package customsort

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	strings = [TESTCOUNT][]string{
		{"sun", "bed", "car"},
		{"abce", "abcd", "cdx"},
	}
	n      = [TESTCOUNT]int{1, 2}
	result = [TESTCOUNT][]string{
		{"car", "bed", "sun"},
		{"abcd", "abce", "cdx"},
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("strings : ", strings[i])
		fmt.Println("n : ", n[i])
		fmt.Println("result : ", result[i])
		fmt.Println("solution result : ", solution(strings[i], n[i]))
	}
}
