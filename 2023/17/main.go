package main

import (
	"container/heap"
	"fmt"
	"math"
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
	seen bool

	in bool
}

type graph []*vertex
type grid struct {
	vertices [][]*vertex
	graph    graph

	xSize int
	ySize int
}

func (h graph) Len() int           { return len(h) }
func (h graph) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h graph) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *graph) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*vertex))
}
func (h *graph) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func (g *grid) getNeighbors(v *vertex) (vertices graph) {
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

func (g *grid) shortestPath() int {
	start := g.vertices[0][0]
	start.dist = 0
	q := &graph{start}
	heap.Init(q)
	for q.Len() > 0 {
		c := heap.Pop(q).(*vertex)
		if c.seen {
			continue
		}
		c.seen = true
		for _, n := range g.getNeighbors(c) {
			distance := c.dist + c.v
			if distance < n.dist {
				n.dist = distance
				n.prev = c
				heap.Push(q, n)
			}
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
			if v.in || (v.x == 0 && v.y == 0) {
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
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`

	grid := importGrid(input)
	for i := 0; i < 1; i++ {
		result = grid.shortestPath()
	}
	grid.print()
	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part0())
}
