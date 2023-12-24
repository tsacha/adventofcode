package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/tsacha/adventofcode/utils"
)

const (
	Left  = 0
	Down  = 1
	Right = 2
	Up    = 3

	Void        = "."
	MirrorLeft  = "/"
	MirrorRight = "\\"
	SplitterH   = "-"
	SplitterV   = "|"
)

type beam struct {
	xCursor   int
	yCursor   int
	direction int
}

type tile struct {
	char      string
	energised bool
}

type grid struct {
	xSize    int
	ySize    int
	position [][]tile
	history  [][3]int
}

func (b *beam) left(g *grid) {
	if b.yCursor > 0 {
		b.yCursor--
		b.direction = Left
		g.piewPiew(*b)
	}
}

func (b *beam) right(g *grid) {
	if b.yCursor < g.ySize-1 {
		b.yCursor++
		b.direction = Right
		g.piewPiew(*b)
	}
}

func (b *beam) up(g *grid) {
	if b.xCursor > 0 {
		b.xCursor--
		b.direction = Up
		g.piewPiew(*b)
	}
}

func (b *beam) down(g *grid) {
	if b.xCursor < g.xSize-1 {
		b.xCursor++
		b.direction = Down
		g.piewPiew(*b)
	}
}

func (b *beam) splitH(g *grid) {
	b.left(g)
	b.right(g)
}

func (b *beam) splitV(g *grid) {
	b.up(g)
	b.down(g)
}

func (g *grid) piewPiew(b beam) {
	for _, h := range g.history[:len(g.history)] {
		if b.xCursor == h[0] && b.yCursor == h[1] && b.direction == h[2] {
			return
		}
	}
	g.history = append(g.history, [3]int{b.xCursor, b.yCursor, b.direction})
	tile := &g.position[b.xCursor][b.yCursor]
	tile.energised = true

	switch b.direction {
	case Left:
		switch tile.char {
		case Void, SplitterH:
			b.left(g)
		case MirrorLeft:
			b.down(g)
		case MirrorRight:
			b.up(g)
		case SplitterV:
			b.splitV(g)
		}
	case Up:
		switch tile.char {
		case Void, SplitterV:
			b.up(g)
		case MirrorLeft:
			b.right(g)
		case MirrorRight:
			b.left(g)
		case SplitterH:
			b.splitH(g)
		}
	case Right:
		switch tile.char {
		case Void, SplitterH:
			b.right(g)
		case MirrorLeft:
			b.up(g)
		case MirrorRight:
			b.down(g)
		case SplitterV:
			b.splitV(g)
		}
	case Down:
		switch tile.char {
		case Void, SplitterV:
			b.down(g)
		case MirrorLeft:
			b.left(g)
		case MirrorRight:
			b.right(g)
		case SplitterH:
			b.splitH(g)
		}
	}
}

func (g *grid) print() {
	energised := color.New(color.Bold, color.BgBlue)
	empty := color.New(color.FgRed, color.Bold)

	for x := range g.position {
		fmt.Println()
		for _, c := range g.position[x] {
			if c.energised {
				_, _ = energised.Printf("%s ", c.char)
			} else {
				_, _ = empty.Printf("%s ", c.char)
			}
		}
	}
	fmt.Println()
}

func (g *grid) count() (solution int) {
	for x := range g.position {
		for _, c := range g.position[x] {
			if c.energised {
				solution++
			}
		}
	}
	return solution
}

func importGrid(input string) (grid grid) {
	for x, row := range strings.Split(strings.Trim(input, "\n"), "\n") {
		grid.position = append(grid.position, []tile{})
		for _, c := range strings.Split(row, "") {
			tile := tile{
				char: c,
			}
			grid.position[x] = append(grid.position[x], tile)
		}
	}
	grid.xSize = len(grid.position)
	grid.ySize = len(grid.position[0])

	return grid
}

func part1() (result int) {
	input := string(utils.PuzzleInput(2023, 16))
	grid := importGrid(input)

	grid.piewPiew(beam{
		xCursor:   0,
		yCursor:   0,
		direction: Right,
	})

	grid.print()
	return grid.count()
}

func part2() (result int) {
	input := string(utils.PuzzleInput(2023, 16))
	inputGrid := importGrid(input)

	var wg sync.WaitGroup
	for x := 0; x < inputGrid.xSize; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			g := importGrid(input)
			g.piewPiew(beam{
				xCursor:   x,
				yCursor:   0,
				direction: Right,
			})
			result = max(g.count(), result)
		}(x)
	}

	for x := 0; x < inputGrid.xSize; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			g := importGrid(input)
			g.piewPiew(beam{
				xCursor:   x,
				yCursor:   g.ySize - 1,
				direction: Left,
			})
			result = max(g.count(), result)
		}(x)
	}

	for y := 0; y < inputGrid.ySize; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			g := importGrid(input)
			g.piewPiew(beam{
				xCursor:   0,
				yCursor:   y,
				direction: Down,
			})
			result = max(g.count(), result)
		}(y)
	}

	for y := 0; y < inputGrid.ySize; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			g := importGrid(input)
			g.piewPiew(beam{
				xCursor:   g.xSize - 1,
				yCursor:   y,
				direction: Up,
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
