package desktopcleanup

const (
	FILE  byte = '#'
	EMPTY byte = '.'

	MIN int = -1
	MAX int = 51
)

func solution(wallpaper []string) []int {
	lux, luy, rdx, rdy := MAX, MAX, MIN, MIN

	for i, rowCells := range wallpaper {
		for j := 0; j < len(rowCells); j++ {
			if rowCells[j] == EMPTY {
				continue
			}

			if lux > i {
				lux = i
			}

			if rdx < i {
				rdx = i
			}

			if luy > j {
				luy = j
			}

			if rdy < j {
				rdy = j
			}
		}
	}
	return []int{lux, luy, rdx + 1, rdy + 1}
}
