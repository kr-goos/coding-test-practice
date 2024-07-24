package parkwalk

import (
	"fmt"
	"log"
	"testing"
)

func TestSetParkPosition(park []string) {
	parkTestcases := [][]string{
		{"SOO", "OOO", "OOO"},
		{"SOO", "OXX", "OOO"},
		{"OSO", "OOO", "OXO", "OOO"},
	}
	r := NewRobot()

	for _, tc := range parkTestcases {
		r.SetParkPosition(tc)
		fmt.Println("Map : ", r.park.Map)
	}

}

func TestSolution(t *testing.T) {
	parkTestcase := []string{
		"SOO", "OXX", "OOO",
	}
	routesTestcase := []string{
		"E 2", "S 2", "W 1",
	}

	successResult := []int{0, 1}
	result := solution(parkTestcase, routesTestcase)
	for i, s := range successResult {

		if s == result[i] {
			fmt.Printf("success : %d\n", s)
		} else {
			log.Printf("failure : %d, successResult : %d\n", result[i], s)
		}
	}
}
