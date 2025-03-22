package main

import (
	"fmt"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type puzzle map[int]map[int]rune

const (
	M = 77
	S = 83
	X = 88
	A = 65
)

func parse() puzzle {
	input := string(utils.PuzzleInput(2024, 4))
	p := puzzle{}
	for xx, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		p[xx] = make(map[int]rune)
		for yy, char := range line {
			p[xx][yy] = char
		}
	}

	return p
}
func part1() string {
	solutions := 0
	p := parse()

	for xx := range p {
		for yy := range p[xx] {
			if p[xx][yy] != X {
				continue
			}

			if p[xx][yy+1] == M && p[xx][yy+2] == A && p[xx][yy+3] == S {
				solutions++
			}

			if p[xx][yy-1] == M && p[xx][yy-2] == A && p[xx][yy-3] == S {
				solutions++
			}

			if p[xx+1][yy+1] == M && p[xx+2][yy+2] == A && p[xx+3][yy+3] == S {
				solutions++
			}

			if p[xx-1][yy-1] == M && p[xx-2][yy-2] == A && p[xx-3][yy-3] == S {
				solutions++
			}

			if p[xx-1][yy+1] == M && p[xx-2][yy+2] == A && p[xx-3][yy+3] == S {
				solutions++
			}

			if p[xx+1][yy-1] == M && p[xx+2][yy-2] == A && p[xx+3][yy-3] == S {
				solutions++
			}

			if p[xx+1][yy] == M && p[xx+2][yy] == A && p[xx+3][yy] == S {
				solutions++
			}

			if p[xx-1][yy] == M && p[xx-2][yy] == A && p[xx-3][yy] == S {
				solutions++
			}

		}
	}

	return fmt.Sprintf("%d\n", solutions)
}

func part2() string {
	solutions := 0
	p := parse()

	for xx := range p {
		for yy := range p[xx] {
			if p[xx][yy] != A {
				continue
			}

			if p[xx-1][yy-1] == M && p[xx+1][yy+1] == S && p[xx-1][yy+1] == S && p[xx+1][yy-1] == M {
				solutions++
			}
			if p[xx-1][yy-1] == S && p[xx+1][yy+1] == M && p[xx-1][yy+1] == M && p[xx+1][yy-1] == S {
				solutions++
			}
			if p[xx-1][yy-1] == S && p[xx+1][yy+1] == M && p[xx-1][yy+1] == S && p[xx+1][yy-1] == M {
				solutions++
			}
			if p[xx-1][yy-1] == M && p[xx+1][yy+1] == S && p[xx-1][yy+1] == M && p[xx+1][yy-1] == S {
				solutions++
			}
		}
	}

	return fmt.Sprintf("%d\n", solutions)
}

func main() {
	fmt.Printf(part1())
	fmt.Printf(part2())
}
