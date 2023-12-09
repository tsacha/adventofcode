package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() (solution int) {
	games := string(utils.PuzzleInput(2023, 4))

	for _, card := range strings.Split(games, "\n") {
		card_points := 0

		l := strings.Split(card, ": ")
		if len(l) < 2 {
			continue
		}
		r := strings.Split(l[1], " | ")

		winning_numbers_str := strings.Split(r[0], " ")
		numbers_str := strings.Split(r[1], " ")

		winning_numbers := []int{}
		for _, n := range winning_numbers_str {
			p, err := strconv.Atoi(n)
			if err == nil {
				winning_numbers = append(winning_numbers, p)
			}
		}
		for _, n := range numbers_str {
			p, err := strconv.Atoi(n)
			if err == nil {
				if slices.Contains(winning_numbers, p) {
					if card_points == 0 {
						card_points++
					} else {
						card_points *= 2
					}
				}
			}
		}
		solution += card_points
	}
	return solution
}

func part2() (solution int) {
	games := string(utils.PuzzleInput(2023, 4))
	scratchboards := make(map[int]int)
	for _, card := range strings.Split(games, "\n") {
		l := strings.Split(card, ": ")
		if len(l) < 2 {
			continue
		}

		card_number_str := strings.Fields(l[0])[1]
		card_number, err := strconv.Atoi(card_number_str)
		if err != nil {
			continue
		}
		scratchboards[card_number]++

		r := strings.Split(l[1], " | ")
		winning_numbers_str := strings.Split(r[0], " ")
		numbers_str := strings.Split(r[1], " ")

		winning_numbers := map[int][]int{}
		numbers := map[int][]int{}
		for _, n := range winning_numbers_str {
			p, err := strconv.Atoi(n)
			if err == nil {
				winning_numbers[card_number] = append(winning_numbers[card_number], p)
			}
		}
		for _, n := range numbers_str {
			p, err := strconv.Atoi(n)
			if err == nil {
				numbers[card_number] = append(numbers[card_number], p)
			}
		}

		matching_numbers := matching_numbers(winning_numbers[card_number], numbers[card_number])
		for c := 1; c <= scratchboards[card_number]; c++ {
			for m := card_number + 1; m <= card_number+matching_numbers; m++ {
				scratchboards[m]++
			}
		}
	}
	for _, n := range scratchboards {
		solution += n
	}

	return solution
}

func matching_numbers(winning_numbers []int, numbers []int) (matching_numbers int) {
	for _, n := range numbers {
		if slices.Contains(winning_numbers, n) {
			matching_numbers++
		}
	}
	return matching_numbers
}

func main() {
	fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("Solution: %d\n", part2())
}
