package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func part1() (result int) {
	input := string(utils.PuzzleInput(2023, 15))
	// examples
	_ = `HASH`
	_ = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	for _, step := range strings.Split(strings.Trim(input, "\n"), ",") {
		solution := 0
		for _, c := range step {
			solution += int(c)
			solution *= 17
			solution = solution % 256
		}
		result += solution

	}
	return result
}

func hash(label string) int {
	solution := 0
	for _, c := range label {
		solution += int(c)
		solution *= 17
		solution = solution % 256
	}
	return solution
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func part2() (result int) {
	input := string(utils.PuzzleInput(2023, 15))
	boxes := [256][]string{}
	lenses := map[string]int{}
	for _, step := range strings.Split(strings.Trim(input, "\n"), ",") {
		s1 := strings.Split(step, "=")
		s2 := strings.Split(s1[0], "-")

		var label string
		var add bool
		var lens int

		if len(s1) > 1 {
			label = s1[0]
			add = true
			lens, _ = strconv.Atoi(s1[1])
		} else if len(s2) > 1 {
			label = s2[0]
			add = false
			lens = -1
		}

		hash := hash(label)

		if add {
			lenses[label] = lens
			exists := false
			for _, l := range boxes[hash] {
				if l == label {
					exists = true
				}
			}
			if !exists {
				boxes[hash] = append(boxes[hash], label)
			}
		} else {
			exists := false
			var box []string
			for nl, l := range boxes[hash] {
				if l == label {
					exists = true
					box = remove(boxes[hash], nl)
				}
			}
			if exists {
				boxes[hash] = box
			}
		}
	}

	for n_box, box := range boxes {
		for n_item, item := range box {
			result += (n_box + 1) * (n_item + 1) * lenses[item]
		}
	}

	return result
}

func main() {
	fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("Solution: %d\n", part2())
}
