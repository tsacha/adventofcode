package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/k0kubun/pp/v3"
	"github.com/tsacha/adventofcode/utils"
)

type position struct {
	x int
	y int
}

type image map[position]string

func run(void int) (solution int) {
	input := string(utils.PuzzleInput(2023, 11))
	_ = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	space := make(image)
	row_galaxies := make(map[int]bool)
	col_galaxies := make(map[int]bool)

	mx, my := 0, 0

	for x, line := range strings.Split(input, "\n") {
		mx = x
		if line == "" {
			continue
		}
		for y, c := range strings.Split(line, "") {
			my = y
			if c == "#" {
				space[position{x: x, y: y}] = c
				row_galaxies[x] = true
				col_galaxies[y] = true
			}
		}
	}

	galaxies := []position{}
	ix, ex := false, 0
	for x := 0; x <= mx; x++ {
		if ix {
			ex = ex + void
			ix = false
		}
		ey := 0
		for y := 0; y <= my; y++ {
			if space[position{x: x, y: y}] == "#" {
				galaxies = append(galaxies, position{x: x + ex, y: y + ey})
			}

			if !col_galaxies[y] {
				ey = ey + void
			}
			if !row_galaxies[x] {
				ix = true
			}
		}
	}

	for n := 0; n < len(galaxies); n++ {
		for p := 0; p < n; p++ {
			dx := int(math.Abs(float64(galaxies[n].x) - float64(galaxies[p].x)))
			dy := int(math.Abs(float64(galaxies[n].y) - float64(galaxies[p].y)))
			solution += dx + dy
		}
	}

	return solution
}

func main() {
	pp.Printf("Je fais des trucs\n")
	fmt.Printf("Solution 1: %d\n", run(1))
	fmt.Printf("Solution 2: %d\n", run(1000000-1))

}
