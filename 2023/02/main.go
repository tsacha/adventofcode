package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() (solution int) {
	games := utils.PuzzleInput(2023, 2)
	solution = 5050 // 1+2+3+..+100
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, line := range strings.Split(string(games), "\n") {
		if line == "" {
			break
		}

		s := strings.Split(string(line), ": ")
		game_name, game_sets := s[0], s[1]
		game_id, err := strconv.Atoi(strings.Split(game_name, "Game ")[1])
		if err != nil {
			fmt.Printf("Wrong game ID")
			break
		}

		sets := strings.Split(game_sets, "; ")
	out:
		for _, set := range sets {
			set_counter := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			subsets := strings.Split(set, ", ")
			for _, subset := range subsets {
				s = strings.Split(subset, " ")
				nb_color_str, color := s[0], s[1]
				color_counter, err := strconv.Atoi(nb_color_str)
				if err != nil {
					fmt.Printf("Wrong cubes number")
					break
				}

				set_counter[color] += color_counter
			}
			for color, color_counter := range set_counter {
				if color_counter > rules[color] {
					//fmt.Printf("Game %d is invalid! Too many %ss (%d > %d)\n", game_id, color, color_counter, rules[color])
					solution -= game_id
					break out
				}
			}
		}
	}
	return solution
}

func part2() (solution int) {
	games := utils.PuzzleInput(2023, 2)
	solution = 0
	for _, line := range strings.Split(string(games), "\n") {
		if line == "" {
			break
		}

		s := strings.Split(string(line), ": ")
		game_name, game_sets := s[0], s[1]
		_, err := strconv.Atoi(strings.Split(game_name, "Game ")[1])
		if err != nil {
			fmt.Printf("Wrong game ID")
			break
		}

		sets := strings.Split(game_sets, "; ")
		set_counter := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range sets {
			subsets := strings.Split(set, ", ")
			for _, subset := range subsets {
				s = strings.Split(subset, " ")
				nb_color_str, color := s[0], s[1]
				color_counter, err := strconv.Atoi(nb_color_str)
				if err != nil {
					fmt.Printf("Wrong cubes number")
					break
				}
				set_counter[color] = max(color_counter, set_counter[color])
			}
		}
		//fmt.Printf("%d: %v#\n", game_id, set_counter)
		solution += (set_counter["red"] * set_counter["green"] * set_counter["blue"])
	}
	return solution
}

func main() {
	fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("Solution: %d\n", part2())
}
