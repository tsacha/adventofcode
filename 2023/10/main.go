package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/fatih/color"
	"github.com/tsacha/adventofcode/utils"
)

func beautiful_maze(maze maze, path []position, insides []position) {
	_ = insides
	x_max := 0
	y_max := 0
	for m := range maze {
		if m.x > x_max {
			x_max = m.x
		}
		if m.y > y_max {
			y_max = m.y
		}
	}

	in_path := color.New(color.FgMagenta, color.Bold)
	in_inside := color.New(color.FgRed, color.Bold)

	for x := 0; x <= x_max; x++ {
		fmt.Printf("\n")
		for y := 0; y <= y_max; y++ {
			p := position{x: x, y: y}
			var char string
			switch maze[p] {
			case connections{0, 0, 0, 0}:
				char = "·"
			case connections{1, 1, 0, 0}:
				char = "║"
			case connections{0, 0, 1, 1}:
				char = "═"
			case connections{1, 0, 0, 1}:
				char = "╚"
			case connections{1, 0, 1, 0}:
				char = "╝"
			case connections{0, 1, 1, 0}:
				char = "╗"
			case connections{0, 1, 0, 1}:
				char = "╔"
			case connections{1, 1, 1, 1}:
				char = "s"
			default:
				char = "·"
			}
			if slices.Contains(path, p) {
				in_path.Printf("%s", char)
			} else if slices.Contains(insides, p) {
				in_inside.Printf("%s", char)
			} else {
				fmt.Printf("%s", char)
			}

		}
	}
	fmt.Print("\n")

}

type position struct {
	x int
	y int
}
type connections [4]int // [U, D, L, R]
type symbols map[string]connections
type maze map[position]connections

func part1() (solution int) {
	input := string(utils.PuzzleInput(2023, 10))
	_ = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	//	beautiful_maze(input)
	maze := make(maze)
	var cursor position

	symbols := make(symbols) // [U, D, L, R]
	symbols["|"] = connections{1, 1, 0, 0}
	symbols["-"] = connections{0, 0, 1, 1}
	symbols["L"] = connections{1, 0, 0, 1}
	symbols["J"] = connections{1, 0, 1, 0}
	symbols["7"] = connections{0, 1, 1, 0}
	symbols["F"] = connections{0, 1, 0, 1}
	symbols["."] = connections{0, 0, 0, 0}
	symbols["S"] = connections{1, 1, 1, 1}

	for x, line := range strings.Split(input, "\n") {
		for y, char := range strings.Split(line, "") {
			if char == "S" {
				cursor = position{x, y}
			}
			maze[position{
				x: x,
				y: y,
			}] = symbols[char]
		}

	}

	cursors := []position{}
	for 1 == 1 {
		cursors = append(cursors, cursor)

		n := position{x: cursor.x - 1, y: cursor.y}
		s := position{x: cursor.x + 1, y: cursor.y}
		w := position{x: cursor.x, y: cursor.y - 1}
		e := position{x: cursor.x, y: cursor.y + 1}

		here := maze[cursor]
		north := maze[n]
		south := maze[s]
		west := maze[w]
		east := maze[e]

		if here[0] == 1 && north[1] == 1 && !slices.Contains(cursors, n) {
			cursor = n
		} else if here[1] == 1 && south[0] == 1 && !slices.Contains(cursors, s) {
			cursor = s
		} else if here[2] == 1 && west[3] == 1 && !slices.Contains(cursors, w) {
			cursor = w
		} else if here[3] == 1 && east[2] == 1 && !slices.Contains(cursors, e) {
			cursor = e
		} else {
			break
		}
	}

	return len(cursors) / 2
}

func part2() (solution int) {
	input := string(utils.PuzzleInput(2023, 10))
	_ = `.....
.S-7.
.|.|.
.L-J.
.....`

	_ = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

	_ = `
......................
..F----7F7F7F7F-7.....
..|F--7||||||||FJ.....
..||.FJ||||||||L7.....
.FJL7L7LJLJ||LJ.L-7...
.L--J.L7...LJS7F-7L7..
.....F-J..F7FJ|L7L7L7.
.....L7.F7||L7|.L7L7|.
......|FJLJ|FJ|F7|.LJ.
.....FJL-7.||.||||....
.....L---J.LJ.LJLJ....
......................
`

	maze := make(maze)
	var cursor position

	symbols := make(symbols) // [U, D, L, R]
	symbols["|"] = connections{1, 1, 0, 0}
	symbols["-"] = connections{0, 0, 1, 1}
	symbols["L"] = connections{1, 0, 0, 1}
	symbols["J"] = connections{1, 0, 1, 0}
	symbols["7"] = connections{0, 1, 1, 0}
	symbols["F"] = connections{0, 1, 0, 1}
	symbols["."] = connections{0, 0, 0, 0}
	symbols["S"] = connections{1, 1, 1, 1}

	for x, line := range strings.Split(input, "\n") {
		for y, char := range strings.Split(line, "") {
			if char == "S" {
				cursor = position{x, y}
			}
			maze[position{
				x: x,
				y: y,
			}] = symbols[char]
		}

	}

	path := []position{}
	for 1 == 1 {
		path = append(path, cursor)

		n := position{x: cursor.x - 1, y: cursor.y}
		s := position{x: cursor.x + 1, y: cursor.y}
		w := position{x: cursor.x, y: cursor.y - 1}
		e := position{x: cursor.x, y: cursor.y + 1}

		here := maze[cursor]
		north := maze[n]
		south := maze[s]
		west := maze[w]
		east := maze[e]

		if here[0] == 1 && north[1] == 1 && !slices.Contains(path, n) {
			cursor = n
		} else if here[1] == 1 && south[0] == 1 && !slices.Contains(path, s) {
			cursor = s
		} else if here[2] == 1 && west[3] == 1 && !slices.Contains(path, w) {
			cursor = w
		} else if here[3] == 1 && east[2] == 1 && !slices.Contains(path, e) {
			cursor = e
		} else {
			break
		}
	}

	x_max := 0
	y_max := 0
	for m := range maze {
		if m.x > x_max {
			x_max = m.x
		}
		if m.y > y_max {
			y_max = m.y
		}
	}

	// https://www.gorillasun.de/blog/an-algorithm-for-polygon-intersections/
	insides := []position{}
	for x := 0; x <= x_max; x++ {
		for y := 0; y <= y_max; y++ {
			p := position{x: x, y: y}

			if slices.Contains(path, p) {
				continue
			}
			inside := false
			loop := true
			i := 0
			j := len(path) - 1
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

	beautiful_maze(maze, path, insides)

	return len(insides)
}

func main() {
	//fmt.Printf("\n\nSolution 1: %d\n", part1())
	fmt.Printf("\n\nSolution 2: %d\n", part2())
}
