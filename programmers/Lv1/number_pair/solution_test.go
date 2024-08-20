package numberpair

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 5

var (
	X = [TESTCOUNT]string{
		"100",
		"100",
		"100",
		"12321",
		"5525",
	}

	Y = [TESTCOUNT]string{
		"2345",
		"203045",
		"123450",
		"42531",
		"1255",
	}
	result = [TESTCOUNT]string{
		"-1",
		"0",
		"10",
		"321",
		"552",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("X[i] : ", X[i])
		fmt.Println("Y[i] : ", Y[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("solution result : ", solution(X[i], Y[i]))
	}
}

func TestCharacterFrequencyArr(t *testing.T) {
	result := []int{
		1, 3, 1, 1, 1, 1, 0, 0, 1, 1,
	}
	fmt.Println(CharacterFrequencyArr("1231541980"))

	for i, v := range CharacterFrequencyArr("1231541980") {
		fmt.Printf("[%d] expected value : %d, function value : %d\n", i, result[i], v)
	}

}
