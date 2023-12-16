package main

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/tsacha/adventofcode/utils"
)

const (
	ASH  = 0
	ROCK = 1

	ASH_STR  = "."
	ROCK_STR = "#"
)

type line []int
type puzzle struct {
	horizontal []line
	vertical   []line
}

func find_reflections(line []line, i int, j int, ret int, is_child bool) (reflections int, err error) {
	reflections = ret
	for x := j; x > i; x-- {
		if reflect.DeepEqual(line[i], line[x]) {
			if x-i == 1 {
				return x, nil
			} else {
				reflections, _ = find_reflections(line, i+1, x-1, reflections, true)
			}
		} else if is_child {
			return reflections, nil
		}
	}

	if reflections == -1 {
		return -1, errors.New("no reflection")
	} else {
		return reflections, nil
	}
}

func (p *puzzle) solve_puzzle() (s int) {
	var err error
	r := 0

	r, err = find_reflections(p.vertical, 0, len(p.vertical)-1, -1, false)
	if err == nil {
		s += r
	}
	slices.Reverse(p.vertical)
	r, err = find_reflections(p.vertical, 0, len(p.vertical)-1, -1, false)
	if err == nil {
		s += len(p.vertical) - r
	}

	r, err = find_reflections(p.horizontal, 0, len(p.horizontal)-1, -1, false)
	if err == nil {
		s += r * 100
	}
	slices.Reverse(p.horizontal)
	r, err = find_reflections(p.horizontal, 0, len(p.horizontal)-1, -1, false)
	if err == nil {
		s += (len(p.horizontal) - r) * 100
	}

	return s
}

func part1() (solution int) {
	solution = 0
	input := string(utils.PuzzleInput(2023, 13))

	// example
	_ = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	for puzzle_nb, puzzle_str := range strings.Split(strings.Trim(input, "\n"), "\n\n") {
		p := puzzle{}
		p.horizontal = []line{}
		p.vertical = []line{}

		for n, line_str := range strings.Split(puzzle_str, "\n") {
			if line_str == "" {
				continue
			}

			p.horizontal = append(p.horizontal, line{})
			for _, char := range line_str {
				ash, _ := utf8.DecodeLastRuneInString(ASH_STR)
				rock, _ := utf8.DecodeLastRuneInString(ROCK_STR)

				switch char {
				case ash:
					p.horizontal[n] = append(p.horizontal[n], ASH)
				case rock:
					p.horizontal[n] = append(p.horizontal[n], ROCK)
				}
			}
		}

		for v := 0; v < len(p.horizontal[0]); v++ {
			p.vertical = append(p.vertical, line{})
			for h := 0; h < len(p.horizontal); h++ {
				p.vertical[v] = append(p.vertical[v], p.horizontal[h][v])
			}
		}

		p_sol := p.solve_puzzle()
		fmt.Printf("[%d] %d\n", puzzle_nb, p_sol)
		solution += p_sol
	}
	return solution
}

func main() {
	fmt.Printf("Solution 1: %d\n", part1())
}
