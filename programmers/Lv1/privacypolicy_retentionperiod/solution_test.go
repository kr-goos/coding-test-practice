package privacypolicyretentionperiod

import (
	"fmt"
	"testing"
	"time"
)

func TestDateToDays(t *testing.T) {
	tc := time.Now().Format("2006.01.02")

	fmt.Println("testcaes : ", tc)
	v := dateToDays(tc)
	fmt.Println("result : ", v)
}

func TestSolution(t *testing.T) {
	const TCCOUNT = 2
	todayTestcases := [TCCOUNT]string{
		"2022.05.19",
		"2020.01.01",
	}
	termsTestcases := [TCCOUNT][]string{
		{"A 6", "B 12", "C 3"},
		{"Z 3", "D 5"},
	}
	privaciesTestcases := [TCCOUNT][]string{
		{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"},
		{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"},
	}

	successResults := [TCCOUNT][]int{
		{1, 3},
		{1, 4, 5},
	}
	for i := 0; i < TCCOUNT; i++ {
		fmt.Println("todayTestcase : ", todayTestcases[i])
		fmt.Println("termsTestcase : ", termsTestcases[i])
		fmt.Println("privaciesTestcase : ", privaciesTestcases[i])
		fmt.Println("successResults : ", successResults[i])
		fmt.Println("result : ", solution(todayTestcases[i], termsTestcases[i], privaciesTestcases[i]))
	}

}
