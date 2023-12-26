package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/tsacha/adventofcode/utils"
)

type vertex struct {
	x int
	y int
	v int

	w int
	p *vertex

	in bool
}
type grid struct {
	vertices [][]vertex

	xSize int
	ySize int
}

func importGrid(input string) grid {
	g := grid{}
	for x, row := range strings.Split(strings.Trim(input, "\n"), "\n") {
		g.vertices = append(g.vertices, []vertex{})
		for y, c := range strings.Split(row, "") {
			v, _ := strconv.Atoi(c)
			g.vertices[x] = append(g.vertices[x], vertex{
				x: x,
				y: y,
				w: math.MaxInt,
				v: v,
			})
		}
	}
	g.xSize = len(g.vertices)
	g.ySize = len(g.vertices[0])

	return g
}

func (g *grid) getNeighbors(v *vertex) (vertices []*vertex) {
	if v.x > 0 {
		vertices = append(vertices, &g.vertices[v.x-1][v.y]) // left
	}
	if v.y > 0 {
		vertices = append(vertices, &g.vertices[v.x][v.y-1]) // up
	}
	if v.x+1 < g.xSize {
		vertices = append(vertices, &g.vertices[v.x+1][v.y]) // right
	}
	if v.y+1 < g.ySize {
		vertices = append(vertices, &g.vertices[v.x][v.y+1]) // down
	}
	return vertices
}

func (g *grid) minNeighbor(q []*vertex) (mn *vertex) {
	min := math.MaxInt
	for _, v := range q {
		if v.w < min {
			min = v.w
			mn = v
		}
	}
	return mn
}

func (g *grid) updateDistance(v1 *vertex, v2 *vertex) {
	if v2.w > v1.w+v1.v {
		v2.w = v1.w + v1.v
		v2.p = v1
	}
}

func (g *grid) shortestPath() int {
	q := []*vertex{}
	start := &g.vertices[0][0]
	start.w = 0

	for x := range g.vertices {
		for y := range g.vertices[x] {
			q = append(q, &g.vertices[x][y])
		}
	}

	for len(q) != 0 {
		v1 := g.minNeighbor(q)
		for n, v := range q {
			if v == v1 {
				q = slices.Delete(q, n, n+1)
				break
			}
		}
		for _, v2 := range g.getNeighbors(v1) {
			g.updateDistance(v1, v2)
		}
	}

	s := []*vertex{}
	cursor := &g.vertices[g.xSize-1][g.ySize-1]

	for start != cursor {
		s = append(s, cursor)
		cursor.in = true
		cursor = cursor.p
	}

	sum := 0
	for _, p := range s {
		sum += p.v
	}

	return sum
}

func (g *grid) print() {
	energised := color.New(color.Bold, color.BgGreen)
	empty := color.New(color.FgRed, color.Bold)
	for x := range g.vertices {
		fmt.Println()
		for _, v := range g.vertices[x] {
			if v.in {
				_, _ = energised.Printf("%d", v.v)

			} else {
				_, _ = empty.Printf("%d", v.v)
			}
		}
	}
	fmt.Println()
}

func part0() (result int) {
	input := string(utils.PuzzleInput(2023, 17))
	grid := importGrid(input)
	result = grid.shortestPath()
	grid.print()
	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part0())
}
