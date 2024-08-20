package maskphonenumber

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 2

var (
	phone_number = [TESTCOUNT]string{
		"01033334444",
		"027778888",
	}

	result = [TESTCOUNT]string{
		"*******4444",
		"*****8888",
	}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Printf("phone_number : %v\n", phone_number[i])
		fmt.Printf("result : %v\n", result[i])
		fmt.Printf("solution result : %v\n", solution(phone_number[i]))
	}
}
