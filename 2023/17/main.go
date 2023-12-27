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

	dist int
	prev *vertex

	in bool
}
type graph []*vertex
type grid struct {
	vertices [][]*vertex
	graph    graph

	xSize int
	ySize int
}

func importGrid(input string) grid {
	g := grid{}
	q := graph{}
	for x, row := range strings.Split(strings.Trim(input, "\n"), "\n") {
		g.vertices = append(g.vertices, []*vertex{})
		for y, c := range strings.Split(row, "") {
			value, _ := strconv.Atoi(c)
			v := vertex{
				x:    x,
				y:    y,
				v:    value,
				dist: math.MaxInt,
			}
			g.vertices[x] = append(g.vertices[x], &v)
			q = append(q, &v)
		}
	}
	g.xSize = len(g.vertices)
	g.ySize = len(g.vertices[0])
	g.graph = q

	return g
}

func (g *grid) getNeighbors(v *vertex) (vertices []*vertex) {
	if v.x > 0 {
		vertices = append(vertices, g.vertices[v.x-1][v.y]) // left
	}
	if v.y > 0 {
		vertices = append(vertices, g.vertices[v.x][v.y-1]) // up
	}
	if v.x+1 < g.xSize {
		vertices = append(vertices, g.vertices[v.x+1][v.y]) // right
	}
	if v.y+1 < g.ySize {
		vertices = append(vertices, g.vertices[v.x][v.y+1]) // down
	}
	return vertices
}

func (g *grid) minVertex(q []*vertex) (mn *vertex) {
	min := math.MaxInt
	for _, v := range q {
		if v.dist < min {
			min = v.dist
			mn = v
		}
	}
	return mn
}

func (g *grid) updateWeight(v1 *vertex, v2 *vertex) {
	if v2.dist > v1.dist+v1.v {
		v2.dist = v1.dist + v1.v
		v2.prev = v1
	}
}

func (g *grid) shortestPath() int {
	q := make(graph, len(g.graph))
	copy(q, g.graph)

	start := g.vertices[0][0]
	start.dist = 0
	for len(q) != 0 {
		v1 := g.minVertex(q)
		for n, v := range q {
			if v == v1 {
				q = slices.Delete(q, n, n+1)
				break
			}
		}
		for _, v2 := range g.getNeighbors(v1) {
			g.updateWeight(v1, v2)
		}
	}

	s := graph{}
	cursor := g.vertices[g.xSize-1][g.ySize-1]

	for start != cursor {
		s = append(s, cursor)
		cursor.in = true
		cursor = cursor.prev
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
	_ = `
123
456
789
`
	grid := importGrid(input)
	result = grid.shortestPath()
	grid.print()
	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part0())
}
