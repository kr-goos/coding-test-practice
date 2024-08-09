package calculateshortfall

import (
	"fmt"
	"testing"
)

var (
	price, money, count int   = 3, 20, 4
	result              int64 = 10
)

func TestSolution(t *testing.T) {
	fmt.Println("price : ", price)
	fmt.Println("money : ", money)
	fmt.Println("count : ", price)
	r := solution(price, money, count)
	if r == result {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}
}

func TestImprovedSolution(t *testing.T) {
	fmt.Println("price : ", price)
	fmt.Println("money : ", money)
	fmt.Println("count : ", price)
	r := improvedSolution(price, money, count)
	if r == result {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(price, money, count)
	}
}

func BenchmarkImprovedSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		improvedSolution(price, money, count)
	}
}
