package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type position struct {
	x int
	y int
}
type point struct {
	v      int
	c      int
	inside bool
}
type maze map[position]point
type path []position

func drawMaze(part int, input string) (m maze) {
	m = maze{}
	path := path{}
	cx, cy := 0, 0
	m[position{x: cx, y: cy}] = point{
		v: 1,
		c: 0x000000,
	}

	for _, row_str := range strings.Split(strings.Trim(input, "\n"), "\n") {
		row := strings.Fields(row_str)
		move := row[0]
		count, _ := strconv.Atoi(row[1])
		if part == 2 {
			color_str, _ := strconv.ParseInt(row[2][2:len(row[2])-1], 16, 32)
			switch color_str & 0xf {
			case 0:
				move = "R"
			case 1:
				move = "D"
			case 2:
				move = "L"
			case 3:
				move = "U"
			}
			count = int(color_str >> 4)
		}
		for i := 0; i < count; i++ {
			switch move {
			case "R":
				cy++
			case "D":
				cx++
			case "U":
				cx--
			case "L":
				cy--
			}
			m[position{x: cx, y: cy}] = point{v: 1, c: 0x000000}
			path = append(path, position{x: cx, y: cy})
		}

	}
	x_min, y_min, x_max, y_max := 0, 0, 0, 0
	for c := range m {
		if c.x < x_min {
			x_min = c.x
		}
		if c.y < y_min {
			y_min = c.y
		}

		if c.x > x_max {
			x_max = c.x
		}
		if c.y > y_max {
			y_max = c.y
		}

	}
	insides := []position{}
	for x := x_min; x <= x_max; x++ {
		for y := y_min; y <= y_max; y++ {
			p := position{x: x, y: y}

			if m[p].v == 1 {
				continue
			}
			inside := false
			loop := true
			i := 0
			j := len(m) - 1
			for loop {
				xi := path[i].x
				yi := path[i].y

				xj := path[j].x
				yj := path[j].y

				var intersect = ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi)
				if intersect {
					inside = !inside
				}

				j = i
				i++
				if i > len(path)-1 {
					loop = false
				}
			}
			if inside {
				insides = append(insides, p)
			}
		}

	}
	for _, inside := range insides {
		m[inside] = point{v: 0, inside: true, c: 0x000000}
	}
	return m
}

func (m *maze) print() {
	x_min, y_min, x_max, y_max := 0, 0, 0, 0
	for p := range *m {
		if p.x < x_min {
			x_min = p.x
		}
		if p.y < y_min {
			y_min = p.y
		}

		if p.x > x_max {
			x_max = p.x
		}
		if p.y > y_max {
			y_max = p.y
		}

	}
	for x := x_min; x <= x_max; x++ {
		fmt.Printf("\n")
		for y := y_min; y <= y_max; y++ {
			p := position{x: x, y: y}
			if (*m)[p].v == 1 {
				fmt.Printf("#")
			} else if (*m)[p].inside == true {
				fmt.Printf("~")
			} else {
				fmt.Printf(".")
			}

		}
	}
	fmt.Println()
}
func (m *maze) solve() (solution int) {
	x_min, y_min, x_max, y_max := 0, 0, 0, 0
	for p := range *m {
		if p.x < x_min {
			x_min = p.x
		}
		if p.y < y_min {
			y_min = p.y
		}

		if p.x > x_max {
			x_max = p.x
		}
		if p.y > y_max {
			y_max = p.y
		}

	}
	for x := x_min; x <= x_max; x++ {
		for y := y_min; y <= y_max; y++ {
			p := position{x: x, y: y}
			if (*m)[p].v == 1 {
				solution++
			} else if (*m)[p].inside == true {
				solution++
			}

		}
	}
	return solution
}

func part1() (result int) {
	input := string(utils.PuzzleInput(2023, 18))
	_ = input
	_ = `
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`
	m := drawMaze(1, input)
	result = m.solve()
	return result
}
func part2() (result int) {
	input := string(utils.PuzzleInput(2023, 18))
	_ = input
	input = `
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`
	m := drawMaze(2, input)
	result = m.solve()
	return result
}

func main() {
	//	fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("Solution: %d\n", part2())

}
