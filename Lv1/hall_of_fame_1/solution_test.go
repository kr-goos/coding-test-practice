package halloffame1

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestLegacySort(t *testing.T) {
	nums := []int{3, 5, 6, 1, 0, 66}
	fmt.Println("before : ", nums)
	sort.Ints(nums)
	fmt.Println("after : ", nums)
}

func TestNewSort(t *testing.T) {
	nums := []int{3, 5, 6, 1, 0, 66}
	fmt.Println("before : ", nums)
	// go version 1.22 ~
	slices.Sort(nums)
	fmt.Println("after : ", nums)
}

func TestSolution(t *testing.T) {
	const TESTCOUNT = 2
	k := [TESTCOUNT]int{3, 4}
	score := [TESTCOUNT][]int{
		{10, 100, 20, 150, 1, 100, 200},
		{0, 300, 40, 300, 20, 70, 150, 50, 500, 1000},
	}
	result := [TESTCOUNT][]int{
		{10, 10, 10, 20, 20, 100, 100},
		{0, 0, 0, 0, 20, 40, 70, 70, 150, 300},
	}

	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("k : ", k[i])
		fmt.Println("score :", score[i])
		fmt.Println("succree result : ", result[i])
		fmt.Println("my solution result : ", solution(k[i], score[i]))
	}

}
