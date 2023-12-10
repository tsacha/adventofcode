package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type path map[string]string
type way map[string]path

func part1() (solution int) {
	input := string(utils.PuzzleInput(2023, 9))
	_ = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`
	prediction := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var dataset [][]int
		history := strings.Fields(line)

		dataset = append(dataset, []int{})
		for _, h := range history {
			n, _ := strconv.Atoi(h)
			dataset[0] = append(dataset[0], n)
		}

		step := 1
		not_zero := true
		for not_zero == true {
			dataset = append(dataset, []int{})
			dataset[step] = []int{}

			not_zero = false
			for p := range dataset[len(dataset)-2] {
				if p >= len(dataset[step-1])-1 {
					break
				}

				dataset[step] = append(dataset[step], dataset[step-1][p+1]-dataset[step-1][p])
				if dataset[step-1][p] != 0 {
					not_zero = true
				}
			}

			if not_zero == false {
				fmt.Print("----\n")
				dataset[step-1] = append(dataset[step-1], 0)

				for s := len(dataset) - 2; s > 0; s-- {
					dataset[s-1] = append(dataset[s-1], dataset[s-1][len(dataset[s-1])-1]+dataset[s][len(dataset[s])-1])
				}
				prediction += dataset[0][len(dataset[0])-1]
			}
			step++
		}
	}
	return prediction
}

func part2() (solution int) {
	input := string(utils.PuzzleInput(2023, 9))
	_ = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	prediction := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var dataset [][]int
		history := strings.Fields(line)

		dataset = append(dataset, []int{})
		for _, h := range history {
			n, _ := strconv.Atoi(h)
			dataset[0] = append(dataset[0], n)
		}

		step := 1
		not_zero := true
		for not_zero == true {
			fmt.Printf("[step %d] \n", step)

			dataset = append(dataset, []int{})
			dataset[step] = []int{}

			not_zero = false
			for p := range dataset[len(dataset)-2] {
				if p >= len(dataset[step-1])-1 {
					break
				}

				dataset[step] = append(dataset[step], dataset[step-1][p+1]-dataset[step-1][p])
				if dataset[step-1][p] != 0 {
					not_zero = true
				}
			}

			if not_zero == false {
				fmt.Print("----\n")
				fmt.Printf("%v\n", dataset)
				dataset[step-1] = append([]int{0}, dataset[step-1]...)

				for s := len(dataset) - 2; s > 0; s-- {
					dataset[s-1] = append([]int{dataset[s-1][0] - dataset[s][0]}, dataset[s-1]...)
				}
				prediction += dataset[0][0]
			}
			step++
		}
	}
	return prediction
}

func main() {
	fmt.Printf("Solution 2: %d\n", part2())
}
