package desktopcleanup

type Desktop struct {
	wallpaper Wallpaper
}

func NewDesktop() *Desktop {
	return &Desktop{}
}

func (d *Desktop) SetWallpaper(wallpaper []string) {
	d.wallpaper = make([][]string, len(wallpaper))

	for xi, yFiles := range wallpaper {
		d.wallpaper[xi] = make([]string, len(yFiles))

		for yi := 0; yi < len(yFiles); yi++ {
			d.wallpaper[xi][yi] = string(yFiles[yi])
		}
	}
}

type Wallpaper [][]string

func solution(wallpaper []string) []int {
	return []int{}
}
