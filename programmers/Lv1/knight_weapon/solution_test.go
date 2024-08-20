package knightweapon

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	number = [TESTCOUNT]int{5, 10}
	limit  = [TESTCOUNT]int{3, 3}
	power  = [TESTCOUNT]int{2, 2}
	result = [TESTCOUNT]int{10, 21}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("number[i] : ", number[i])
		fmt.Println("limit[i] : ", limit[i])
		fmt.Println("power[i] : ", power[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("solution result : ", solution(number[i], limit[i], power[i]))
	}
}

func TestSolution2(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("number[i] : ", number[i])
		fmt.Println("limit[i] : ", limit[i])
		fmt.Println("power[i] : ", power[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("solution result : ", solution(number[i], limit[i], power[i]))
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(number[1], limit[1], power[1])
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2(number[1], limit[1], power[1])
	}
}
