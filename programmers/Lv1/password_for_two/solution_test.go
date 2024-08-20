package passwordfortwo

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	sTestcase := "aukks"
	skipTestcase := "wbqd"
	index := 5
	successResult := "happy"

	result := solution(sTestcase, skipTestcase, index)

	if result == successResult {
		fmt.Println("test success")
	} else {
		t.Fatal("test failed")
	}
}
