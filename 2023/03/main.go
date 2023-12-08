package main

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() (solution int) {
	games := string(utils.PuzzleInput(2023, 3))

	matrix := make(map[int]map[int]rune)
	neighbors := map[string]rune{}

	for h, line := range strings.Split(games, "\n") {
		matrix[h] = make(map[int]rune)

		for v, char := range line {
			matrix[h][v] = char
		}
	}

	current_number := 0
	match := false
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			center := matrix[x][y]

			neighbors["top_left"] = matrix[x-1][y-1]
			neighbors["top"] = matrix[x-1][y]
			neighbors["top_right"] = matrix[x-1][y+1]
			neighbors["left"] = matrix[x][y-1]
			neighbors["right"] = matrix[x][y+1]
			neighbors["bottom_left"] = matrix[x+1][y-1]
			neighbors["bottom"] = matrix[x+1][y]
			neighbors["bottom_right"] = matrix[x+1][y+1]

			// Flush number if we are at the beginning of a line
			if neighbors["left"] == 0 {
				current_number = 0
				match = false
			}

			if n, err := strconv.Atoi(string(center)); err == nil {
				// Parse digits

				if current_number > 0 {
					current_number *= 10
				}
				current_number = current_number + n

				for _, n := range neighbors {
					if (n < 48 || n > 57) && n != 46 && n != 0 {
						match = true
					}
				}

				if _, err := strconv.Atoi(string(neighbors["right"])); err != nil {
					// End of number
					if match {
						solution += current_number
					}
				}
			} else {
				// Parse symbols

				current_number = 0
				match = false
			}
		}
	}

	return solution
}

type neighbor struct {
	value rune
	x     int
	y     int
}
type matched_number struct {
	number int
	digits int
	x      int
	y      int
}
type position struct {
	x int
	y int
}

func get_neighbors(matrix map[int]map[int]rune, x int, y int) (neighbors map[string]neighbor) {
	neighbors = make(map[string]neighbor)

	neighbors["top_left"] = neighbor{matrix[x-1][y-1], x - 1, y - 1}
	neighbors["top"] = neighbor{matrix[x-1][y], x - 1, y}
	neighbors["top_right"] = neighbor{matrix[x-1][y+1], x - 1, y + 1}
	neighbors["left"] = neighbor{matrix[x][y-1], x, y - 1}
	neighbors["right"] = neighbor{matrix[x][y+1], x, y + 1}
	neighbors["bottom_left"] = neighbor{matrix[x+1][y-1], x + 1, y - 1}
	neighbors["bottom"] = neighbor{matrix[x+1][y], x + 1, y}
	neighbors["bottom_right"] = neighbor{matrix[x+1][y+1], x + 1, y + 1}

	return neighbors
}

func part2() (solution int) {
	games := string(utils.PuzzleInput(2023, 3))

	matrix := make(map[int]map[int]rune)
	numbers := []matched_number{}
	gears := []position{}

	for h, line := range strings.Split(games, "\n") {
		matrix[h] = make(map[int]rune)

		for v, char := range line {
			matrix[h][v] = char
		}
	}

	current_number, digits := 0, 0
	match := false
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			center := matrix[x][y]
			neighbors := get_neighbors(matrix, x, y)

			// Flush number if we are at the beginning of a line
			if neighbors["left"].value == 0 {
				current_number, digits = 0, 0
				match = false
			}

			if n, err := strconv.Atoi(string(center)); err == nil {
				// Parse digits

				if current_number > 0 {
					current_number *= 10
					digits++
				}
				current_number = current_number + n

				for _, n := range neighbors {
					if (n.value < 48 || n.value > 57) && n.value != 46 && n.value != 0 {
						match = true
					}
				}

				if _, err := strconv.Atoi(string(neighbors["right"].value)); err != nil {
					// End of number
					if match {
						numbers = append(numbers, matched_number{current_number, digits + 1, x, y - digits})
					}
				}
			} else {
				// Parse symbols

				current_number, digits = 0, 0
				match = false
				if center == 42 {
					gears = append(gears, position{x, y})
				}
			}
		}
	}

	// Loop through gears
	for _, g := range gears {
		gear_ratio := 0
		adjacent_numbers := make(map[position]int)
		neighbors := get_neighbors(matrix, g.x, g.y)

		for _, n := range neighbors {
			if n.value >= 48 && n.value <= 57 {
				for _, m := range numbers {
					for i := 0; i < m.digits; i++ {
						if m.x == n.x && m.y+i == n.y {
							adjacent_numbers[position{m.x, m.y}] = m.number
						}
					}
				}
			}
		}

		if len(adjacent_numbers) > 1 {
			for _, n := range adjacent_numbers {
				if gear_ratio == 0 {
					gear_ratio = n
				} else {
					gear_ratio *= n
				}

			}
			solution += gear_ratio
		}
	}

	return solution
}

func main() {
	fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("Solution: %d\n", part2())
}
