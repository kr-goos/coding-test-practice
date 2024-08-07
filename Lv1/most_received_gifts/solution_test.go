package mostreceivedgifts

import (
	"fmt"
	"testing"
)

const TESTCOUNT = 3

var (
	friends = [TESTCOUNT][]string{
		{"muzi", "ryan", "frodo", "neo"},
		{"joy", "brad", "alessandro", "conan", "david"},
		{"a", "b", "c"},
	}

	gifts = [TESTCOUNT][]string{
		{"muzi frodo", "muzi frodo", "ryan muzi", "ryan muzi", "ryan muzi", "frodo muzi", "frodo ryan", "neo muzi"},
		{"alessandro brad", "alessandro joy", "alessandro conan", "david alessandro", "alessandro david"},
		{"a b", "b a", "c a", "a c", "a c", "c a"},
	}

	result = []int{2, 4, 0}
)

func TestSolution(t *testing.T) {
	for i := 0; i < TESTCOUNT; i++ {
		fmt.Println("friends[i] : ", friends[i])
		fmt.Println("gifts[i] : ", gifts[i])
		fmt.Println("result[i] : ", result[i])
		fmt.Println("my solution result : ", solution(friends[i], gifts[i]))
	}

}
