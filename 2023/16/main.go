package main

import (
	"fmt"
	"strings"
	"sync"

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

func (g *grid) piewpiew(b beam) {
	for _, h := range g.history[:len(g.history)] {
		if b.x_cursor == h[0] && b.y_cursor == h[1] && b.direction == h[2] {
			return
		}
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

func import_grid(input string) (grid grid) {
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

	return grid
}

func part1() (result int) {
	input := string(utils.PuzzleInput(2023, 16))
	grid := import_grid(input)

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
	input_grid := import_grid(input)

	var wg sync.WaitGroup
	for x := 0; x < input_grid.x_size; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			g := import_grid(input)
			g.piewpiew(beam{
				x_cursor:  x,
				y_cursor:  0,
				direction: RIGHT,
			})
			result = max(g.count(), result)
		}(x)
	}

	for x := 0; x < input_grid.x_size; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			g := import_grid(input)
			g.piewpiew(beam{
				x_cursor:  x,
				y_cursor:  g.y_size - 1,
				direction: LEFT,
			})
			result = max(g.count(), result)
		}(x)
	}

	for y := 0; y < input_grid.y_size; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			g := import_grid(input)
			g.piewpiew(beam{
				x_cursor:  0,
				y_cursor:  y,
				direction: DOWN,
			})
			result = max(g.count(), result)
		}(y)
	}

	for y := 0; y < input_grid.y_size; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			g := import_grid(input)
			g.piewpiew(beam{
				x_cursor:  g.x_size - 1,
				y_cursor:  y,
				direction: UP,
			})
			result = max(g.count(), result)
		}(y)
	}
	wg.Wait()

	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part2())

}
