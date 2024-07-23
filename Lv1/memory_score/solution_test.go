package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	const TEST_COUNT = 3
	nameTestCases := [TEST_COUNT][]string{
		{"may", "kein", "kain", "radi"},
		{"kali", "mari", "don"},
		{"may", "kein", "kain", "radi"},
	}
	yearningTestCases := [TEST_COUNT][]int{
		{5, 10, 1, 3},
		{11, 1, 55},
		{5, 10, 1, 3},
	}

	photoTestCases := [TEST_COUNT][][]string{
		{{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"}},
		{{"kali", "mari", "don"}, {"pony", "tom", "teddy"}, {"con", "mona", "don"}},
		{{"may"}, {"kein", "deny", "may"}, {"kon", "coni"}},
	}

	successResult := [TEST_COUNT][]int{
		{19, 15, 6},
		{67, 0, 55},
		{5, 15, 0},
	}

	for i := 0; i < TEST_COUNT; i++ {
		fmt.Println("nameTestCase :", nameTestCases[i])
		fmt.Println("yearningTestCase :", yearningTestCases[i])
		fmt.Println("photoTestCase :", photoTestCases[i])
		result := solution(nameTestCases[i], yearningTestCases[i], photoTestCases[i])
		fmt.Println("result :=", result)
		fmt.Println("successResult :=", successResult[i])
	}

}

func TestSolution2(t *testing.T) {
	const TEST_COUNT = 3
	nameTestCases := [TEST_COUNT][]string{
		{"may", "kein", "kain", "radi"},
		{"kali", "mari", "don"},
		{"may", "kein", "kain", "radi"},
	}
	yearningTestCases := [TEST_COUNT][]int{
		{5, 10, 1, 3},
		{11, 1, 55},
		{5, 10, 1, 3},
	}

	photoTestCases := [TEST_COUNT][][]string{
		{{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"}},
		{{"kali", "mari", "don"}, {"pony", "tom", "teddy"}, {"con", "mona", "don"}},
		{{"may"}, {"kein", "deny", "may"}, {"kon", "coni"}},
	}

	successResult := [TEST_COUNT][]int{
		{19, 15, 6},
		{67, 0, 55},
		{5, 15, 0},
	}

	for i := 0; i < TEST_COUNT; i++ {
		fmt.Println("nameTestCase :", nameTestCases[i])
		fmt.Println("yearningTestCase :", yearningTestCases[i])
		fmt.Println("photoTestCase :", photoTestCases[i])
		result := solution2(nameTestCases[i], yearningTestCases[i], photoTestCases[i])
		fmt.Println("result :=", result)
		fmt.Println("successResult :=", successResult[i])
	}

}

func BenchmarkSolution(b *testing.B) {
	nameTestCase := []string{
		"may", "kein", "kain", "radi",
	}
	yearningTestCase := []int{
		5, 10, 1, 3,
	}

	photoTestCase := [][]string{
		{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"},
	}

	for i := 0; i < b.N; i++ {
		_ = solution(nameTestCase, yearningTestCase, photoTestCase)
	}
}

func BenchmarkSolution2(b *testing.B) {
	nameTestCase := []string{
		"may", "kein", "kain", "radi",
	}
	yearningTestCase := []int{
		5, 10, 1, 3,
	}

	photoTestCase := [][]string{
		{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"},
	}

	for i := 0; i < b.N; i++ {
		_ = solution2(nameTestCase, yearningTestCase, photoTestCase)
	}
}
