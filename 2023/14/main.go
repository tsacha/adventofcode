package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/tsacha/adventofcode/utils"
)

const (
	NORTH = 0
	WEST  = 1
	SOUTH = 2
	EAST  = 3

	VOID       = 0
	CUBE_ROCK  = 1
	ROUND_ROCK = 2

	VOID_STR       = "."
	ROUND_ROCK_STR = "O"
	CUBE_ROCK_STR  = "#"
)

type line []int
type round_rocks map[int][2]int
type puzzle struct {
	lines       []line
	round_rocks round_rocks
}

func (p *puzzle) print_puzzle() {
	fmt.Print("\n")
	for v := range p.lines {
		for h := range p.lines[v] {
			switch p.lines[v][h] {
			case VOID:
				fmt.Printf(VOID_STR)
			case ROUND_ROCK:
				fmt.Printf(ROUND_ROCK_STR)
			case CUBE_ROCK:
				fmt.Printf(CUBE_ROCK_STR)
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func (p *puzzle) solve_puzzle() {
	for n := 1; n <= len(p.lines); n++ {
		for nrock, rock := range p.round_rocks {
			if rock[0] == 0 {
				continue
			}

			neighbor := &p.lines[rock[0]-1][rock[1]]
			cursor := &p.lines[rock[0]][rock[1]]
			if *neighbor == VOID {
				*neighbor = ROUND_ROCK
				*cursor = VOID
				p.round_rocks[nrock] = [2]int{rock[0] - 1, rock[1]}
			}
		}
	}
	p.print_puzzle()
}

func copy_puzzle(input []line) (output []line) {
	for v := range input {
		output = append(output, line{})
		output[v] = append(output[v], input[v]...)
	}
	return output
}

func same_puzzle(a []line, b []line) bool {
	is_equal := true
	for v := range a {
		for h := range a[v] {
			if a[v][h] != b[v][h] {
				is_equal = false
			}
		}
	}
	return is_equal
}

func (p *puzzle) solve2_puzzle() int {
	puzzles := [][]line{
		copy_puzzle(p.lines),
	}

	loop := false
	last_counts := []int{p.count_puzzle()}
	offset := 0
	neighbors := round_rocks{
		NORTH: {-1, 0},
		WEST:  {0, -1},
		SOUTH: {1, 0},
		EAST:  {0, 1},
	}

	direction := NORTH
	for c := 0; c <= (CYCLES)*4-1; c++ {
		direction = c % 4
		moved := true
		for moved {
			moved = false
			for nrock, rock := range p.round_rocks {
				switch direction {
				case NORTH:
					if rock[0] == 0 {
						continue
					}
				case WEST:
					if rock[1] == 0 {
						continue
					}
				case SOUTH:
					if rock[0] == len(p.lines)-1 {
						continue
					}
				case EAST:
					if rock[1] == len(p.lines[0])-1 {
						continue
					}
				}

				neighbor := &p.lines[rock[0]+neighbors[direction][0]][rock[1]+neighbors[direction][1]]
				cursor := &p.lines[rock[0]][rock[1]]
				if *neighbor == VOID {
					*neighbor = ROUND_ROCK
					*cursor = VOID
					p.round_rocks[nrock] = [2]int{rock[0] + neighbors[direction][0], rock[1] + neighbors[direction][1]}
					moved = true
				}
			}
		}

		if direction != EAST {
			continue
		}

		for np, puzzle := range puzzles {
			if same_puzzle(p.lines, puzzle) {
				last_counts = last_counts[np:]
				offset = np
				loop = true
				puzzles = [][]line{copy_puzzle(p.lines)}
				break
			}
		}
		if !loop {
			last_counts = append(last_counts, p.count_puzzle())
			puzzles = append(puzzles, copy_puzzle(p.lines))
		} else {
			break
		}

	}

	if loop {
		return last_counts[(CYCLES-offset)%len(last_counts)]
	} else {
		return p.count_puzzle()
	}
}

const CYCLES = 1000000000

func (p *puzzle) count_puzzle() (solution int) {
	for _, r := range p.round_rocks {
		solution += len(p.lines) - r[0]
	}

	return solution
}

func part(part int) (solution int) {
	solution = 0

	input := string(utils.PuzzleInput(2023, 14))
	// examples
	_ = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	p := puzzle{
		lines:       []line{},
		round_rocks: make(round_rocks),
	}
	for v, puzzle_str := range strings.Split(strings.Trim(input, "\n"), "\n") {
		p.lines = append(p.lines, line{})
		for h, char := range puzzle_str {
			void, _ := utf8.DecodeLastRuneInString(VOID_STR)
			round, _ := utf8.DecodeLastRuneInString(ROUND_ROCK_STR)
			cube, _ := utf8.DecodeLastRuneInString(CUBE_ROCK_STR)

			switch char {
			case void:
				p.lines[v] = append(p.lines[v], VOID)
			case round:
				p.lines[v] = append(p.lines[v], ROUND_ROCK)
				p.round_rocks[len(p.round_rocks)] = [2]int{v, h}
			case cube:
				p.lines[v] = append(p.lines[v], CUBE_ROCK)
			}
		}
	}

	if part == 1 {
		p.solve_puzzle()
		solution = p.count_puzzle()
	} else if part == 2 {
		solution = p.solve2_puzzle()
	}

	return solution
}

func main() {
	fmt.Printf("Solution: %d\n", part(2))
}
