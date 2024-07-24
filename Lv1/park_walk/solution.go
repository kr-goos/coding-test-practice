package parkwalk

import (
	"strconv"
	"strings"
)

const (
	// park constant
	START      = "S"
	POSSIBLE   = "O"
	IMPOSSIBLE = "X"

	// moving direction constant
	EAST  = "E"
	WEST  = "W"
	NORTH = "N"
	SOUTH = "S"
)

type PositionIndex struct {
	row, col int
}

type Park struct {
	Map           [][]string
	rowLiMitIndex int
	colLiMitIndex int
}

type Robot struct {
	current PositionIndex
	park    Park
}

func NewRobot() *Robot {
	return &Robot{current: PositionIndex{row: -1, col: -1}}
}

func (r *Robot) GetPositionIndex() []int {
	return []int{r.current.row, r.current.col}
}

func (r *Robot) SetParkPosition(park []string) {
	r.park.Map = make([][]string, len(park))
	r.park.rowLiMitIndex = len(park) - 1

	for rowIdx, rowStr := range park {
		if r.park.colLiMitIndex <= 0 {
			r.park.colLiMitIndex = len(rowStr) - 1
		}
		r.park.Map[rowIdx] = make([]string, len(rowStr))

		for colIdx := 0; colIdx < len(rowStr); colIdx++ {
			v := string(rowStr[colIdx])
			if v == START {
				r.current.row = rowIdx
				r.current.col = colIdx
			}
			r.park.Map[rowIdx][colIdx] = v
		}
	}
}

func (r *Robot) isMovePossible(cmd string, distance int) bool {
	if r.current.row == -1 || r.current.col == -1 {
		return false
	}

	switch cmd {
	case EAST:
		if r.current.col+distance > r.park.colLiMitIndex {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.row][r.current.col+i] == IMPOSSIBLE {
				return false
			}
		}

	case WEST:
		if r.current.col-distance < 0 {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.row][r.current.col-i] == IMPOSSIBLE {
				return false
			}
		}
	case SOUTH:
		if r.current.row+distance > r.park.rowLiMitIndex {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.row+i][r.current.col] == IMPOSSIBLE {
				return false
			}
		}

	case NORTH:
		if r.current.row-distance < 0 {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.row-i][r.current.col] == IMPOSSIBLE {
				return false
			}
		}

	default:
		// just in case..
		return false
	}
	return true

}

func (r *Robot) move(cmd string, distance int) {
	switch cmd {
	case EAST:
		r.current.col += distance
	case WEST:
		r.current.col -= distance
	case SOUTH:
		r.current.row += distance
	case NORTH:
		r.current.row -= distance
	}
}

func solution(park []string, routes []string) []int {
	r := NewRobot()
	r.SetParkPosition(park)

	for _, route := range routes {
		s := strings.Split(route, " ")
		distance, _ := strconv.Atoi(s[1])

		if r.isMovePossible(s[0], distance) {
			r.move(s[0], distance)
		}
	}

	return r.GetPositionIndex()
}
