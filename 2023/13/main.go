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

func find_reflections(line []line, i int, j int, ret int, reverse bool, mult int, skip_result int, is_child bool) (reflections int, err error) {
	reflections = ret
	for x := j; x > i; x-- {
		if reflect.DeepEqual(line[i], line[x]) {
			if x-i == 1 {
				var result int
				if !reverse {
					result = mult * x
				} else {
					result = mult * (len(line) - x)
				}

				if skip_result != result {
					return result, nil
				} else {
					skip_result = -1
				}
			}
			reflections, _ = find_reflections(line, i+1, x-1, reflections, reverse, mult, skip_result, true)
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

func (p *puzzle) solve_puzzle(skip_result int) (s int) {
	var err error

	s, err = find_reflections(p.vertical, 0, len(p.vertical)-1, -1, false, 1, skip_result, false)
	if err == nil {
		return s
	}
	slices.Reverse(p.vertical)
	s, err = find_reflections(p.vertical, 0, len(p.vertical)-1, -1, true, 1, skip_result, false)
	slices.Reverse(p.vertical)
	if err == nil {
		return s
	}

	s, err = find_reflections(p.horizontal, 0, len(p.horizontal)-1, -1, false, 100, skip_result, false)
	if err == nil {
		return s
	}
	slices.Reverse(p.horizontal)
	s, err = find_reflections(p.horizontal, 0, len(p.horizontal)-1, -1, true, 100, skip_result, false)
	slices.Reverse(p.horizontal)
	if err == nil {
		return s
	}

	return 0
}

func (p *puzzle) solve_smudge_puzzle() (s int) {
	initial_solve := p.solve_puzzle(-1)
	for ph := 0; ph < len(p.horizontal); ph++ {
		for pv := 0; pv < len(p.horizontal[0]); pv++ {
			s := puzzle{}
			s.horizontal = []line{}
			s.vertical = []line{}

			for h := 0; h < len(p.horizontal); h++ {
				s.horizontal = append(s.horizontal, line{})
				for v := 0; v < len(p.horizontal[0]); v++ {
					s.horizontal[h] = append(s.horizontal[h], p.horizontal[h][v])
				}
			}

			if s.horizontal[ph][pv] == ASH {
				s.horizontal[ph][pv] = ROCK
			} else if s.horizontal[ph][pv] == ROCK {
				s.horizontal[ph][pv] = ASH
			}

			for v := 0; v < len(s.horizontal[0]); v++ {
				s.vertical = append(s.vertical, line{})
				for h := 0; h < len(s.horizontal); h++ {
					s.vertical[v] = append(s.vertical[v], s.horizontal[h][v])
				}
			}

			solve := s.solve_puzzle(initial_solve)
			if solve > 0 {
				fmt.Printf("[%d:%d] %d (was %d)\n", ph, pv, solve, initial_solve)
				return solve
			}
		}
	}

	return initial_solve
}

func part(part int) (solution int) {
	solution = 0

	input := string(utils.PuzzleInput(2023, 13))

	// examples
	_ = `#.##..##.
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

	for _, puzzle_str := range strings.Split(strings.Trim(input, "\n"), "\n\n") {
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

		if part == 1 {
			p_sol := p.solve_puzzle(-1)
			solution += p_sol
		} else if part == 2 {
			p_sol := p.solve_smudge_puzzle()
			solution += p_sol
		}
	}
	return solution
}

func main() {
	//fmt.Printf("Solution 1: %d\n", part(1))
	fmt.Printf("Solution: %d\n", part(2))
}
