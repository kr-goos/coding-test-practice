package parkwalk

import (
	"strconv"
	"strings"
)

const (
	// park map
	START      = "S"
	POSSIBLE   = "O"
	IMPOSSIBLE = "X"

	// moving direction
	EAST  = "E"
	WEST  = "W"
	NORTH = "N"
	SOUTH = "S"
)

type Position struct {
	x, y int
}

type Park struct {
	Map           [][]string
	rowLiMitIndex int
	colLiMitIndex int
}

type Robot struct {
	current Position
	park    Park
}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) GetPosition() []int {
	return []int{r.current.x, r.current.y}
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
				r.current.x = rowIdx
				r.current.y = colIdx
			}
			r.park.Map[rowIdx][colIdx] = v
		}
	}
}

func (r *Robot) isMovePossible(cmd string, distance int) bool {
	switch cmd {
	case EAST:
		if r.current.x+distance > r.park.colLiMitIndex {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.x+i][r.current.y] == IMPOSSIBLE {
				return false
			}
		}

	case WEST:
		if r.current.x-distance < 0 {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.x-i][r.current.y] == IMPOSSIBLE {
				return false
			}
		}
	case SOUTH:
		if r.current.y+distance > r.park.rowLiMitIndex {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.x][r.current.y+i] == IMPOSSIBLE {
				return false
			}
		}

	case NORTH:
		if r.current.y-distance < 0 {
			return false
		}

		for i := 1; i <= distance; i++ {
			if r.park.Map[r.current.x][r.current.y-i] == IMPOSSIBLE {
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
		r.current.x += distance
	case WEST:
		r.current.x -= distance
	case SOUTH:
		r.current.y += distance
	case NORTH:
		r.current.y -= distance
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

	return r.GetPosition()
}
