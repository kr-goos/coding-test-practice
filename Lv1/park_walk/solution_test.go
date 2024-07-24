package parkwalk

import (
	"fmt"
	"log"
	"testing"
)

func TestSetParkPosition(t *testing.T) {
	parkTestcases := [][]string{
		{"SOO", "OOO", "OOO"},
		{"SOO", "OXX", "OOO"},
		{"OSO", "OOO", "OXO", "OOO"},
	}

	r := NewRobot()

	for _, tc := range parkTestcases {
		fmt.Println("testcase : ", tc)
		r.SetParkPosition(tc)
		fmt.Println("Map : ", r.park.Map)
	}

}

func TestIsMovePossible(t *testing.T) {
	r := NewRobot()

	cmdTestcase := []string{"E", "N", "S", "W"}
	distanceTestcase := []int{1, 3, 2, 5}

	for i := range cmdTestcase {
		fmt.Printf("%s %d is move possible %t\n", cmdTestcase[i], distanceTestcase[i], r.isMovePossible(cmdTestcase[i], distanceTestcase[i]))
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
