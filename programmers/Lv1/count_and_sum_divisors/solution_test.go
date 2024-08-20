package countandsumdivisors

import (
	"fmt"
	"log"
	"testing"
)

const TESTCOUNT = 2

var (
	left   = [TESTCOUNT]int{13, 24}
	right  = [TESTCOUNT]int{17, 27}
	result = [TESTCOUNT]int{43, 52}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("left : ", left[i])
		fmt.Println("right : ", right[i])

		if solution(left[i], right[i]) == result[i] {
			fmt.Printf("test %d success\n", i)
		} else {
			log.Printf("test %d failure\n", i)
		}
	}
}

func TestCountDivisors(t *testing.T) {

	const TESTCOUNT = 5
	v := [TESTCOUNT]int{13, 14, 15, 16, 17}
	result := [TESTCOUNT]int{2, 4, 4, 5, 2}

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("v : ", v[i])
		fmt.Println("result : ", result[i])
		fmt.Println("countDivisors : ", countDivisors(v[i]))
	}
}
