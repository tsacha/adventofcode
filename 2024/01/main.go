package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type locations struct {
	left       []int
	right      []int
	occurences map[int]int

	distances  int
	similarity int
}

func part1() string {
	input := utils.PuzzleInput(2024, 1)
	locations := locations{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		s := strings.Split(line, "   ")
		left, _ := strconv.Atoi(s[0])
		right, _ := strconv.Atoi(s[1])

		locations.left = append(locations.left, left)
		locations.right = append(locations.right, right)
	}

	sort.Ints(locations.left)
	sort.Ints(locations.right)

	for n := range locations.left {
		locations.distances += int(math.Abs(float64(locations.left[n] - locations.right[n])))
	}

	return fmt.Sprintf("%d\n", locations.distances)
}

func part2() string {
	input := utils.PuzzleInput(2024, 1)
	locations := locations{}
	locations.occurences = make(map[int]int)

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		s := strings.Split(line, "   ")
		left, _ := strconv.Atoi(s[0])
		right, _ := strconv.Atoi(s[1])

		locations.left = append(locations.left, left)
		locations.occurences[right]++
	}

	sort.Ints(locations.left)
	sort.Ints(locations.right)

	for _, n := range locations.left {
		locations.similarity += n * locations.occurences[n]
	}

	return fmt.Sprintf("%d\n", locations.similarity)
}

func main() {
	fmt.Printf(part1())
	fmt.Printf(part2())
}
