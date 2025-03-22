package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() string {
	input := utils.PuzzleInput(2024, 3)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		for _, instructions := range re.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(instructions[1])
			b, _ := strconv.Atoi(instructions[2])

			result += a * b
		}
	}

	return fmt.Sprintf("%d\n", result)
}

type areas []validArea
type validArea struct {
	begin int
	end   int
}

func part2() string {
	input := utils.PuzzleInput(2024, 3)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	validRe := regexp.MustCompile(`^(.+?)don't\(\)|do\(\)(.+?)don't\(\)|do\(\)(.+?)$`)
	result := 0

	for _, line := range strings.Split(string(input), "\n") {
		areas := areas{}
		if line == "" {
			break
		}

		for _, idx := range validRe.FindAllStringSubmatchIndex(line, -1) {
			for i := 2; i < len(idx); i += 2 {
				if idx[i] > -1 {
					areas = append(areas, validArea{begin: idx[i], end: idx[i+1]})
				}

			}
		}
		for _, instructions := range re.FindAllStringSubmatchIndex(line, -1) {
			validInstruction := false
			for _, area := range areas {
				if instructions[0]-area.begin >= 0 && area.end-instructions[1] >= 0 {
					validInstruction = true
				}
			}

			if validInstruction {
				a, _ := strconv.Atoi(string(line[instructions[2]:instructions[3]]))
				b, _ := strconv.Atoi(string(line[instructions[4]:instructions[5]]))

				result += a * b
			}
		}
	}
	return fmt.Sprintf("%d\n", result)
}

func main() {
	fmt.Printf(part1())
	fmt.Printf(part2())
}
