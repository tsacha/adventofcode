package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/tsacha/adventofcode/utils"
)

const (
	LEFT  = 0
	DOWN  = 1
	RIGHT = 2
	UP    = 3

	VOID         = "."
	MIRROR_LEFT  = "/"
	MIRROR_RIGHT = "\\"
	SPLITTER_H   = "-"
	SPLITTER_V   = "|"
)

type beam struct {
	x_cursor  int
	y_cursor  int
	direction int
}

type tile struct {
	char      string
	energised bool
}

type grid struct {
	x_size   int
	y_size   int
	position [][]tile
	history  [][3]int
}

func (b *beam) left(g *grid) {
	if b.y_cursor > 0 {
		b.y_cursor--
		b.direction = LEFT
		g.piewpiew(*b)
	}
}
func (b *beam) right(g *grid) {
	if b.y_cursor < g.y_size-1 {
		b.y_cursor++
		b.direction = RIGHT
		g.piewpiew(*b)
	}
}
func (b *beam) up(g *grid) {
	if b.x_cursor > 0 {
		b.x_cursor--
		b.direction = UP
		g.piewpiew(*b)
	}
}
func (b *beam) down(g *grid) {
	if b.x_cursor < g.x_size-1 {
		b.x_cursor++
		b.direction = DOWN
		g.piewpiew(*b)
	}
}

func (b *beam) split_h(g *grid) {
	b.left(g)
	b.right(g)
}

func (b *beam) split_v(g *grid) {
	b.up(g)
	b.down(g)
}

func (b *beam) stop_the_beam(g *grid) bool {
	for _, h := range g.history[:len(g.history)] {
		if b.x_cursor == h[0] && b.y_cursor == h[1] && b.direction == h[2] {
			return true
		}
	}

	return false
}

func (g *grid) piewpiew(b beam) {
	if b.stop_the_beam(g) {
		return
	}

	g.history = append(g.history, [3]int{b.x_cursor, b.y_cursor, b.direction})
	tile := &g.position[b.x_cursor][b.y_cursor]
	tile.energised = true

	switch b.direction {
	case LEFT:
		switch tile.char {
		case VOID, SPLITTER_H:
			b.left(g)
		case MIRROR_LEFT:
			b.down(g)
		case MIRROR_RIGHT:
			b.up(g)
		case SPLITTER_V:
			b.split_v(g)
		}
	case UP:
		switch tile.char {
		case VOID, SPLITTER_V:
			b.up(g)
		case MIRROR_LEFT:
			b.right(g)
		case MIRROR_RIGHT:
			b.left(g)
		case SPLITTER_H:
			b.split_h(g)
		}
	case RIGHT:
		switch tile.char {
		case VOID, SPLITTER_H:
			b.right(g)
		case MIRROR_LEFT:
			b.up(g)
		case MIRROR_RIGHT:
			b.down(g)
		case SPLITTER_V:
			b.split_v(g)
		}
	case DOWN:
		switch tile.char {
		case VOID, SPLITTER_V:
			b.down(g)
		case MIRROR_LEFT:
			b.left(g)
		case MIRROR_RIGHT:
			b.right(g)
		case SPLITTER_H:
			b.split_h(g)
		}
	}
}

func (grid grid) print() {
	energised := color.New(color.Bold, color.BgBlue)
	empty := color.New(color.FgRed, color.Bold)

	for x := range grid.position {
		fmt.Println()
		for _, c := range grid.position[x] {
			if c.energised {
				energised.Printf("%s ", c.char)
			} else {
				empty.Printf("%s ", c.char)
			}
		}
	}
	fmt.Println()
}

func (grid *grid) clear() {
	grid.history = [][3]int{}
	for x := range grid.position {
		for y := range grid.position[x] {
			grid.position[x][y].energised = false
		}
	}
}

func (grid grid) count() (solution int) {
	for x := range grid.position {
		for _, c := range grid.position[x] {
			if c.energised {
				solution++
			}
		}
	}
	return solution
}

func part1() (result int) {
	input := string(utils.PuzzleInput(2023, 16))

	grid := grid{
		position: [][]tile{},
	}
	for x, row := range strings.Split(strings.Trim(input, "\n"), "\n") {
		grid.position = append(grid.position, []tile{})
		for _, c := range strings.Split(row, "") {
			tile := tile{
				char: c,
			}
			grid.position[x] = append(grid.position[x], tile)
		}
	}
	grid.x_size = len(grid.position)
	grid.y_size = len(grid.position[0])

	grid.piewpiew(beam{
		x_cursor:  0,
		y_cursor:  0,
		direction: RIGHT,
	})

	grid.print()
	return grid.count()
}

func part2() (result int) {
	input := string(utils.PuzzleInput(2023, 16))

	grid := grid{
		position: [][]tile{},
	}
	for x, row := range strings.Split(strings.Trim(input, "\n"), "\n") {
		grid.position = append(grid.position, []tile{})
		for _, c := range strings.Split(row, "") {
			tile := tile{
				char: c,
			}
			grid.position[x] = append(grid.position[x], tile)
		}
	}
	grid.x_size = len(grid.position)
	grid.y_size = len(grid.position[0])

	for x := 0; x < grid.x_size; x++ {
		grid.piewpiew(beam{
			x_cursor:  x,
			y_cursor:  0,
			direction: RIGHT,
		})
		result = max(grid.count(), result)
		grid.clear()
	}

	for x := 0; x < grid.x_size; x++ {
		grid.piewpiew(beam{
			x_cursor:  x,
			y_cursor:  grid.y_size - 1,
			direction: LEFT,
		})
		result = max(grid.count(), result)
		grid.clear()
	}

	for y := 0; y < grid.y_size; y++ {
		grid.piewpiew(beam{
			x_cursor:  0,
			y_cursor:  y,
			direction: DOWN,
		})
		result = max(grid.count(), result)
		grid.clear()
	}

	for y := 0; y < grid.y_size; y++ {
		grid.piewpiew(beam{
			x_cursor:  grid.x_size - 1,
			y_cursor:  y,
			direction: UP,
		})
		result = max(grid.count(), result)
		grid.clear()
	}

	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part2())

}
