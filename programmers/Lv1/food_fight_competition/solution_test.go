package foodfightcompetition

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	food = [][]int{
		{1, 3, 4, 6},
		{1, 7, 1, 2},
	}
	result = []string{
		"1223330333221",
		"111303111",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("food[i] : ", food[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("my solution result : ", solution(food[i]))
	}
}

func TestImprovedSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("food[i] : ", food[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("my solution result : ", ImprovedSolution(food[i]))
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(food[0])
	}
}

func BenchmarkImprovedSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImprovedSolution(food[0])
	}
}
