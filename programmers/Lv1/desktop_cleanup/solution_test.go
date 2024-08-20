package desktopcleanup

import (
	"fmt"
	"testing"
)

func TestSetWallpaper(t *testing.T) {

	wallpaperTestCases := [][]string{
		{".#...", "..#..", "...#."},
		{"..........", ".....#....", "......##..", "...##.....", "....#....."},
		{".##...##.", "#..#.#..#", "#...#...#", ".#.....#.", "..#...#..", "...#.#...", "....#...."},
		{"..", "#."},
	}

	successResults := [][]int{
		{0, 1, 3, 4},
		{1, 3, 5, 8},
		{0, 0, 7, 9},
		{1, 0, 2, 1},
	}

	for i, wcase := range wallpaperTestCases {
		fmt.Println("testcase : ", wcase)
		result := solution(wcase)
		for j, v := range result {
			if v != successResults[i][j] {
				t.Fatalf("unmatch result %d == %d", v, successResults[i][j])
			}
		}
		fmt.Println("result : ", result)
	}

}
