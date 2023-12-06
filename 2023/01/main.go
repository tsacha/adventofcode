package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() string {
	calibration_values := utils.PuzzleInput(2023, 1)
	sum := 0

	for _, line := range strings.Split(string(calibration_values), "\n") {
		first_digit := 0
		last_digit := 0
		for _, c := range line {
			if n, err := strconv.Atoi(string(c)); err == nil {
				if first_digit == 0 {
					first_digit = n
				}
				last_digit = n
			}
		}
		n, err := strconv.Atoi(fmt.Sprintf("%d%d", first_digit, last_digit))
		if err != nil {
			fmt.Println("Wrong conversion")
		}
		sum += n
	}
	return fmt.Sprintf("%d\n", sum)
}

func part2() string {
	calibration_values := utils.PuzzleInput(2023, 1)
	//calibration_values := `twone`
	sum := 0
	var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, line := range strings.Split(string(calibration_values), "\n") {
		first_digit := 0
		last_digit := 0
		var word []rune

		for _, char := range line {
			if n, err := strconv.Atoi(string(char)); err == nil {
				if first_digit == 0 {
					first_digit = n
				}
				last_digit = n
			} else {
				word = append(word, char)
				for n, digit_string := range digits {
					// Compare buffer right-to-left to catch some exceptions like "twone"
					if len(word) >= len(digit_string) && string(word[len(word)-len(digit_string):]) == digit_string {
						if first_digit == 0 {
							first_digit = n + 1
						}
						last_digit = n + 1
					}
				}
			}
		}
		n, err := strconv.Atoi(fmt.Sprintf("%d%d", first_digit, last_digit))
		if err != nil {
			fmt.Println("Wrong conversion")
		}
		sum += n
	}
	return fmt.Sprintf("%d\n", sum)
}

func main() {
	fmt.Printf(part1())
	fmt.Printf(part2())
}
